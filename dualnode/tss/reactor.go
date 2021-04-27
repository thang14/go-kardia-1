package tss

import (
	"sync"

	"github.com/binance-chain/tss-lib/tss"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/lib/p2p"
)

const (
	TssChannel = byte(0x40)
	MaxMsgSize = 1048576 // 1MB
)

type Reactor struct {
	p2p.BaseReactor
	logger        log.Logger
	privValidator types.PrivValidator
	state         *State

	signingParties map[common.Hash]tss.Party
	signingPartyM  sync.Mutex
}

// Receive implements Reactor by handling different message types.
func (r *Reactor) Receive(chID byte, src p2p.Peer, msgBytes []byte) {

}

// GetChannels implements Reactor
func (r *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	return []*p2p.ChannelDescriptor{
		{
			ID:                  TssChannel,
			Priority:            5,
			SendQueueCapacity:   100,
			RecvBufferCapacity:  100 * 100,
			RecvMessageCapacity: MaxMsgSize,
		},
	}
}
