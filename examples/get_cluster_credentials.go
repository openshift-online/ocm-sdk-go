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

// This example shows how to retrieve the credentials of a cluster.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := sdk.NewGoLoggerBuilder().
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

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the credentials for the cluster that we are
	// looking for. Note that this will not send any request to the server yet, so it will
	// succeed even if that cluster doesn't exist.
	resource := collection.Cluster("123").Credentials()

	// Send the request to retrieve the credentials:
	response, err := resource.Get().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve cluster credentials: %v\n", err)
		os.Exit(1)
	}

	// Print the credentials:
	credentials := response.Body()
	fmt.Printf("Admin user: %s\n", credentials.Admin().User())
	fmt.Printf("Admin password: %s\n", credentials.Admin().Password())
	fmt.Printf("SSH public key:\n%s\n", credentials.SSH().PublicKey())
	fmt.Printf("SSH private key:\n%s\n", credentials.SSH().PrivateKey())
	fmt.Printf("Kubeconfig:\n%s\n", credentials.Kubeconfig())
}
