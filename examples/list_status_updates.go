/*
Copyright (c) 2021 Red Hat, Inc.

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

// This example shows how to retrieve the collection of status updates
// from the Status Board project.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	sb "github.com/openshift-online/ocm-sdk-go/v2/statusboard/v1"
)

func listStatusUpdates(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		URL("http://localhost:8000").
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		return err
	}

	defer connection.Close()

	// Get the client for the resource that manages the collection of statuses:
	collection := connection.StatusBoard().V1().StatusUpdates()

	// Update as needed.
	product_ids := "4aff22fc-b5b9-4863-adfd-92dc92974cd5,33496a9f-f2bc-408a-b503-be726ab04976"

	response, err := collection.List().ProductIds(product_ids).Send(ctx)

	if err != nil {
		return err
	}

	// Display the page:
	response.Items().Each(func(status *sb.Status) bool {
		fmt.Printf("%s - %s - %s\n", status.ID(), status.Status(), status.Metadata())
		return true
	})

	fmt.Println("Size:", response.Size())

	return nil
}
