package kardiachain

import (
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Router interface {
	SendProposal(proposal *dproto.Proposal)
	Register(chainID string, handler *Handler)
}
