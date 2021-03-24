/*
 *  Copyright 2021 KardiaChain
 *  This file is part of the go-kardia library.
 *
 *  The go-kardia library is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Lesser General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  The go-kardia library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU Lesser General Public License for more details.
 *
 *  You should have received a copy of the GNU Lesser General Public License
 *  along with the go-kardia library. If not, see <http://www.gnu.org/licenses/>.
 */

package kaiclient

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/kardiachain/go-kardia/rpc"

	"github.com/kardiachain/go-kardia"
	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	"github.com/kardiachain/go-kardia/kai/storage/kvstore"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/crypto"
	kai "github.com/kardiachain/go-kardia/mainchain"
	"github.com/kardiachain/go-kardia/mainchain/genesis"
	"github.com/kardiachain/go-kardia/mainchain/tx_pool"
	"github.com/kardiachain/go-kardia/node"
	"github.com/kardiachain/go-kardia/types"
)

func TestToFilterArg(t *testing.T) {
	blockHashErr := fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
	addresses := []common.Address{
		common.HexToAddress("0xD36722ADeC3EdCB29c8e7b5a47f352D701393462"),
	}
	blockHash := common.HexToHash(
		"0xeb94bb7d78b73657a9d7a99792413f50c0a45c51fc62bdcb08a53f18e9a2b4eb",
	)

	for _, testCase := range []struct {
		name   string
		input  kardia.FilterQuery
		output interface{}
		err    error
	}{
		{
			"without BlockHash",
			kardia.FilterQuery{
				Addresses: addresses,
				FromBlock: 1,
				ToBlock:   2,
				Topics:    [][]common.Hash{},
			},
			map[string]interface{}{
				"address":   addresses,
				"fromBlock": uint64(1),
				"toBlock":   uint64(2),
				"topics":    [][]common.Hash{},
			},
			nil,
		},
		{
			"with nil fromBlock and nil toBlock",
			kardia.FilterQuery{
				Addresses: addresses,
				Topics:    [][]common.Hash{},
			},
			map[string]interface{}{
				"address":   addresses,
				"fromBlock": uint64(1),
				"toBlock":   "latest",
				"topics":    [][]common.Hash{},
			},
			nil,
		},
		{
			"with blockhash",
			kardia.FilterQuery{
				Addresses: addresses,
				BlockHash: &blockHash,
				Topics:    [][]common.Hash{},
			},
			map[string]interface{}{
				"address":   addresses,
				"blockHash": blockHash,
				"topics":    [][]common.Hash{},
			},
			nil,
		},
		{
			"with blockhash and from block",
			kardia.FilterQuery{
				Addresses: addresses,
				BlockHash: &blockHash,
				FromBlock: 1,
				Topics:    [][]common.Hash{},
			},
			nil,
			blockHashErr,
		},
		{
			"with blockhash and to block",
			kardia.FilterQuery{
				Addresses: addresses,
				BlockHash: &blockHash,
				ToBlock:   1,
				Topics:    [][]common.Hash{},
			},
			nil,
			blockHashErr,
		},
		{
			"with blockhash and both from / to block",
			kardia.FilterQuery{
				Addresses: addresses,
				BlockHash: &blockHash,
				FromBlock: 1,
				ToBlock:   2,
				Topics:    [][]common.Hash{},
			},
			nil,
			blockHashErr,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			output, err := toFilterArg(testCase.input)
			if (testCase.err == nil) != (err == nil) {
				t.Fatalf("expected error %v but got %v", testCase.err, err)
			}
			if testCase.err != nil {
				if testCase.err.Error() != err.Error() {
					t.Fatalf("expected error %v but got %v", testCase.err, err)
				}
			} else if !reflect.DeepEqual(output, testCase.output) {
				t.Fatalf("expected filter arg %v but got %v", testCase.output, output)
			}
		})
	}
}

var (
	testKey, _       = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddr         = common.HexToAddress("0x4dAe614b2eA2FaeeDDE7830A2e7fcEDdAE9f9161")
	testBalance, _   = new(big.Int).SetString("1250000000000000000000000000", 10)
	valBalance, _    = new(big.Int).SetString("1125000000000000000000000000", 10)
	genesisContracts = make(map[string]string)
)

type MemoryDbInfo struct{}

func NewMemoryDbInfo() *MemoryDbInfo {
	return &MemoryDbInfo{}
}

func (db *MemoryDbInfo) Name() string {
	return "Memory"
}

func (db *MemoryDbInfo) Start() (types.StoreDB, error) {
	return kvstore.NewStoreDB(memorydb.New()), nil
}

func newTestBackend(t *testing.T) (*node.Node, []*types.Block) {
	db := NewMemoryDbInfo()
	// Generate test chain.
	alloc := map[string]*big.Int{
		"0x14191195F9BB6e54465a341CeC6cce4491599ccC": testBalance,
		testAddr.Hex(): testBalance,
	}

	configs.AddDefaultContract()
	for key, contract := range configs.GetContracts() {
		configs.LoadGenesisContract(key, contract.Address, contract.ByteCode, contract.ABI)
		genesisContracts[contract.Address] = contract.ByteCode
	}

	genesisData := genesis.DefaulTestnetFullGenesisBlock(alloc, genesisContracts)
	genesisData.Validators = []*genesis.GenesisValidator{
		{
			Name:             "val1",
			Address:          "0x4dAe614b2eA2FaeeDDE7830A2e7fcEDdAE9f9161",
			CommissionRate:   "10",
			MaxRate:          "20",
			MaxChangeRate:    "5",
			SelfDelegate:     "125000000000000000000000000",
			StartWithGenesis: true,
			Delegators:       nil,
		},
	}
	// Create node
	config := &node.DefaultConfig
	config.Name = "Test node"
	config.DataDir = ""
	config.Genesis = genesisData
	config.Genesis.ConsensusParams = configs.TestConsensusParams()
	config.MainChainConfig = node.MainChainConfig{
		DBInfo:    db,
		Genesis:   genesisData,
		TxPool:    tx_pool.TxPoolConfig{},
		Consensus: configs.TestConsensusConfig(),
		FastSync:  configs.TestFastSyncConfig(),
	}

	n, err := node.New(config)
	if err != nil {
		t.Fatalf("can't create new node: %v", err)
	}
	if err := n.Register(kai.NewKardiaService); err != nil {
		t.Fatalf("error while adding kardia service: %v", err)
	}
	// Import the test chain.
	if err = n.Start(); err != nil {
		t.Fatalf("can't start test node: %v", err)
	}
	return n, nil
}

//func TestHeader(t *testing.T) {
//	backend, chain := newTestBackend(t)
//	client, _ := backend.Attach()
//	defer backend.Close()
//	defer client.Close()
//
//	tests := map[string]struct {
//		block   *big.Int
//		want    *types.Header
//		wantErr error
//	}{
//		"genesis": {
//			block: big.NewInt(0),
//			want:  chain[0].Header(),
//		},
//		"first_block": {
//			block: big.NewInt(1),
//			want:  chain[1].Header(),
//		},
//		"future_block": {
//			block: big.NewInt(1000000000),
//			want:  nil,
//		},
//	}
//	for name, tt := range tests {
//		t.Run(name, func(t *testing.T) {
//			ec := NewClient(client)
//			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
//			defer cancel()
//
//			got, err := ec.HeaderByNumber(ctx, tt.block)
//			if tt.wantErr != nil && (err == nil || err.Error() != tt.wantErr.Error()) {
//				t.Fatalf("HeaderByNumber(%v) error = %q, want %q", tt.block, err, tt.wantErr)
//			}
//			if got != nil && got.Number.Sign() == 0 {
//				got.Number = big.NewInt(0) // hack to make DeepEqual work
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Fatalf("HeaderByNumber(%v)\n   = %v\nwant %v", tt.block, got, tt.want)
//			}
//		})
//	}
//}

func TestBalanceAt(t *testing.T) {
	backend, _ := newTestBackend(t)
	client, _ := backend.Attach()
	defer backend.Close()
	defer client.Close()

	tests := map[string]struct {
		account common.Address
		block   rpc.BlockNumber
		want    *big.Int
		wantErr error
	}{
		"valid_account": {
			account: common.HexToAddress("0x14191195F9BB6e54465a341CeC6cce4491599ccC"),
			block:   rpc.BlockNumber(0),
			want:    testBalance,
		},
		"valid_validator": {
			account: testAddr,
			block:   rpc.BlockNumber(0),
			want:    valBalance,
		},
		"non_existent_account": {
			account: common.Address{1},
			block:   rpc.BlockNumber(0),
			want:    big.NewInt(0),
		},
		"future_block": {
			account: testAddr,
			block:   rpc.BlockNumber(1000000000),
			want:    big.NewInt(0),
			wantErr: kai.ErrHeaderNotFound,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ec := NewClient(client)
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			got, err := ec.BalanceAt(ctx, tt.account, tt.block)
			if tt.wantErr != nil && (err == nil || err.Error() != tt.wantErr.Error()) {
				t.Fatalf("BalanceAt(%x, %v) error = %q, want %q", tt.account, tt.block, err, tt.wantErr)
			}
			if got.Cmp(tt.want) != 0 {
				t.Fatalf("BalanceAt(%x, %v) = %v, want %v err %v", tt.account, tt.block, got, tt.want, err)
			}
		})
	}
}

func TestBlockNumber(t *testing.T) {
	backend, _ := newTestBackend(t)
	client, _ := backend.Attach()
	defer backend.Close()
	defer client.Close()
	ec := NewClient(client)

	blockNumber, err := ec.BlockNumber(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if blockNumber != 0 {
		t.Fatalf("BlockNumber returned wrong number: %d", blockNumber)
	}
}
