package types

import (
	"bytes"

	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type VoteSet struct {
	votes []*dproto.Vote
}

func (vs *VoteSet) Add(vote *dproto.Vote) {
	if !vs.Has(vote.ValidatorAddress) {
		vs.votes = append(vs.votes, vote)
	}
}

func (vs *VoteSet) Has(address []byte) bool {
	for _, v := range vs.votes {
		if bytes.Equal(v.ValidatorAddress, address) {
			return true
		}
	}
	return false
}

func (vs *VoteSet) Remove(vote *dproto.Vote) {
	votes := make([]*dproto.Vote, 0)
	for _, v := range vs.votes {
		if !bytes.Equal(vote.ValidatorAddress, v.ValidatorAddress) {
			votes = append(votes, v)
		}
	}
	vs.votes = votes
}

func (vs *VoteSet) Signs() [][]byte {
	signs := make([][]byte, 0)
	for _, vote := range vs.votes {
		signs = append(signs, vote.Signature)
	}
	return signs
}

func NewVoteSet(votes []*dproto.Vote) *VoteSet {
	return &VoteSet{votes: votes}
}
