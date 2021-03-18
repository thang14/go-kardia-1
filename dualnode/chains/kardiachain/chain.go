package kardiachain

import "github.com/kardiachain/go-kardia/dualnode/config"

type Chain struct {
	watcher *Watcher
	handler *Handler
	router  Router
	config  config.ChainConfig
}

func NewChain() *Chain {
	return &Chain{
		watcher: newWatcher(),
		handler: newHandler(),
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

func (c *Chain) SetRouter(r Router) {
	r.Register(c.config.ChainID, c.handler)
	c.router = r
	c.watcher.SetRouter(r)
}
