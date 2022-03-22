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

// This example shows how to use the export control review resource.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	azv1 "github.com/openshift-online/ocm-sdk-go/v2/authorizations/v1"
)

func exportControlReview(ctx context.Context, args []string) error {
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

	// Get the client for the resource that manages export control review:
	resource := connection.Authorizations().V1().ExportControlReview()

	// Build the request:
	reviewRequest, err := azv1.NewExportControlReviewRequest().
		AccountUsername("alice").
		Build()
	if err != nil {
		return err
	}

	// Send the request:
	postResponse, err := resource.Post().
		Request(reviewRequest).
		SendContext(ctx)
	if err != nil {
		return err
	}

	// Print the results:
	reviewResponse := postResponse.Response()
	fmt.Printf("Restricted: %v\n", reviewResponse.Restricted())

	return nil
}
