package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/kardiachain/go-kardia/lib/common"

	"github.com/kardiachain/go-kardia/dualnode/chains/ethereum/bridge"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	kaiTypes "github.com/kardiachain/go-kardia/types"
)

type Watcher struct {
	quit chan struct{}

	client *ETHLightClient

	minConfirmations uint64
	pendingLocks     map[ethCommon.Hash]*types.DualEvent
	pendingLocksMtx  sync.Mutex

	store     *store.Store
	depositC  chan *dproto.Deposit
	withdrawC chan types.Withdraw
	vsChan    chan *types.ValidatorSet
	internalC struct {
		depositedC chan *bridge.BridgeDeposited
		withdrawC  chan *bridge.BridgeWithdraw
	}
}

func newWatcher(client *ETHLightClient, s *store.Store, depositC chan *dproto.Deposit, withdrawC chan types.Withdraw, vsC chan *types.ValidatorSet) *Watcher {
	return &Watcher{
		quit:         make(chan struct{}, 1),
		client:       client,
		store:        s,
		pendingLocks: make(map[ethCommon.Hash]*types.DualEvent),
		depositC:     depositC,
		withdrawC:    withdrawC,
		vsChan:       vsC,
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
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	f, err := bridge.NewBridgeFilterer(ethCommon.HexToAddress(w.client.ChainConfig.BridgeSmcAddr), w.client.ETHClient)
	if err != nil {
		return fmt.Errorf("could not create ETH bridge filter: %w", err)
	}

	// Subscribe to token deposited events
	depositedSub, err := f.WatchDeposited(&bind.WatchOpts{Context: timeout}, w.internalC.depositedC)
	if err != nil {
		return fmt.Errorf("failed to subscribe to token deposited events: %w", err)
	}

	// Subscribe to token withdraw events
	withdrawSub, err := f.WatchWithdraw(&bind.WatchOpts{Context: timeout}, w.internalC.withdrawC)
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

			case ev := <-w.internalC.depositedC:
				b, err := w.client.ETHClient.BlockByNumber(ctx, new(big.Int).SetUint64(ev.Raw.BlockNumber))
				if err != nil {
					lgr.Error("error while getting block from ETH client", "err", err)
					errC <- fmt.Errorf("error while getting block from ETH client: %w", err)
					return
				}
				dualEv := &types.DualEvent{
					Source:      1,
					Destination: ev.DestChainId.Int64(),
					Arguments:   ev,
					Timestamp:   b.ReceivedAt,
					RawName:     dualCfg.DepositedEventRawName,
					Raw:         convertToKAILog(ev.Raw),
				}
				w.pendingLocksMtx.Lock()
				w.pendingLocks[ev.Raw.TxHash] = dualEv
				w.pendingLocksMtx.Unlock()
			case ev := <-w.internalC.withdrawC:
				b, err := w.client.ETHClient.BlockByNumber(ctx, new(big.Int).SetUint64(ev.Raw.BlockNumber))
				if err != nil {
					lgr.Error("error while getting block from ETH client", "err", err)
					errC <- fmt.Errorf("error while getting block from ETH client: %w", err)
					return
				}
				dualEv := &types.DualEvent{
					Source:      ev.SourceChainId.Int64(),
					Destination: 1,
					Arguments:   ev,
					Timestamp:   b.ReceivedAt,
					RawName:     dualCfg.WithdrawEventRawName,
					Raw:         convertToKAILog(ev.Raw),
				}
				w.pendingLocksMtx.Lock()
				w.pendingLocks[ev.Raw.TxHash] = dualEv
				w.pendingLocksMtx.Unlock()
			}
		}
	}()

	// Watch for new headers
	headSink := make(chan *ethTypes.Header, 2)
	headerSubscription, err := w.client.ETHClient.SubscribeNewHead(ctx, headSink)
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
				lgr.Info("processing new header", "block", ev.Number)

				w.pendingLocksMtx.Lock()

				blockNumberU := ev.Number.Uint64()
				for hash, pLock := range w.pendingLocks {

					// Transaction was dropped and never picked up again
					if pLock.Raw.BlockHeight+4*w.minConfirmations <= blockNumberU {
						lgr.Debug("lockup timed out", "tx", pLock.Raw.TxHash, "block", ev.Number)
						delete(w.pendingLocks, hash)
						continue
					}

					// Transaction is now ready
					if pLock.Raw.BlockHeight+w.minConfirmations <= ev.Number.Uint64() {
						lgr.Debug("lockup confirmed", "tx", pLock.Raw.TxHash, "block", ev.Number)
						delete(w.pendingLocks, hash)
						if pLock.RawName == dualCfg.DepositedEventRawName {
							//w.depositC <- pLock
						} else if pLock.RawName == dualCfg.DepositedEventRawName {
							//w.withdrawC <- pLock
						}
					}
				}

				w.pendingLocksMtx.Unlock()
				lgr.Info("processed new header", "block", ev.Number,
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

func convertToKAILog(ethLog ethTypes.Log) kaiTypes.Log {
	topics := make([]common.Hash, len(ethLog.Topics))
	for i := range ethLog.Topics {
		topics[i] = common.BytesToHash(ethLog.Topics[i].Bytes())
	}
	return kaiTypes.Log{
		Address:     common.HexToAddress(ethLog.Address.Hex()),
		Topics:      topics,
		Data:        ethLog.Data,
		BlockHeight: ethLog.BlockNumber,
		TxHash:      common.HexToHash(ethLog.TxHash.Hex()),
		TxIndex:     ethLog.TxIndex,
		BlockHash:   common.HexToHash(ethLog.BlockHash.Hex()),
		Index:       ethLog.Index,
		Removed:     ethLog.Removed,
	}
}
