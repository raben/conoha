package models

import (
	"time"
)

type AccountBillingInvoices struct {
	BillingInvoices []AccountBillingInvoicesValue `json:"billing_invoices"`
}
type AccountBillingInvoicesValue struct {
	InvoiceId         int       `json:"invoice_id"`
	PaymentMethodType string    `json:"payment_method_type"`
	BillPlasTax       int       `json:"bill_plus_tax"`
	InvoiceDate       time.Time `json:"invoice_date"`
	DueDate           time.Time `json:"due_date"`
}
