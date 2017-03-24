package cmd

import (
	"fmt"
	"runtime"

	"github.com/jawher/mow.cli"
	conoha "github.com/raben/conoha/lib"
)

func (c *CLI) RegisterCommands() {
	c.Command("init", "Get Token And Create Config File", Init)
	c.Command("compute", "Compute Service", func(cmd *cli.Cmd) {
		cmd.Command("version", "Compute version", ComputeVersion)
		cmd.Command("flavors", "Compute flavor list", ComputeFlavor)
		cmd.Command("servers", "Compute server list", ComputeServers)
		cmd.Command("images", "Compute Image list", ComputeImages)
	})

	c.Command("identity", "Identity Service", func(cmd *cli.Cmd) {
		cmd.Command("version", "Identity version", IdentityVersion)
		cmd.Command("new", "Get Identity Token ( Not Set Config File )", IdentityToken)
	})

	c.Command("account", "Billing Service", func(cmd *cli.Cmd) {
		cmd.Command("version", "Account version", AccountVersion)
		cmd.Command("order-items", "Billing Item List", AccountOrderItems)
		cmd.Command("payment-history", "Payment History List", AccountPaymentHistory)
		cmd.Command("payment-summary", "Payment Summary", AccountPaymentSummary)
		cmd.Command("notifications", "Notification List", AccountNotifications)
	})

	c.Command("version", "CLI version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			fmt.Print("Client version:")
			fmt.Print("\n")
			fmt.Print(conoha.Version)
			fmt.Print("\n\n")
			fmt.Print("Os/Arch (client):")
			fmt.Print("\n")
			fmt.Print(fmt.Sprintf("%v/%v", runtime.GOOS, runtime.GOARCH))
			fmt.Print("\n\n")
			fmt.Print("Go version:")
			fmt.Print("\n")
			fmt.Print(runtime.Version())
			fmt.Print("\n\n")
		}
	})
}
