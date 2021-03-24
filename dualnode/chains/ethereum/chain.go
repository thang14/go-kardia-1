package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/kardiachain/go-kardia/configs"
	dualn "github.com/kardiachain/go-kardia/dualnode"
	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/log"
)

type Chain struct {
	watcher *Watcher
	router  *Router
	handler *Handler
}

type ETHLightClient struct {
	ChainConfig *dualCfg.ChainConfig
	ETHClient   *ethclient.Client

	ctx    context.Context
	logger log.Logger
}

func NewETHLightClient(chainCfg *dualCfg.ChainConfig) (*ETHLightClient, error) {
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

func (e *ETHLightClient) Start() error {

}

func (e *ETHLightClient) Stop() error {

}

func (e *ETHLightClient) SetRouter(router *dualn.Router) {

}
