package config

import (
	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/node"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
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

type ChainConfig struct {
	Type     string `yaml:"type"`
	ChainID  int64  `yaml:"chainId"`
	Endpoint string `yaml:"endpoint"`

	BridgeSmcAddr string `yaml:"bridgeSmcAddr"`
}

type ChainManagerConfig struct {
	Cfg *Config
	S   *store.Store

	DepositC  chan *dproto.Deposit
	WithdrawC chan types.Withdraw
	VsChan    chan *types.ValidatorSet
}

func RopstenDualETHChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  2,
		Endpoint: "https://ropsten.infura.io/v3/ccb2e224843840dc99f3261937eb1900",

		BridgeSmcAddr: "0x",
	}
}

func TestDualETHChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.ETHSymbol,
		ChainID:  1,
		Endpoint: "wss://mainnet.infura.io/ws/v3/ccb2e224843840dc99f3261937eb1900",

		BridgeSmcAddr: "0x",
	}
}

func TestDualKardiaChainConfig() *ChainConfig {
	return &ChainConfig{
		Type:     configs.KAISymbol,
		ChainID:  0,
		Endpoint: "wss://ws-dev.kardiachain.io",

		BridgeSmcAddr: "0x",
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

				BridgeSmcAddr: "0x",
			},
			{
				Type:     configs.ETHSymbol,
				ChainID:  1,
				Endpoint: "1",

				BridgeSmcAddr: "0x",
			},
		},
	}
}
