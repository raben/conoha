package models

type ComputeServerMin struct {
	Server ComputeServerMinValue `json:"server"`
}
type ComputeServerMinValue struct {
	Id        string `json:"id"`
	AdminPass string `json:"adminPass"`
}
