package dualn

import (
	"github.com/kardiachain/go-kardia/lib/p2p"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Router struct {
}

func newRouter() *Router {
	return &Router{}
}

func (r *Router) AddVote(vote *dproto.Vote) {

}

func (r *Router) SendVote(proposal *dproto.Proposal, chID byte, src p2p.Peer) {

}
