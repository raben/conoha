package lib

import (
	"github.com/raben/conoha/lib/models"
	"strings"
)

const (
	ComputeAPIVersion = "v2"
	ComputeEndpoint   = "https://compute.tyo1.conoha.io/"
)

func (c *Client) GetComputeVersion() (computeVersion models.ComputeVersion, err error) {
	if err := c.get(ComputeEndpoint, &computeVersion); err != nil {
		return models.ComputeVersion{}, err
	}
	return computeVersion, nil
}

func (c *Client) GetComputeFlavors() (computeFlavors models.ComputeFlavors, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/flavors/detail", &computeFlavors); err != nil {
		return models.ComputeFlavors{}, err
	}
	return computeFlavors, nil
}

func (c *Client) GetComputeServers() (computeServers models.ComputeServers, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/detail", &computeServers); err != nil {
		return models.ComputeServers{}, err
	}
	return computeServers, nil
}

func (c *Client) GetComputeImages(computeImagesValue models.ComputeImagesValue) (computeImages models.ComputeImages, err error) {

	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/images/detail", &computeImages); err != nil {
		return models.ComputeImages{}, err
	}

	filterdValue := []models.ComputeImagesValue{}
	for _, d := range computeImages.Images {
		if strings.Contains(d.Name, computeImagesValue.Name) {
			filterdValue = append(filterdValue, d)
		}
	}
	computeImages.Images = filterdValue

	return computeImages, nil
}
