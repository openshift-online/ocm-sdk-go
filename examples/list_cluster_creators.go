/*
Copyright (c) 2019 Red Hat, Inc.

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

// This example shows how to get the e-mail addresses of the creators of managed clusters.

package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger:
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(false).
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

	// Get the client for the collections that we will be using:
	clustersCollection := connection.ClustersMgmt().V1().Clusters()
	subscriptionsCollection := connection.AccountsMgmt().V1().Subscriptions()
	accountsCollection := connection.AccountsMgmt().V1().Accounts()

	// Create a tab writer that will be used to print the results in columns and print the
	// header line:
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(writer, "ID\tNAME\tUSER\tMAIL\n")

	// Retrieve the list of clusters using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	size := 10
	page := 1
	for {
		// Retrieve the page:
		listClustersResponse, err := clustersCollection.List().
			Search("managed = 't'").
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't retrieve clusters page %d: %s\n", page, err)
			os.Exit(1)
		}

		// Process the page:
		listClustersResponse.Items().Each(func(cluster *cmv1.Cluster) bool {
			// Get the cluster data:
			clusterID := cluster.ID()
			clusterName := cluster.Name()

			// If the cluster doesn't have a link to the subscription then ignore it,
			// otherwise get the identifier of the subscription from the link:
			subscriptionLink := cluster.Subscription()
			if subscriptionLink == nil {
				return true
			}
			subscriptionID := subscriptionLink.ID()

			// Retrieve the details of the subscription, so that we can follow the link
			// to the account that created it:
			subscriptionResource := subscriptionsCollection.Subscription(subscriptionID)
			subscriptionGetResponse, err := subscriptionResource.Get().Send()
			if err != nil {
				fmt.Fprintf(
					os.Stderr,
					"Can't retrieve details of subscription '%s' for cluster '%s': %v\n",
					subscriptionID, clusterID, err,
				)
				os.Exit(1)
			}
			subscription := subscriptionGetResponse.Body()

			// If the subscription doesn't have a link to the account that created it
			// then ignore it, otherwise get the identifier of the account from the
			// link:
			creatorLink := subscription.Creator()
			if creatorLink == nil {
				return true
			}
			creatorID := creatorLink.ID()

			// Retrieve the details of the account:
			accountResource := accountsCollection.Account(creatorID)
			accountGetResponse, err := accountResource.Get().Send()
			if err != nil {
				fmt.Fprintf(
					os.Stderr,
					"Can't retrieve details of creator account '%s' for "+
						"subscription '%s' and cluster '%s': %v\n",
					creatorID, subscriptionID, clusterID,
				)
				os.Exit(1)
			}
			account := accountGetResponse.Body()

			// Get the account data:
			creatorFirst := account.FirstName()
			creatorLast := account.LastName()
			creatorMail := account.Email()

			// Print the results:
			fmt.Fprintf(
				writer,
				"%s\t%s\t%s %s\t%s\n",
				clusterID,
				clusterName,
				creatorFirst,
				creatorLast,
				creatorMail,
			)

			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if listClustersResponse.Size() < size {
			break
		}
		page++
	}

	// Flush the tab writer, otherwise the results may not be displayed:
	writer.Flush()
}
