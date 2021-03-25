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

func newChainState() *ChainState {
	return &ChainState{
		votes:      make([]dproto.Vote, 0),
		validators: []common.Address{},
	}
}

func (cv *ChainState) AddVote(vote *dproto.Vote) {
	cv.votes = append(cv.votes, *vote)
}

func (cv *ChainState) HasVote(vote *dproto.Vote) bool {
	for _, v := range cv.votes {
		if bytes.Equal(v.Hash, vote.Hash) {
			return true
		}
	}
	return false
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
	chains          map[int64]*ChainState
	vpool           *Pool
	store           *store.Store
	privValidator   types.PrivValidator
	pendingDeposits map[string]*dproto.Deposit
}

func NewState(vpool *Pool, store *store.Store) (*State, error) {
	state := &State{
		vpool:  vpool,
		store:  store,
		chains: make(map[int64]*ChainState),
	}

	pendingDeposits, err := store.PendingDeposit()
	if err != nil {
		return nil, err
	}

	for _, d := range pendingDeposits {
		state.pendingDeposits[string(d.Hash)] = d
	}

	return state, nil
}

func (s *State) getChain(chainID int64) *ChainState {
	if s.chains[chainID] == nil {
		s.chains[chainID] = newChainState()
	}
	return s.chains[chainID]
}

func (s *State) addVote(vote *dproto.Vote) error {
	if s.pendingDeposits[string(vote.Hash)] == nil {
		return nil
	}

	chain := s.getChain(vote.Destination)
	if chain.HasVote(vote) == true {
		return nil
	}

	chain.AddVote(vote)
	s.vpool.AddVote(vote)
	return nil
}

func (s *State) signVote(vote *dproto.Vote) error {
	return s.privValidator.SignVote(vote)
}

func (s *State) AddDeposit(d *dproto.Deposit) error {
	if s.pendingDeposits[string(d.Hash)] == nil {
		return nil
	}

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
	s.getChain(d.Destination).RemoveVoteByHash(d.Hash)
	delete(s.pendingDeposits, string(d.Hash))
	return s.store.MarkDepositCompleted(d)
}

func (s *State) GetDepositSignatures(d *dproto.Deposit) [][]byte {
	return s.getChain(d.Destination).Signatures(d.Hash)
}
