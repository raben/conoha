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
	name := cmd.String(cli.StringOpt{
		Name:      "n name",
		Value:     "",
		Desc:      "Flavor Name",
		HideValue: true,
	})

	cmd.Action = func() {
		computeFlavors, err := GetAuthorizedClient().GetComputeFlavors(*name)
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

func ComputeServer(cmd *cli.Cmd) {
	serverId := cmd.String(cli.StringOpt{
		Name:      "i serverId",
		Value:     "",
		Desc:      "Server Id",
		HideValue: true,
	})

	cmd.Spec = "-i"
	cmd.Action = func() {
		computeServer, err := GetAuthorizedClient().GetComputeServer(*serverId)
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap([]conoha.ComputeServersValue{computeServer.Server})
	}
}

func ComputeImages(cmd *cli.Cmd) {
	name := cmd.String(cli.StringOpt{
		Name:      "n name",
		Value:     "",
		Desc:      "Image Name",
		HideValue: true,
	})

	cmd.Action = func() {
		computeImages, err := GetAuthorizedClient().GetComputeImages(*name)
		if err != nil {
			log.Fatal(err)
		}

		SliceToMap(computeImages.Images)
	}
}

func ComputeServerCreate(cmd *cli.Cmd) {
	image := cmd.String(cli.StringOpt{
		Name:      "i image",
		Value:     "",
		Desc:      "Image Name",
		HideValue: true,
	})
	flavor := cmd.String(cli.StringOpt{
		Name:      "f flavor",
		Value:     "",
		Desc:      "Flavor Name",
		HideValue: true,
	})
	cmd.Spec = "-i -f"
	cmd.Action = func() {
		err := GetAuthorizedClient().CreateComputeServer(*image, *flavor)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Accepted Create Request.\n\n")
	}
}

func ComputeServerRemove(cmd *cli.Cmd) {
	id := cmd.String(cli.StringOpt{
		Name:      "i id",
		Value:     "",
		Desc:      "Server Id",
		HideValue: true,
	})
	cmd.Spec = "-i"
	cmd.Action = func() {
		err := GetAuthorizedClient().RemoveComputeServer(*id)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Accepted Remove Request.\n\n")
	}
}

func ComputeServerStart(cmd *cli.Cmd) {
	id := cmd.String(cli.StringOpt{
		Name:      "i id",
		Value:     "",
		Desc:      "Server Id",
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
		Desc:      "Server Id",
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

func ComputeServerRestart(cmd *cli.Cmd) {
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
		err := GetAuthorizedClient().RestartComputeServer(*id, *force)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Accepted Restart Request.\n\n")
	}
}
