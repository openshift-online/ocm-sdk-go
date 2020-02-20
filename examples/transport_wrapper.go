/*
Copyright (c) 2020 Red Hat, Inc.

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

// This example shows how to use a transport wrapper for the connection.

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

type LoggingTransport struct {
	logger  sdk.Logger
	wrapped http.RoundTripper
}

// NewLoggingTransport creates a transport that sends basic details of requests to the
// given logger. The wrapped transport will be used actually send the requests.
func NewLoggingTransport(logger sdk.Logger, wrapped http.RoundTripper) http.RoundTripper {
	return &LoggingTransport{
		logger:  logger,
		wrapped: wrapped,
	}
}

func (t *LoggingTransport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	t.logger.Info(request.Context(), "Sending request %s '%s'", request.Method, request.URL.String())
	response, err = t.wrapped.RoundTrip(request)
	if err != nil {
		t.logger.Error(response.Request.Context(), "Got error sending request")
	} else {
		t.logger.Info(response.Request.Context(), "Got response status code %d", response.StatusCode)
	}
	return response, err
}

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger:
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(false).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		TransportWrapper(func(wrapped http.RoundTripper) http.RoundTripper {
			return NewLoggingTransport(logger, wrapped)
		}).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of cloud providers:
	providersCollection := connection.ClustersMgmt().V1().CloudProviders()

	// Retrieve the first page of cloud providers:
	_, err = providersCollection.List().SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve providers: %v\n", err)
		os.Exit(1)
	}
}
