package ethereum

import (
	"context"
	"fmt"
	"math/big"

	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/kardiachain/go-kardia/dualnode/chains/ethereum/bridge"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kardiachain/go-kardia/configs"
	dualCmn "github.com/kardiachain/go-kardia/dualnode/common"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/log"
)

type Chain struct {
	watcher dualCmn.IWatcher

	config           *config.ChainConfig
	client           *ETHLightClient
	signer           types.Signer
	eipSigner        etypes.Signer
	bridgeTransactor *bridge.BridgeTransactor
	depositC         chan *dproto.Deposit
	withdrawC        chan types.Withdraw
	quit             chan bool
	router           types.Router
}

type SwapSMC struct {
	Address common.Address
	ABI     *abi.ABI
}

type ETHLightClient struct {
	ChainConfig *config.ChainConfig
	ETHClient   *ethclient.Client
	ctx         context.Context
	logger      log.Logger
}

func NewETHLightClient(chainCfg *config.ChainConfig) (*ETHLightClient, error) {
	logger := log.New()
	logger.AddTag("DUAL-" + configs.ETHSymbol)
	client, err := ethclient.Dial(chainCfg.Endpoint)
	if err != nil {
		logger.Error("Cannot connect to ETH light client", "error", err)
		return nil, err
	}
	logger.Info("Successfully connected to ETH light client", "endpoint", chainCfg.Endpoint)
	return &ETHLightClient{
		ChainConfig: chainCfg,
		ETHClient:   client,

		ctx:    context.Background(),
		logger: logger,
	}, nil
}

func NewChain(chainCfg *config.ChainConfig, s *store.Store) *Chain {
	if chainCfg == nil {
		panic("ETH light client is not available")
	}
	ethClient, err := NewETHLightClient(chainCfg)
	if err != nil {
		panic(fmt.Errorf("cannot setup ETH light client, err: %v", err))
	}

	depositC := make(chan *dproto.Deposit)
	withdrawC := make(chan types.Withdraw)

	return &Chain{
		watcher: newWatcher(ethClient, s, depositC, withdrawC),

		client: ethClient,
		config: chainCfg,
	}
}

func (c *Chain) SetSigner(signer types.Signer) {
	c.signer = signer
}

func (c *Chain) SetRouter(router types.Router) {
	c.router = router
}

func (c *Chain) Start() error {
	if err := c.watcher.Start(); err != nil {
		return err
	}
	go c.processHandleEvents()
	return nil
}

func (c *Chain) Stop() error {
	if err := c.watcher.Stop(); err != nil {
		return err
	}
	return nil
}

func (c *Chain) SubmitTransaction(ctx context.Context, tx *etypes.Transaction) error {
	txHash := c.eipSigner.Hash(tx)
	signature, err := c.signer.Sign(txHash.Bytes())
	if err != nil {
		return err
	}
	tx.WithSignature(c.eipSigner, signature)
	return c.client.ETHClient.SendTransaction(ctx, tx)
}

func (c *Chain) processHandleEvents() {
	for {
		select {
		case d := <-c.depositC:
			c.router.Send(d.DestChainId, d)
		case <-c.quit:
			return
		}
	}
}

func (c *Chain) Receive(msgI interface{}) error {
	var err error
	var tx *etypes.Transaction
	ctx := context.Background()

	opts, err := c.newTransactOpts(ctx)
	if err != nil {
		return err
	}

	switch msg := msgI.(type) {
	case dproto.Deposit:
		tx, err = buildWithdrawTransaction(c.bridgeTransactor, opts, msg)
		if err != nil {
			return err
		}
	case types.TransferOwnershipMsg:
		tx, err = c.bridgeTransactor.TransferOwnership(opts, common.BytesToAddress(msg.Owner))
		if err != nil {
			return err
		}
	case types.AddTokenMsg:
		tx, err = c.bridgeTransactor.CreateOrEditToken(opts, msg.Symbol, common.BytesToAddress(msg.TokenAddr), msg.Locktype, big.NewInt(0))
		if err != nil {
			return err
		}
	case types.RemoveToken:
		tx, err = c.bridgeTransactor.RemoveToken(opts, msg.Symbol)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return c.SubmitTransaction(ctx, tx)
}

func (c *Chain) newTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	pub, err := crypto.UnmarshalPubkey(c.signer.GetPubKey())
	if err != nil {
		return nil, err
	}

	addr := crypto.PubkeyToAddress(*pub)

	nonce, err := c.client.ETHClient.PendingNonceAt(ctx, addr)
	if err != nil {
		return nil, err
	}
	return &bind.TransactOpts{
		From:     addr,
		Nonce:    big.NewInt(int64(nonce)),
		Value:    big.NewInt(0),
		GasPrice: big.NewInt(c.config.GasPrice),
		GasLimit: 100000,
	}, nil
}

// utils
func buildWithdrawTransaction(b *bridge.BridgeTransactor, opts *bind.TransactOpts, deposit dproto.Deposit) (*etypes.Transaction, error) {
	token := [32]byte{}
	copy(token[:], deposit.Token)

	depositor := [32]byte{}
	copy(depositor[:], deposit.Depositor)

	return b.Withdraw(
		opts,
		big.NewInt(deposit.SourceChainId),
		big.NewInt(deposit.DepositId),
		depositor,
		common.BytesToAddress(deposit.Recipient),
		token,
		big.NewInt(0),
	)
}
