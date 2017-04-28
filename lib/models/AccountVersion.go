package models

import (
	"time"
)

type AccountVersion struct {
	Versions []AccountVersionValue `json:"versions"`
}

type AccountVersionValue struct {
	Id      string    `json:"id"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}
