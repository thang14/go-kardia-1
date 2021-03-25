package consensus

import (
	"fmt"
	"time"

	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/clist"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/lib/p2p"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	ktypes "github.com/kardiachain/go-kardia/types"
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
}

func newReactor(state *State) *Reactor {
	r := &Reactor{}
	r.BaseReactor = *p2p.NewBaseReactor("DualReactor", r)
	return r
}

// NewReactor creates a new reactor instance.
func NewReactor(state *State) *Reactor {
	return newReactor(state)
}

func (r *Reactor) broadcastNewDeposit(deposit *dproto.Deposit) {
	msg := &dproto.Message{
		Sum: &dproto.Message_NewDeposit{
			NewDeposit: &dproto.NewDeposit{ChainId: deposit.Destination, DepositId: deposit.DepositId},
		},
	}

	msgBytes, err := EncodeMsg(msg)
	if err != nil {
		panic(err)
	}
	r.Switch.Broadcast(DualChannel, msgBytes)
}

// SetLogger sets the Logger on the reactor and the underlying Evidence.
func (r *Reactor) SetLogger(l log.Logger) {
	r.Logger = l
	r.vpool.SetLogger(l)
}

// InitPeer implements Reactor by creating a state for the peer.
func (r *Reactor) InitPeer(peer p2p.Peer) p2p.Peer {
	peerState := NewPeerState(peer).SetLogger(r.Logger)
	peer.Set(ktypes.PeerStateKey, peerState)
	return peer
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

	// Get peer states
	ps, ok := src.Get(ktypes.PeerStateKey).(*PeerState)
	if !ok {
		panic(fmt.Sprintf("Peer %v has no state", src))
	}

	switch msg := msg.(type) {
	case *dproto.Vote:
		if err := r.state.addVote(msg); err != nil {
			r.Switch.StopPeerForError(src, err)
			return
		}
	case *dproto.NewDeposit:
		ps.Deposit[msg.ChainId] = msg.DepositId
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

		vote := r.prepareVoteMsg(peer, next.Value.(*dproto.Vote))
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

func (evR Reactor) prepareVoteMsg(
	peer p2p.Peer,
	vote *dproto.Vote,
) *dproto.Vote {
	peerState, ok := peer.Get(ktypes.PeerStateKey).(PeerState)
	if !ok {
		// Peer does not have a state yet. We set it in the consensus reactor, but
		// when we add peer in Switch, the order we call reactors#AddPeer is
		// different every time due to us using a map. Sometimes other reactors
		// will be initialized before the consensus reactor. We should wait a few
		// milliseconds and retry.
		return nil
	}

	if peerState.Deposit[vote.Destination] >= vote.DepositId {
		return nil
	}
	return vote
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

type PeerState struct {
	peer    p2p.Peer
	logger  log.Logger
	Deposit map[int64]int64
}

// NewPeerState returns a new PeerState for the given Peer
func NewPeerState(peer p2p.Peer) *PeerState {
	return &PeerState{
		peer:    peer,
		logger:  log.NewNopLogger(),
		Deposit: make(map[int64]int64, 0),
	}
}

// SetLogger allows to set a logger on the peer state. Returns the peer state
// itself.
func (ps *PeerState) SetLogger(logger log.Logger) *PeerState {
	ps.logger = logger
	return ps
}
