package config

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

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
	Address common.Address `yaml:"address"`
	ABI     *abi.ABI
}

type ChainConfig struct {
	Type     string `yaml:"type"`
	ChainID  int64  `yaml:"chainId"`
	Endpoint string `yaml:"endpoint"`

	SwapSMC *Contract `json:"swapSMC"`
}

func RopstenDualETHChainConfig() *ChainConfig {
	abi, err := abi.JSON(strings.NewReader(SwapSMCAbi))
	if err != nil {
		panic("cannot read swap smc abi")
	}
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  2,
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",

		SwapSMC: &Contract{
			Address: common.Address{},
			ABI:     &abi,
		},
	}
}

func TestDualETHChainConfig() *ChainConfig {
	abi, err := abi.JSON(strings.NewReader(TestSwapSMCABI))
	if err != nil {
		panic("cannot read swap smc abi")
	}
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  3,
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",

		SwapSMC: &Contract{
			Address: common.HexToAddress(TestSwapSMCAddress),
			ABI:     &abi,
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
