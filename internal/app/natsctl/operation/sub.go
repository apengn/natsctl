package operation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/upengs/natsctl/pkg/constans"

	nats "github.com/nats-io/nats.go"
)

type NatsOperation struct {
	conn        *nats.Conn
	inputReader *bufio.Reader
}

func NewNatsOperation(conn *nats.Conn) *NatsOperation {
	return &NatsOperation{conn: conn, inputReader: bufio.NewReader(os.Stdin)}
}

func (n *NatsOperation) Stdin() {
	input, err := n.inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("please re-enter")
		n.Stdin()
	}
	args := strings.Split(strings.TrimSpace(input), " ")
	if len(args) == 0 {
		n.Stdin()
		return
	}
	if len(args) == 1 && args[0] == "" {
		n.Stdin()
		return
	}
	for k, v := range args {
		args[k] = strings.TrimSpace(v)
	}
	n.run(args...)
}

func (n *NatsOperation) run(args ...string) {
	switch args[0] {
	case "subasync":
		// input: sub key
		n.subscribeAsync(args[1])
	case "subqueue":
		// input subqueue subj queue
		n.queueSubscribe(args[1], args[2])
	case "pub":
		n.publish(args...)
	case "help", "h":
		fmt.Print(constans.Example)
		n.Stdin()
	case "exit", "quit":
		os.Exit(1)
	default:
		fmt.Printf("No command [%s]\n", args[0])
		fmt.Println("See examples usage command: help,h")
		n.Stdin()
	}
}

func (n *NatsOperation) publish(args ...string) {
	if len(args) < 3 {
		fmt.Println("参数不正确！example: pub key value")
		n.Stdin()
	}
	n.publishBytes(args[1:]...)
}

func (n *NatsOperation) publishBytes(args ...string) {
	if err := n.conn.Publish(args[0], []byte(args[1])); err != nil {
		fmt.Println("publish failed. please retry!")
		n.Stdin()
		return
	}
	fmt.Println("send success!")
	n.Stdin()
}

func (n *NatsOperation) subscribeAsync(subj string) {
	defer n.conn.Close()
	fmt.Println("subscribe success...")
	wg := sync.WaitGroup{}
	wg.Add(1)
	if _, err := n.conn.Subscribe(subj, func(msg *nats.Msg) {
		wg.Add(1)
		defer wg.Done()
		fmt.Printf("subj:[%s] received:[%s]\n", msg.Subject, string(msg.Data))
	}); err != nil {
		fmt.Println(err)
	}
	wg.Wait()
}

func (n *NatsOperation) queueSubscribe(subj, queue string) {
	defer n.conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	if _, err := n.conn.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
		wg.Add(1)
		defer wg.Done()
		fmt.Printf("subj:[%s]  queue: [%s] received:[%s]\n", msg.Subject, queue, string(msg.Data))
	}); err != nil {
		n.Stdin()
		fmt.Println(err)
		return
	}
	wg.Wait()
}
