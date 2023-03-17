/*
Copyright (c) 2023 Red Hat, Inc.

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

// This example shows how to list the cloud providers and their regions.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
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
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	client := connection.ClustersMgmt().V1().OidcConfigs()

	//Registers a Managed oidc configuration
	oidcConfig, err := cmv1.NewOidcConfig().Managed(true).Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create host oidc config description: %v\n", err)
		os.Exit(1)
	}

	addResponse, err := client.Add().Body(oidcConfig).SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't add hosted oidc config: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	oidcConfig = addResponse.Body()
	fmt.Printf("%s\n", oidcConfig.ID())
	fmt.Printf("%s\n", oidcConfig.IssuerUrl())
	fmt.Printf("%s\n", oidcConfig.SecretArn())
	fmt.Printf("%v\n", oidcConfig.Managed())
	fmt.Printf("%v\n", oidcConfig.Reusable())

	listResponse, err := client.List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't list hosted oidc configs: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d", listResponse.Total())

	getResponse, err := client.OidcConfig(listResponse.Items().Get(0).ID()).Get().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get hosted oidc configs: %v\n", err)
		os.Exit(1)
	}
	oidcConfig = getResponse.Body()
	fmt.Printf("%s\n", oidcConfig.ID())
	fmt.Printf("%s\n", oidcConfig.IssuerUrl())
	fmt.Printf("%s\n", oidcConfig.SecretArn())

	_, err = client.OidcConfig(oidcConfig.ID()).Delete().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't delete hosted oidc configs: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
