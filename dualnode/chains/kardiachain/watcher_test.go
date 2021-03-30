package kardiachain

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/abi/bind"
	"github.com/kardiachain/go-kardia/lib/crypto"
	"github.com/kardiachain/go-kardia/types"
)

type LockParams struct {
	token       [32]byte
	destination *big.Int
	amount      uint64
	recipient   [32]byte
}

type UnlockParams struct {
	source      *big.Int
	destination *big.Int
	token       [32]byte
	amount      uint64
	depositor   [32]byte
	depositId   *big.Int
	recipient   [32]byte
	signs       []byte
}

func initChain() (*Watcher, *bind.BoundContract, error) {
	client, err := NewKardiaClient(dualCfg.TestDualKardiaChainConfig())
	if err != nil {
		return nil, nil, fmt.Errorf("cannot init ETH light client, err %v", err)
	}
	watcher := newWatcher(client)
	err = watcher.start()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot start ETH watcher, err %v", err)
	}
	swapABI, err := abi.JSON(strings.NewReader(dualCfg.TestSwapSMCABI))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse swap SMC ABI, err %v", err)
	}
	return watcher, bind.NewBoundContract(watcher.client.Address, swapABI, client.KAIClient, client.KAIClient, client.KAIClient), nil
}

func (w *Watcher) callLockFunctionWithParams(boundSwapSMC *bind.BoundContract, params *LockParams) (*types.Transaction, error) {
	auth, err := w.setupSender()
	if err != nil {
		return nil, fmt.Errorf("cannot setup sender, err %v", err)
	}
	tx, err := boundSwapSMC.Transact(auth, "lock", params.token, params.destination, params.amount, params.recipient)
	if err != nil {
		return nil, fmt.Errorf("cannot create a lock transaction to swap SMC, err %v", err)
	}
	return tx, nil
}

func (w *Watcher) callUnlockFunctionWithParams(boundSwapSMC *bind.BoundContract, params *UnlockParams) (*types.Transaction, error) {
	auth, err := w.setupSender()
	if err != nil {
		return nil, fmt.Errorf("cannot setup sender, err %v", err)
	}
	tx, err := boundSwapSMC.Transact(auth, "unlock", params.source, params.destination,
		params.token, params.amount, params.depositor, params.depositId, params.recipient, params.signs)
	if err != nil {
		return nil, fmt.Errorf("cannot create a unlock transaction to swap SMC, err %v", err)
	}
	return tx, nil
}

func (w *Watcher) setupSender() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(dualCfg.TestRopstenPrivKey)
	if err != nil {
		return nil, fmt.Errorf("error casting private key %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := w.client.KAIClient.PendingNonceAt(w.client.ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("cannot get pending nonce at %s err %v", fromAddress.Hex(), err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nonce
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = new(big.Int).SetUint64(1)

	return auth, nil
}

func (w *Watcher) getLatestBlockNumber() (uint64, error) {
	height, err := w.client.KAIClient.BlockHeight(w.client.ctx)
	if err != nil {
		return 0, fmt.Errorf("cannot get latest block, err %v", err)
	}
	return height, err
}

func TestCaptureLockEvent(t *testing.T) {
	watcher, boundSwapSMC, err := initChain()
	if err != nil {
		t.Fatal(err)
	}
	startHeight, err := watcher.getLatestBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	watcher.checkpoint = startHeight + 1
	t.Logf("current checkpoint %v latest block number %v", watcher.checkpoint, startHeight)
	tx, err := watcher.callLockFunctionWithParams(boundSwapSMC, &LockParams{
		token:       [32]byte{0x1},
		destination: new(big.Int).SetInt64(1),
		amount:      100,
		recipient:   [32]byte{0x2},
	})
	t.Logf("transaction sent %v", tx)
	if err != nil {
		t.Fatal(err)
	}
LOOP:
	for {
		select {
		case event, ok := <-watcher.events:
			assert.Equal(t, ok, true, "stopping event capturing")
			t.Logf("dual event captured %+v, checkpoint %v", event, watcher.checkpoint)
			assert.Equal(t, dualCfg.LockEventRawName, event.RawName, "captured event must be a lock event")
			startHeight, err = watcher.getLatestBlockNumber()
			assert.NotEqual(t, 0, startHeight, "cannot get latest block")
			assert.GreaterOrEqualf(t, watcher.checkpoint, startHeight, "checkpoint (%v) must be increased to greater than latest block number (%v)",
				watcher.checkpoint, startHeight)
			err = watcher.stop()
			assert.Nil(t, err, "cannot stop watcher")
			break LOOP
		default:
			continue LOOP
		}
	}
}

func TestCaptureUnlockEvent(t *testing.T) {
	watcher, boundSwapSMC, err := initChain()
	if err != nil {
		t.Fatal(err)
	}
	startHeight, err := watcher.getLatestBlockNumber()
	if err != nil {
		t.Fatal(err)
	}
	watcher.checkpoint = startHeight + 1
	t.Logf("current checkpoint %v latest block number %v", watcher.checkpoint, startHeight)
	tx, err := watcher.callUnlockFunctionWithParams(boundSwapSMC, &UnlockParams{
		source:      new(big.Int).SetInt64(1),
		destination: new(big.Int).SetInt64(2),
		token:       [32]byte{0x1},
		amount:      100,
		depositor:   [32]byte{0x2},
		depositId:   new(big.Int).SetInt64(3),
		recipient:   [32]byte{0x3},
		signs:       []byte{0x4},
	})
	t.Logf("transaction sent %v", tx)
	if err != nil {
		t.Fatal(err)
	}
LOOP:
	for {
		select {
		case event, ok := <-watcher.events:
			assert.Equal(t, ok, true, "stopping event capturing")
			t.Logf("dual event captured %+v, checkpoint %v", event, watcher.checkpoint)
			assert.Equal(t, dualCfg.UnlockEventRawName, event.RawName, "captured event must be a unlock event")
			startHeight, err = watcher.getLatestBlockNumber()
			assert.NotEqual(t, 0, startHeight, "cannot get latest block")
			assert.GreaterOrEqualf(t, watcher.checkpoint, startHeight, "checkpoint (%v) must be increased to greater than latest block number (%v)",
				watcher.checkpoint, startHeight)
			err = watcher.stop()
			assert.Nil(t, err, "cannot stop watcher")
			break LOOP
		default:
			continue LOOP
		}
	}
}
