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

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
)

func getService(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		URL("http://localhost:8000").
		Tokens(token).
		Build()
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of services:
	collection := connection.StatusBoard().V1().Services()

	// Get the client for the resource that manages the service that we are looking for. Note
	// that this will not send any request to the server yet, so it will succeed even if that
	// service doesn't exist.
	resource := collection.Service("dc440fc4-db27-40eb-8180-765bc4e28620") // Update as needed

	// Send the request to retrieve the service:
	response, err := resource.Get().Send(ctx)
	if err != nil {
		return err
	}

	// Print the result:
	service := response.Body()
	fmt.Printf("%s - %s - %s\n", service.ID(), service.Name(), service.Owners())

	return nil
}
