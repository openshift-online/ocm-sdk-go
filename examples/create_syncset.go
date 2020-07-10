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

// This example shows how to create a k8s configmap on the cluster using a syncset.

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		log.Fatalf("Can't build logger: %v", err)
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
	clustersResource := connection.ClustersMgmt().V1().Clusters()

	// create an array of resources to add to the syncset (this example only has 1 item)
	resources := []interface{}{
		&corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "ConfigMap",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
			Data: map[string]string{
				"foo": "bar",
			},
		},
	}

	// Build the syncset - "ext-" prefix is required
	syncsetBuilder := cmv1.NewSyncset()
	syncsetBuilder = syncsetBuilder.ID("ext-foo2").Resources(resources...)
	syncset, err := syncsetBuilder.Build()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Create the syncset on the cluster
	response, err := clustersResource.Cluster("1Jam7Ejgpm7AbZshbgaA9TsM1SQ").
		ExternalConfiguration().
		Syncsets().
		Add().
		Body(syncset).
		Send()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Print the result:
	syncset = response.Body()
	fmt.Printf("%s - %s\n", syncset.ID(), syncset.Resources())

}
