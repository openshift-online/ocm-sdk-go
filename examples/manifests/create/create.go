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

// This example shows how to update the display name of a cluster.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
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
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Client(clientId, clientSecret).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the cluster that we want to update. Note
	// that this will not yet send any request to the server, so it will succeed even if the
	// cluster doesn't exist.
	resource := collection.Cluster("<cluster_id>")

	manifest, err := v1.NewManifest().
		ID("example").
		Spec(
			map[string]interface{}{
				"deleteOption": map[string]string{
					"propagationPolicy": "Foreground",
				},
				"workload": map[string]interface{}{
					"manifests": []interface{}{
						map[string]interface{}{
							"apiVersion": "apps/v1",
							"kind":       "Deployment",
							"metadata": map[string]string{
								"name":      "hello",
								"namespace": "default",
							},
							"spec": map[string]interface{}{
								"selector": map[string]interface{}{
									"matchLabels": map[string]string{
										"app": "hello",
									},
								},
								"template": map[string]interface{}{
									"metadata": map[string]interface{}{
										"labels": map[string]string{
											"app": "hello",
										},
									},
									"spec": map[string]interface{}{
										"containers": []interface{}{
											map[string]interface{}{
												"name":  "hello",
												"image": "quay.io/asmacdo/busybox",
												"command": []string{
													"/bin/sh",
													"-c",
													"echo \"Hello, Kubernetes!\" && sleep 300",
												},
											},
										},
									},
								},
							},
						},
						map[string]interface{}{
							"apiVersion": "v1",
							"kind":       "Service",
							"metadata": map[string]interface{}{
								"labels": map[string]string{
									"app": "hello",
								},
								"name":      "hello",
								"namespace": "default",
							},
							"spec": map[string]interface{}{
								"selector": map[string]string{
									"app": "hello",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"port":       8000,
										"protocol":   "TCP",
										"targetPort": 8000,
									},
								},
							},
						},
					},
				},
			},
		).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Issue building manifest body: %v\n", err)
		os.Exit(1)
	}
	_, err = resource.ExternalConfiguration().Manifests().Add().Body(manifest).SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create manifest: %v\n", err)
		os.Exit(1)
	}
}
