package main

import (
	"fmt"
	"github.com/upengs/natsctl/internal/app/natsctl/cmd"
	"os"

	_ "github.com/upengs/natsctl/version"
)

func main() {
	if err := cmd.NewNatsCtlCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
