package kardiachain

import (
	"fmt"
	"time"

	dualCmn "github.com/kardiachain/go-kardia/dualnode/common"

	"github.com/kardiachain/go-kardia"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	dualTypes "github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/common"
)

type Watcher struct {
	quit chan struct{}

	client *KardiaClient
	events chan *dualTypes.DualEvent

	checkpoint uint64
	dualTopics [][]common.Hash
}

func newWatcher(client *KardiaClient) *Watcher {
	return &Watcher{
		quit:   make(chan struct{}, 1),
		client: client,
		events: make(chan *dualTypes.DualEvent, dualCfg.DualEventChanSize),
	}
}

func (w *Watcher) Start() error {
	// update checkpoint
	if w.checkpoint <= 0 {
		latestBlockHeight, err := w.client.KAIClient.BlockHeight(w.client.ctx)
		if err != nil {
			w.client.logger.Error("Cannot get latest Kardia block", "err", err)
			return err
		}
		w.checkpoint = latestBlockHeight
	}
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
	latestBlock, err := w.client.KAIClient.BlockHeight(w.client.ctx)
	if err != nil {
		w.client.logger.Error("Cannot get latest ETH block", "err", err)
		return nil, err
	}
	if w.checkpoint > latestBlock {
		// prevent grabbing events of a block multiple times
		return nil, nil
	}
	topics, err := w.getDualEventTopics()
	if err != nil {
		w.client.logger.Error("Cannot get dual event topics", "err", err)
		return nil, err
	}
	query := kardia.FilterQuery{
		FromBlock: w.checkpoint,
		ToBlock:   latestBlock,
		Addresses: []common.Address{w.client.SwapSMC.Address},
		Topics:    topics,
	}
	w.checkpoint = latestBlock + 1 // increase checkpoint to prevent grabbing events of a block multiple times
	w.client.logger.Debug("Dual events query", "query", query)
	logs, err := w.client.KAIClient.FilterLogs(w.client.ctx, query)
	if err != nil {
		w.client.logger.Error("Cannot get dual event", "err", err, "FromBlock", query.FromBlock, "ToBlock", query.ToBlock)
		return nil, err
	}
	result := make([]*dualTypes.DualEvent, len(logs))
	for i := range logs {
		result[i] = &dualTypes.DualEvent{
			Source:      1,
			Destination: 2,

			Address:     common.HexToAddress(logs[i].Address.Hex()),
			Topics:      logs[i].Topics,
			Data:        logs[i].Data,
			BlockHeight: logs[i].BlockHeight,
			TxHash:      common.BytesToHash(logs[i].TxHash.Bytes()),
		}
	}
	return result, err
}

func (w *Watcher) getDualEventTopics() ([][]common.Hash, error) {
	if w.dualTopics != nil {
		return w.dualTopics, nil
	}
	var (
		swapSMCABI = w.client.SwapSMC.ABI
		topics     []common.Hash
	)
	for i := range swapSMCABI.Events {
		topics = append(topics, swapSMCABI.Events[i].ID)
		w.client.logger.Debug("Appending dual topics...", "topic", swapSMCABI.Events[i].ID)
	}
	w.dualTopics = [][]common.Hash{topics}
	w.client.logger.Info("Dual topics", "topics", topics)
	return w.dualTopics, nil
}
