package cmd

import (
	"log"

	"github.com/jawher/mow.cli"
	conoha "github.com/raben/conoha/lib/models"
)

func AccountVersion(cmd *cli.Cmd) {
	cmd.Action = func() {
		accountVersion, err := GetClient().GetAccountVersion()
		if err != nil {
			log.Fatal(err)
		}
		if len(accountVersion.Versions) == 0 {
			log.Fatal(err)
		}
		SliceToMap(accountVersion.Versions)

	}
}

func AccountOrderItems(cmd *cli.Cmd) {
	cmd.Action = func() {
		accountOrderItems, err := GetAuthorizedClient().GetAccountOrderItems()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(accountOrderItems.OrderItems)
	}
}

func AccountPaymentHistory(cmd *cli.Cmd) {
	cmd.Action = func() {
		accountPaymentHistory, err := GetAuthorizedClient().GetAccountPaymentHistory()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(accountPaymentHistory.PaymentHistory)
	}
}

func AccountPaymentSummary(cmd *cli.Cmd) {
	cmd.Action = func() {
		accountPaymentSummary, err := GetAuthorizedClient().GetAccountPaymentSummary()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap([]conoha.AccountPaymentSummaryValue{accountPaymentSummary.PaymentSummary})
	}
}

func AccountBillingInvoices(cmd *cli.Cmd) {
	cmd.Action = func() {
		accountBillingInvoices, err := GetAuthorizedClient().GetAccountBillingInvoices()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(accountBillingInvoices.BillingInvoices)
	}
}

func AccountBillingInvoice(cmd *cli.Cmd) {
	invoiceId := cmd.String(cli.StringOpt{
		Name:      "i invoiceId",
		Value:     "",
		Desc:      "Invoice Id",
		HideValue: true,
	})
	cmd.Spec = "-i"
	cmd.Action = func() {
		accountBillingInvoice, err := GetAuthorizedClient().GetAccountBillingInvoice(*invoiceId)
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(accountBillingInvoice.BillingInvoice.Items)
	}
}

func AccountNotifications(cmd *cli.Cmd) {
	cmd.Action = func() {
		accountNotifications, err := GetAuthorizedClient().GetAccountNotifications()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(accountNotifications.Notifications)
	}
}
