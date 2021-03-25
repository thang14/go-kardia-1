package consensus

import (
	"bytes"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/common"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type ChainState struct {
	votes      []dproto.Vote
	validators []common.Address
}

func (cv *ChainState) AddVote(vote *dproto.Vote) {
	cv.votes = append(cv.votes, *vote)
}

func (cv *ChainState) RemoveVoteByHash(hash []byte) {
	votes := make([]dproto.Vote, 0)
	for _, vote := range cv.votes {
		if !bytes.Equal(vote.Hash, hash) {
			votes = append(votes, vote)
		}
	}
	cv.votes = votes
}

func (cv *ChainState) Signatures(hash []byte) [][]byte {
	signs := make([][]byte, 0)
	for _, vote := range cv.votes {
		if bytes.Equal(vote.Hash, hash) {
			signs = append(signs, vote.Signature)
		}
	}
	return signs
}

type State struct {
	chains        map[int64]*ChainState
	vpool         *Pool
	store         *store.Store
	privValidator types.PrivValidator
}

func NewState(vpool *Pool, store *store.Store) *State {
	return &State{
		vpool:  vpool,
		store:  store,
		chains: make(map[int64]*ChainState),
	}
}

func (s *State) addVote(vote *dproto.Vote) error {
	if s.chains[vote.Destination] == nil {
		s.chains[vote.Destination] = &ChainState{}
	}
	s.chains[vote.Destination].AddVote(vote)
	s.vpool.AddVote(vote)
	return nil
}

func (s *State) signVote(vote *dproto.Vote) error {
	return s.privValidator.SignVote(vote)
}

func (s *State) AddDeposit(d *dproto.Deposit) error {
	if err := s.store.SetDeposit(d); err != nil {
		return err
	}

	vote := &dproto.Vote{}
	if err := s.signVote(vote); err != nil {
		return err
	}

	return s.addVote(vote)
}

func (s *State) MarkDepositComplete(d *dproto.Deposit) error {
	s.vpool.MakeDepositCompleted(d)
	s.chains[d.Destination].RemoveVoteByHash(d.Hash)
	return s.store.MarkDepositCompleted(d)
}

func (s *State) GetDepositSignatures(d *dproto.Deposit) [][]byte {
	return s.chains[d.Destination].Signatures(d.Hash)
}
