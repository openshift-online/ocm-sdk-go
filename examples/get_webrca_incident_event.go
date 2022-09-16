/*
Copyright (c) 2022 Red Hat, Inc.

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

// This example shows how to retrieve an incident event from the web-rca project.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/renan-campos/ocm-sdk-go"
	"github.com/renan-campos/ocm-sdk-go/logging"
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
		URL("http://localhost:8000").
		Tokens(token).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Update as needed
	incident_id := "0d6fcdde-e9d8-4dc4-beb0-67ed975d71b7"
	event_id := "61ac5f1d-c425-48b2-b87a-dc5ef2dfe192"

	// Get the client for the resource that manages the collection of incidents:
	incidents_collection := connection.WebRCA().V1().Incidents()
	events_collection := incidents_collection.Incident(incident_id).Events()

	// Get the client for the resource that manages the incident that we are looking for. Note
	// that this will not send any request to the server yet, so it will succeed even if that
	// incident doesn't exist.
	resource := events_collection.Event(event_id)

	// Send the request to retrieve the incident:
	response, err := resource.Get().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve event: %v\n", err)
		os.Exit(1)
	}

	// Print the result:
	event := response.Body()
	fmt.Printf("%s - %s - %s\n", event.ID(), event.EventType(), event.Note())
}
