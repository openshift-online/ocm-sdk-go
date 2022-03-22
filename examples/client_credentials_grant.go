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

// This example shows how create a connection using the OpenID client credentials grant for
// authentication instead of the resource owner password grant.

package main

import (
	"context"
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	cmv1 "github.com/openshift-online/ocm-sdk-go/v2/clustersmgmt/v1"
)

func clientCredentialsGrant(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	connection, err := sdk.NewConnection().
		Logger(logger).
		Client("myclientid", "myclientsecret").
		Build()
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the service that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Retrieve the collection of clusters:
	response, err := collection.List().
		Search("name like 'my%'").
		Page(1).
		Size(10).
		Send(ctx)
	if err != nil {
		return err
	}

	// Print the result:
	response.Items().Each(func(cluster *cmv1.Cluster) bool {
		fmt.Printf("%s - %s\n", cluster.ID(), cluster.Name())
		return true
	})

	return nil
}
