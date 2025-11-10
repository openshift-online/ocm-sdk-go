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

// This example shows how to create a log forwarder for a cluster.

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

	// Create a log forwarder configuration with CloudWatch destination:
	logForwarder, err := cmv1.NewLogForwarder().
		ID("example-forwarder").
		Applications("default", "kube-system").
		CloudWatch(
			cmv1.NewLogForwarderCloudWatchConfig().
				LogGroupName("/aws/rosa/example-cluster").
				LogDistributionRoleArn("arn:aws:iam::123456789012:role/log-distribution-role"),
		).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create log forwarder description: %v\n", err)
		os.Exit(1)
	}

	// Send a request to create the log forwarder:
	response, err := logForwardersResource.Add().
		Body(logForwarder).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create log forwarder: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	logForwarder = response.Body()
	fmt.Printf("Log forwarder created:\n")
	fmt.Printf("- ID: %s\n", logForwarder.ID())
	fmt.Printf("- Cluster ID: %s\n", logForwarder.ClusterID())
	if len(logForwarder.Applications()) > 0 {
		fmt.Printf("- Applications: %v\n", logForwarder.Applications())
	}
	if logForwarder.CloudWatch() != nil {
		fmt.Printf("- CloudWatch log group: %s\n", logForwarder.CloudWatch().LogGroupName())
	}
	if logForwarder.S3() != nil {
		fmt.Printf("- S3 bucket: %s\n", logForwarder.S3().BucketName())
	}
}
