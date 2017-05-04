package lib

import (
	"encoding/json"
	"errors"
	"github.com/raben/conoha/lib/models"
	"github.com/raben/conoha/spec"
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

func (c *Client) GetComputeFlavors(name string) (computeFlavors models.ComputeFlavors, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/flavors/detail", &computeFlavors); err != nil {
		return models.ComputeFlavors{}, err
	}

	filterdValue := []models.ComputeFlavorsValue{}
	for _, d := range computeFlavors.Flavors {
		if strings.Contains(d.Name, name) {
			filterdValue = append(filterdValue, d)
		}
	}
	computeFlavors.Flavors = filterdValue

	return computeFlavors, nil
}

func (c *Client) GetComputeServers() (computeServers models.ComputeServers, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/detail", &computeServers); err != nil {
		return models.ComputeServers{}, err
	}
	return computeServers, nil
}

func (c *Client) GetComputeServer(serverId string) (computeServer models.ComputeServer, err error) {
	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/"+serverId, &computeServer); err != nil {
		return models.ComputeServer{}, err
	}
	return computeServer, nil
}

func (c *Client) GetComputeServerAddresses(serverId string) (addresses []models.ComputeServerValueAddress, err error) {
	computeServer, err := c.GetComputeServer(serverId)
	if err != nil {
		return []models.ComputeServerValueAddress{}, err
	}
	for _, value := range computeServer.Server.Addresses {
		for _, v := range value {
			addresses = append(addresses, v)
		}
	}
	return addresses, nil
}

func (c *Client) GetComputeImages(name string) (computeImages models.ComputeImages, err error) {

	if err := c.get(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/images/detail", &computeImages); err != nil {
		return models.ComputeImages{}, err
	}

	filterdValue := []models.ComputeImagesValue{}
	for _, d := range computeImages.Images {
		if strings.Contains(d.Name, name) {
			filterdValue = append(filterdValue, d)
		}
	}
	computeImages.Images = filterdValue

	return computeImages, nil
}

func (c *Client) CreateComputeServer(spec spec.ConohaServerConfig) (err error) {
	computeFlavors, err := c.GetComputeFlavors(spec.Flavor)
	if err != nil {
		return err
	}
	if len(computeFlavors.Flavors) != 1 {
		return errors.New("Not Found Flavors [ " + spec.Flavor + " ]")
	}

	computeImages, err := c.GetComputeImages(spec.Image)
	if err != nil {
		return err
	}
	if len(computeImages.Images) != 1 {
		return errors.New("Not Found Images [ " + spec.Image + " ]")
	}

	info := map[string]interface{}{
		"server": map[string]interface{}{
			"imageRef":  computeImages.Images[0].Id,
			"flavorRef": computeFlavors.Flavors[0].Id,
		},
	}
	input, err := json.Marshal(info)
	if err := c.post(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers", input, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveComputeServer(serverId string) (err error) {
	if err := c.delete(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/"+serverId, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) StartComputeServer(serverId string) (err error) {
	actioninfo := map[string]interface{}{
		"os-start": nil,
	}
	input, err := json.Marshal(actioninfo)
	if err := c.post(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/"+serverId+"/action", input, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) StopComputeServer(serverId string, force bool) (err error) {
	actioninfo := map[string]interface{}{
		"os-stop": nil,
	}
	if force {
		actioninfo["os-stop"] = map[string]interface{}{
			"force-stop": true,
		}
	}
	input, err := json.Marshal(actioninfo)
	if err := c.post(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/"+serverId+"/action", input, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) RestartComputeServer(serverId string, force bool) (err error) {
	stopType := "SOFT"

	if force {
		stopType = "HARD"
	}

	actioninfo := map[string]interface{}{
		"reboot": map[string]interface{}{
			"type": stopType,
		},
	}
	input, err := json.Marshal(actioninfo)
	if err := c.post(ComputeEndpoint+ComputeAPIVersion+"/"+c.AuthConfig.TenantId+"/servers/"+serverId+"/action", input, nil); err != nil {
		return err
	}
	return nil
}
