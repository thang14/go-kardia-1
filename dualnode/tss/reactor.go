package tss

import (
	"math/big"
	"sync"

	"github.com/binance-chain/tss-lib/tss"
	"github.com/gogo/protobuf/proto"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/lib/p2p"
)

const (
	TssChannel = byte(0x50)
	MaxMsgSize = 1048576 // 1MB
)

type Reactor struct {
	p2p.BaseReactor
	logger        log.Logger
	privValidator types.PrivValidator
	state         *State

	pendingParties map[string]tss.Party
	pendingPartyM  sync.Mutex
}

// Receive implements Reactor by handling different message types.
func (r *Reactor) Receive(chID byte, src p2p.Peer, msgBytes []byte) {
	msg, err := DecodeMsg(msgBytes)
	if err != nil {
		r.logger.Error("decode signing message err", err)
		return
	}

	_msgBytes, err := proto.Marshal(msg.wire)
	if err != nil {
		r.logger.Error("marshal err", err)
		return
	}

	wireFrom := msg.wire.GetFrom()
	partyId := tss.NewPartyID(wireFrom.Id, wireFrom.Moniker, new(big.Int).SetBytes(wireFrom.Key))
	r.pendingParties[msg.ID].UpdateFromBytes(_msgBytes, partyId, true)
}

func (r *Reactor) addOutMsg(msg tss.Message) {
	// add message
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
