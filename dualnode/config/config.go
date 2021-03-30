package config

import (
	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/node"
)

var (
	DefaultDir = ".dualnode"
)

type Config struct {
	Chains   []ChainConfig `yaml:"chains"`
	Node     *node.Config  `yaml:"node"`
	LogLevel string        `yaml:"logLevel"` // crit, error, warn, info, debug, trace
	Key      string        `yaml:"key"`
}

type Contract struct {
	Address string `yaml:"address"`
	ABI     string `yaml:"abi"`
}

type ChainConfig struct {
	Type     string `yaml:"type"`
	ChainID  int64  `yaml:"chainId"`
	Endpoint string `yaml:"endpoint"`

	SwapSMC *Contract `json:"swapSMC"`
}

func RopstenDualETHChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  2,
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",

		SwapSMC: &Contract{
			Address: "0x0",
			ABI:     SwapSMCAbi,
		},
	}
}

func TestDualETHChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  3,
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",

		SwapSMC: &Contract{
			Address: TestSwapSMCAddress,
			ABI:     TestSwapSMCABI,
		},
	}
}

func TestDualKardiaChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.KAISymbol,
		ChainID:  0,
		Endpoint: "http://10.10.0.251:8545",

		SwapSMC: &Contract{
			Address: TestSwapSMCAddressKAITestnet,
			ABI:     TestSwapSMCABI,
		},
	}
}

func DefaultConfig() *Config {
	return &Config{
		LogLevel: "info",
		Node:     node.GetDefaultConfig(),
		Chains: []ChainConfig{
			{
				Type:     configs.KAISymbol,
				ChainID:  1,
				Endpoint: "1",
			},
			{
				Type:     configs.ETHSymbol,
				ChainID:  1,
				Endpoint: "1",
			},
		},
	}
}
