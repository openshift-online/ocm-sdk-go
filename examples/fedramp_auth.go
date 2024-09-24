/*
Copyright (c) 2024 Red Hat, Inc.

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

// This example shows how create a connection using the OpenID refresh token grant for
// FedRAMP authentication.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/authentication"
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

	// Create the connection, and remember to close it:
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		URL(sdk.FedRAMPURL).
		TokenURL(authentication.FedRAMPTokenURL).
		Tokens(os.Getenv("OCM_REFRESH_TOKEN")).
		Client(authentication.FedRAMPClientID, "").
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the service that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Retrieve the collection of clusters:
	response, err := collection.List().
		Search("name like 'my%'").
		Page(1).
		Size(10).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve clusters: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	response.Items().Each(func(cluster *cmv1.Cluster) bool {
		fmt.Printf("%s - %s\n", cluster.ID(), cluster.Name())
		return true
	})
}
