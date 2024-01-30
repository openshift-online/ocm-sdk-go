// This example shows how to retrieve the collection of capabilities.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openshift-online/ocm-sdk-go/authentication"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/logging"
)

func main() {
	// Create a context:
	ctx := context.Background()
	clientId := "ocm-cli"

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token, err := authentication.InitiateAuthCode(clientId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get token: %v\n", err)
		os.Exit(1)
	}
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Client(clientId, ""). // Required
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	collection := connection.AccountsMgmt().V1().CurrentAccount()

	// Retrieve current account information
	_, err = collection.Get().
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve account: %s\n", err)
		os.Exit(1)
	}
}
