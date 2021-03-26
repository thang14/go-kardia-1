package main

import (
	cmd "github.com/kardiachain/go-kardia/dualnode/cmd/dualnode/commands"
)

func main() {
	rootCmd := cmd.RootCmd

	// Create & start node
	rootCmd.AddCommand(cmd.NewRunNodeCmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
