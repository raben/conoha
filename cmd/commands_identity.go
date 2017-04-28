package cmd

import (
	"fmt"
	"log"
	"syscall"

	"github.com/jawher/mow.cli"
	conoha "github.com/raben/conoha/lib/models"
	terminal "golang.org/x/crypto/ssh/terminal"
)

var region *string
var tenantId *string
var userName *string
var password *string

func IdentityVersion(cmd *cli.Cmd) {
	cmd.Action = func() {
		identityVersion, err := GetClient().GetIdentityVersion()
		if err != nil {
			log.Fatal(err)
		}
		if len(identityVersion.Versions.Values) == 0 {
			log.Fatal(err)
		}
		SliceToMap(identityVersion.Versions.Values)
	}
}

func IdentityToken(cmd *cli.Cmd) {

	tenantId = cmd.String(cli.StringOpt{
		Name:      "t tenantId",
		Value:     "",
		Desc:      "tenant id",
		HideValue: true,
	})
	userName = cmd.String(cli.StringOpt{
		Name:      "u userName",
		Value:     "",
		Desc:      "user name",
		HideValue: true,
	})

	cmd.Spec = "-t -u"

	cmd.Action = func() {
		fmt.Print("Password:\n")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Cannot Empty Password.")
		}

		var auth = conoha.Auth{
			UserName: *userName,
			Password: string(password),
			TenantId: *tenantId,
		}
		identity, err := GetClient().GetIdentityToken(auth)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("\n")
		fmt.Print(identity.Access.Token.Id)
		fmt.Print("\n\n")
	}
}
