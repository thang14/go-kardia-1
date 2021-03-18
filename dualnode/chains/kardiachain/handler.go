package kardiachain

import (
	"github.com/kardiachain/go-kardia/lib/p2p"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Handler struct {
}

func newHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ReceiveProposal(proposal *dproto.Proposal) {

}

func (h *Handler) SendVote(proposal *dproto.Proposal, chID byte, src p2p.Peer) {

}

func (h *Handler) AddVote(vote *dproto.Vote) {

}
