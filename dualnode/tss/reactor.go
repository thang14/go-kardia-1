package tss

import (
	"sync"

	"github.com/binance-chain/tss-lib/tss"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/lib/p2p"
)

const (
	SigningChannel = byte(0x50)
	KeygenChannel  = byte(0x51)
	MaxMsgSize     = 1048576 // 1MB
)

type Reactor struct {
	p2p.BaseReactor
	logger        log.Logger
	privValidator types.PrivValidator
	state         *State

	signingParties map[string]tss.Party
	signingPartyM  sync.Mutex
}

// Receive implements Reactor by handling different message types.
func (r *Reactor) Receive(chID byte, src p2p.Peer, msgBytes []byte) {
	switch chID {
	case SigningChannel:
		r.handlerSigning(msgBytes)
	}
}

func (r *Reactor) addOutMsg(msg tss.Message) {
	// add message
}

// GetChannels implements Reactor
func (r *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	return []*p2p.ChannelDescriptor{
		{
			ID:                  SigningChannel,
			Priority:            5,
			SendQueueCapacity:   100,
			RecvBufferCapacity:  100 * 100,
			RecvMessageCapacity: MaxMsgSize,
		},
		{
			ID:                  KeygenChannel,
			Priority:            5,
			SendQueueCapacity:   100,
			RecvBufferCapacity:  100 * 100,
			RecvMessageCapacity: MaxMsgSize,
		},
	}
}
