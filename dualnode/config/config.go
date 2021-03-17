package config

type Config struct {
	Chains []ChainConfig
}

type ChainConfig struct {
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}
