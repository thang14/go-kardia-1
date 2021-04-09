package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/consensus"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	"github.com/kardiachain/go-kardia/lib/p2p"
	"github.com/kardiachain/go-kardia/node"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/kardiachain/go-kardia/rpc"
	ktypes "github.com/kardiachain/go-kardia/types"
)

type Service struct {
	state    *consensus.State
	cReactor *consensus.Reactor
	chainM   *ChainManager

	// deposit channel
	depositC chan *dproto.Deposit
	// withdraw channel
	withdrawC chan bool
	// validator set channel
	vsChan chan *types.ValidatorSet
}

func New(ctx *node.ServiceContext, cfg *config.Config) (*Service, error) {
	db := memorydb.New()
	s := store.New(db)
	vpool := consensus.NewPool()
	cState, err := consensus.NewState(vpool, s)
	if err != nil {
		panic(err)
	}

	sv := &Service{
		state:     cState,
		depositC:  make(chan *dproto.Deposit),
		withdrawC: make(chan bool),
		vsChan:    make(chan *types.ValidatorSet),
	}

	sv.cReactor = consensus.NewReactor(
		cState,
		cfg,
		sv.depositC,
		sv.withdrawC,
		sv.vsChan,
	)

	sv.chainM = newChainManager(
		cfg,
		s,
		sv.depositC,
		sv.withdrawC,
		sv.vsChan,
	)

	return sv, nil
}

// Start implements Service, starting all internal goroutines needed by the
// Kardia protocol implementation.
func (s *Service) Start(srvr *p2p.Switch) error {
	srvr.AddReactor("BLOCKCHAIN", s.cReactor)
	return s.chainM.Start()
}

func (s *Service) Stop() error {
	_ = s.cReactor.Stop()
	return s.chainM.Stop()
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

func (s *Service) DB() ktypes.StoreDB {
	return nil
}

func (s *Service) Signs(chainId int64, depositID int64) [][]byte {
	d := s.state.GetDepositByID(chainId, depositID)
	if d == nil {
		return nil
	}

	return s.state.Signs(d)
}
