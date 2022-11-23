// This example shows how to retrieve the collection of capabilities.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	search := fmt.Sprint("type='Account' and internal='true'")
	labelsResource := connection.AccountsMgmt().V1().Labels()
	labelsResponse, err := labelsResource.List().Size(5).Search(search).Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve labels: %s\n", err)
		os.Exit(1)
	}

	labelsResponse.Items().Each(func(l *v1.Label) bool {
		fmt.Println(l.AccountID())
		return true
	})

}
