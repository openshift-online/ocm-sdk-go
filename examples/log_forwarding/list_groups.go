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

// This example shows how to list log forwarding group versions.

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

	// Get the client for log forwarding groups:
	groupsResource := connection.ClustersMgmt().V1().
		LogForwarding().
		Groups()

	// Send a request to list the log forwarding group versions:
	response, err := groupsResource.List().
		Size(100).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve log forwarding groups: %v\n", err)
		os.Exit(1)
	}

	// Print the results:
	groups := response.Items()
	fmt.Printf("Available log forwarding group versions (%d):\n", groups.Len())

	groups.Each(func(group *cmv1.LogForwarderGroupVersions) bool {
		fmt.Printf("- Group ID: %s\n", group.ID())
		if group.State() != "" {
			fmt.Printf("  State: %s\n", group.State())
		}
		if len(group.Versions()) > 0 {
			fmt.Printf("  Available versions:\n")
			for _, version := range group.Versions() {
				fmt.Printf("    - %s\n", version.ID())
				if len(version.Applications()) > 0 {
					fmt.Printf("      Applications: %v\n", version.Applications())
				}
			}
		}
		fmt.Println()
		return true
	})

	if groups.Len() == 0 {
		fmt.Println("No log forwarding group versions found.")
	}
}
