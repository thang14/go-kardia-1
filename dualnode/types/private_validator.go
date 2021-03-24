package types

import (
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type PrivValidator interface {
	SignVote(vote *dproto.Vote) error
}

type privValidator struct {
}

func (p *privValidator) SignVote(vote *dproto.Vote) error {
	return nil
}
