package main

import (
	"fmt"

	"github.com/openshift-online/ocm-sdk-go/authentication/securestore"
)

func main() {
	// Get a list of available backends on the current OS
	available := securestore.AvailableBackends()
	fmt.Printf("Available backends: %v\n", available)

	// Create bytes
	config := []byte("mybytestring")

	// Upsert to keyring
	securestore.UpsertConfigToKeyring(config)

	// Upsert again to keyring
	config = []byte("mybytestringagain")
	securestore.UpsertConfigToKeyring(config)

	// Read bytes back from Keyring
	readVal, _ := securestore.GetConfigFromKeyring()
	fmt.Printf("Read from keyring: %s\n", string(readVal))
}
