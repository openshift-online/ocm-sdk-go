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

// This example shows how to get the quota summary.

package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	sdk "github.com/openshift-online/ocm-sdk-go"
	amv1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
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

	// Get the client for the resource that returns the description of the current account:
	accountResource := connection.AccountsMgmt().V1().CurrentAccount()

	// Send the request to get the details of the current account, as it contains the
	// organization identifier that we will need in order to find the quota summary:
	accountResponse, err := accountResource.Get().Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get current account: %v\n", err)
		os.Exit(1)
	}

	// Extract the organization identifier from the account details:
	orgID := accountResponse.Body().Organization().ID()
	fmt.Printf("Organization ID is '%s'\n", orgID)

	// Get the client for the resource that manages the quota summary:
	summaryResource := connection.AccountsMgmt().V1().
		Organizations().
		Organization(orgID).
		QuotaSummary()

	// Send the request to retrieve the quota summary:
	summaryResponse, err := summaryResource.List().Send()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Can't get quota summary: %v\n", err)
		os.Exit(1)
	}

	// Create a tab writer that will be used to print the results in columns and print the
	// header line:
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(writer, "TYPE\tNAME\tALLOWED\tRESERVED\n")

	// Print the results. Note that this will print only the first page of results. To print all
	// the results this will need to be repeated using the paging mechanism until the returned
	// page has less items than requested. See the `list_clusters.go` example for more details.
	summaryResponse.Items().Each(func(item *amv1.QuotaSummary) bool {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%d\t%d\n",
			item.ResourceType(),
			item.ResourceName(),
			item.Allowed(),
			item.Reserved(),
		)
		return true
	})

	// Flush the tab writer, otherwise the results may not be displayed:
	writer.Flush()
}
