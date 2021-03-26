package commands

import (
	"github.com/kardiachain/go-kardia/dualnode"
	"github.com/kardiachain/go-kardia/node"
	"github.com/spf13/cobra"
)

func NewRunNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"node", "run"},
		Short:   "Run the dual node",
		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := node.NewBasic(cfg.Node)
			if err != nil {
				return err
			}
			n.Register(func(ctx *node.ServiceContext) (node.Service, error) {
				return dualnode.New(cfg)
			})
			forver := make(chan bool, 0)
			if err := n.Start(); err != nil {
				return err
			}
			<-forver
			return nil
		},
	}
	return cmd
}
