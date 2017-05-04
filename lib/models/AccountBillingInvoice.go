package models

import (
	"time"
)

type AccountBillingInvoice struct {
	BillingInvoice AccountBillingInvoiceValue `json:"billing_invoice"`
}
type AccountBillingInvoiceValue struct {
	InvoiceId         int                          `json:"invoice_id"`
	PaymentMethodType string                       `json:"payment_method_type"`
	BillPlasTax       int                          `json:"bill_plus_tax"`
	InvoiceDate       time.Time                    `json:"invoice_date"`
	DueDate           time.Time                    `json:"due_date"`
	Items             []AccountBillingInvoiceItems `json:"items"`
}
type AccountBillingInvoiceItems struct {
	InvoiceDetailId int        `json:"invoice_detail_id"`
	ProductName     string     `json:"product_name"`
	UnitPrice       float64    `json:"unit_price"`
	Quantity        int        `json:"quantity"`
	StartDate       time.Time  `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
}
