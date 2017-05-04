package models

import (
	"time"
)

type AuthConfig struct {
	Region    string    `json:region`
	TenantId  string    `json:tenant_id`
	UserName  string    `json:username`
	AuthToken string    `json:auth_token`
	ExpiredAt time.Time `json:expired_at`
}
