package types

import (
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/math"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/kardiachain/go-kardia/signer/core"
)

var types = core.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},

	"Deposit": {
		{
			Name: "sourceChainId",
			Type: "uint256",
		},
		{
			Name: "destChainId",
			Type: "uint256",
		},
		{
			Name: "depositId",
			Type: "uint256",
		},
		{
			Name: "depositor",
			Type: "bytes32",
		},
		{
			Name: "recipient",
			Type: "address",
		},
		{
			Name: "token",
			Type: "bytes32",
		},
		{
			Name: "amount",
			Type: "uint256",
		},
	},
}

func WithDepositHash(d *dproto.Deposit, bridgeAddr common.Address) error {

	var domain = core.TypedDataDomain{
		Name:              "KAI",
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(d.DestChainId),
		VerifyingContract: bridgeAddr.String(),
		Salt:              "",
	}

	var msg = map[string]interface{}{
		"sourceChainId": math.NewHexOrDecimal256(d.SourceChainId),
		"destChainId":   math.NewHexOrDecimal256(d.DestChainId),
		"depositId":     math.NewHexOrDecimal256(d.DepositId),
		"depositor":     common.BytesToHash(d.Depositor).String(),
		"recipient":     common.BytesToHash(d.Recipient).String(),
		"amount":        math.NewHexOrDecimal256(d.Amount),
	}

	var typedData = core.TypedData{
		Types:       types,
		PrimaryType: "Deposit",
		Domain:      domain,
		Message:     msg,
	}

	hash := typedData.TypeHash(typedData.PrimaryType)
	d.Hash = hash[:]
	return nil
}
