package models

import (
	"time"
)

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
