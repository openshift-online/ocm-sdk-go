/*
Copyright (c) 2024 Red Hat, Inc.

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
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Client(clientId, clientSecret).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().DNSDomains()

	// Prepare the description of the cluster to create:
	dns, err := cmv1.NewDNSDomain().
		ClusterArch(cmv1.ClusterArchitectureHcp).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create dns spec: %v\n", err)
		os.Exit(1)
	}

	// Send a request to create the cluster:
	response, err := collection.Add().
		Body(dns).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create dns: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	dns = response.Body()
	fmt.Printf("%s - %s\n", dns.ID())
}
