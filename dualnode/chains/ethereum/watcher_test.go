package ethereum

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/stretchr/testify/assert"

	"github.com/kardiachain/go-kardia/dualnode/chains/ethereum/bridge"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

var (
	depositedC = make(chan *bridge.BridgeDeposited, 2)
	withdrawC  = make(chan *bridge.BridgeWithdraw, 2)
)

func initChain() (*Watcher, error) {
	client, err := NewETHLightClient(dualCfg.TestDualETHChainConfig())
	if err != nil {
		return nil, fmt.Errorf("cannot init ETH light client, err %v", err)
	}
	db := memorydb.New()
	s := store.New(db)
	chainManagerCfg := &dualCfg.ChainManagerConfig{
		Cfg: &dualCfg.Config{
			Chains:   []dualCfg.ChainConfig{*dualCfg.TestDualETHChainConfig()},
			LogLevel: "info",
		},
		S:         s,
		DepositC:  make(chan *dproto.Deposit),
		WithdrawC: make(chan types.Withdraw),
		VsChan:    make(chan *types.ValidatorSet),
	}
	watcher := newWatcher(client, chainManagerCfg)
	err = watcher.Start()
	if err != nil {
		return nil, fmt.Errorf("cannot start ETH watcher, err %v", err)
	}
	return watcher, nil
}

func TestWatcher_watch(t *testing.T) {
	var (
		testDepositedEvent = &bridge.BridgeDeposited{
			DestChainId: new(big.Int).SetUint64(0),
			DepositId:   new(big.Int).SetUint64(1),
			Depositor:   ethCmn.Address{},
			Recipient:   [32]byte{},
			Token:       [32]byte{},
			Amount:      new(big.Int).SetUint64(69),
			Raw: ethTypes.Log{
				TxHash: ethCmn.HexToHash("1"),
			},
		}
		testWithdrawEvent = &bridge.BridgeWithdraw{
			SourceChainId: new(big.Int).SetUint64(1),
			DepositId:     new(big.Int).SetUint64(1),
			Depositor:     [32]byte{},
			Recipient:     ethCmn.Address{},
			Token:         [32]byte{},
			Amount:        new(big.Int).SetUint64(70),
			Raw: ethTypes.Log{
				TxHash: ethCmn.HexToHash("2"),
			},
		}
	)
	tests := []struct {
		name    string
		ev      interface{}
		wantErr bool
	}{
		{
			name:    "Test deposited event",
			ev:      testDepositedEvent,
			wantErr: false,
		},
		{
			name:    "Test withdraw event",
			ev:      testWithdrawEvent,
			wantErr: false,
		},
	}

	// init watcher
	w, err := initChain()
	assert.Nilf(t, err, "cannot init watcher for testing, err: %v", err)
	go func() {
		if err := w.watch(depositedC, withdrawC); err != nil {
			t.Fatalf("watch blocks error: %s", err)
		}
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ev, ok := tt.ev.(*bridge.BridgeDeposited); ok {
				depositedC <- ev
				time.Sleep(2 * time.Second)
				assert.NotNilf(t, w.pendingLocks[ev.Raw.TxHash], "pendingLocks must not be nil")
			} else if ev, ok := tt.ev.(*bridge.BridgeWithdraw); ok {
				withdrawC <- ev
				time.Sleep(2 * time.Second)
				assert.NotNilf(t, w.pendingLocks[ev.Raw.TxHash], "pendingLocks must not be nil")
			} else {
				t.Fatal("cannot recognize event type")
			}
		})
	}
}
