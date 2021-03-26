package commands

import (
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg = config.DefaultConfig()
)

// RootCmd is the root command for Duanode core.
var RootCmd = &cobra.Command{
	Use: "dualnode",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		cfg, err = ParseConfig()
		if err != nil {
			return err
		}
		return nil
	},
}

func ParseConfig() (*config.Config, error) {
	conf := config.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
