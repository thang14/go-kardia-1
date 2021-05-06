package ethereum

import (
	"context"
	"fmt"

	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/kardiachain/go-kardia/dualnode/chains/kardiachain/bridge"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

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

func NewChain(chainCfg *config.ChainConfig, s *store.Store, depositC chan *dproto.Deposit, withdrawC chan types.Withdraw, vsC chan *types.ValidatorSet) *Chain {
	if chainCfg == nil {
		panic("ETH light client is not available")
	}
	ethClient, err := NewETHLightClient(chainCfg)
	if err != nil {
		panic(fmt.Errorf("cannot setup ETH light client, err: %v", err))
	}
	return &Chain{
		watcher: newWatcher(ethClient, s, depositC, withdrawC, vsC),

		client: ethClient,
		config: chainCfg,
	}
}

func (c *Chain) SetSigner(signer types.Signer) {
	c.signer = signer
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

func (c *Chain) buildWithdrawTransaction(deposit dproto.Deposit) (*etypes.Transaction, error) {
	return nil, nil
}

func (c *Chain) processHandleEvents() {
	for {
		select {
		case d := <-c.depositC:
			c.router.SendDeposit(*d)
		case <-c.quit:
			return
		}
	}
}

func (c *Chain) ReceiveDepositEvent(ctx context.Context, deposit dproto.Deposit) error {
	tx, err := c.buildWithdrawTransaction(deposit)
	if err != nil {
		return err
	}
	return c.SubmitTransaction(ctx, tx)
}

func (c *Chain) ReceiveTransferOwnershipEvent(newOwner []byte) error {
	return nil
}
