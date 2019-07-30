package main

import (
	"fmt"
	"os"

	"github.com/upengs/natsctl/cmd/nats/app/options"
	_ "github.com/upengs/natsctl/version"
)

func main() {
	if err := options.NewNatsCtlCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
