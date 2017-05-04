package models

type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Region   string `json:"region"`
	TenantId string `json:"tenantId"`
}
