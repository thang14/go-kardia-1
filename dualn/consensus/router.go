package consensus

import (
	"github.com/kardiachain/go-kardia/lib/p2p"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Router interface {
	SendVote(proposal *dproto.Proposal, chID byte, src p2p.Peer)
	AddVote(vote *dproto.Vote)
}
