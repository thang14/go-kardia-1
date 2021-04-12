package consensus

import (
	"bytes"

	"github.com/kardiachain/go-kardia/lib/clist"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/log"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Pool struct {
	logger   log.Logger
	voteList *clist.CList // concurrent linked-list of evidence
}

func NewPool() *Pool {
	return &Pool{
		voteList: clist.New(),
	}
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
	vpool.voteList.PushBack(vote)
	return nil
}

func (vpool *Pool) MakeDepositCompleted(hash common.Hash) {
	for e := vpool.voteList.Front(); e != nil; e = e.Next() {
		v := e.Value.(*dproto.Vote)
		if bytes.Equal(hash.Bytes(), v.Hash) {
			vpool.voteList.Remove(e)
		}
	}
}
