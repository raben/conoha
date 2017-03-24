package main

import (
	"github.com/raben/conoha/cmd"
	"os"
)

func main() {
	cli := cmd.NewCLI()
	cli.RegisterCommands()
	cli.Run(os.Args)
}
