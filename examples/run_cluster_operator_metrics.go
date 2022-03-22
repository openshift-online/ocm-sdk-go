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

// This example shows how to use the metric queries provided for clusters management service.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
)

func runClusterOperatorMetrics(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the resource that manages the the metrics query that we want to use:
	resource := connection.ClustersMgmt().V1().
		Clusters().
		Cluster("...").
		MetricQueries().
		ClusterOperators()

	// Send the request to run the query:
	response, err := resource.Get().SendContext(ctx)
	if err != nil {
		return err
	}

	// Print the results:
	for _, operator := range response.Body().Operators() {
		fmt.Printf("Name: %v\n", operator.Name())
		fmt.Printf("Reason: %s\n", operator.Reason())
		fmt.Printf("Time: %f\n", operator.Time())
		fmt.Printf("Version: %f\n", operator.Version())
		fmt.Printf("Condition: %f\n", operator.Condition())
	}

	return nil
}
