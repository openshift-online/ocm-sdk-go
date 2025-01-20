package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
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
		_, _ = fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can't create connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of migrations for a cluster. Note
	// that this will not send any request to the server yet, so it will succeed even if the cluster
	// doesn't exist.
	collection := connection.ClustersMgmt().
		V1().
		Clusters().
		Cluster("2gaog4ns6phso6q17v07rj3rre9dm24e").
		Migrations()

	// Create a request body for the specific cluster migration.
	requestBuilder := v1.ClusterMigrationBuilder{}
	requestBuilder.Type(v1.ClusterMigrationTypeSdnToOvn) // Type is required

	// Create a builder for the specific migration type's configuration if necessary
	sdnToOvnBuilder := &v1.SdnToOvnClusterMigrationBuilder{}
	sdnToOvnBuilder.TransitIpv4("192.168.255.0/24")
	sdnToOvnBuilder.MasqueradeIpv4("192.168.255.0/24")
	sdnToOvnBuilder.JoinIpv4("192.168.255.0/24")
	requestBuilder.SdnToOvn(sdnToOvnBuilder)

	requestBody, err := requestBuilder.Build()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to create a cluster migration request: %v\n", err)
		os.Exit(1)
	}

	// Send the request to add a cluster migration.
	response, err := collection.Add().Body(requestBody).SendContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to create cluster migration: %v\n", err)
		os.Exit(1)
	}

	// Perform next actions with the response
	fmt.Println(response)
}
