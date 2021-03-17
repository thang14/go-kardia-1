package dualn

import "github.com/kardiachain/go-kardia/lib/p2p"

type Service struct {
}

func New() *Service {
	return &Service{}
}

// Start implements Service, starting all internal goroutines needed by the
// Kardia protocol implementation.
func (s *Service) Start(srvr *p2p.Switch) error {
	return nil
}

func (s *Service) Stop() error {
	return nil
}
