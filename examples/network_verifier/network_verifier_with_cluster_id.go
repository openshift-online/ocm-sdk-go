package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
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
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:

	clusterId := "2726d3ceb9q7qso3nceqdlm3j2fh7oqg"
	body, err := cmv1.NewNetworkVerification().ClusterId(clusterId).Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build request body: %v\n", err)
		os.Exit(1)
	}

	networkVerifyResponse, err := connection.ClustersMgmt().
		V1().NetworkVerifications().Add().Body(body).SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Can't send request to network verifier using cluster id: %v\n", clusterId)
		os.Exit(1)
	}

	// Print the result
	fmt.Printf("%#v", networkVerifyResponse.Body())

	// Check the inflight check (egress type) after a couple of minutes
	// to see the update applied from the network verifier
	inflightResponse, err := connection.ClustersMgmt().
		V1().Clusters().Cluster(clusterId).InflightChecks().List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Can't retrieve inflight request using cluster id: %v\n", clusterId)
		os.Exit(1)
	}
	inflights := inflightResponse.Items()
	// Print the result
	fmt.Printf("%#v", inflights)

}
