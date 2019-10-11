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

// This example shows how to retrieve an access token that can then be used to
// install a cluster with the _OpenShift_ command line installer.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openshift-online/ocm-sdk-go"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := sdk.NewGoLoggerBuilder().
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
		Tokens(token).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the cient for the resource that manages access tokens:
	resource := connection.AccountsMgmt().V1().AccessToken()

	// Send the request to get the access token:
	response, err := resource.Post().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get access token: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	accessToken := response.Body()
	for registry, auth := range accessToken.Auths() {
		fmt.Printf("Mail for registry '%s' is '%s'\n", registry, auth.Email())
		fmt.Printf("Authorization for registry '%s' is '%s'\n", registry, auth.Auth())
	}
}
