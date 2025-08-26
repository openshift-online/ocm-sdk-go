/*
Copyright (c) 2025 Red Hat, Inc.

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
	clusterId := "test-cluster-id"
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
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of image mirrors:
	collection := connection.ClustersMgmt().V1().Clusters().Cluster(clusterId).ImageMirrors()

	// Create the image mirror object:
	imageMirror, err := cmv1.NewImageMirror().
		Source("quay.io/example/source").
		Mirrors("mirror.example.com/example/source").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build image mirror: %v\n", err)
		os.Exit(1)
	}

	// Send the request to create the image mirror:
	response, err := collection.Add().
		Body(imageMirror).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create image mirror: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	createdImageMirror := response.Body()
	fmt.Printf("Created image mirror ID: %s\n", createdImageMirror.ID())
	fmt.Printf("Source: %s\n", createdImageMirror.Source())
	fmt.Printf("Mirrors: %v\n", createdImageMirror.Mirrors())
}
