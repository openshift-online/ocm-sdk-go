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

	regValue, err := sdk.GetRhRegion("https://api.stage.openshift.com", "ap-southeast-1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't find region: %v", err)
		os.Exit(1)
	}
	token := os.Getenv("OCM_TOKEN")

	// Build a regionalized connection based on the desired shard
	connection, err := sdk.NewConnectionBuilder().
		Client("ocm-cli", "").
		Logger(logger).
		Tokens(token).
		URL(fmt.Sprintf("https://%s", regValue.URL)). // Apply the region URL
		Insecure(false).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	collection := connection.AccountsMgmt().V1().CurrentAccount()

	// Even though we've provided a regional url, the SDK should redirect to the global API for AccountMgmt
	_, err = collection.Get().
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve current account : %s\n", err)
		os.Exit(1)
	}
}
