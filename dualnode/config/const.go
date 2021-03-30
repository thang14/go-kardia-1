package config

import (
	"time"
)

const (
	DualEventFreq      = 5 * time.Second
	DualEventChanSize  = 100
	UnlockEventRawName = "Unlock"
	LockEventRawName   = "Lock"

	SwapSMCAbi = `[
	{
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"internalType": "uint256",
				"name": "destination",
				"type": "uint256"
			},
			{
				"internalType": "uint64",
				"name": "amount",
				"type": "uint64"
			},
			{
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			}
		],
		"name": "lock",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "safeTransfer",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "safeTransferFrom",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "source",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "destination",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"internalType": "uint64",
				"name": "amount",
				"type": "uint64"
			},
			{
				"internalType": "bytes32",
				"name": "depositor",
				"type": "bytes32"
			},
			{
				"internalType": "uint256",
				"name": "depositId",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signs",
				"type": "bytes"
			}
		],
		"name": "unlock",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
)

var (
	TestSwapSMCABI = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "destination",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "amount",
				"type": "uint64"
			},
			{
				"indexed": true,
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			}
		],
		"name": "Lock",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "source",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "destination",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "amount",
				"type": "uint64"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "depositor",
				"type": "bytes32"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "depositId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			}
		],
		"name": "Unlock",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"internalType": "uint256",
				"name": "destination",
				"type": "uint256"
			},
			{
				"internalType": "uint64",
				"name": "amount",
				"type": "uint64"
			},
			{
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			}
		],
		"name": "lock",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "safeTransfer",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "safeTransferFrom",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "source",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "destination",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "token",
				"type": "bytes32"
			},
			{
				"internalType": "uint64",
				"name": "amount",
				"type": "uint64"
			},
			{
				"internalType": "bytes32",
				"name": "depositor",
				"type": "bytes32"
			},
			{
				"internalType": "uint256",
				"name": "depositId",
				"type": "uint256"
			},
			{
				"internalType": "bytes32",
				"name": "recipient",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signs",
				"type": "bytes"
			}
		],
		"name": "unlock",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
	TestSwapSMCAddress = "0x47621b9dbf7fe0ee5ed89eb36e0f15b4078f9d84" // On Ropsten network at https://ropsten.etherscan.io/tx/0xd8aa2d69027dca25004d879056ea6394fcc398cf3b4c5d45e2e427c6bfb3418c
	TestRopstenPrivKey = "fedc1b2f43e18fb7c2e12f993cee187c8e162dd01d67a08323d9344fb49ef157"

	TestSwapSMCAddressKAITestnet = "0xcBeAE3FC1B8000ae88453C12887d1C16aA111cfA"
)
