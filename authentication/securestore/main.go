package securestore

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/99designs/keyring"
)

const (
	SecureStoreConfigKey = "securestore"       // OCM_CONFIG key to enable secure OS store
	KindInternetPassword = "Internet password" // MacOS Keychain item kind
	ItemKey              = "RedHatSSO"
	CollectionName       = "login" // Common OS default collection name
	MaxWindowsByteSize   = 2500    // Windows Credential Manager has a 2500 byte limit
)

func getKeyringConfig() keyring.Config {
	return keyring.Config{
		// The order of the backends is important. The first backend in the list is the first one
		// that will attempt to be used.
		AllowedBackends: []keyring.BackendType{
			keyring.WinCredBackend,
			keyring.KeychainBackend,
			keyring.SecretServiceBackend,
			keyring.PassBackend,
		},
		// Generic
		ServiceName: ItemKey,
		// MacOS
		KeychainName:                   CollectionName,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
		// Windows
		WinCredPrefix: ItemKey,
		// Secret Service
		LibSecretCollectionName: CollectionName,
	}
}

// AvailableBackends provides a slice of all available backend keys on the current OS.
//
// Note: CGO_ENABLED=1 is required for OSX Keychain and darwin builds
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
//
// Note: CGO_ENABLED=1 is required for OSX Keychain and darwin builds
func UpsertConfigToKeyring(creds []byte) error {
	ring, err := keyring.Open(getKeyringConfig())
	if err != nil {
		return err
	}

	compressed, err := compressConfig(creds)
	if err != nil {
		return err
	}

	// check if available backend contains windows credential manager and exceeds the byte limit
	if len(compressed) > MaxWindowsByteSize &&
		keyring.AvailableBackends()[0] == keyring.WinCredBackend {
		return fmt.Errorf("credentials are too large for Windows Credential Manager: %d bytes (max %d)", len(compressed), MaxWindowsByteSize)
	}

	err = ring.Set(keyring.Item{
		Label:       ItemKey,
		Key:         ItemKey,
		Description: KindInternetPassword,
		Data:        compressed,
	})

	return err
}

// GetConfigFromKeyring will retrieve the credentials from the first priority OS secure store.
//
// Note: CGO_ENABLED=1 is required for OSX Keychain and darwin builds
func GetConfigFromKeyring() ([]byte, error) {
	credentials := []byte("")

	ring, err := keyring.Open(getKeyringConfig())
	if err != nil {
		return nil, err
	}

	i, err := ring.Get(ItemKey)
	if err != nil && err != keyring.ErrKeyNotFound {
		return credentials, err
	} else if err == keyring.ErrKeyNotFound {
		// Not found, continue
	} else {
		credentials = i.Data
	}

	if len(credentials) == 0 {
		// No creds to decompress, return early
		return credentials, nil
	}

	creds, err := decompressConfig(credentials)
	if err != nil {
		return nil, err
	}

	return creds, nil

}

// Compresses credential bytes to help ensure all OS secure stores can store the data.
// Windows Credential Manager has a 2500 byte limit.
func compressConfig(creds []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err := gz.Write(creds)
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// Decompresses credential bytes
func decompressConfig(creds []byte) ([]byte, error) {
	reader := bytes.NewReader(creds)
	gzreader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	output, err := io.ReadAll(gzreader)
	if err != nil {
		return nil, err
	}

	return output, err
}
