package types

import (
	"strings"

	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/crypto"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

var depositAbiJSON = `
[
	{
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "deposit_hash",
				"type": "bytes32"
			},
			{
				"internalType": "uint256",
				"name": "source",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "dest",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "depositor",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "deposit",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`

func DepositHash(deposit *dproto.Deposit) error {
	depositABI, err := abi.JSON(strings.NewReader(depositAbiJSON))
	if err != nil {
		return err
	}

	hash, err := depositABI.Pack("deposit")
	if err != nil {
		return err
	}

	deposit.Hash = crypto.Keccak256Hash(hash).Bytes()
	return nil
}
