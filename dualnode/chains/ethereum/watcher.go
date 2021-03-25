package ethereum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/kardiachain/go-kardia/configs"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
)

type Watcher struct {
	quit chan struct{}

	client *ETHLightClient
	router Router

	checkpoint uint64
}

func newWatcher(client *ETHLightClient) *Watcher {
	return &Watcher{
		quit:   make(chan struct{}, 1),
		client: client,
	}
}

func (w *Watcher) SetRouter(r Router) {
	w.router = r
}

func (w *Watcher) start() error {
	go func() {
		if err := w.watch(); err != nil {
			fmt.Printf("watch blocks error: %s", err)
		}
	}()
	return nil
}

func (w *Watcher) stop() error {
	close(w.quit)
	return nil
}

func (w *Watcher) watch() error {
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
	if w.checkpoint <= 0 {
		w.checkpoint = latestBlock.NumberU64()
	}
	topics, err := w.getDualEventTopics()
	if err != nil {
		w.client.logger.Error("Cannot get dual event topics", "err", err)
		return nil, err
	}
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(w.checkpoint),
		ToBlock:   new(big.Int).SetUint64(latestBlock.NumberU64()),
		Addresses: []ethCommon.Address{w.client.ChainConfig.SwapSMCs[configs.ETHSymbol].Address},
		Topics:    topics,
	}
	w.client.logger.Debug("Dual events query", "query", query)
	logs, err := w.client.ETHClient.FilterLogs(w.client.ctx, query)
	if err != nil {
		w.client.logger.Error("Cannot get dual event", "err", err, "FromBlock", query.FromBlock.Uint64(), "ToBlock", query.ToBlock.Uint64())
		return nil, err
	}
	return logs, err
}

func (w *Watcher) getDualEventTopics() ([][]ethCommon.Hash, error) {
	var (
		swapSMCABI = w.client.ChainConfig.SwapSMCs[configs.ETHSymbol].ABI
		topics     []ethCommon.Hash
	)
	for i := range swapSMCABI.Events {
		topics = append(topics, swapSMCABI.Events[i].ID)
	}
	w.client.logger.Debug("Dual topics", "topics", topics)
	return [][]ethCommon.Hash{topics}, nil
}
