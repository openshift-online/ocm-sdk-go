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

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Prepare the description of the cluster to create:
	cluster, err := cmv1.NewCluster().
		Name("classic").
		Flavour(
			cmv1.NewFlavour().
				ID("osd-4"),
		).
		Region(
			cmv1.NewCloudRegion().
				ID("us-east-1"),
		).
		CloudProvider(cmv1.NewCloudProvider().ID("aws")).
		CCS(cmv1.NewCCS().Enabled(true)).
		Version(
			cmv1.NewVersion().
				ID("openshift-v4.13.1"),
		).
		AWS(
			cmv1.NewAWS().AccountID("xxxxxx").
				STS(cmv1.NewSTS().
					RoleARN("arn:aws:iam::xxxxxx:role/ManagedOpenShift-Installer-Role").
					SupportRoleARN("arn:aws:iam::xxxxxx:role/ManagedOpenShift-Support-Role").
					InstanceIAMRoles(cmv1.NewInstanceIAMRoles().MasterRoleARN("arn:aws:iam::xxxxxx:role/ManagedOpenShift-ControlPlane-Role").
						WorkerRoleARN("arn:aws:iam::xxxxxx:role/ManagedOpenShift-Worker-Role")).
					OperatorRolePrefix("classic-abc")),
		).
		Properties(map[string]string{
			"rosa_creator_arn": "arn:aws:iam::xxxxxx:user/user",
		}).
		Product(cmv1.NewProduct().ID("rosa")).
		Ingresses(cmv1.NewIngressList().Items(cmv1.NewIngress().
			Default(true).
			ExcludedNamespaces("stage", "int", "dev").RouteSelectors(
			map[string]string{
				"foo": "bar",
			}).
			RouteWildcardPolicy(cmv1.WildcardPolicyWildcardsAllowed)),
		).Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create cluster description: %v\n", err)
		os.Exit(1)
	}

	// Send a request to create the cluster:
	response, err := collection.Add().
		Body(cluster).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create cluster: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	cluster = response.Body()
	fmt.Printf("%s - %s\n", cluster.ID(), cluster.Name())
}
