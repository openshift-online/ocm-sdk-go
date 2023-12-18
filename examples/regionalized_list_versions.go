/*
Copyright (c) 2023 Red Hat, Inc.

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

// This example shows how to retrieve the collection of clusters from a specific OCM region

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

	regValue, err := sdk.GetRhRegion("https://api.integration.openshift.com", "singapore")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't find region: %v", err)
		os.Exit(1)
	}
	token := os.Getenv("OCM_TOKEN")

	// Build a regionalized connection based on the desired shard
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		URL(fmt.Sprintf("https://%s", regValue.URL)). // Apply the region URL
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Versions()

	// Retrieve the list of clusters using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	// Retrieve the page:
	response, err := collection.List().
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve versions : %s\n", err)
		os.Exit(1)
	}

	// Display the page:
	response.Items().Each(func(version *cmv1.Version) bool {
		fmt.Printf("%s\n", version.ID())
		return true
	})

}
