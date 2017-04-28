package models

type AccountPaymentSummary struct {
	PaymentSummary AccountPaymentSummaryValue `json:"payment_summary"`
}
type AccountPaymentSummaryValue struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
}
