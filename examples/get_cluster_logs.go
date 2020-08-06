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

// This example shows how to retrieve the contents of the logs of a cluster.

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		log.Fatalf("Can't build logger: %v", err)
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

	// Get the client for the resource that manages the collection of clusters:
	clustersResource := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the collection of logs for the cluster that
	// we are looking for. Note that this will not send any request to the server yet, so it
	// will succeed even if that cluster doesn't exist.
	logsCollection := clustersResource.Cluster("1Jam7Ejgpm7AbZshbgaA9TsM1SQ").Logs()

	// Send the request to retrieve the collection of logs:
	listResponse, err := logsCollection.List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve list of logs: %v\n", err)
		os.Exit(1)
	}

	// The response obtained from the above list operation will contain the identifier of each
	// log, but not the content. To obtain the content it is necessary to send a request for
	// that specific log.
	listResponse.Items().Each(func(log *cmv1.Log) bool {
		logID := log.ID()
		logResource := logsCollection.Install()
		getResponse, err := logResource.Get().SendContext(ctx)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Can't retrive details of log '%s': %v\n",
				logID, err,
			)
			os.Exit(1)
		}
		log = getResponse.Body()
		logContent := log.Content()
		fmt.Printf("%s:\n%s\n", logID, logContent)
		return true
	})
}
