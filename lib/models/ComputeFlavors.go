package models

type ComputeFlavors struct {
	Flavors []ComputeFlavorsValue `json:"flavors"`
}
type ComputeFlavorsValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Ram  int    `json:"ram"`
	Cpus int    `json:"vcpus"`
	Disk int    `json:"disk"`
}
