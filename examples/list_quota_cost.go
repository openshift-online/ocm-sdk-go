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

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	amv1 "github.com/openshift-online/ocm-sdk-go/v2/accountsmgmt/v1"
)

func listQuotaCost(ctx context.Context, args []string) error {
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

	organizationId := "" // update with your own organizationId

	// Get the client for the resource that manages the collection of quota cost:
	collection := connection.AccountsMgmt().V1().Organizations().Organization(organizationId).QuotaCost()

	// Search quota cost items where quota_id starts with 'add-on':
	response, err := collection.List().
		Search("quota_id like 'add-on%'").
		SendContext(ctx)
	if err != nil {
		return err
	}

	// Prints quota cost items that were found in the previous step:
	response.Items().Each(func(quota *amv1.QuotaCost) bool {
		fmt.Printf("%s - %d - %d\n", quota.QuotaID(), quota.Allowed(), quota.Consumed())
		return true
	})

	return nil
}
