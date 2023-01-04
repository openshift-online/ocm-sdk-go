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

	// Get the client for the resource that manages the collection of capabilities:
	collection := connection.AccountsMgmt().V1().Capabilities()

	// Retrieve all capabilities
	response, err := collection.List().
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve capabilities: %s\n", err)
		os.Exit(1)
	}

	// Display the capabilities:
	response.Items().Each(func(capability *v1.Capability) bool {
		fmt.Printf("%s - %s\n", capability.Name(), capability.Value())
		return true
	})

	// Retrieve first page of 10 capabilities
	size := 10
	page := 1
	// Retrieve the page:
	response, err = collection.List().
		Size(size).
		Page(page).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve page %d: %s\n", page, err)
		os.Exit(1)
	}

	// Display the page:
	response.Items().Each(func(capability *v1.Capability) bool {
		fmt.Printf("%s - %s\n", capability.Name(), capability.Value())
		return true
	})

	// Search capabilities
	response, err = collection.List().Search("name like 'Organization%'").
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve capabilities: %s\n", err)
		os.Exit(1)
	}

	// Display the capabilities:
	response.Items().Each(func(capability *v1.Capability) bool {
		fmt.Printf("%s - %s\n", capability.Name(), capability.Value())
		return true
	})

}
