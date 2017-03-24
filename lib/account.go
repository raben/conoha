package lib

import (
	"time"
)

const (
	AccountAPIVersion = "v1"
	AccountEndpoint   = "https://account.tyo1.conoha.io/"
)

type AccountVersion struct {
	Versions []AccountVersionValue `json:"versions"`
}
type AccountVersionValue struct {
	Id      string    `json:"id"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}

type AccountOrderItems struct {
	OrderItems []AccountOrderItemsValue `json:"order_items"`
}
type AccountOrderItemsValue struct {
	UuId             string    `json:"uu_id"`
	ServiceName      string    `json:"service_name"`
	ServiceStartDate time.Time `json:"service_start_date"`
	ItemStatus       string    `json:"item_status"`
}

type AccountPaymentHistory struct {
	PaymentHistory []AccountPaymentHistoryValue `json:"payment_history"`
}
type AccountPaymentHistoryValue struct {
	MoneyType     string    `json:"money_type"`
	DepositAmount int       `json:"deposit_amount"`
	RecievedDate  time.Time `json:"received_date"`
}

type AccountPaymentSummary struct {
	PaymentSummary AccountPaymentSummaryValue `json:"payment_summary"`
}
type AccountPaymentSummaryValue struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
}

type AccountNotifications struct {
	Notifications []AccountNotificationsValue `json:notifications`
}
type AccountNotificationsValue struct {
	NotificationCode int       `json:"notification_code"`
	Title            string    `json:"title"`
	Contents         string    `json:"contents"`
	ReadStatus       string    `json:"read_status"`
	StartDate        time.Time `json:"start_date"`
}

func (c *Client) GetAccountVersion() (accountVersion AccountVersion, err error) {
	if err := c.get(AccountEndpoint, &accountVersion); err != nil {
		return AccountVersion{}, err
	}
	return accountVersion, nil
}

func (c *Client) GetAccountOrderItems() (accountOrderItems AccountOrderItems, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/order-items", &accountOrderItems); err != nil {
		return AccountOrderItems{}, err
	}
	return accountOrderItems, nil
}

func (c *Client) GetAccountPaymentHistory() (accountPaymentHistory AccountPaymentHistory, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/payment-history", &accountPaymentHistory); err != nil {
		return AccountPaymentHistory{}, err
	}
	return accountPaymentHistory, nil
}

func (c *Client) GetAccountPaymentSummary() (accountPaymentSummary AccountPaymentSummary, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/payment-summary", &accountPaymentSummary); err != nil {
		return AccountPaymentSummary{}, err
	}
	return accountPaymentSummary, nil
}

func (c *Client) GetAccountNotifications() (accountNotifications AccountNotifications, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/notifications", &accountNotifications); err != nil {
		return AccountNotifications{}, err
	}
	return accountNotifications, nil
}
