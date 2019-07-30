package cmd

import (
	"fmt"
	connect2 "github.com/upengs/natsctl/internal/app/natsctl/connect"

	"github.com/upengs/natsctl/internal/app/natsctl/constans"

	nats "github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

type tls struct {
	ca   []string
	cert string
	key  string
}

// NewCmdRollout returns a Command instance for 'rollout' sub command
func NewCmdConnect() *cobra.Command {
	options := &nats.Options{}
	tls := tls{}
	cmd := &cobra.Command{
		Use: "connect",
		DisableFlagsInUseLine: true,
		Example:               constans.Example,
		Run: func(cmd *cobra.Command, args []string) {
			connect(options, tls)
		},
	}
	cmd.Flags().StringArrayVarP(&options.Servers, "servers", "s", options.Servers, "Servers is a configured set of servers which this client will use when attempting to connect.")
	cmd.Flags().StringVar(&options.Name, "name", options.Name, "Name is an optional name label which will be sent to the server on CONNECT to identify the client.")
	cmd.Flags().StringVar(&options.Token, "token", options.Token, "Token sets the token to be used when connecting to a server.")
	cmd.Flags().StringVar(&options.User, "user", options.User, "User sets the username to be used when connecting to the server.")
	cmd.Flags().StringVar(&options.Password, "password", options.Password, "Password sets the password to be used when connecting to a server.")

	cmd.Flags().DurationVar(&options.PingInterval, "ping-interval", options.PingInterval, "PingInterval is the period at which the client will be sending ping commands to the server, disabled if 0 or negative.")
	cmd.Flags().DurationVarP(&options.ReconnectWait, "reconnect-wait", "r", options.ReconnectWait, "ReconnectWait sets the time to back off after attempting a reconnect to a server that we were already connected to previously.")

	cmd.Flags().IntVarP(&options.MaxReconnect, "max-reconnect", "m", options.MaxReconnect, "MaxReconnect sets the number of reconnect attempts that will be tried before giving up. If negative, then it will never give up trying to reconnect.")
	cmd.Flags().IntVarP(&options.ReconnectBufSize, "reconnect-buffer-size", "b", options.ReconnectBufSize, "ReconnectBufSize is the size of the backing bufio during reconnect. Once this has been exhausted publish operations will return an error.")

	cmd.Flags().StringVar(&tls.key, "key", tls.key, "key.pem file path")
	cmd.Flags().StringVar(&tls.cert, "cert", tls.cert, "cert.pem file path")
	cmd.Flags().StringArrayVar(&tls.ca, "ca", tls.ca, "ca.pem file path")
	return cmd
}

func connect(options *nats.Options, tls tls) {
	if err := connect2.NewNatsConnectOptions().Connect(*options, tls.ca, tls.cert, tls.key); err != nil {
		fmt.Println(err)
		return
	}
}
