package config

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/kardiachain/go-kardia/configs"
)

type Config struct {
	Chains []ChainConfig
	API    string
}

type Contract struct {
	Address common.Address
	ABI     *abi.ABI
}

type ChainConfig struct {
	Type     string `json:"type"`
	ChainID  string `json:"chain_id"`
	Endpoint string `json:"endpoint"`

	SwapSMCs map[string]*Contract `json:"swapSMCs"`
}

func TestDualETHChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  "0x3",
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",
	}
}
