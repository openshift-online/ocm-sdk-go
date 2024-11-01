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

	// Prepare the description of the cluster to create:
	cluster, err := cmv1.NewCluster().
		Name("hcp").
		Flavour(
			cmv1.NewFlavour().
				ID("osd-4"),
		).
		Region(
			cmv1.NewCloudRegion().
				ID("xxx"),
		).
		DNS(cmv1.NewDNS().
			BaseDomain("xxx")).
		CloudProvider(cmv1.
			NewCloudProvider().ID("aws")).
		MultiAZ(true).
		Version(
			cmv1.NewVersion().ID("xxx")).
		AWS(
			cmv1.NewAWS().
				AccountID("xxx").
				BillingAccountID("xxx").
				AdditionalAllowedPrincipals(
					"arn:aws:iam::xxx:role/ca-route53",
					"arn:aws:iam::xxx:role/ca-vpc-endpoint-service").
				PrivateHostedZoneID("xxx").
				PrivateHostedZoneRoleARN("arn:aws:iam::xxx:role/ca-route53").
				HcpInternalCommunicationHostedZoneId("xxx").
				VpcEndpointRoleArn("arn:aws:iam::xxx:role/ca-vpc-endpoint-service").
				SubnetIDs(
					"subnet-xxx",
					"subnet-xxx").
				STS(cmv1.NewSTS().
					OidcConfig(cmv1.NewOidcConfig().ID("xxx")).
					RoleARN("arn:aws:iam::xxx:role/ManagedOpenShift-HCP-ROSA-Installer-Role").
					SupportRoleARN("arn:aws:iam::xxx:role/ManagedOpenShift-HCP-ROSA-Support-Role").
					InstanceIAMRoles(cmv1.NewInstanceIAMRoles().
						WorkerRoleARN("arn:aws:iam::xxx:role/ManagedOpenShift-HCP-ROSA-Worker-Role")).
					OperatorRolePrefix("hcp")),
		).
		Nodes(cmv1.NewClusterNodes().
			AvailabilityZones("xxx").
			Compute(2).
			ComputeMachineType(cmv1.NewMachineType().ID("m5.xlarge"))).
		Properties(map[string]string{
			"rosa_creator_arn": "arn:aws:iam::xxx:user/user",
		}).
		Product(cmv1.NewProduct().ID("rosa")).
		Hypershift(cmv1.NewHypershift().Enabled(true)).
		Build()
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
