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

// This example shows how to list the Job Queues.

package main

import (
	"context"
	"fmt"
	sdk "github.com/openshift-online/ocm-sdk-go"
	jqv1 "github.com/openshift-online/ocm-sdk-go/jobqueue/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
	"os"
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
	queuesResponse, err := jobQueues.Queues().List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve Job Queues: %v\n", err)
		os.Exit(1)
	}

	queuesResponse.Items().Each(func(queue *jqv1.Queue) bool {
		var (
			id          string
			name        string
			maxAttempts int
			maxRunTime  int
			ok          bool
		)
		id = queue.ID()
		if name, ok = queue.GetName(); !ok {
			name = "No name"
		}
		if maxAttempts, ok = queue.GetMaxAttempts(); !ok {
			maxAttempts = 0
		}
		if maxRunTime, ok = queue.GetMaxRunTime(); !ok {
			maxRunTime = 0
		}

		fmt.Printf("id: %s\nname: %s\nmax_run_time: %d\nmax_attempts: %d\n", id, name, maxAttempts, maxRunTime)
		return true
	})
}
