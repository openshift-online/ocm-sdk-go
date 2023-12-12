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
	"encoding/json"
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

	// USE CASE 1:
	// You do not have the desired region ID and need to query the global region to find URL

	// Create the global connection
	token := os.Getenv("OCM_TOKEN")
	globalConnection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		RhRegion("integration"). // Define the integration global region
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer globalConnection.Close()

	// Fetch the shards from the global region
	response, err := globalConnection.Get().Path("/static/ocm-shards.json").SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve shards: %s\n", err)
		os.Exit(1)
	}

	// Turn response into an interface with regions
	data := response.Bytes()

	// turn region bytes into a map
	var regions map[string]interface{}
	err = json.Unmarshal(data, &regions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't unmarshal shards: %s\n", err)
		os.Exit(1)
	}

	// Grab the singapore region URL
	regionURL := regions["rh-singapore"].(map[string]interface{})["url"].(string)

	// Build a regionalized connection based on the desired shard
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		URL(fmt.Sprintf("https://%s", regionURL)). // Apply the region URL
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Retrieve the list of clusters using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	size := 10
	page := 1
	for {
		// Retrieve the page:
		response, err := collection.List().
			Search("name like 'my%'").
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't retrieve page %d: %s\n", page, err)
			os.Exit(1)
		}

		// Display the page:
		response.Items().Each(func(cluster *cmv1.Cluster) bool {
			fmt.Printf("%s - %s - %s\n", cluster.ID(), cluster.Name(), cluster.State())
			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if response.Size() < size {
			break
		}
		page++
	}

	// USE CASE 2:
	// You have a specific region ID and want to query that region directly

	// Build the regionalized connection based on the desired region ID
	connection, err = sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		RhRegion("aws.xcm.integration"). // Apply the rh region URL
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	collection = connection.ClustersMgmt().V1().Clusters()

	// Retrieve the list of clusters using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	for {
		// Retrieve the page:
		response, err := collection.List().
			Search("name like 'my%'").
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't retrieve page %d: %s\n", page, err)
			os.Exit(1)
		}

		// Display the page:
		response.Items().Each(func(cluster *cmv1.Cluster) bool {
			fmt.Printf("%s - %s - %s\n", cluster.ID(), cluster.Name(), cluster.State())
			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if response.Size() < size {
			break
		}
		page++
	}
}
