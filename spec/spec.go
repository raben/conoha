package spec

import (
	"github.com/docker/machine/libmachine/engine"
	"github.com/docker/machine/libmachine/mcnflag"
)

const (
	defaultRegion = "tyo1"
	defaultImage  = "vmi-ubuntu-16.04-amd64-unified-20gb"
	defaultFlavor = "g-512mb"
	defaultSSHKey = ""
)

type ConohaServerConfig struct {
	HostName   string
	Region     string
	Image      string
	Flavor     string
	SSHKey     string
	EnginePort int
}

var DefaultServerConfig = &ConohaServerConfig{
	Region:     defaultRegion,
	Image:      defaultImage,
	Flavor:     defaultFlavor,
	SSHKey:     defaultSSHKey,
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

	mcnflag.StringFlag{
		EnvVar: "CONOHA_SSH_KEY",
		Name:   "conoha-ssh-key",
		Usage:  "SSH Private Key Path",
		Value:  "",
	},

	mcnflag.IntFlag{
		EnvVar: "CONOHA_ENGINE_PORT",
		Name:   "conoha-engine-port",
		Usage:  "Conoha Engine Port",
		Value:  engine.DefaultPort,
	},
}
