package dualn

import (
	"github.com/kardiachain/go-kardia/lib/p2p"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Handler interface {
	AddVote(vote *dproto.Vote)
	SendVote(proposal *dproto.Proposal, chID byte, src p2p.Peer)
	ReceiveProposal(proposal *dproto.Proposal)
}

type Router struct {
	handlers map[string]Handler
}

func newRouter() *Router {
	return &Router{}
}

func (r *Router) AddVote(vote *dproto.Vote) {
	r.handlers[vote.Destination].AddVote(vote)
}

func (r *Router) SendVote(proposal *dproto.Proposal, chID byte, src p2p.Peer) {
	r.handlers[proposal.Destination].SendVote(proposal, chID, src)
}

func (r *Router) SendProposal(proposal *dproto.Proposal) {
	r.handlers[proposal.Destination].ReceiveProposal(proposal)
}
