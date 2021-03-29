package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/consensus"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	"github.com/kardiachain/go-kardia/lib/p2p"
	"github.com/kardiachain/go-kardia/node"
	"github.com/kardiachain/go-kardia/rpc"
	"github.com/kardiachain/go-kardia/types"
)

type Chain interface {
	Start() error
	Stop() error
	SetState(state *consensus.State)
}

type Service struct {
	state    *consensus.State
	chains   []Chain
	cReactor *consensus.Reactor
}

func New(ctx *node.ServiceContext, cfg *config.Config) (*Service, error) {
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
	}, nil
}

func (s *Service) AddChain(chains ...Chain) {
	for _, c := range chains {
		c.SetState(s.state)
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

func (s *Service) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "bridge",
			Version:   "1.0",
			Service:   NewBridgeAPI(s),
			Public:    true,
		},
	}
}

func (s *Service) DB() types.StoreDB {
	return nil
}

func (s *Service) Signs(chainId int64, depositID int64) [][]byte {
	d := s.state.GetDepositByID(chainId, depositID)
	if d == nil {
		return nil
	}

	return s.state.Signs(d)
}
