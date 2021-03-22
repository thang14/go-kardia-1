package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/consensus"
	"github.com/kardiachain/go-kardia/lib/p2p"
)

type Chain interface {
	Start() error
	Stop() error
	SetRouter(router *Router)
}

type Service struct {
	router   *Router
	chains   []Chain
	cReactor *consensus.Reactor
}

func New() *Service {
	r := newRouter()
	cReacter := consensus.NewReactor(r)

	return &Service{
		router:   r,
		cReactor: cReacter,
	}
}

func (s *Service) AddChain(chains ...Chain) {
	for _, c := range chains {
		c.SetRouter(s.router)
		s.chains = append(s.chains, c)
	}
}

// Start implements Service, starting all internal goroutines needed by the
// Kardia protocol implementation.
func (s *Service) Start(srvr *p2p.Switch) error {
	for _, c := range s.chains {
		if err := c.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Stop() error {
	for _, c := range s.chains {
		if err := c.Stop(); err != nil {
			return err
		}
	}
	return nil
}
