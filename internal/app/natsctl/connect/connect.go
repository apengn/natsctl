package connect

import (
	"fmt"
	"strings"

	"github.com/upengs/natsctl/internal/app/natsctl/operation"

	nats "github.com/nats-io/nats.go"
)

type NatsConnectOptions struct {
}

func NewNatsConnectOptions() *NatsConnectOptions {
	return &NatsConnectOptions{}
}

func (n *NatsConnectOptions) Connect(option nats.Options, ca []string, cert, key string) error {
	var options []nats.Option

	if option.Name != "" {
		options = append(options, nats.Name(option.Name))
	}

	if option.PingInterval != 0 {
		options = append(options, nats.PingInterval(option.PingInterval))
	}
	if option.User != "" && option.Password != "" {
		options = append(options, nats.UserInfo(option.User, option.Password))
	}
	if option.Token != "" {
		options = append(options, nats.Token(option.Token))
	}
	if len(ca) != 0 {
		options = append(options, nats.RootCAs(ca...))
	}
	if cert != "" && key != "" {
		options = append(options, nats.ClientCert(cert, key))
	}
	if option.MaxReconnect != 0 {
		options = append(options, nats.MaxReconnects(option.MaxReconnect))
	}
	if option.ReconnectBufSize != 0 {
		options = append(options, nats.ReconnectBufSize(option.ReconnectBufSize))
	}
	if option.ReconnectWait != 0 {
		options = append(options, nats.ReconnectWait(option.ReconnectWait))
	}

	nc, err := nats.Connect(strings.Join(option.Servers, ","), options...)
	if err != nil {
		return err
	}

	fmt.Printf("connect nats server:%s,status:%s\n", nc.ConnectedUrl(), getStatusText(nc))
	operation.NewNatsOperation(nc).Stdin()
	return nil
}

func getStatusText(nc *nats.Conn) string {
	switch nc.Status() {
	case nats.CONNECTED:
		return "Connected"
	case nats.CLOSED:
		return "Closed"
	default:
		return "Other"
	}
}
