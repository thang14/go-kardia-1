package ethereum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"

	dualCmn "github.com/kardiachain/go-kardia/dualnode/common"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	dualTypes "github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/common"
)

type Watcher struct {
	quit chan struct{}

	client *ETHLightClient
	events chan *dualTypes.DualEvent

	store      *store.Store
	checkpoint uint64
	dualTopics [][]ethCommon.Hash
}

func newWatcher(client *ETHLightClient, store *store.Store) *Watcher {
	return &Watcher{
		quit:   make(chan struct{}, 1),
		client: client,
		store:  store,
		events: make(chan *dualTypes.DualEvent, dualCfg.DualEventChanSize),
	}
}

func (w *Watcher) GetDualEventsChannel() chan *dualTypes.DualEvent {
	return w.events
}

func (w *Watcher) Start() error {
	checkpoint, err := w.store.GetCheckpoint(w.client.ChainConfig.ChainID)
	if err != nil {
		w.client.logger.Error("Cannot get old checkpoint", "chainID", w.client.ChainConfig.ChainID, "err", err)
	}
	// update checkpoint
	if w.checkpoint <= 0 {
		latestBlock, err := w.client.ETHClient.BlockByNumber(w.client.ctx, nil)
		if err != nil {
			w.client.logger.Error("Cannot get latest ETH block", "err", err)
			return err
		}
		w.checkpoint = latestBlock.NumberU64()
		err = w.store.SetCheckpoint(w.checkpoint, w.client.ChainConfig.ChainID)
		if err != nil {
			w.client.logger.Error("Cannot store checkpoint to store", "err", err, "chainID", w.client.ChainConfig.ChainID, "checkpoint", w.checkpoint)
		}
	} else {
		w.checkpoint = checkpoint
	}
	w.client.logger.Info("Ethereum watcher started at", "height", w.checkpoint)
	go func() {
		if err := w.watch(); err != nil {
			fmt.Printf("watch blocks error: %s", err)
		}
	}()
	return nil
}

func (w *Watcher) Stop() error {
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
			dualEvents, err := w.GetLatestDualEvents()
			if err != nil {
				w.client.logger.Warn("Cannot get latest dual events", "err", err, "checkpoint", w.checkpoint)
				continue
			}
			// send dual events to events channel
			for i := range dualEvents {
				decodedEvent, err := w.client.SwapSMC.ABI.EventByID(dualEvents[i].Topics[0])
				if err != nil {
					w.client.logger.Warn("Cannot decode dual event", "err", err, "checkpoint", w.checkpoint, "event", dualEvents[i])
					continue
				}
				// extend event data to current dual events
				dualEvents[i].ID = decodedEvent.ID
				dualEvents[i].RawName = decodedEvent.RawName
				dualEvents[i].Inputs = decodedEvent.Inputs
				dualEvents[i].Sig = decodedEvent.Sig

				// unpack event arguments
				dualEvents[i].Arguments, dualEvents[i].Source, dualEvents[i].Destination, err = dualCmn.UnpackDualEventIntoMap(w.client.ABI,
					dualEvents[i], w.client.ChainConfig.ChainID)
				if err != nil {
					w.client.logger.Warn("Cannot unpack dual event", "err", err, "dualEvent", dualEvents[i])
					continue
				}

				// send qualified to dual event channel for further actions
				if dualEvents[i].Source != -1 && dualEvents[i].Destination != -1 {
					w.events <- dualEvents[i]
				}
			}
		default:
		}
	}
}

func (w *Watcher) GetLatestDualEvents() ([]*dualTypes.DualEvent, error) {
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
	w.checkpoint, err = w.store.GetCheckpoint(w.client.ChainConfig.ChainID)
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(w.checkpoint),
		ToBlock:   new(big.Int).SetUint64(latestBlock.NumberU64()),
		Addresses: []ethCommon.Address{w.client.SwapSMC.Address},
		Topics:    topics,
	}
	w.checkpoint = latestBlock.NumberU64() + 1 // increase and store checkpoint to prevent grabbing events of a block multiple times
	err = w.store.SetCheckpoint(w.checkpoint, w.client.ChainConfig.ChainID)
	if err != nil {
		w.client.logger.Error("Cannot store checkpoint to store", "err", err, "chainID", w.client.ChainConfig.ChainID, "checkpoint", w.checkpoint)
	}
	w.client.logger.Debug("Dual events query", "query", query)
	logs, err := w.client.ETHClient.FilterLogs(w.client.ctx, query)
	if err != nil {
		w.client.logger.Error("Cannot get dual event", "err", err, "FromBlock", query.FromBlock.Uint64(), "ToBlock", query.ToBlock.Uint64())
		return nil, err
	}
	result := make([]*dualTypes.DualEvent, len(logs))
	for i := range logs {
		result[i] = &dualTypes.DualEvent{
			// extract log data to current dual event
			Address:     common.HexToAddress(logs[i].Address.Hex()),
			Topics:      convertTopics(logs[i].Topics),
			Data:        logs[i].Data,
			BlockHeight: logs[i].BlockNumber,
			TxHash:      common.BytesToHash(logs[i].TxHash.Bytes()),
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
