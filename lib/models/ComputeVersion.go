package models

import (
	"time"
)

type ComputeVersion struct {
	Versions []ComputeVersionValue `json:"versions"`
}
type ComputeVersionValue struct {
	Id      string    `json:"id"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}
