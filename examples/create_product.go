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

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	sb "github.com/openshift-online/ocm-sdk-go/v2/statusboard/v1"
)

func createProduct(ctx context.Context, args []string) error {
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

	// Prepare the description of the product to create:
	product, err := sb.NewProduct().
		Name("myproduct").
		Fullname("myproduct's fullname").
		Build()
	if err != nil {
		return err
	}

	// Send a request to create the product:
	response, err := collection.Add().
		Body(product).
		SendContext(ctx)
	if err != nil {
		return err
	}

	// Print the result:
	product = response.Body()
	fmt.Printf("%s - %s\n", product.ID(), product.Name())

	return nil
}
