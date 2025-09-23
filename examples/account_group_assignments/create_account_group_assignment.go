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

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	amv1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
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

	// Get required environment variables
	orgID := os.Getenv("OCM_ORG_ID")
	if orgID == "" {
		fmt.Fprintf(os.Stderr, "OCM_ORG_ID environment variable is required\n")
		os.Exit(1)
	}

	accountID := os.Getenv("OCM_ACCOUNT_ID")
	if accountID == "" {
		fmt.Fprintf(os.Stderr, "OCM_ACCOUNT_ID environment variable is required\n")
		os.Exit(1)
	}

	accountGroupID := os.Getenv("OCM_ACCOUNT_GROUP_ID")
	if accountGroupID == "" {
		fmt.Fprintf(os.Stderr, "OCM_ACCOUNT_GROUP_ID environment variable is required\n")
		os.Exit(1)
	}

	// Get the client for the resource that manages the collection of account group assignments:
	collection := connection.AccountsMgmt().V1().Organizations().Organization(orgID).AccountGroupAssignments()

	// Prepare the description of the account group assignment to create:
	assignment, err := amv1.NewAccountGroupAssignment().
		AccountID(accountID).
		AccountGroupID(accountGroupID).
		ManagedBy(amv1.AccountGroupAssignmentManagedByRBAC).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create account group assignment description: %v\n", err)
		os.Exit(1)
	}

	// Send a request to create the account group assignment:
	response, err := collection.Add().
		Body(assignment).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create account group assignment: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	assignment = response.Body()
	fmt.Printf("%s - Account %s assigned to Group %s\n", assignment.ID(), assignment.AccountID(), assignment.AccountGroupID())
}
