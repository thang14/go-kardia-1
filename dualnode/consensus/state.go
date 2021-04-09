package consensus

import (
	"fmt"
	"time"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/crypto"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type depositState struct {
	deposit    *dproto.Deposit
	signatures map[common.Address][]byte
	submitted  bool
	createdAt  time.Time
}

type depositMap map[common.Hash]*depositState

type State struct {
	vpool         *Pool
	store         *store.Store
	privValidator types.PrivValidator

	// deposit map
	dmap depositMap
	// validator set
	vs             *types.ValidatorSet
	depositHashMap map[string]common.Hash
}

func NewState(vpool *Pool, store *store.Store) (*State, error) {
	s := &State{
		vpool:          vpool,
		store:          store,
		dmap:           depositMap{},
		depositHashMap: make(map[string]common.Hash),
		vs:             types.NewValidatorSet(nil),
	}
	return s, nil
}

func (s *State) AddVote(vote *dproto.Vote) error {
	valAddr := common.BytesToAddress(vote.Addr)
	if !s.vs.Has(valAddr) {
		return nil
	}

	if crypto.VerifySignature(valAddr, vote.Hash, vote.Signature) {
		return fmt.Errorf("invalid signature")
	}
	return s.addVote(vote)
}

func (s *State) addVote(vote *dproto.Vote) error {
	hash := common.BytesToHash(vote.Hash)
	valAddr := common.BytesToAddress(vote.Addr)
	if s.dmap[hash] == nil {
		s.dmap[hash] = &depositState{
			signatures: make(map[common.Address][]byte),
			createdAt:  time.Now(),
		}
	}

	if s.dmap[hash].signatures[valAddr] == nil {
		s.dmap[hash].signatures[valAddr] = vote.Signature
		s.vpool.AddVote(vote)
	}

	return nil
}

func (s *State) signVote(vote *dproto.Vote) error {
	return s.privValidator.SignVote(vote)
}

func (s *State) AddDeposit(d *dproto.Deposit) error {
	hash := common.BytesToHash(d.Hash)

	if s.dmap[hash] == nil {
		s.dmap[hash] = &depositState{
			signatures: make(map[common.Address][]byte),
			createdAt:  time.Now(),
		}
	}

	s.dmap[hash].deposit = d
	s.depositHashMap[depositKey(d.DestChainId, d.DepositId)] = hash

	if err := s.store.SetDeposit(d); err != nil {
		return err
	}

	if !s.IsValidator() {
		return nil
	}

	// sign and send vote
	vote := &dproto.Vote{
		Hash: d.Hash,
		Addr: s.privValidator.GetAddress().Bytes(),
	}
	if err := s.signVote(vote); err != nil {
		return err
	}
	return s.addVote(vote)
}

func (s *State) MarkDepositComplete(d *dproto.Deposit) error {
	hash := common.BytesToHash(d.Hash)
	s.vpool.MakeDepositCompleted(d)
	delete(s.dmap, hash)
	return s.store.MarkDepositCompleted(d)
}

func (s *State) Signs(d *dproto.Deposit) [][]byte {
	signs := make([][]byte, 0)
	hash := common.BytesToHash(d.Hash)

	if s.dmap[hash] == nil {
		return signs
	}

	for _, sign := range s.dmap[hash].signatures {
		signs = append(signs, sign)
	}

	return signs
}

func (s *State) GetDepositByID(chainID, depositID int64) *dproto.Deposit {
	k := s.depositHashMap[depositKey(chainID, depositID)]
	return s.dmap[k].deposit
}

func (s *State) SetPrivValidator(priv types.PrivValidator) {
	s.privValidator = priv
}

func (s *State) AddValidator(chainID int64, addr common.Address) {
	s.vs.Add(addr)
}

func (s *State) RemoveValidator(chainID int64, addr common.Address) {
	s.vs.Remove(addr)
}

func (s *State) SetValidatorSet(vs *types.ValidatorSet) error {
	s.vs = vs
	return nil
}

func (s *State) IsValidator() bool {
	return s.vs.Has(s.privValidator.GetAddress())
}

func depositKey(chainID, depositID int64) string {
	return fmt.Sprintf("%d:%d", chainID, depositID)
}
