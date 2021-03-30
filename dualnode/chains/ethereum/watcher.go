package ethereum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"

	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/types"
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

func (w *Watcher) Start() error {
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

func (w *Watcher) Stop() error {
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
			logs, err := w.GetLatestDualEvents()
			if err != nil {
				w.client.logger.Warn("Cannot get latest dual events", "err", err, "checkpoint", w.checkpoint)
				continue
			}
			// send dual events to events channel
			for _, log := range logs {
				decodedLog, err := w.client.SwapSMC.ABI.EventByID(log.Topics[0])
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

func (w *Watcher) GetLatestDualEvents() ([]types.Log, error) {
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
		Addresses: []ethCommon.Address{w.client.SwapSMC.Address},
		Topics:    topics,
	}
	w.checkpoint = latestBlock.NumberU64() + 1 // increase checkpoint to prevent grabbing events of a block multiple times
	w.client.logger.Debug("Dual events query", "query", query)
	logs, err := w.client.ETHClient.FilterLogs(w.client.ctx, query)
	if err != nil {
		w.client.logger.Error("Cannot get dual event", "err", err, "FromBlock", query.FromBlock.Uint64(), "ToBlock", query.ToBlock.Uint64())
		return nil, err
	}
	result := make([]types.Log, len(logs))
	for i := range logs {
		result[i] = types.Log{
			Address:     common.HexToAddress(logs[i].Address.Hex()),
			Topics:      convertTopics(logs[i].Topics),
			Data:        logs[i].Data,
			BlockHeight: logs[i].BlockNumber,
			TxHash:      common.BytesToHash(logs[i].TxHash.Bytes()),
			TxIndex:     logs[i].TxIndex,
			BlockHash:   common.BytesToHash(logs[i].BlockHash.Bytes()),
			Index:       logs[i].Index,
			Removed:     logs[i].Removed,
		}
	}
	return result, err
}

func (w *Watcher) getDualEventTopics() ([][]ethCommon.Hash, error) {
	if w.dualTopics != nil {
		return w.dualTopics, nil
	}
	var (
		swapSMCABI = w.client.SwapSMC.ABI
		topics     []ethCommon.Hash
	)
	for i := range swapSMCABI.Events {
		topics = append(topics, ethCommon.BytesToHash(swapSMCABI.Events[i].ID.Bytes()))
		w.client.logger.Debug("Appending dual topics...", "topic", swapSMCABI.Events[i].ID)
	}
	w.dualTopics = [][]ethCommon.Hash{topics}
	w.client.logger.Info("Dual topics", "topics", topics)
	return w.dualTopics, nil
}

func convertTopics(ethTopics []ethCommon.Hash) []common.Hash {
	result := make([]common.Hash, len(ethTopics))
	for i := range ethTopics {
		result[i] = common.BytesToHash(ethTopics[i].Bytes())
	}
	return result
}
