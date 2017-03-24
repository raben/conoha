package cmd

import (
	conoha "github.com/raben/conoha/lib"
	"log"
)

func GetClient() *conoha.Client {
	return conoha.NewClient()
}
func GetAuthorizedClient() *conoha.Client {
	auth, err := Verify()
	if err != nil {
		log.Fatal(err)
	}
	return GetClient().SetAuth(auth)
}
