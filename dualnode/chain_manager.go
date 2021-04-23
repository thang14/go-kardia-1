package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/chains/ethereum"
	"github.com/kardiachain/go-kardia/dualnode/chains/kardiachain"
	"github.com/kardiachain/go-kardia/dualnode/config"
)

type Chain interface {
	Stop() error
	Start() error
}

type ChainManager struct {
	chains []Chain
}

func newChainManager(cfg *config.ChainManagerConfig) *ChainManager {
	m := &ChainManager{}
	for _, chainConfig := range cfg.Cfg.Chains {
		var chain Chain
		if chainConfig.Type == "eth" {
			chain = ethereum.NewChain(cfg)
		} else {
			chain = kardiachain.NewChain(cfg)
		}
		m.chains = append(m.chains, chain)
	}
	return m
}

func (c *ChainManager) Start() error {
	for _, c := range c.chains {
		if err := c.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (c *ChainManager) Stop() error {
	for _, c := range c.chains {
		if err := c.Stop(); err != nil {
			return err
		}
	}
	return nil
}
