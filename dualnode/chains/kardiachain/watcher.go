package kardiachain

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kardiachain/go-kardia/dualnode/chains/kardiachain/bridge"
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
}

func newWatcher(client *KardiaClient, cfg *dualCfg.ChainManagerConfig) *Watcher {
	return &Watcher{
		quit:         make(chan struct{}, 1),
		client:       client,
		pendingLocks: make(map[common.Hash]*dualTypes.DualEvent),
		store:        cfg.S,
		depositC:     cfg.DepositC,
		withdrawC:    cfg.WithdrawC,
		vsChan:       cfg.VsChan,
	}
}

func (w *Watcher) Start() error {
	depositedC := make(chan *bridge.BridgeDeposited, 2)
	withdrawC := make(chan *bridge.BridgeWithdraw, 2)
	go func() {
		if err := w.watch(depositedC, withdrawC); err != nil {
			fmt.Printf("watch blocks error: %s", err)
		}
	}()
	return nil
}

func (w *Watcher) Stop() error {
	close(w.quit)
	return nil
}

func (w *Watcher) watch(depositedC chan *bridge.BridgeDeposited, withdrawC chan *bridge.BridgeWithdraw) error {
	lgr := w.client.logger
	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	f, err := bridge.NewBridgeFilterer(common.HexToAddress(w.client.ChainConfig.BridgeSmcAddr), w.client.KAIClient)
	if err != nil {
		return fmt.Errorf("could not create KAI bridge filter: %w", err)
	}

	// Subscribe to token deposited events
	depositedSub, err := f.WatchDeposited(&bind.WatchOpts{Context: timeout}, depositedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token deposited events: %w", err)
	}

	// Subscribe to token withdraw events
	withdrawSub, err := f.WatchWithdraw(&bind.WatchOpts{Context: timeout}, withdrawC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token withdraw events: %w", err)
	}

	errC := make(chan error)

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

			case ev := <-depositedC:
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
			case ev := <-withdrawC:
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

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}
