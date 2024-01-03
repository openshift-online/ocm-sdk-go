/*
Copyright (c) 2024 Red Hat, Inc.

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

// This example shows how create a connection using the OpenID client credentials grant for
// authentication instead of the resource owner password grant.

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

	token := os.Getenv("OCM_TOKEN")

	// Create the connection, and remember to close it:
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	cloudProviderData, err := cmv1.NewCloudProviderData().Region(cmv1.NewCloudRegion().ID("us-east-1")).AWS(
		cmv1.NewAWS().
			AccountID("123456789010").
			AccessKeyID("xxx").
			SecretAccessKey("xxx")).Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build AWS Request object: %v\n", err)
		os.Exit(1)
	}

	// Get the client for the service that manages creds validation
	collection := connection.ClustersMgmt().V1().AWSInquiries().ValidateCredentials()

	response, err := collection.Post().Body(cloudProviderData).Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to validate AWS creds: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(response.Status())
}
