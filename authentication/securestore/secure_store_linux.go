//go:build linux
// +build linux

package securestore

import (
	"github.com/docker/docker-credential-helpers/secretservice"
)

var nativeStore = secretservice.SecretService{}

func upsertCredentials(labelKey string, credentials []byte) error {
	c := &credentials.Credentials{
		ServerURL: labelKey,
		Username:  labelKey,
		Secret:    credentials,
	}

	err := nativeStore.Add(c)
	return nil
}

func getCredentials(labelKey string) ([]byte, error) {
	secret, err := nativeStore.Get(labelKey)
	return secret, err
}
