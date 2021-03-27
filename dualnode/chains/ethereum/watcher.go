package ethereum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
)

type Watcher struct {
	quit chan struct{}

	client *ETHLightClient
	events chan *abi.Event

	checkpoint uint64
	dualTopics [][]ethCommon.Hash
}

func newWatcher(client *ETHLightClient) *Watcher {
	return &Watcher{
		quit:   make(chan struct{}, 1),
		client: client,
		events: make(chan *abi.Event, dualCfg.DualEventChanSize),
	}
}

func (w *Watcher) start() error {
	// update checkpoint
	if w.checkpoint <= 0 {
		latestBlock, err := w.client.ETHClient.BlockByNumber(w.client.ctx, nil)
		if err != nil {
			w.client.logger.Error("Cannot get latest ETH block", "err", err)
			return err
		}
		w.checkpoint = latestBlock.NumberU64()
	}
	go func() {
		if err := w.watch(); err != nil {
			fmt.Printf("watch blocks error: %s", err)
		}
	}()
	return nil
}

func (w *Watcher) stop() error {
	//w.quit <- struct{}{}
	close(w.quit)
	return nil
}

func (w *Watcher) watch() error {
	// init a ticker for polling dual events
	var (
		pollingEventsFreq = dualCfg.DualEventFreq
		pollingEventsCh   = make(chan struct{}, 1)
		pollingEventsTk   = time.NewTicker(pollingEventsFreq)
	)
	defer pollingEventsTk.Stop()
	for {
		select {
		case <-w.quit:
			return nil
		case <-pollingEventsTk.C:
			select {
			case pollingEventsCh <- struct{}{}:
			default:
			}

		case <-pollingEventsCh:
			// read dual events from filtered logs
			logs, err := w.getLatestDualEvents()
			if err != nil {
				w.client.logger.Warn("Cannot get latest dual events", "err", err, "checkpoint", w.checkpoint)
				continue
			}
			// send dual events to events channel
			for _, log := range logs {
				decodedLog, err := w.client.ChainConfig.SwapSMC.ABI.EventByID(log.Topics[0])
				if err != nil {
					w.client.logger.Warn("Cannot decode dual event", "err", err, "checkpoint", w.checkpoint, "log", log)
					continue
				}
				w.events <- decodedLog
			}
		default:
		}
	}
}

func (w *Watcher) getLatestDualEvents() ([]ethTypes.Log, error) {
	latestBlock, err := w.client.ETHClient.BlockByNumber(w.client.ctx, nil)
	if err != nil {
		w.client.logger.Error("Cannot get latest ETH block", "err", err)
		return nil, err
	}
	if w.checkpoint > latestBlock.NumberU64() {
		// prevent grabbing events of a block multiple times
		return nil, nil
	}
	topics, err := w.getDualEventTopics()
	if err != nil {
		w.client.logger.Error("Cannot get dual event topics", "err", err)
		return nil, err
	}
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(w.checkpoint),
		ToBlock:   new(big.Int).SetUint64(latestBlock.NumberU64()),
		Addresses: []ethCommon.Address{w.client.ChainConfig.SwapSMC.Address},
		Topics:    topics,
	}
	w.checkpoint = latestBlock.NumberU64() + 1 // increase checkpoint to prevent grabbing events of a block multiple times
	fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@ query %+v\n", query)
	w.client.logger.Debug("Dual events query", "query", query)
	logs, err := w.client.ETHClient.FilterLogs(w.client.ctx, query)
	if err != nil {
		w.client.logger.Error("Cannot get dual event", "err", err, "FromBlock", query.FromBlock.Uint64(), "ToBlock", query.ToBlock.Uint64())
		return nil, err
	}
	return logs, err
}

func (w *Watcher) getDualEventTopics() ([][]ethCommon.Hash, error) {
	if w.dualTopics != nil {
		return w.dualTopics, nil
	}
	var (
		swapSMCABI = w.client.ChainConfig.SwapSMC.ABI
		topics     []ethCommon.Hash
	)
	for i := range swapSMCABI.Events {
		topics = append(topics, swapSMCABI.Events[i].ID)
		w.client.logger.Debug("Appending dual topics...", "topic", swapSMCABI.Events[i].ID)
	}
	w.dualTopics = [][]ethCommon.Hash{topics}
	w.client.logger.Info("Dual topics", "topics", topics)
	return w.dualTopics, nil
}
