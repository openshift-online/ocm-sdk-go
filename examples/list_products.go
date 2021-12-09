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

// This example shows how to retrieve the collection of products
// from the Status Board project.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/logging"
	sb "github.com/openshift-online/ocm-sdk-go/statusboard/v1"
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

	// Retrieve the list of products using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	size := 10
	page := 1
	for {
		// Retrieve the page:
		response, err := collection.List().
			//Fullname("exhibentur").
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't retrieve page %d: %s\n", page, err)
			os.Exit(1)
		}

		// Display the page:
		response.Items().Each(func(product *sb.Product) bool {
			fmt.Printf("%s - %s - %s\n", product.ID(), product.Name(), product.Owners())
			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if response.Size() < size {
			break
		}
		page++
	}
}
