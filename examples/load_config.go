/*
Copyright (c) 2020 Red Hat, Inc.

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

// This example shows how to load the configuration of the connection from an
// external file.

package main

import (
	"context"
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	cmv1 "github.com/openshift-online/ocm-sdk-go/v2/clustersmgmt/v1"
)

func loadConfig(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	connection, err := sdk.NewConnection().
		Logger(logger).
		Load("load_config.yaml").
		Build()
	if err != nil {
		return err
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
			Send(ctx)
		if err != nil {
			return err
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

	return nil
}
