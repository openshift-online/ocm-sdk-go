/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may you may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This example shows how to list available applications for log forwarding.

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

	// Get the client for log forwarding applications:
	applicationsResource := connection.ClustersMgmt().V1().
		LogForwarding().
		Applications()

	// Send a request to list the available applications:
	response, err := applicationsResource.List().
		Size(100).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve log forwarding applications: %v\n", err)
		os.Exit(1)
	}

	// Print the results:
	applications := response.Items()
	fmt.Printf("Available log forwarding applications (%d):\n", applications.Len())

	applications.Each(func(application *cmv1.LogForwarderApplication) bool {
		fmt.Printf("- Application ID: %s\n", application.ID())
		if application.State() != "" {
			fmt.Printf("  State: %s\n", application.State())
		}
		return true
	})

	if applications.Len() == 0 {
		fmt.Println("No log forwarding applications found.")
	}
}
