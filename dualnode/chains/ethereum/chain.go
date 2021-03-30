package ethereum

import (
	"context"
	"fmt"
	"strings"

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

	config *config.ChainConfig
	client *ETHLightClient
}

type SwapSMC struct {
	Address common.Address
	ABI     *abi.ABI
}

type ETHLightClient struct {
	ChainConfig *config.ChainConfig
	ETHClient   *ethclient.Client
	*SwapSMC

	ctx    context.Context
	logger log.Logger
}

func NewETHLightClient(chainCfg *config.ChainConfig) (*ETHLightClient, error) {
	swapSMCAbi, err := abi.JSON(strings.NewReader(chainCfg.SwapSMC.ABI))
	if err != nil {
		panic("cannot read swap smc abi")
	}
	swapSMC := &SwapSMC{
		Address: common.HexToAddress(chainCfg.SwapSMC.Address),
		ABI:     &swapSMCAbi,
	}
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
		SwapSMC:     swapSMC,

		ctx:    context.Background(),
		logger: logger,
	}, nil
}

func NewChain(chainCfg *config.ChainConfig) *Chain {
	ethClient, err := NewETHLightClient(chainCfg)
	if err != nil {
		panic(fmt.Errorf("cannot setup ETH client, err: %v", err))
	}
	return &Chain{
		watcher: newWatcher(ethClient),

		client: ethClient,
		config: chainCfg,
	}
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
