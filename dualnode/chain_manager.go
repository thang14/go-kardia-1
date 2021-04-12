package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/chains/ethereum"
	"github.com/kardiachain/go-kardia/dualnode/chains/kardiachain"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Chain interface {
	Stop() error
	Start() error
}

type ChainManager struct {
	chains []Chain
}

func newChainManager(cfg *config.Config, s *store.Store, depositC chan *dproto.Deposit, withdrawC chan types.Withdraw, vsC chan *types.ValidatorSet) *ChainManager {
	m := &ChainManager{}
	for _, chainConfig := range cfg.Chains {
		var chain Chain
		if chainConfig.Type == "eth" {
			chain = ethereum.NewChain(&chainConfig, s)
		} else {
			chain = kardiachain.NewChain(&chainConfig, s)
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
