/*
Copyright (c) 2021 Red Hat, Inc.

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

// This example shows how to PUSH into/POP from a Job Queue.

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

	// Get the client for the resource that manages the Job Queues:
	jobQueues := connection.JobQueue().V1()

	// Retrieve the first page of Job Queues and print them:
	queueID := "ocm-test-queue.fifo"
	pushResponse, err := jobQueues.Queues().Queue(queueID).Push().Arguments("foo bar").SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't push: %v\n", err)
		os.Exit(1)
	}
	pushID := pushResponse.ID()
	pushArguments := pushResponse.Arguments()
	fmt.Printf("Pushed:\nid: %s\narguments: %s\n\n", pushID, pushArguments)

	popResponse, err := jobQueues.Queues().Queue(queueID).Pop().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't pop: %v\n", err)
		os.Exit(1)
	}
	popID := popResponse.Body().ID()
	popArguments := popResponse.Body().Arguments()
	popAttempts := popResponse.Body().Attempts()
	abandonedAt := popResponse.Body().AbandonedAt()
	receiptID := popResponse.Body().ReceiptId()
	fmt.Printf("Popped:\nid: %s\narguments: %s\nattempts: %d\nabandoned_at: %s\nreceipt_id: %s\n",
		popID, popArguments, popAttempts, abandonedAt, receiptID)
}
