package lib

import (
	"encoding/json"
	"github.com/raben/conoha/lib/models"
)

const (
	IdentityAPIVersion = "v2.0"
	IdentityEndpoint   = "https://identity.tyo1.conoha.io/"
)

func (c *Client) GetIdentityVersion() (identityVersion models.IdentityVersion, err error) {
	if err := c.get(IdentityEndpoint, &identityVersion); err != nil {
		return models.IdentityVersion{}, err
	}
	return identityVersion, nil
}

func (c *Client) GetIdentityToken(auth models.Auth) (identityToken models.IdentityToken, err error) {

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
		return models.IdentityToken{}, err
	}
	if err := c.post(IdentityEndpoint+"/"+IdentityAPIVersion+"/tokens", input, &identityToken); err != nil {
		return models.IdentityToken{}, err
	}
	return identityToken, nil
}
