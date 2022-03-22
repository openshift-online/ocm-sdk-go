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

// This example shows how two create a connection using an access token and a refresh token that
// have been previously created by some other means.

package main

import (
	"context"
	"fmt"
	"io/ioutil"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	cmv1 "github.com/openshift-online/ocm-sdk-go/v2/clustersmgmt/v1"
)

func existingToken(ctx context.Context, args []string) error {
	// Get the tokens, maybe receiving it from some other part of the application, or reading
	// them from a file where they have been previously saved:
	accessToken, err := ioutil.ReadFile("access.token")
	if err != nil {
		return err
	}
	refreshToken, err := ioutil.ReadFile("refresh.token")
	if err != nil {
		return err
	}

	// Create the connection, and remember to close it. Note that this connection will stop
	// working when both tokens expire. This can happen, for example, if the connection isn't
	// used for period of time longer than the life of the refresh token.
	connection, err := sdk.NewConnection().
		Logger(logger).
		Tokens(string(accessToken), string(refreshToken)).
		BuildContext(ctx)
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Retrieve the collection of clusters:
	response, err := collection.List().
		Search("name like 'my%'").
		Page(1).
		Size(10).
		SendContext(ctx)
	if err != nil {
		return err
	}

	// Print the result:
	response.Items().Each(func(cluster *cmv1.Cluster) bool {
		fmt.Printf("%s - %s\n", cluster.ID(), cluster.Name())
		return false
	})

	return nil
}
