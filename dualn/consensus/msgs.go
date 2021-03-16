package dualn

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

func EncodeMsg(pb proto.Message) ([]byte, error) {
	msg := dproto.Message{}

	switch pb := pb.(type) {
	case *dproto.Vote:
		msg.Sum = &dproto.Message_Vote{Vote: pb}
	}

	bz, err := proto.Marshal(&msg)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal %T: %w", pb, err)
	}

	return bz, nil
}

// DecodeMsg decodes a Protobuf message.
func DecodeMsg(bz []byte) (proto.Message, error) {
	pb := &dproto.Message{}

	err := proto.Unmarshal(bz, pb)
	if err != nil {
		return nil, err
	}
	switch msg := pb.Sum.(type) {
	case *dproto.Message_Vote:
		return msg.Vote, nil
	default:
		return nil, fmt.Errorf("unknown message type %T", msg)
	}
}
