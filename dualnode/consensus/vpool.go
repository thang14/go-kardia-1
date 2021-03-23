package consensus

import (
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/lib/clist"
	"github.com/kardiachain/go-kardia/lib/log"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
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
	//vote, err := vpool.store.GetPendingDeposit(vote.Hash)

	// verify signature

	//

	vpool.voteList.PushBack(vote)
	return nil
}

func (vpool *Pool) MakeDepositCompleted(deposit *dproto.Deposit) error {
	return vpool.store.MarkDepositCompleted(deposit, 1)
}

func (vpool *Pool) AddDeposit(deposit *dproto.Deposit) error {
	return vpool.store.SetDeposit(deposit)
}
