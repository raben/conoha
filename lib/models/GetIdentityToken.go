package models

import (
	"time"
)

type IdentityToken struct {
	Access IdentityTokenAccess `json:"access"`
}
type IdentityTokenAccess struct {
	Token IdentityTokenAccessToken `json:"token"`
}
type IdentityTokenAccessToken struct {
	IssuedAt string    `json:"issued_at"`
	Expires  time.Time `json:"expires"`
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	Name     string    `json:"name"`
}
