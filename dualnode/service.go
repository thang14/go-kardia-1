package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/tss"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/p2p"
	"github.com/kardiachain/go-kardia/node"
	"github.com/kardiachain/go-kardia/rpc"
	ktypes "github.com/kardiachain/go-kardia/types"
)

type Service struct {
	tssReactor   *tss.Reactor
	router       types.Router
	chainManager *ChainManager
}

func New(ctx *node.ServiceContext, cfg *config.Config, s *store.Store) (*Service, error) {
	service := &Service{}

	tssReactor := tss.NewReactor()
	router := types.NewRouter()
	service.tssReactor = tssReactor
	service.router = router
	service.chainManager = newChainManager(cfg, tssReactor, router, s)
	return service, nil
}

// Start implements Service, starting all internal goroutines needed by the
// Kardia protocol implementation.
func (s *Service) Start(srvr *p2p.Switch) error {
	srvr.AddReactor("TSS", s.tssReactor)
	return nil
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) DB() ktypes.StoreDB {
	return nil
}

func (s *Service) APIs() []rpc.API {
	return []rpc.API{}
}
