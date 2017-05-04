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
		cmd.Command("version", "Compute Version", ComputeVersion)
		cmd.Command("flavors", "Compute Flavor List", ComputeFlavor)
		cmd.Command("servers", "Compute Server List", ComputeServers)
		cmd.Command("server", "Compute Server Detail", ComputeServer)
		cmd.Command("images", "Compute Image List", ComputeImages)
		cmd.Command("up", "Compute Server Create", ComputeServerCreate)
		cmd.Command("rm", "Compute Server Remove", ComputeServerRemove)
		cmd.Command("start", "Compute Server Start", ComputeServerStart)
		cmd.Command("stop", "Compute Server Stop", ComputeServerStop)
		cmd.Command("restart", "Compute Server Restart", ComputeServerRestart)
	})

	c.Command("identity", "Identity Service", func(cmd *cli.Cmd) {
		cmd.Command("version", "Identity Version", IdentityVersion)
		cmd.Command("new", "Get Identity Token ( Not Set Config File )", IdentityToken)
	})

	c.Command("account", "Billing Service", func(cmd *cli.Cmd) {
		cmd.Command("version", "Account Version", AccountVersion)
		cmd.Command("order-items", "Billing Item List", AccountOrderItems)
		cmd.Command("payment-history", "Payment History List", AccountPaymentHistory)
		cmd.Command("payment-summary", "Payment Summary", AccountPaymentSummary)
		cmd.Command("billing-invoices", "Billing Invoices", AccountBillingInvoices)
		cmd.Command("billing-invoice", "Billing Invoices", AccountBillingInvoice)
		cmd.Command("notifications", "Notification List", AccountNotifications)
	})

	c.Command("version", "CLI Version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			fmt.Print("Client Version:")
			fmt.Print("\n")
			fmt.Print(conoha.Version)
			fmt.Print("\n\n")
			fmt.Print("Os/Arch (client):")
			fmt.Print("\n")
			fmt.Print(fmt.Sprintf("%v/%v", runtime.GOOS, runtime.GOARCH))
			fmt.Print("\n\n")
			fmt.Print("Go Version:")
			fmt.Print("\n")
			fmt.Print(runtime.Version())
			fmt.Print("\n\n")
		}
	})
}
