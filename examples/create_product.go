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

	sdk "github.com/renan-campos/ocm-sdk-go"
	"github.com/renan-campos/ocm-sdk-go/logging"
	sb "github.com/renan-campos/ocm-sdk-go/statusboard/v1"
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
		URL("http://localhost:8000").
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "Can't create product description: %v\n", err)
		os.Exit(1)
	}

	// Send a request to create the product:
	response, err := collection.Add().
		Body(product).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create product: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	product = response.Body()
	fmt.Printf("%s - %s\n", product.ID(), product.Name())
}
