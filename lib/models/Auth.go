package models

type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	TenantId string `json:"tenantId"`
}
