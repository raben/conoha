package models

import (
	"time"
)

type AccountOrderItems struct {
	OrderItems []AccountOrderItemsValue `json:"order_items"`
}
type AccountOrderItemsValue struct {
	UuId             string    `json:"uu_id"`
	ServiceName      string    `json:"service_name"`
	ServiceStartDate time.Time `json:"service_start_date"`
	ItemStatus       string    `json:"item_status"`
}
