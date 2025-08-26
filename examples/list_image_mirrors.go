/*
Copyright (c) 2025 Red Hat, Inc.

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
	ctx := context.Background()
	clusterId := "test-cluster-id"

	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

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

	collection := connection.ClustersMgmt().V1().Clusters().Cluster(clusterId).ImageMirrors()

	response, err := collection.List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve image mirrors: %v\n", err)
		os.Exit(1)
	}

	imageMirrors := response.Items()
	fmt.Printf("Found %d image mirrors:\n", imageMirrors.Len())

	imageMirrors.Each(func(imageMirror *cmv1.ImageMirror) bool {
		fmt.Printf("  ID: %s\n", imageMirror.ID())
		if imageMirror.Source() != "" {
			fmt.Printf("  Source: %s\n", imageMirror.Source())
		}
		if mirrors := imageMirror.Mirrors(); len(mirrors) > 0 {
			fmt.Printf("  Mirrors: %v\n", mirrors)
		}
		fmt.Println()
		return true
	})
}