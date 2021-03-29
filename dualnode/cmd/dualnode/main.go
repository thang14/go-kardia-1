package main

import (
	"path/filepath"

	cmd "github.com/kardiachain/go-kardia/dualnode/cmd/dualnode/commands"
	"github.com/kardiachain/go-kardia/dualnode/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	HomeFlag = "home"
)

// Bind all flags and read the config into viper
func bindFlagsLoadViper(cmd *cobra.Command, args []string) error {
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	homeDir := viper.GetString(HomeFlag)
	viper.Set(HomeFlag, homeDir)
	viper.SetConfigName("config")                         // name of config file (without extension)
	viper.AddConfigPath(homeDir)                          // search root directory
	viper.AddConfigPath(filepath.Join(homeDir, "config")) // search root directory /config

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// stderr, so if we redirect output to json file, this doesn't appear
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// ignore not found error, return other errors
		return err
	}
	return nil
}

func prepareBaseCmd(cmd *cobra.Command, defaultHome string) *cobra.Command {
	cmd.PersistentFlags().StringP(HomeFlag, "", config.DefaultDir, "directory for config and data")
	cmd.PersistentPreRunE = concatCobraCmdFuncs(bindFlagsLoadViper, cmd.PersistentPreRunE)
	return cmd
}

type cobraCmdFunc func(cmd *cobra.Command, args []string) error

// Returns a single function that calls each argument function in sequence
// RunE, PreRunE, PersistentPreRunE, etc. all have this same signature
func concatCobraCmdFuncs(fs ...cobraCmdFunc) cobraCmdFunc {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range fs {
			if f != nil {
				if err := f(cmd, args); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func main() {
	rootCmd := cmd.RootCmd
	// Create & start node
	rootCmd.AddCommand(cmd.NewRunNodeCmd())
	rootCmd = prepareBaseCmd(rootCmd, config.DefaultDir)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
