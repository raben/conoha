package models

import (
	"time"
)

type IdentityVersion struct {
	Versions IdentityVersionValues `json:"versions"`
}
type IdentityVersionValues struct {
	Values []IdentityVersionValue `json:"values"`
}
type IdentityVersionValue struct {
	Id      string    `json:"id"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}
