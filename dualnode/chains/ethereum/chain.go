package ethereum

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/log"
)

type Chain struct {
	watcher *Watcher

	config *config.ChainConfig
	client *ETHLightClient
}

type ETHLightClient struct {
	ChainConfig *config.ChainConfig
	ETHClient   *ethclient.Client

	ctx    context.Context
	logger log.Logger
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
