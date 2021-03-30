package types

import (
	"math/big"

	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
)

type DualEvent struct {
	Source      int64       `json:"source"`
	Destination int64       `json:"destination"`
	Arguments   interface{} `json:"arguments"`

	// log info
	Address     common.Address `json:"address"`
	Topics      []common.Hash  `json:"topics"`
	Data        common.Bytes   `json:"data"`
	BlockHeight uint64         `json:"blockHeight"`
	TxHash      common.Hash    `json:"transactionHash"`

	// event info
	ID      common.Hash   `json:"ID"`
	RawName string        `json:"rawName"`
	Inputs  abi.Arguments `json:"inputs"`
	Sig     string        `json:"sig"`
}

type LockParams struct {
	Token       [32]byte `json:"token"`
	Destination *big.Int `json:"destination"`
	Amount      uint64   `json:"amount"`
	Recipient   [32]byte `json:"recipient"`
}

type UnlockParams struct {
	Source      *big.Int `json:"source"`
	Destination *big.Int `json:"destination"`
	Token       [32]byte `json:"token"`
	Amount      uint64   `json:"amount"`
	Depositor   [32]byte `json:"depositor"`
	DepositId   *big.Int `json:"depositId"`
	Recipient   [32]byte `json:"recipient"`
	Signs       []byte   `json:"signs"`
}

type AddValidatorParams struct{}
type RemoveValidatorParams struct{}
