package tss

import (
	"github.com/binance-chain/tss-lib/tss"
	"github.com/gogo/protobuf/proto"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Message struct {
	ID   string
	wire *tss.MessageWrapper
}

func EncodeMsg(msg Message) ([]byte, error) {
	msgBytes, err := proto.Marshal(msg.wire)
	if err != nil {
		return nil, err
	}
	msgP := &dproto.Message{
		Id:      msg.ID,
		Message: msgBytes,
	}

	return msgP.Marshal()
}

func DecodeMsg(msgBytes []byte) (*Message, error) {
	msgP := &dproto.Message{}
	if err := proto.Unmarshal(msgBytes, msgP); err != nil {
		return nil, err
	}

	msg := &tss.MessageWrapper{}

	if err := proto.Unmarshal(msgP.Message, msg); err != nil {
		return nil, err
	}

	return &Message{
		wire: msg,
		ID:   msgP.Id,
	}, nil
}
