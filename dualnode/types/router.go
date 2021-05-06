package types

import (
	"context"

	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type Router interface {
	Deposit(ctx context.Context, d dproto.Deposit) error
	TransferOwnership(ctx context.Context, chainId int64, newOwner []byte) error
	AddToken(ctx context.Context, chainId int64, tokenAddr []byte, locktype int) error
	RemoveToken(ctx context.Context, chainId int64, tokenAddr []byte) error
}

type Chain interface {
	ReceiveDepositEvent(ctx context.Context, deposit dproto.Deposit) error
	ReceiveTransferOwnershipEvent(ctx context.Context, newOwner []byte) error
	ReceiveAddTokenEvent(ctx context.Context, tokenAddr []byte, locktype int) error
	ReceiveRemoveTokenEvent(ctx context.Context, tokenAddr []byte) error
}

type router struct {
	chains map[int64]Chain
}

func NewRouter() Router {
	return &router{}
}

func (r *router) Deposit(ctx context.Context, d dproto.Deposit) error {
	return r.chains[d.DestChainId].ReceiveDepositEvent(ctx, d)
}

func (r *router) TransferOwnership(ctx context.Context, chainId int64, newOwner []byte) error {
	return r.chains[chainId].ReceiveTransferOwnershipEvent(ctx, newOwner)
}

func (r *router) AddToken(ctx context.Context, chainId int64, tokenAddr []byte, locktype int) error {
	return r.chains[chainId].ReceiveAddTokenEvent(ctx, tokenAddr, locktype)
}

func (r *router) RemoveToken(ctx context.Context, chainId int64, tokenAddr []byte) error {
	return r.chains[chainId].ReceiveRemoveTokenEvent(ctx, tokenAddr)
}
