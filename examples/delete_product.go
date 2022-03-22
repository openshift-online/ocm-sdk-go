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

package main

// This example shows how to delete a product.

import (
	"context"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
)

func deleteProduct(ctx context.Context, args []string) error {
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

	// Get the client for the resource that manages the collection of products:
	collection := connection.StatusBoard().V1().Products()

	// Get the client for the resource that manages the product that we want to delete. Note
	// that this will not send any request to the server yet, so it will succeed even if the
	// product doesn't exist.
	resource := collection.Product("ea7ee64f-978d-4705-a271-85b072bc5241")

	// Send the request to delete the product:
	_, err = resource.Delete().SendContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
