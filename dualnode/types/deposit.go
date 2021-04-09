package types

import (
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/math"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/kardiachain/go-kardia/signer/core"
)

var typesStandard = core.Types{
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

	var domainStandard = core.TypedDataDomain{
		Name:              "KAI",
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(d.DestChainId),
		VerifyingContract: bridgeAddr.String(),
		Salt:              "",
	}

	var messageStandard = map[string]interface{}{}

	var typedData = core.TypedData{
		Types:       typesStandard,
		PrimaryType: "Deposit",
		Domain:      domainStandard,
		Message:     messageStandard,
	}

	hash := typedData.TypeHash(typedData.PrimaryType)
	d.Hash = hash[:]
	return nil
}
