package types

import dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"

type Router interface {
	SendDeposit(dproto.Deposit) error
}
