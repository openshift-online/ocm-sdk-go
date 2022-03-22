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
)

func pushpopJobQueue(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		return err
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
	pushResponse, err := client.Push().Arguments("foo bar").Send(ctx)
	if err != nil {
		return err
	}
	pushID := pushResponse.ID()
	pushArguments := pushResponse.Arguments()
	fmt.Printf("Pushed:\n\tid: %s\n\targuments: %s\n\n", pushID, pushArguments)

	// Retrieve this job back
	popResponse, err := client.Pop().Send(ctx)
	if err != nil {
		return err
	}
	popID := popResponse.ID()
	popAttempts := popResponse.Attempts()
	abandonedAt := popResponse.AbandonedAt()
	receiptID := popResponse.ReceiptId()
	popArguments := popResponse.Arguments()
	fmt.Printf("Popped:\n\tid: %s\n\targuments: %s\n\tattempts: %d\n\tabandoned_at: %s\n\treceipt_id: %s\n",
		popID, popArguments, popAttempts, abandonedAt, receiptID)

	// Mark it as success
	_, err = client.Jobs().Job(popID).Success().ReceiptId(receiptID).Send(ctx)
	if err != nil {
		return err
	}

	// To mark it as Failure use
	// _, err = client.Jobs().Job(popID).Failure().FailureReason("Failure reason").ReceiptId(receiptID).Send(ctx)

	return nil
}
