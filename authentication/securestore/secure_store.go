package securestore

import (
	"github.com/99designs/keyring"
)

const (
	CollectionLabel = "RedHat" // Secret Service prefers no spaces
)

func getKeyringConfig() keyring.Config {
	return keyring.Config{
		// Generic
		ServiceName: CollectionLabel,
		// MacOS
		KeychainName:                   CollectionLabel,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: true,
		KeychainPasswordFunc:           keyring.TerminalPrompt,
		// Windows
		WinCredPrefix: CollectionLabel,
		// Secret Service
		LibSecretCollectionName: CollectionLabel,
		// KWallet
		KWalletFolder: CollectionLabel,
		KWalletAppID:  CollectionLabel,
		// KeyCtl
		KeyCtlScope: "user",
		// Pass
		PassDir:    "~/", // TODO: use default ocm config
		PassPrefix: CollectionLabel,
		// Encrypted File
		FilePasswordFunc: keyring.TerminalPrompt,
		FileDir:          "~/", // TODO: use default ocm config
	}
}
func upsertCredentials(labelKey string, creds []byte) error {
	ring, err := keyring.Open(getKeyringConfig())
	if err != nil {
		// TODO
	}

	err = ring.Set(keyring.Item{
		Key:  labelKey,
		Data: creds,
	})

	return err
}

func getCredentials(labelKey string) ([]byte, error) {
	credentials := []byte("")

	ring, err := keyring.Open(getKeyringConfig())
	if err != nil {
		// TODO
	}

	i, err := ring.Get(labelKey)
	if err != nil && err != keyring.ErrKeyNotFound {
		return credentials, err
	} else if err == keyring.ErrKeyNotFound {
		// Not Found
	} else {
		credentials = i.Data
	}

	return credentials, nil

}
