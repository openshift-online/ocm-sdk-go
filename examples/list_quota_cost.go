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

// This example shows how to retrieve a list of quota cost

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

	token := os.Getenv("OCM_TOKEN")

	const url = "https://api.openshift.com" // use "https://api.stage.openshift.com" for stage

	// Create the connection, and remember to close it:
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

	// we need to get the organizationID before calling the quota cost endpoint
	accountResp, err := connection.AccountsMgmt().V1().CurrentAccount().Get().Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve current account: %s\n", err)
		os.Exit(1)
	}
	organizationId := accountResp.Body().Organization().ID()

	// Get the client for the resource that manages the collection of quota cost:
	collection := connection.AccountsMgmt().V1().Organizations().Organization(organizationId).QuotaCost()

	// List quota cost items and their related cloud accounts:
	response, err := collection.List().Parameter("fetchCloudAccounts", true).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve quota cost list for organization %s: %s\n", organizationId, err)
		os.Exit(1)
	}

	// Prints quota cost items that were found in the previous step:
	fmt.Printf("id,allowed,consumed,number_of_accounts\n")
	response.Items().Each(func(quota *amv1.QuotaCost) bool {
		fmt.Printf("%s,%d,%d,%d\n", quota.QuotaID(), quota.Allowed(), quota.Consumed(), len(quota.CloudAccounts()))
		return true
	})

}
