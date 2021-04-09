package consensus

import (
	"time"

	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/clist"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/lib/p2p"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

const (
	// DualChannel is a channel for vote dualnode transaction
	DualChannel = byte(0x40)

	MaxMsgSize = 1048576 // 1MB

	broadcastEvidenceIntervalS = 10

	// If a message fails wait this much before sending it again
	peerRetryMessageIntervalMS = 100
)

type Reactor struct {
	p2p.BaseReactor
	logger        log.Logger
	vpool         *Pool
	state         *State
	privValidator types.PrivValidator

	depositC  chan *dproto.Deposit
	valSetC   chan *types.ValidatorSet
	withdrawC chan bool
}

func newReactor(state *State, cfg *config.Config, depositC chan *dproto.Deposit, withdrawC chan bool, vsC chan *types.ValidatorSet) *Reactor {
	r := &Reactor{
		depositC:  depositC,
		valSetC:   vsC,
		withdrawC: withdrawC,
	}
	r.BaseReactor = *p2p.NewBaseReactor("DualReactor", r)
	return r
}

// NewReactor creates a new reactor instance.
func NewReactor(state *State, cfg *config.Config, depositC chan *dproto.Deposit, withdrawC chan bool, vsC chan *types.ValidatorSet) *Reactor {
	return newReactor(state, cfg, depositC, withdrawC, vsC)
}

func (r *Reactor) OnStart() error {
	return r.run()
}

func (r *Reactor) run() error {
	cleanup := time.NewTicker(30 * time.Second)
	for {
		select {
		case depositRecord := <-r.depositC:
			return r.handleDeposit(depositRecord)
		case valSet := <-r.valSetC:
			return r.handleUpdateValSet(valSet)
		case withdraw := <-r.withdrawC:
			return r.handleWithdraw(withdraw)
		case <-cleanup.C:
			return r.handleCleanup()
		case <-r.Quit():
			return nil
		}
	}
}

func (r *Reactor) handleDeposit(d *dproto.Deposit) error {
	return r.state.AddDeposit(d)
}

func (r *Reactor) handleWithdraw(withdraw bool) error {
	return nil
}

func (r *Reactor) handleUpdateValSet(vs *types.ValidatorSet) error {
	return r.state.SetValidatorSet(vs)
}

// SetLogger sets the Logger on the reactor and the underlying Evidence.
func (r *Reactor) SetLogger(l log.Logger) {
	r.Logger = l
	r.vpool.SetLogger(l)
}

// AddPeer implements Reactor.
func (r *Reactor) AddPeer(peer p2p.Peer) error {
	go r.broadcastVoteRoutine(peer)
	return nil
}

// Receive implements Reactor by handling different message types.
func (r *Reactor) Receive(chID byte, src p2p.Peer, msgBytes []byte) {
	msg, err := DecodeMsg(msgBytes)
	if err != nil {
		r.logger.Error("error decoding message",
			"src", src.ID(), "chId", chID, "msg", msg, "err", err)
		return
	}

	switch msg := msg.(type) {
	case *dproto.Vote:
		if err := r.state.AddVote(msg); err != nil {
			r.Switch.StopPeerForError(src, err)
			return
		}
	}

}

func (r *Reactor) broadcastVoteRoutine(peer p2p.Peer) {
	var next *clist.CElement

	for {
		// This happens because the CElement we were looking at got garbage
		// collected (removed). That is, .NextWait() returned nil. Go ahead and
		// start from the beginning.
		if next == nil {
			select {
			case <-r.vpool.VoteWaitChan(): // Wait until evidence is available
				if next = r.vpool.VoteFront(); next == nil {
					continue
				}
			case <-peer.Quit():
				return
			case <-r.Quit():
				return
			}
		} else if !peer.IsRunning() || !r.IsRunning() {
			return
		}

		vote := next.Value.(*dproto.Vote)
		if vote != nil {
			voteBytes, err := vote.Marshal()
			if err != nil {
				panic(err)
			}
			success := peer.Send(DualChannel, voteBytes)
			if !success {
				time.Sleep(peerRetryMessageIntervalMS * time.Millisecond)
				continue
			}
		}

		afterCh := time.After(time.Second * broadcastEvidenceIntervalS)
		select {
		case <-afterCh:
			// start from the beginning every tick.
			// TODO: only do this if we're at the end of the list!
			next = nil
		case <-next.NextWaitChan():
			// see the start of the for loop for nil check
			next = next.Next()
		case <-peer.Quit():
			return
		case <-r.Quit():
			return
		}
	}
}

// GetChannels implements Reactor
func (r *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	return []*p2p.ChannelDescriptor{
		{
			ID:                  DualChannel,
			Priority:            5,
			SendQueueCapacity:   100,
			RecvBufferCapacity:  100 * 100,
			RecvMessageCapacity: MaxMsgSize,
		},
	}
}

func (r *Reactor) handleCleanup() error {
	return nil
}
