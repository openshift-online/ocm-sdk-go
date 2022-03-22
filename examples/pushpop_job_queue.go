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

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	"github.com/openshift-online/ocm-sdk-go/v2/logging"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer func(connection *sdk.Connection) {
		_ = connection.Close()
	}(connection)

	// Get the client for the resource that manages the Job Queues:
	jobQueues := connection.JobQueue().V1()

	// Client for the queue
	queueID := "job-queue-service-heartbeat-staging.fifo"
	client := jobQueues.Queues().Queue(queueID)

	// Push a new job
	pushResponse, err := client.Push().Arguments("foo bar").SendContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Can't push: %v\n", err)
		os.Exit(1)
	}
	pushID := pushResponse.ID()
	pushArguments := pushResponse.Arguments()
	fmt.Printf("Pushed:\n\tid: %s\n\targuments: %s\n\n", pushID, pushArguments)

	// Retrieve this job back
	popResponse, err := client.Pop().SendContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Can't pop: %v\n", err)
		os.Exit(1)
	}
	popID := popResponse.ID()
	popAttempts := popResponse.Attempts()
	abandonedAt := popResponse.AbandonedAt()
	receiptID := popResponse.ReceiptId()
	popArguments := popResponse.Arguments()
	fmt.Printf("Popped:\n\tid: %s\n\targuments: %s\n\tattempts: %d\n\tabandoned_at: %s\n\treceipt_id: %s\n",
		popID, popArguments, popAttempts, abandonedAt, receiptID)

	// Mark it as success
	_, err = client.Jobs().Job(popID).Success().ReceiptId(receiptID).SendContext(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: Can't success: %v\n", err)
		os.Exit(1)
	}

	// To mark it as Failure use
	// _, err = client.Jobs().Job(popID).Failure().FailureReason("Failure reason").ReceiptId(receiptID).SendContext(ctx)
}
