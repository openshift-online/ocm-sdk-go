/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This example shows how to use the support for Prometheus metrics.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create and start a Prometheus metric server that will respond to any request in port
	// 8000, regardless of the request path.
	go http.ListenAndServe(":8000", promhttp.Handler())

	// Create a logger that has the debug level enabled:
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, specifying the `api_outbound` subsystem so that metrics are
	// enabled and available with the `api_outbound_` prefix.
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		Metrics("api_outbound").
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	clustersCollection := connection.ClustersMgmt().V1().Clusters()

	// Send requests to retrieve the first page of clusters, the details of the cluster, the
	// logs, and the credentials, in a loop, to accumulate metrics. To see the metrics point
	// your browser to http://localhost:8000.
	for {
		// Get the list of clusters:
		clustersListResponse, err := clustersCollection.List().SendContext(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't send list request: %v\n", err)
			os.Exit(1)
		}

		// For each cluster:
		clustersListResponse.Items().Each(func(cluster *cmv1.Cluster) bool {
			// Get the details:
			clusterResource := clustersCollection.Cluster(cluster.ID())
			clusterResource.Get().SendContext(ctx)

			// Get the list of logs:
			logsResource := clusterResource.Logs()
			logsResource.List().SendContext(ctx)

			// Get the credentials:
			credentialsResource := clusterResource.Credentials()
			credentialsResource.Get().SendContext(ctx)

			return true
		})

		// Wait a bit:
		time.Sleep(1 * time.Second)
	}
}
