/*
Copyright (c) 2018 Red Hat, Inc.

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

// This example shows how two create a connection using an access token and a refresh token that
// have been previously created by some other means.

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/openshift-online/uhc-sdk-go/pkg/client"
)

func main() {
	// Create a logger that has the debug level enabled:
	logger, err := client.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		log.Fatalf("Can't build logger: %v", err)
	}

	// Get the tokens, maybe receiving it from some other part of the application, or reading
	// them from a file where they have been previously saved:
	accessToken, err := ioutil.ReadFile("access.token")
	if err != nil {
		log.Fatalf("Can't read access token: %v", err)
	}
	refreshToken, err := ioutil.ReadFile("refresh.token")
	if err != nil {
		log.Fatalf("Can't read refresh token: %v", err)
	}

	// Create the connection, and remember to close it. Note that this connection will stop
	// working when both tokens expire. This can happen, for example, if the connection isn't
	// used for period of time longer than the life of the refresh token.
	connection, err := client.NewConnectionBuilder().
		Logger(logger).
		Tokens(string(accessToken), string(refreshToken)).
		Build()
	if err != nil {
		log.Fatalf("Can't build client: %v", err)
	}
	defer connection.Close()

	// Retrieve the collection of clusters:
	response, err := connection.Get().
		Path("/api/clusters_mgmt/v1/clusters").
		Parameter("search", "name like 'my%'").
		Parameter("page", 1).
		Parameter("size", 10).
		Send()
	if err != nil {
		log.Fatalf("Can't retrieve clusters: %s", err)
	}

	// Print the result:
	fmt.Printf("%d\n", response.Status())
	fmt.Printf("%s\n", response.String())
}
