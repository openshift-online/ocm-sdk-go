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

// This example shows how to list the cloud providers and their regions.

package main

import (
	"context"
	"fmt"
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
	providersCollection := connection.ClustersMgmt().V1().CloudProviders()

	// Retrieve the first page of cloud providers and print them:
	providersResponse, err := providersCollection.List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve providers: %v\n", err)
		os.Exit(1)
	}
	providersResponse.Items().Each(func(provider *cmv1.CloudProvider) bool {
		providerID := provider.ID()
		regionsCollection := providersCollection.CloudProvider(providerID).Regions()
		regionsResponse, err := regionsCollection.List().SendContext(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't retrieve regions: %v\n", err)
			os.Exit(1)
		}
		regionsResponse.Items().Each(func(region *cmv1.CloudRegion) bool {
			regionID := region.ID()
			fmt.Printf("%s - %s\n", providerID, regionID)
			return true
		})
		return true
	})
}
