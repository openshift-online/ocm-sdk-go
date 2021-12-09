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

// This example shows how to retrieve a service from the Status Board project.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift-online/ocm-sdk-go/logging"
	//sb "github.com/openshift-online/ocm-sdk-go/statusboard/v1"
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
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of services:
	collection := connection.StatusBoard().V1().Services()

	// Get the client for the resource that manages the service that we are looking for. Note
	// that this will not send any request to the server yet, so it will succeed even if that
	// service doesn't exist.
	resource := collection.Service("dc440fc4-db27-40eb-8180-765bc4e28620") // Update as needed

	// Send the request to retrieve the service:
	response, err := resource.Get().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve service: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	service := response.Body()
	fmt.Printf("%s - %s - %s\n", service.ID(), service.Name(), service.Owners())
}
