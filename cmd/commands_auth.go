package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
	"time"

	"github.com/jawher/mow.cli"
	conoha "github.com/raben/conoha/lib/models"
	terminal "golang.org/x/crypto/ssh/terminal"
)

const (
	ConfigDir  = ".conoha"
	ConfigFile = "config"
)

func Init(cmd *cli.Cmd) {
	cmd.Action = func() {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Region: ")
		region, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Cannot Empty Region.")
		}
		fmt.Print("TenantId: ")
		tenantId, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Cannot Empty TenantId.")
		}
		fmt.Print("UserName: ")
		userName, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error Empty UserName.")
		}
		fmt.Print("Password: ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Error Empty Password.")
		}

		var auth = conoha.Auth{
			UserName: strings.Trim(string(userName), "\n"),
			Password: strings.Trim(string(password), "\n"),
			Region:   strings.Trim(string(region), "\n"),
			TenantId: strings.Trim(string(tenantId), "\n"),
		}
		CreateConfig(auth)
	}
}

func Verify() (config conoha.AuthConfig, err error) {
	config, err = ReadConfig()
	if err != nil {
		log.Fatal("Cannot Read Config File")
		return
	}
	if config.ExpiredAt.Before(time.Now()) {
		fmt.Print("Password: ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Error Empty Password.")
		}

		var auth = conoha.Auth{
			UserName: strings.Trim(config.UserName, "\n"),
			Password: strings.Trim(string(password), "\n"),
			Region:   strings.Trim(config.Region, "\n"),
			TenantId: strings.Trim(config.TenantId, "\n"),
		}
		config, err := CreateConfig(auth)
		if err != nil {
			log.Fatal(err)
		}
		return config, err
	}
	return config, err
}

func ReadConfig() (config conoha.AuthConfig, err error) {
	usr, err := user.Current()
	if err != nil {
		return config, err
	}

	dir := usr.HomeDir + "/" + ConfigDir
	file, e := ioutil.ReadFile(dir + "/" + ConfigFile)
	if e != nil {
		log.Fatal("Not Found Config File. Please exec 'conoha init'.")
	}
	json.Unmarshal(file, &config)

	return config, nil
}

func CreateConfig(auth conoha.Auth) (config conoha.AuthConfig, err error) {
	identity, err := GetClient().GetIdentityToken(auth)
	if err != nil {
		return config, err
	}

	config = conoha.AuthConfig{
		Region:    auth.Region,
		TenantId:  auth.TenantId,
		UserName:  auth.UserName,
		AuthToken: identity.Access.Token.Id,
		ExpiredAt: identity.Access.Token.Expires,
	}
	data, err := json.Marshal(config)
	if err != nil {
		return config, err
	}

	usr, err := user.Current()
	if err != nil {
		return config, err
	}
	dir := usr.HomeDir + "/" + ConfigDir
	os.MkdirAll(dir, 0755)
	os.Create(dir + "/" + ConfigFile)
	err = ioutil.WriteFile(dir+"/"+ConfigFile, data, 0755)
	if err != nil {
		return config, err
	}

	return config, nil
}
