package kardiachain

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kardiachain/go-kardia/dualnode/chains/kardiachain/bridge"
	tssAbi "github.com/kardiachain/go-kardia/dualnode/chains/kardiachain/tss"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	dualTypes "github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/abi/bind"
	"github.com/kardiachain/go-kardia/lib/common"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/kardiachain/go-kardia/rpc"
	"github.com/kardiachain/go-kardia/types"
)

type Watcher struct {
	quit chan struct{}

	client *KardiaClient

	minConfirmations uint64
	pendingLocks     map[common.Hash]*dualTypes.DualEvent
	pendingLocksMtx  sync.Mutex

	store     *store.Store
	depositC  chan *dproto.Deposit
	withdrawC chan dualTypes.Withdraw
	vsChan    chan *dualTypes.ValidatorSet
	internalC struct {
		depositedC chan *bridge.BridgeDeposited
		withdrawC  chan *bridge.BridgeWithdraw

		vaultCreatedC      chan *tssAbi.TssVaultCreated
		vaultUpdatedC      chan *tssAbi.TssVaultUpdated
		vaultChainEditedC  chan *tssAbi.TssVaultChainEdited
		vaultChainRemovedC chan *tssAbi.TssVaultChainRemoved
		tokenAddedC        chan *tssAbi.TssTokenAdded
		tokenRemovedC      chan *tssAbi.TssTokenRemoved
	}
}

func newWatcher(client *KardiaClient, s *store.Store, depositC chan *dproto.Deposit, withdrawC chan dualTypes.Withdraw) *Watcher {
	return &Watcher{
		quit:         make(chan struct{}, 1),
		client:       client,
		pendingLocks: make(map[common.Hash]*dualTypes.DualEvent),
		store:        s,
		depositC:     depositC,
		withdrawC:    withdrawC,
	}
}

func (w *Watcher) Start() error {
	go func() {
		if err := w.Watch(); err != nil {
			fmt.Printf("watch blocks error: %s", err)
		}
	}()
	return nil
}

func (w *Watcher) Stop() error {
	close(w.quit)
	return nil
}

func (w *Watcher) Watch() error {
	lgr := w.client.logger
	ctx := context.Background()
	errC := make(chan error)

	err := w.handleBridgeEvents(ctx, errC)
	if err != nil {
		lgr.Error("error while handling bridge events", "err", err)
		errC <- fmt.Errorf("error while handling bridge events: %w", err)
		return err
	}
	err = w.handleTssEvents(ctx, errC)
	if err != nil {
		lgr.Error("error while handling tssAbi events", "err", err)
		errC <- fmt.Errorf("error while handling tssAbi events: %w", err)
		return err
	}
	err = w.handleNewHeaders(ctx, errC)
	if err != nil {
		lgr.Error("error while handling new headers", "err", err)
		errC <- fmt.Errorf("error while handling new headers: %w", err)
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}

func (w *Watcher) handleBridgeEvents(ctx context.Context, errC chan error) error {
	lgr := w.client.logger
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	bridgeF, err := bridge.NewBridgeFilterer(common.HexToAddress(w.client.ChainConfig.BridgeSmcAddr), w.client.KAIClient)
	if err != nil {
		return fmt.Errorf("could not create KAI bridge filter: %w", err)
	}
	// Subscribe to token deposited events
	depositedSub, err := bridgeF.WatchDeposited(&bind.WatchOpts{Context: timeout}, w.internalC.depositedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token deposited events: %w", err)
	}
	// Subscribe to token withdraw events
	withdrawSub, err := bridgeF.WatchWithdraw(&bind.WatchOpts{Context: timeout}, w.internalC.withdrawC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}

	// watch for dual events
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-depositedSub.Err():
				lgr.Error("error while processing deposited event", "err", err)
				errC <- fmt.Errorf("error while processing deposited event: %w", err)
				return
			case err := <-withdrawSub.Err():
				lgr.Error("error while processing withdraw event", "err", err)
				errC <- fmt.Errorf("error while processing withdraw event: %w", err)
				return

			case ev := <-w.internalC.depositedC:
				blockHeight := rpc.BlockNumber(ev.Raw.BlockHeight)
				b, err := w.client.KAIClient.BlockByHeight(ctx, &blockHeight)
				if err != nil {
					lgr.Error("error while getting block from KAI client", "err", err)
					errC <- fmt.Errorf("error while getting block from KAI client: %w", err)
					return
				}
				dualEv := &dualTypes.DualEvent{
					Source:      1,
					Destination: ev.DestChainId.Int64(),
					Arguments:   ev,
					Timestamp:   b.Time(),
					RawName:     dualCfg.DepositedEventRawName,
					Raw:         ev.Raw,
				}
				w.pendingLocksMtx.Lock()
				w.pendingLocks[ev.Raw.TxHash] = dualEv
				w.pendingLocksMtx.Unlock()
			case ev := <-w.internalC.withdrawC:
				blockHeight := rpc.BlockNumber(ev.Raw.BlockHeight)
				b, err := w.client.KAIClient.BlockByHeight(ctx, &blockHeight)
				if err != nil {
					lgr.Error("error while getting block from KAI client", "err", err)
					errC <- fmt.Errorf("error while getting block from KAI client: %w", err)
					return
				}
				dualEv := &dualTypes.DualEvent{
					Source:      ev.SourceChainId.Int64(),
					Destination: 1,
					Arguments:   ev,
					Timestamp:   b.Time(),
					RawName:     dualCfg.WithdrawEventRawName,
					Raw:         ev.Raw,
				}
				w.pendingLocksMtx.Lock()
				w.pendingLocks[ev.Raw.TxHash] = dualEv
				w.pendingLocksMtx.Unlock()
			}
		}
	}()
	return nil
}

func (w *Watcher) handleTssEvents(ctx context.Context, errC chan error) error {
	lgr := w.client.logger
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	tssF, err := tssAbi.NewTssFilterer(common.HexToAddress(w.client.ChainConfig.BridgeSmcAddr), w.client.KAIClient)
	if err != nil {
		return fmt.Errorf("could not create KAI bridge filter: %w", err)
	}
	// Subscribe to token added events
	tokenAddedSub, err := tssF.WatchTokenAdded(&bind.WatchOpts{Context: timeout}, w.internalC.tokenAddedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token added events: %w", err)
	}
	// Subscribe to token removed events
	tokenRemovedSub, err := tssF.WatchTokenRemoved(&bind.WatchOpts{Context: timeout}, w.internalC.tokenRemovedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}
	// Subscribe to vault created events
	vaultCreatedSub, err := tssF.WatchVaultCreated(&bind.WatchOpts{Context: timeout}, w.internalC.vaultCreatedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}
	// Subscribe to vault updated events
	vaultUpdatedSub, err := tssF.WatchVaultUpdated(&bind.WatchOpts{Context: timeout}, w.internalC.vaultUpdatedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}
	// Subscribe to vault edited events
	vaultChainEditedSub, err := tssF.WatchVaultChainEdited(&bind.WatchOpts{Context: timeout}, w.internalC.vaultChainEditedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}
	// Subscribe to vault removes events
	vaultChainRemovedSub, err := tssF.WatchVaultChainRemoved(&bind.WatchOpts{Context: timeout}, w.internalC.vaultChainRemovedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-tokenAddedSub.Err():
				lgr.Error("error while processing token added event", "err", err)
				errC <- fmt.Errorf("error while processing token added event: %w", err)
				return
			case err := <-tokenRemovedSub.Err():
				lgr.Error("error while processing token removed event", "err", err)
				errC <- fmt.Errorf("error while processing token removed event: %w", err)
				return
			case err := <-vaultCreatedSub.Err():
				lgr.Error("error while processing vault created event", "err", err)
				errC <- fmt.Errorf("error while processing vault created event: %w", err)
				return
			case err := <-vaultUpdatedSub.Err():
				lgr.Error("error while processing vault created event", "err", err)
				errC <- fmt.Errorf("error while processing vault created event: %w", err)
				return
			case err := <-vaultChainEditedSub.Err():
				lgr.Error("error while processing vault created event", "err", err)
				errC <- fmt.Errorf("error while processing vault created event: %w", err)
				return
			case err := <-vaultChainRemovedSub.Err():
				lgr.Error("error while processing vault created event", "err", err)
				errC <- fmt.Errorf("error while processing vault created event: %w", err)
				return

			case <-w.internalC.tokenAddedC:
			case <-w.internalC.tokenRemovedC:
			case <-w.internalC.vaultCreatedC:

			case <-w.internalC.vaultUpdatedC:
			case <-w.internalC.vaultChainEditedC:
			case <-w.internalC.vaultChainRemovedC:
			}
		}
	}()
	return nil
}

func (w *Watcher) handleNewHeaders(ctx context.Context, errC chan error) error {
	lgr := w.client.logger
	// Watch for new headers
	headSink := make(chan *types.Header, 2)
	headerSubscription, err := w.client.KAIClient.SubscribeNewHead(ctx, headSink)
	if err != nil {
		return fmt.Errorf("failed to subscribe to header events: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-headerSubscription.Err():
				errC <- fmt.Errorf("error while processing header subscription: %w", e)
				return
			case ev := <-headSink:
				start := time.Now()
				lgr.Info("processing new header", "block", ev.Height)

				w.pendingLocksMtx.Lock()

				for hash, pLock := range w.pendingLocks {

					// Transaction was dropped and never picked up again
					if pLock.Raw.BlockHeight+4*w.minConfirmations <= ev.Height {
						lgr.Debug("lockup timed out", "tx", pLock.Raw.TxHash, "block", ev.Height)
						delete(w.pendingLocks, hash)
						continue
					}

					// Transaction is now ready
					if pLock.Raw.BlockHeight+w.minConfirmations <= ev.Height {
						lgr.Debug("lockup confirmed", "tx", pLock.Raw.TxHash, "block", ev.Height)
						delete(w.pendingLocks, hash)
						if pLock.RawName == dualCfg.DepositedEventRawName {
							//w.depositC <- pLock
						} else if pLock.RawName == dualCfg.DepositedEventRawName {
							//w.withdrawC <- pLock
						}
					}
				}

				w.pendingLocksMtx.Unlock()
				lgr.Info("processed new header", "block", ev.Height,
					"took", time.Since(start))
			}
		}
	}()
	return nil
}
