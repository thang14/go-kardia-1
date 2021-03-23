package kardiachain

import (
	"fmt"

	"github.com/kardiachain/go-kardia/dualnode/consensus"
)

// type Client interface {
// 	FilterLogs(ctx context.Context, q kardia.FilterQuery) ([]types.Log, error)
// 	LatestBlockNumber(ctx context.Context) (uint64, error)
// }

type Watcher struct {
	quit   chan struct{}
	router Router
	//client Client
	//store  *store.Store
	vpool *consensus.Pool
}

func newWatcher() *Watcher {
	return &Watcher{
		quit: make(chan struct{}, 1),
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
	//retry := 0
	for {
		select {
		case <-w.quit:
			return nil
		default:
			// if retry == 0 {
			// 	return fmt.Errorf("max retry")
			// }

			// if err := w.handleEventsForBlock(100); err != nil {
			// 	retry--
			// 	continue
			// }
		}
	}
}

func (w *Watcher) handleEventsForBlock(latestBlock uint64) error {
	// logs, err := w.client.FilterLogs(context.Background(), kardia.FilterQuery{})
	// if err != nil {
	// 	return err
	// }

	// depositTopicHash := crypto.Keccak256Hash([]byte("Deposit(uint256 destination)"))
	// withdrawTopicHash := crypto.Keccak256Hash([]byte("Withdraw(uint256 destination)"))

	// for _, log := range logs {
	// 	if log.Topics[0].Equal(depositTopicHash) {
	// 		deposit := &dproto.Deposit{
	// 			Destination: log.Topics[1].Big().Int64(),
	// 		}
	// 		w.pool.AddDeposit(deposit)
	//      w.pool.AddVote(1)
	// 	} else if log.Topics[0].Equal(withdrawTopicHash) {
	// 		w.pool.MarkDepositCompleted(&dproto.Deposit{})
	// 	}
	// }

	return nil
}
