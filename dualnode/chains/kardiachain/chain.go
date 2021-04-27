package kardiachain

import (
	"context"
	"fmt"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"

	"github.com/kardiachain/go-kardia/configs"
	dualCmn "github.com/kardiachain/go-kardia/dualnode/common"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/kaiclient"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/log"
)

type Chain struct {
	watcher dualCmn.IWatcher

	config *config.ChainConfig
	client *KardiaClient
}

func NewChain(chainCfg *config.ChainConfig, s *store.Store, depositC chan *dproto.Deposit, withdrawC chan types.Withdraw, vsC chan *types.ValidatorSet) *Chain {
	if chainCfg == nil {
		panic("ETH light client is not available")
	}
	kaiClient, err := NewKardiaClient(chainCfg)
	if err != nil {
		panic(fmt.Errorf("cannot setup KAI client, err: %v", err))
	}
	return &Chain{
		watcher: newWatcher(kaiClient, s, depositC, withdrawC, vsC),

		client: kaiClient,
		config: chainCfg,
	}
}

type SwapSMC struct {
	Address common.Address
	ABI     *abi.ABI
}

type KardiaClient struct {
	ChainConfig *config.ChainConfig
	KAIClient   *kaiclient.Client
	*SwapSMC

	ctx    context.Context
	logger log.Logger
}

func NewKardiaClient(chainCfg *config.ChainConfig) (*KardiaClient, error) {
	logger := log.New()
	logger.AddTag("DUAL-" + configs.KAISymbol)
	client, err := kaiclient.Dial(chainCfg.Endpoint)
	if err != nil {
		logger.Error("Cannot connect to Kardia client", "error", err)
		return nil, err
	}
	logger.Info("Successfully connected to Kardia client", "endpoint", chainCfg.Endpoint)
	return &KardiaClient{
		ChainConfig: chainCfg,
		KAIClient:   client,

		ctx:    context.Background(),
		logger: logger,
	}, nil
}

func (c *Chain) Start() error {
	if err := c.watcher.Start(); err != nil {
		return err
	}
	return nil
}

func (c *Chain) Stop() error {
	if err := c.watcher.Stop(); err != nil {
		return err
	}
	return nil
}
