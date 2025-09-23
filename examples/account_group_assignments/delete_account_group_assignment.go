/*
Copyright (c) 2025 Red Hat, Inc.

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

// This example shows how to delete an account group assignment.

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
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
	url := os.Getenv("OCM_ENV")
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		URL(url).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get organization ID from environment variable
	orgID := os.Getenv("OCM_ORG_ID")
	if orgID == "" {
		fmt.Fprintf(os.Stderr, "OCM_ORG_ID environment variable is required\n")
		os.Exit(1)
	}

	// Get the client for the resource that manages the collection of account group assignments:
	collection := connection.AccountsMgmt().V1().Organizations().Organization(orgID).AccountGroupAssignments()

	// Get the client for the resource that manages the account group assignment that we want to delete. Note
	// that this will not send any request to the server yet, so it will succeed even if the
	// assignment doesn't exist.
	resource := collection.AccountGroupAssignment("32qBymq4rqkf0P1YIfzdwCMwvvR")

	// Send the request to delete the account group assignment:
	_, err = resource.Delete().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't delete account group assignment: %v\n", err)
		os.Exit(1)
	}
}
