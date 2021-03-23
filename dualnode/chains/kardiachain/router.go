package kardiachain

import (
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Router interface {
	SendDeposit(deposit *dproto.Deposit)
	Register(chainID string, handler *Handler)
}
