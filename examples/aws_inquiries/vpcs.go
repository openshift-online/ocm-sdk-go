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

	cloudProviderData, err := cmv1.NewCloudProviderData().
		AWS(
			cmv1.NewAWS().
				STS(
					cmv1.NewSTS().RoleARN("arn:aws:iam::xxx:role/ManagedOpenShift-Installer-Role"),
				),
		).
		Region(
			cmv1.NewCloudRegion().
				ID("us-east-1"),
		).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build cloud provider data: %v\n", err)
		os.Exit(1)
	}
	// Get the client for the resource that manages the collection of clusters:
	response, err := connection.ClustersMgmt().V1().AWSInquiries().Vpcs().Search().Body(cloudProviderData).Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't send request for vpcs: %v\n", err)
		os.Exit(1)
	}
	response.Items().Each(func(item *cmv1.CloudVPC) bool {
		fmt.Println(item.Name())
		return true
	})
}
