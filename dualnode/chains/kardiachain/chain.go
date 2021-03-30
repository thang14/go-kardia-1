package kardiachain

import (
	"context"
	"fmt"
	"strings"

	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/kaiclient"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/log"
)

type Chain struct {
	watcher *Watcher

	config *config.ChainConfig
	client *KardiaClient
}

func NewChain(chainCfg *config.ChainConfig) *Chain {
	kaiClient, err := NewKardiaClient(chainCfg)
	if err != nil {
		panic(fmt.Errorf("cannot setup ETH client, err: %v", err))
	}
	return &Chain{
		watcher: newWatcher(kaiClient),

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
	swapSMCAbi, err := abi.JSON(strings.NewReader(chainCfg.SwapSMC.ABI))
	if err != nil {
		panic("cannot read swap smc abi")
	}
	swapSMC := &SwapSMC{
		Address: common.HexToAddress(chainCfg.SwapSMC.Address),
		ABI:     &swapSMCAbi,
	}
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
		SwapSMC:     swapSMC,

		ctx:    context.Background(),
		logger: logger,
	}, nil
}

func (c *Chain) Start() error {
	if err := c.watcher.start(); err != nil {
		return err
	}
	return nil
}

func (c *Chain) Stop() error {
	if err := c.watcher.stop(); err != nil {
		return err
	}
	return nil
}
