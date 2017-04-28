package models

import (
	"time"
)

type ComputeServers struct {
	Servers []ComputeServersValue `json:"servers"`
}
type ComputeServersValue struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}
