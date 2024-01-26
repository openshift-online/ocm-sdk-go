package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/logging"
)

// This example retrieves a cluster using an OCM connection that does
// not perform authentication. The example assumes that there is a
// local Cluster Service process running at http://localhost:9000 and
// that no Authentication is enabled in it.
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

	// Create the connection without using authentication,
	// and remember to close it:
	connection, err := sdk.NewUnauthenticatedConnectionBuilder().
		Logger(logger).
		URL("http://localhost:9000").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the cluster that we are looking for. Note
	// that this will not send any request to the server yet, so it will succeed even if that
	// cluster doesn't exist.
	resource := collection.Cluster("1i1rics0s96htk98gmh50rkakdljoiau")

	// Send the request to retrieve the cluster:
	response, err := resource.Get().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve cluster: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	cluster := response.Body()
	fmt.Printf("%s - %s\n", cluster.ID(), cluster.Name())

}
