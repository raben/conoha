package lib

import (
	"github.com/raben/conoha/lib/models"
)

const (
	AccountAPIVersion = "v1"
	AccountEndpoint   = "https://account.tyo1.conoha.io/"
)

func (c *Client) GetAccountVersion() (accountVersion models.AccountVersion, err error) {
	if err := c.get(AccountEndpoint, &accountVersion); err != nil {
		return models.AccountVersion{}, err
	}
	return accountVersion, nil
}

func (c *Client) GetAccountOrderItems() (accountOrderItems models.AccountOrderItems, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/order-items", &accountOrderItems); err != nil {
		return models.AccountOrderItems{}, err
	}
	return accountOrderItems, nil
}

func (c *Client) GetAccountPaymentHistory() (accountPaymentHistory models.AccountPaymentHistory, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/payment-history", &accountPaymentHistory); err != nil {
		return models.AccountPaymentHistory{}, err
	}
	return accountPaymentHistory, nil
}

func (c *Client) GetAccountPaymentSummary() (accountPaymentSummary models.AccountPaymentSummary, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/payment-summary", &accountPaymentSummary); err != nil {
		return models.AccountPaymentSummary{}, err
	}
	return accountPaymentSummary, nil
}

func (c *Client) GetAccountBillingInvoices() (accountBillingInvoices models.AccountBillingInvoices, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/billing-invoices/", &accountBillingInvoices); err != nil {
		return models.AccountBillingInvoices{}, err
	}
	return accountBillingInvoices, nil
}

func (c *Client) GetAccountBillingInvoice(invoiceId string) (accountBillingInvoice models.AccountBillingInvoice, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/billing-invoices/"+invoiceId, &accountBillingInvoice); err != nil {
		return models.AccountBillingInvoice{}, err
	}
	return accountBillingInvoice, nil
}

func (c *Client) GetAccountNotifications() (accountNotifications models.AccountNotifications, err error) {
	if err := c.get(AccountEndpoint+AccountAPIVersion+"/"+c.AuthConfig.TenantId+"/notifications", &accountNotifications); err != nil {
		return models.AccountNotifications{}, err
	}
	return accountNotifications, nil
}
