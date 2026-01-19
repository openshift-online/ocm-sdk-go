package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
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

	// Create the connection builder with logger and tokens.
	token := os.Getenv("OCM_TOKEN")
	connectionBuilder := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token)

	// In order to test against another URL we're using OCM_API_URL environment, by default it'll be production.
	url := os.Getenv("OCM_API_URL")
	if url != "" {
		connectionBuilder.URL(url)
	}

	// Create the connection, and remember to close it:
	connection, err := connectionBuilder.BuildContext(ctx)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can't create connection: %v\n", err)
		os.Exit(1)
	}

	defer connection.Close()

	deletedClustersClient := connection.ClustersMgmt().V1().DeletedClusters()

	deletedClusters, err := deletedClustersClient.List().SendContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Can't retrieve deleted clusters: %v\n", err)
		os.Exit(1)
	}

	// Perform next actions with the retrieved deleted clusters
	fmt.Println(deletedClusters)
}
