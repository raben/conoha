package spec

import (
	"github.com/docker/machine/libmachine/engine"
	"github.com/docker/machine/libmachine/mcnflag"
)

const (
	defaultRegion = "tyo1"
	defaultImage  = "vmi-ubuntu-16.04-amd64-unified-20gb"
	defaultFlavor = "g-512mb"
)

type ConohaServerConfig struct {
	HostName   string
	Region     string
	Image      string
	Flavor     string
	EnginePort int
}

var DefaultServerConfig = &ConohaServerConfig{
	Region:     defaultRegion,
	Image:      defaultImage,
	Flavor:     defaultFlavor,
	EnginePort: engine.DefaultPort,
}

var McnFlags = []mcnflag.Flag{
	mcnflag.StringFlag{
		EnvVar: "CONOHA_REGION",
		Name:   "conoha-region",
		Usage:  "Conoha Region Name[tyo1/sin1/sjc1]",
		Value:  defaultRegion,
	},

	mcnflag.StringFlag{
		EnvVar: "CONOHA_IMAGE",
		Name:   "conoha-image",
		Usage:  "Conoha Image Name",
		Value:  defaultImage,
	},

	mcnflag.StringFlag{
		EnvVar: "CONOHA_FLAVOR",
		Name:   "conoha-flavor",
		Usage:  "Conoha Flavor Name",
		Value:  defaultFlavor,
	},

	mcnflag.IntFlag{
		EnvVar: "CONOHA_ENGINE_PORT",
		Name:   "conoha-engine-port",
		Usage:  "Conoha Engine Port",
		Value:  engine.DefaultPort,
	},
}
