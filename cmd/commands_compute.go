package cmd

import (
	"fmt"
	"log"

	"github.com/jawher/mow.cli"
	conoha "github.com/raben/conoha/lib/models"
)

func ComputeVersion(cmd *cli.Cmd) {
	cmd.Action = func() {
		computeVersion, err := GetClient().GetComputeVersion()
		if err != nil {
			log.Fatal(err)
		}
		if len(computeVersion.Versions) == 0 {
			log.Fatal(err)
		}
		SliceToMap(computeVersion.Versions)

	}
}

func ComputeFlavor(cmd *cli.Cmd) {
	cmd.Action = func() {
		computeFlavors, err := GetAuthorizedClient().GetComputeFlavors()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(computeFlavors.Flavors)
	}

}

func ComputeServers(cmd *cli.Cmd) {
	cmd.Action = func() {
		computeServers, err := GetAuthorizedClient().GetComputeServers()
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(computeServers.Servers)
	}
}

func ComputeImages(cmd *cli.Cmd) {
	name := cmd.String(cli.StringOpt{
		Name:      "n name",
		Value:     "",
		Desc:      "image name",
		HideValue: true,
	})

	cmd.Action = func() {
		computeImages, err := GetAuthorizedClient().GetComputeImages(conoha.ComputeImagesValue{Name: *name})
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(computeImages.Images)
	}
}

func ComputeServerStart(cmd *cli.Cmd) {
	id := cmd.String(cli.StringOpt{
		Name:      "i id",
		Value:     "",
		Desc:      "server id",
		HideValue: true,
	})
	cmd.Spec = "-i"
	cmd.Action = func() {
		err := GetAuthorizedClient().StartComputeServer(*id)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Accepted Start Request.\n\n")
	}
}

func ComputeServerStop(cmd *cli.Cmd) {
	id := cmd.String(cli.StringOpt{
		Name:      "i id",
		Value:     "",
		Desc:      "server id",
		HideValue: true,
	})
	force := cmd.Bool(cli.BoolOpt{
		Name:      "f force",
		Value:     false,
		Desc:      "force stop",
		HideValue: true,
	})
	cmd.Spec = "-if"
	cmd.Action = func() {
		err := GetAuthorizedClient().StopComputeServer(*id, *force)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Accepted Stop Request.\n\n")
	}
}
