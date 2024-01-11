// //go:build windows
// // +build windows

package securestore

// import (
// 	"fmt"

// 	"github.com/docker/docker-credential-helpers/credentials"
// 	"github.com/docker/docker-credential-helpers/wincred"
// )

// var nativeStore = wincred.WinCred{}

// func upsertCredentials() {
// 	c := &credentials.Credentials{
// 		ServerURL: "https://sso.redhat.com",
// 		Username:  "token",
// 		Secret:    "",
// 	}

// 	nativeStore.Add(c)
// 	secret, username, err := nativeStore.Get("https://sso.redhat.com")
// 	fmt.Println(secret, username, err)
// }

// func getCreds() (string, string, error) {
// 	secret, username, err := nativeStore.Get("https://sso.redhat.com")
// 	return secret, username, err
// }
