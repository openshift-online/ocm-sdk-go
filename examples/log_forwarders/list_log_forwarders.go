/*
Copyright (c) 2020 Red Hat, Inc.

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

// This example shows how to list log forwarders for a cluster.

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

	// Get the client for the resource that manages the collection of log forwarders:
	clusterID := os.Getenv("CLUSTER_ID")
	if clusterID == "" {
		fmt.Fprintf(os.Stderr, "CLUSTER_ID environment variable is required\n")
		os.Exit(1)
	}

	logForwardersResource := connection.ClustersMgmt().V1().
		Clusters().
		Cluster(clusterID).
		ControlPlane().
		LogForwarders()

	// Send a request to list the log forwarders:
	response, err := logForwardersResource.List().
		Size(100).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve log forwarders: %v\n", err)
		os.Exit(1)
	}

	// Print the results:
	logForwarders := response.Items()
	fmt.Printf("Found %d log forwarders:\n", len(logForwarders.Slice()))

	logForwarders.Each(func(logForwarder *cmv1.LogForwarder) bool {
		fmt.Printf("- ID: %s\n", logForwarder.ID())
		fmt.Printf("  Cluster ID: %s\n", logForwarder.ClusterID())

		if len(logForwarder.Applications()) > 0 {
			fmt.Printf("  Applications: %v\n", logForwarder.Applications())
		}

		if logForwarder.CloudWatch() != nil {
			fmt.Printf("  CloudWatch:\n")
			fmt.Printf("    Log group: %s\n", logForwarder.CloudWatch().LogGroupName())
			if logForwarder.CloudWatch().LogDistributionRoleArn() != "" {
				fmt.Printf("    Role ARN: %s\n", logForwarder.CloudWatch().LogDistributionRoleArn())
			}
		}

		if logForwarder.S3() != nil {
			fmt.Printf("  S3:\n")
			fmt.Printf("    Bucket: %s\n", logForwarder.S3().BucketName())
			if logForwarder.S3().BucketPrefix() != "" {
				fmt.Printf("    Prefix: %s\n", logForwarder.S3().BucketPrefix())
			}
		}

		if logForwarder.Status() != nil {
			fmt.Printf("  Status: %s\n", logForwarder.Status().State())
			if logForwarder.Status().Message() != "" {
				fmt.Printf("  Message: %s\n", logForwarder.Status().Message())
			}
		}

		fmt.Println()
		return true
	})
}
