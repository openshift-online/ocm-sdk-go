/*
Copyright (c) 2018 Red Hat, Inc.

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

// This example shows how to perform an resource review request.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	azv1 "github.com/openshift-online/ocm-sdk-go/authorizations/v1"
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

	// Get the client for the resource that manages resource reviews:
	resource := connection.Authorizations().V1().ResourceReview()

	// Prepare the request:
	request, err := azv1.NewResourceReviewRequest().
		AccountUsername("alice").
		Action("get").
		ResourceType("Cluster").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build request: %v\n", err)
		os.Exit(1)
	}

	// Send a request to get the list identifiers of clusters that we can see:
	response, err := resource.Post().
		Request(request).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't review resource: %v\n", err)
		os.Exit(1)
	}

	// Print the results:
	ids := response.Review().ClusterIDs()
	for _, id := range ids {
		fmt.Printf("%s\n", id)
	}
}
