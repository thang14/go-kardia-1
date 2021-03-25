package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/consensus"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	"github.com/kardiachain/go-kardia/lib/p2p"
)

type Chain interface {
	Start() error
	Stop() error
	setState(state *consensus.State)
}

type Service struct {
	state    *consensus.State
	chains   []Chain
	cReactor *consensus.Reactor
}

func New() *Service {
	db := memorydb.New()
	s := store.New(db)
	vpool := consensus.NewPool()
	cState, err := consensus.NewState(vpool, s)
	if err != nil {
		panic(err)
	}
	cReacter := consensus.NewReactor(cState)
	return &Service{
		cReactor: cReacter,
		state:    cState,
	}
}

func (s *Service) AddChain(chains ...Chain) {
	for _, c := range chains {
		c.setState(s.state)
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
