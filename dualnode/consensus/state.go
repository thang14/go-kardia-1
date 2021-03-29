package consensus

import (
	"fmt"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/crypto"
	kevents "github.com/kardiachain/go-kardia/lib/events"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type State struct {
	vpool           *Pool
	store           *store.Store
	privValidator   types.PrivValidator
	pendingDeposits map[string]*dproto.Deposit
	depositHashMap  map[string]string
	// Synchronous pubsub between consensus state and manager.
	// State only emits EventNewRoundStep, EventVote and EventProposalHeartbeat
	evsw kevents.EventSwitch

	validatorSet map[int64]*types.ValidatorSet
	voteSets     map[string]*types.VoteSet
	lastDeposit  map[int64]int64
}

func NewState(vpool *Pool, store *store.Store) (*State, error) {
	state := &State{
		vpool:           vpool,
		store:           store,
		pendingDeposits: make(map[string]*dproto.Deposit),
		depositHashMap:  make(map[string]string),
		evsw:            kevents.NewEventSwitch(),
		validatorSet:    make(map[int64]*types.ValidatorSet),
		voteSets:        make(map[string]*types.VoteSet),
		lastDeposit:     make(map[int64]int64),
	}

	pendingDeposits, err := store.PendingDeposit()
	if err != nil {
		return nil, err
	}

	// restore all pending deposit
	for _, d := range pendingDeposits {
		state.AddDeposit(d)
	}

	return state, nil
}

func (s *State) AddVote(vote *dproto.Vote) error {
	valAddr := common.BytesToAddress(vote.ValidatorAddress)
	if !s.validatorSet[vote.Destination].Has(valAddr) {
		return nil
	}

	if crypto.VerifySignature(valAddr, vote.Hash, vote.Signature) {
		return fmt.Errorf("invalid signature")
	}
	return s.addVote(vote)
}

func (s *State) addVote(vote *dproto.Vote) error {
	if s.pendingDeposits[string(vote.Hash)] == nil {
		return nil
	}
	depositHash := string(vote.Hash)
	if s.voteSets[depositHash] == nil {
		s.voteSets[depositHash] = types.NewVoteSet(make([]*dproto.Vote, 0))
	}

	if s.voteSets[depositHash].Has(vote.ValidatorAddress) {
		return nil
	}

	s.voteSets[depositHash].Add(vote)
	s.vpool.AddVote(vote)
	return nil
}

func (s *State) signVote(vote *dproto.Vote) error {
	return s.privValidator.SignVote(vote)
}

func (s *State) AddDeposit(d *dproto.Deposit) error {
	hash := string(d.Hash)

	if s.pendingDeposits[hash] != nil {
		return nil
	}

	s.depositHashMap[depositKey(d.Destination, d.DepositId)] = hash
	s.pendingDeposits[hash] = d

	if err := s.store.SetDeposit(d); err != nil {
		return err
	}

	vote := &dproto.Vote{
		Hash:             d.Hash,
		Destination:      d.Destination,
		DepositId:        d.DepositId,
		ValidatorAddress: s.privValidator.GetAddress().Bytes(),
	}
	if err := s.signVote(vote); err != nil {
		return err
	}
	s.lastDeposit[d.Source] = d.DepositId
	s.evsw.FireEvent("deposit", d)
	return s.addVote(vote)
}

func (s *State) MarkDepositComplete(d *dproto.Deposit) error {
	s.vpool.MakeDepositCompleted(d)
	s.voteSets[string(d.Hash)] = nil
	delete(s.pendingDeposits, string(d.Hash))
	delete(s.depositHashMap, depositKey(d.Destination, d.DepositId))
	return s.store.MarkDepositCompleted(d)
}

func (s *State) Signs(d *dproto.Deposit) [][]byte {
	if s.voteSets[string(d.Hash)] == nil {
		return nil
	}
	return s.voteSets[string(d.Hash)].Signs()
}

func (s *State) GetDepositByID(chainID, depositID int64) *dproto.Deposit {
	k := s.depositHashMap[depositKey(chainID, depositID)]
	return s.pendingDeposits[k]
}

func (s *State) SetPrivValidator(priv types.PrivValidator) {
	s.privValidator = priv
}

func (s *State) GetDepositState() map[int64]int64 {
	return s.lastDeposit
}

func (s *State) AddValidator(chainID int64, addr common.Address) {
	if s.validatorSet[chainID] == nil {
		s.validatorSet[chainID] = types.NewValidatorSet([]common.Address{addr})
		return
	}
	s.validatorSet[chainID].Add(addr)
}

func (s *State) RemoveValidator(chainID int64, addr common.Address) {
	s.validatorSet[chainID].Remove(addr)
}

func depositKey(chainID, depositID int64) string {
	return fmt.Sprintf("%d:%d", chainID, depositID)
}
