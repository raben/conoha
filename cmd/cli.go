package cmd

import "github.com/jawher/mow.cli"

// CLI struct for main
type CLI struct {
	*cli.Cli
}

// NewCLI initializes new command line interface
func NewCLI() *CLI {
	c := &CLI{cli.App("conoha", "A Conoha CLI")}
	return c
}
