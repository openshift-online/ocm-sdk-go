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

	billingModelsResource := connection.AccountsMgmt().V1().BillingModels()
	billingModelsResoponse, err := billingModelsResource.List().Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve billing models: %s\n", err)
		os.Exit(1)
	}

	billingModelsResoponse.Items().Each(func(l *v1.BillingModelItem) bool {
		fmt.Println(l.Id())
		return true
	})

}
