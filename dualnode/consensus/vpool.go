package consensus

import (
	"bytes"
	"fmt"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/lib/clist"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/log"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/kardiachain/go-kardia/types"
)

type Pool struct {
	logger   log.Logger
	store    *store.Store
	voteList *clist.CList // concurrent linked-list of evidence
}

// SetLogger sets the Logger.
func (vpool *Pool) SetLogger(l log.Logger) {
	vpool.logger = l
}

func (vpool *Pool) VoteWaitChan() <-chan struct{} {
	return vpool.voteList.WaitChan()
}

func (vpool *Pool) VoteFront() *clist.CElement {
	return vpool.voteList.Front()
}

func (vpool *Pool) AddVote(vote *dproto.Vote) error {
	if !vpool.store.HasValidator(vote.Destination, vote.ValidatorAddress) {
		return fmt.Errorf("invalid validator")
	}

	if !types.VerifySignature(common.BytesToAddress(vote.ValidatorAddress), vote.Hash, vote.Signature) {
		return fmt.Errorf("invalid signature")
	}

	if err := vpool.store.AddVote(vote); err != nil {
		return err
	}

	vpool.voteList.PushBack(vote)
	return nil
}

func (vpool *Pool) MakeDepositCompleted(deposit *dproto.Deposit) error {
	for e := vpool.voteList.Front(); e != nil; e = e.Next() {
		v := e.Value.(*dproto.Vote)
		if bytes.Equal(deposit.Hash, v.Hash) {
			vpool.voteList.Remove(e)
		}
	}
	return vpool.store.MarkDepositCompleted(deposit, 1)
}

func (vpool *Pool) AddDeposit(deposit *dproto.Deposit) error {
	return vpool.store.SetDeposit(deposit)
}

func (vpool *Pool) PendingDeposit() ([]*dproto.Deposit, error) {
	return vpool.store.PendingDeposit()
}
