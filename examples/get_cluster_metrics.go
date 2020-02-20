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

// This example shows how to retrieve the basic metrics of a cluster.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
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

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the the cluster that we are looking for.
	// Note that this will not send any request to the server yet, so it will succeed even if
	// that cluster doesn't exist.
	resource := collection.Cluster("1Jam7Ejgpm7AbZshbgaA9TsM1SQ")

	// Send the request to retrieve the details of the cluster:
	response, err := resource.Get().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve cluster details: %v\n", err)
		os.Exit(1)
	}

	// Print the metrics:
	cluster := response.Body()
	metrics := cluster.Metrics()
	cpu := metrics.CPU()
	memory := metrics.Memory()
	fmt.Printf("CPU total: %s\n", valueString(cpu.Total()))
	fmt.Printf("CPU used: %s\n", valueString(cpu.Used()))
	fmt.Printf("Memory total: %s\n", valueString(memory.Total()))
	fmt.Printf("CPU used: %s\n", valueString(memory.Used()))
}

// valueString converts an API value to an string, taking into account that it may be nil, and that
// it may not have an unit.
func valueString(v *cmv1.Value) string {
	if v == nil {
		return "N/A"
	}
	value, ok := v.GetValue()
	if !ok {
		return "N/A"
	}
	unit, ok := v.GetUnit()
	if !ok {
		return fmt.Sprintf("%.2f", value)
	}
	return fmt.Sprintf("%.2f %s", value, unit)
}
