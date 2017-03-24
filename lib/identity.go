package lib

import (
	"encoding/json"
	//"fmt"
	"time"
)

const (
	IdentityAPIVersion = "v2.0"
	IdentityEndpoint   = "https://identity.tyo1.conoha.io/"
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
type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	TenantId string `json:"tenantId"`
}
type AuthConfig struct {
	TenantId  string    `json:tenant_id`
	UserName  string    `json:username`
	AuthToken string    `json:auth_token`
	ExpiredAt time.Time `json:expired_at`
}
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

func (c *Client) GetIdentityVersion() (identityVersion IdentityVersion, err error) {
	if err := c.get(IdentityEndpoint, &identityVersion); err != nil {
		return IdentityVersion{}, err
	}
	return identityVersion, nil
}

func (c *Client) GetIdentityToken(auth Auth) (identityToken IdentityToken, err error) {

	authinfo := map[string]interface{}{
		"auth": map[string]interface{}{
			"tenantId": auth.TenantId,
			"passwordCredentials": map[string]interface{}{
				"username": auth.UserName,
				"password": auth.Password,
			},
		},
	}
	input, err := json.Marshal(authinfo)
	if err != nil {
		return IdentityToken{}, err
	}
	if err := c.post(IdentityEndpoint+"/"+IdentityAPIVersion+"/tokens", input, &identityToken); err != nil {
		return IdentityToken{}, err
	}
	return identityToken, nil
}
