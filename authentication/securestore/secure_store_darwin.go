//go:build darwin
// +build darwin

package securestore

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

// var nativeStore = osxkeychain.Osxkeychain{}

func upsertCredentials(labelKey string, credentials []byte) error {
	// b, err := json.Marshal(cred)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Option 1
	// c := &credentials.Credentials{
	// 	ServerURL: "https://sso.redhat.com",
	// 	Username:  "token",
	// 	Secret:    cred,
	// }

	// nativeStore.Add(c)
	// secret, username, err := nativeStore.Get("https://sso.redhat.com")
	// fmt.Println(secret, username, err)
	// End: Option 1

	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassInternetPassword)
	item.SetLabel(labelKey)
	item.SetData(credentials)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleAfterFirstUnlockThisDeviceOnly)

	err := keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		// Item already exists, update it
		err = keychain.UpdateItem(item, item)
		if err != nil {
			return fmt.Errorf("error updating keychain item: %v", err)
		}
	}

	return nil
}

func getCredentials(labelKey string) ([]byte, error) {
	// secret, username, err := nativeStore.Get("https://sso.redhat.com")
	// secretCreds := Credentials{}
	// json.Unmarshal([]byte(secret), &secretCreds)
	// return secretCreds, username, err
	credentials := []byte("")

	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassInternetPassword)
	query.SetLabel(labelKey)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return credentials, fmt.Errorf("error fetching keychain item: %v", err)
	} else if len(results) != 1 {
		// Not found
	} else {
		credentials = results[0].Data
	}
	return credentials, nil
}
