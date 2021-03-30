package kardiachain

import (
	"fmt"
	"time"

	"github.com/kardiachain/go-kardia"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/types"
)

type Watcher struct {
	quit chan struct{}

	client *KardiaClient
	events chan *abi.Event

	checkpoint uint64
	dualTopics [][]common.Hash
}

func newWatcher(client *KardiaClient) *Watcher {
	return &Watcher{
		quit:   make(chan struct{}, 1),
		client: client,
		events: make(chan *abi.Event, dualCfg.DualEventChanSize),
	}
}

func (w *Watcher) start() error {
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
	fmt.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@ query %+v\n", query)
	w.client.logger.Debug("Dual events query", "query", query)
	logs, err := w.client.KAIClient.FilterLogs(w.client.ctx, query)
	if err != nil {
		w.client.logger.Error("Cannot get dual event", "err", err, "FromBlock", query.FromBlock, "ToBlock", query.ToBlock)
		return nil, err
	}
	return logs, err
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
