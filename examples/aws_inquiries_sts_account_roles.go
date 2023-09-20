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

	// Get the client for the service that manages inquiries to AWS STS Account Roles
	collection := connection.ClustersMgmt().V1().AWSInquiries().STSAccountRoles()

	// Specify the AWS Account Id to search for STS Account Roles
	awsAccountId := os.Getenv("AWS_ACCOUNT_ID")

	if awsAccountId == "" {
		fmt.Fprintln(os.Stderr, "Please set the AWS Account Id to search via the 'AWS_ACCOUNT_ID' environment property.")
		os.Exit(1)
	}

	// Retrieve the collection of STS Account Roles:
	builder := &cmv1.AWSBuilder{}
	aws, err := builder.AccountID(awsAccountId).Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build AWS Request object: %v\n", err)
		os.Exit(1)
	}

	response, err := collection.Search().Body(aws).Send()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve STS Account Roles from AWS: %v\n", err)
		os.Exit(1)
	}

	// Print the AWS Account Id
	fmt.Printf("Listing STS Account Roles for AWS Account '%s':\n", response.AwsAccountId())
	// Print the result:
	response.Items().Each(func(role *cmv1.AWSSTSAccountRole) bool {
		fmt.Printf("Account Role with prefix '%s' has the following roles:\n", role.Prefix())

		for _, r := range role.Items() {
			fmt.Printf(" - Arn: '%s' Type: '%s', Version: '%s', Admin: '%t', HCP Managed Policies: '%t', Managed Policies: '%t'\n", r.RoleARN(), r.RoleType(), r.RoleVersion(), r.IsAdmin(), r.HcpManagedPolicies(), r.ManagedPolicies())
		}

		return true
	})
}
