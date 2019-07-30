package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const tips = `nats controls the connect nats cluster

Usage:
  natsctl [flags]
  natsctl [command]

Available Commands:
  connect     
  help        Help about any command

Flags:
  -h, --help   help for natsctl

Use "natsctl [command] --help" for more information about a command.`

func NewNatsCtlCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "natsctl",
		Short: "nats controls the connect nats cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(tips)
		},
	}
	cmds.AddCommand(NewCmdConnect())
	return cmds
}
