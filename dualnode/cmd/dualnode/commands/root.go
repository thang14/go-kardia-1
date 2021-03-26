package commands

import (
	"fmt"
	"os"

	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/kardiachain/go-kardia/lib/log"
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

		level, err := log.LvlFromString(cfg.LogLevel)
		if err != nil {
			fmt.Printf("invalid log level argument, default to INFO: %v \n", err)
			level = log.LvlInfo
		}
		log.Root().SetHandler(log.LvlFilterHandler(level,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

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
