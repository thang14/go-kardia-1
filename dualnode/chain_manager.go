package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/chains/ethereum"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/tss"
	"github.com/kardiachain/go-kardia/dualnode/types"
)

type Chain interface {
	Stop() error
	Start() error
	SetSigner(types.Signer)
	SetRouter(types.Router)
}

type ChainManager struct {
	chains []Chain
	router types.Router
}

func newChainManager(cfg *config.Config, tssReactor *tss.Reactor, r types.Router, s *store.Store) *ChainManager {
	m := &ChainManager{}
	for _, chainConfig := range cfg.Chains {
		var chain Chain
		if chainConfig.Type == "eth" {
			chain = ethereum.NewChain(&chainConfig, s)
		} else {
			//chain = kardiachain.NewChain(&chainConfig, s, depositC, withdrawC, vsC)
		}
		chain.SetSigner(tssReactor)
		chain.SetRouter(r)
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
