package main

import (
	"fmt"
	"os"

	"github.com/pokoetea/gomemcache/cli"
)

func main() {
	cmd, err := cli.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
