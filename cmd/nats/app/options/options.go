package options

import (
	"fmt"

	"github.com/upengs/natsctl/internal/app/natsctl/connect"

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
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(tips)
		},
		// Hook before and after Run initialize and write profiles to disk,
		// respectively.
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		BashCompletionFunction: "",
	}
	cmds.AddCommand(connect.NewCmdConnect())
	return cmds
}
