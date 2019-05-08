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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openshift-online/uhc-sdk-go/pkg/client"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := client.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("UHC_TOKEN")
	connection, err := client.NewConnectionBuilder().
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

	// Prepare the description of the cluster to create:
	cluster, err := v1.NewCluster().
		Name("mycluster").
		Flavour(
			v1.NewFlavour().
				ID("4"),
		).
		Region(
			v1.NewCloudRegion().
				ID("us-east-1"),
		).
		DNS(
			v1.NewDNS().
				BaseDomain("example.com"),
		).
		AWS(
			v1.NewAWS().
				AccessKeyID("...").
				SecretAccessKey("..."),
		).
		Version(
			v1.NewVersion().
				ID("openshift-v4.0-beta4"),
		).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create cluster description: %v\n", err)
		os.Exit(1)
	}

	// Send a request to create the cluster:
	response, err := collection.Add().
		Body(cluster).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create cluster: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	cluster = response.Body()
	fmt.Printf("%s - %s\n", cluster.ID(), cluster.Name())
}
