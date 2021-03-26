package config

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/node"
)

type Config struct {
	Chains   []ChainConfig `yaml:"chains"`
	Node     *node.Config  `yaml:"node"`
	LogLevel string        `yaml:"log_level"` // crit, error, warn, info, debug, trace
}

type Contract struct {
	Address common.Address `yaml:"address"`
	ABI     *abi.ABI
}

type ChainConfig struct {
	Type     string `yaml:"type"`
	ChainID  int64  `yaml:"chain_id"`
	Endpoint string `yaml:"endpoint"`

	SwapSMCs map[string]*Contract `yaml:"swapSMCs"`
}

func TestDualETHChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  3,
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",
	}
}

func DefaultConfig() *Config {
	return &Config{
		LogLevel: "info",
		Node:     node.GetDefaultConfig(),
		Chains: []ChainConfig{
			ChainConfig{
				Type:     configs.KAISymbol,
				ChainID:  1,
				Endpoint: "1",
				SwapSMCs: make(map[string]*Contract),
			},
			ChainConfig{
				Type:     configs.ETHSymbol,
				ChainID:  1,
				Endpoint: "1",
				SwapSMCs: make(map[string]*Contract),
			},
		},
	}
}
