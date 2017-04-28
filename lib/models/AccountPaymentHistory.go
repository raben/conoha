package models

import (
	"time"
)

type AccountPaymentHistory struct {
	PaymentHistory []AccountPaymentHistoryValue `json:"payment_history"`
}
type AccountPaymentHistoryValue struct {
	MoneyType     string    `json:"money_type"`
	DepositAmount int       `json:"deposit_amount"`
	RecievedDate  time.Time `json:"received_date"`
}
