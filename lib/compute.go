package lib

import (
	//"fmt"
	"strings"
	"time"
)

const (
	ComputeAPIVersion = "v2"
	ComputeEndpoint   = "https://compute.tyo1.conoha.io/"
)

type ComputeVersion struct {
	Versions []ComputeVersionValue `json:"versions"`
}
type ComputeVersionValue struct {
	Id      string    `json:"id"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}
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

type ComputeServers struct {
	Servers []ComputeServersValue `json:"servers"`
}
type ComputeServersValue struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}

type ComputeImages struct {
	Images []ComputeImagesValue `json:"images"`
}

type ComputeImagesValue struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	MinDisk  int    `json:"minDisk"`
	MinRam   int    `json:"minRam"`
	Progress int    `json:"progress"`
	//	Created  time.Time `json:"created"`
	//	Updated  time.Time `json:"updated"`
}

func (c *Client) GetComputeVersion() (computeVersion ComputeVersion, err error) {
	if err := c.get(ComputeEndpoint, &computeVersion); err != nil {
		return ComputeVersion{}, err
	}
	return computeVersion, nil
}

func (c *Client) GetComputeFlavors() (computeFlavors ComputeFlavors, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/flavors/detail", &computeFlavors); err != nil {
		return ComputeFlavors{}, err
	}
	return computeFlavors, nil
}

func (c *Client) GetComputeServers() (computeServers ComputeServers, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/detail", &computeServers); err != nil {
		return ComputeServers{}, err
	}
	return computeServers, nil
}

func (c *Client) GetComputeImages(computeImagesValue ComputeImagesValue) (computeImages ComputeImages, err error) {

	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/images/detail", &computeImages); err != nil {
		return ComputeImages{}, err
	}

	filterdValue := []ComputeImagesValue{}
	for _, d := range computeImages.Images {
		if strings.Contains(d.Name, computeImagesValue.Name) {
			filterdValue = append(filterdValue, d)
		}
	}
	computeImages.Images = filterdValue

	return computeImages, nil
}
