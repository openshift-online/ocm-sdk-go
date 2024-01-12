package securestore

import (
	"github.com/99designs/keyring"
)

const (
	SecureStoreConfigKey = "securestore"       // OCM_CONFIG key to enable secure OS store
	KindInternetPassword = "Internet password" // MacOS Keychain item kind
	LabelKey             = "RedHatSSO"
	CollectionLabel      = "OCM"
	KeychainName         = "login" // MacOS Keychain name
	DefaultFilePath      = "~/.config/ocm/ocm.json"
)

func getKeyringConfig() keyring.Config {
	return keyring.Config{
		// The order of the backends is important. The first backend in the list is the first one
		// that will attempt to be used.
		AllowedBackends: []keyring.BackendType{
			keyring.WinCredBackend,
			keyring.KeychainBackend,      // Tested
			keyring.SecretServiceBackend, // Tested
			keyring.KWalletBackend,
			keyring.KeyCtlBackend,
			keyring.PassBackend, // Tested
			// The FileBackend is a last resort and will store credentials in an encrypted file. This has
			// the worst user experience as the user will have to enter a password every time they attempt
			// to access the file.
			keyring.FileBackend, // Tested
		},
		// Generic
		ServiceName: CollectionLabel,
		// MacOS
		KeychainName:                   KeychainName,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
		// Windows
		WinCredPrefix: CollectionLabel,
		// Secret Service
		LibSecretCollectionName: CollectionLabel,
		// KWallet
		KWalletFolder: CollectionLabel,
		KWalletAppID:  CollectionLabel,
		// KeyCtl
		KeyCtlScope: "user",
		// Encrypted File
		FilePasswordFunc: keyring.TerminalPrompt,
		FileDir:          DefaultFilePath,
	}
}

// AvailableBackends provides a slice of all available backend keys on the current OS.
//
// The first backend in the slice is the first one that will be used.
func AvailableBackends() []string {
	b := []string{}
	for _, k := range keyring.AvailableBackends() {
		b = append(b, string(k))
	}
	return b
}

// UpsertConfigToKeyring will upsert the provided credentials to first priority OS secure store.
func UpsertConfigToKeyring(creds []byte, debug bool) error {
	ring, err := keyring.Open(getKeyringConfig())
	if err != nil {
		return err
	}

	err = ring.Set(keyring.Item{
		Label:       LabelKey,
		Key:         LabelKey,
		Description: KindInternetPassword,
		Data:        creds,
	})

	return err
}

// GetConfigFromKeyring will retrieve the credentials from the first priority OS secure store.
func GetConfigFromKeyring(debug bool) ([]byte, error) {
	credentials := []byte("")

	ring, err := keyring.Open(getKeyringConfig())
	if err != nil {
		return nil, err
	}

	i, err := ring.Get(LabelKey)
	if err != nil && err != keyring.ErrKeyNotFound {
		return credentials, err
	} else if err == keyring.ErrKeyNotFound {
		// Not found, continue
	} else {
		credentials = i.Data
	}

	return credentials, nil

}
