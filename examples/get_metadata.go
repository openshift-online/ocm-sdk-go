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

// This example shows how to retrieve the metadata of a service version.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
)

func getMetadata(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		Tokens(token).
		Build()
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the resource that manages the metadata:
	client := connection.ClustersMgmt().V1()

	// Send the request to retrieve the metadata:
	response, err := client.Get().SendContext(ctx)
	if err != nil {
		return err
	}
	metadata := response.Body()

	// Print the details:
	fmt.Printf("Server version: %s\n", metadata.ServerVersion())

	return nil
}
