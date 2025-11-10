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

// This example shows how to get details of a specific log forwarder.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
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

	// Get the required parameters:
	clusterID := os.Getenv("CLUSTER_ID")
	if clusterID == "" {
		fmt.Fprintf(os.Stderr, "CLUSTER_ID environment variable is required\n")
		os.Exit(1)
	}

	logForwarderID := os.Getenv("LOG_FORWARDER_ID")
	if logForwarderID == "" {
		fmt.Fprintf(os.Stderr, "LOG_FORWARDER_ID environment variable is required\n")
		os.Exit(1)
	}

	// Get the client for the specific log forwarder:
	logForwarderResource := connection.ClustersMgmt().V1().
		Clusters().
		Cluster(clusterID).
		ControlPlane().
		LogForwarders().
		LogForwarder(logForwarderID)

	// Send a request to get the log forwarder details:
	response, err := logForwarderResource.Get().
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve log forwarder: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	logForwarder := response.Body()
	fmt.Printf("Log forwarder details:\n")
	fmt.Printf("- ID: %s\n", logForwarder.ID())
	fmt.Printf("- Cluster ID: %s\n", logForwarder.ClusterID())

	if len(logForwarder.Applications()) > 0 {
		fmt.Printf("- Applications: %v\n", logForwarder.Applications())
	}

	if logForwarder.CloudWatch() != nil {
		fmt.Printf("- CloudWatch Configuration:\n")
		fmt.Printf("  Log group: %s\n", logForwarder.CloudWatch().LogGroupName())
		if logForwarder.CloudWatch().LogDistributionRoleArn() != "" {
			fmt.Printf("  Role ARN: %s\n", logForwarder.CloudWatch().LogDistributionRoleArn())
		}
	}

	if logForwarder.S3() != nil {
		fmt.Printf("- S3 Configuration:\n")
		fmt.Printf("  Bucket: %s\n", logForwarder.S3().BucketName())
		if logForwarder.S3().BucketPrefix() != "" {
			fmt.Printf("  Prefix: %s\n", logForwarder.S3().BucketPrefix())
		}
	}

	if len(logForwarder.Groups()) > 0 {
		fmt.Printf("- Groups (%d):\n", len(logForwarder.Groups()))
		for _, group := range logForwarder.Groups() {
			fmt.Printf("  * ID: %s, Version: %s\n", group.ID(), group.Version())
		}
	}

	if logForwarder.Status() != nil {
		fmt.Printf("- Status:\n")
		fmt.Printf("  State: %s\n", logForwarder.Status().State())
		if logForwarder.Status().Message() != "" {
			fmt.Printf("  Message: %s\n", logForwarder.Status().Message())
		}
		if len(logForwarder.Status().ResolvedApplications()) > 0 {
			fmt.Printf("  Resolved applications: %v\n", logForwarder.Status().ResolvedApplications())
		}
	}
}
