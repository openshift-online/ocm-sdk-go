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

// This example shows how to retrieve the collection of applications
// from the Status Board project.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
	sb "github.com/openshift-online/ocm-sdk-go/v2/statusboard/v1"
)

func listApplications(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		URL("http://localhost:8000").
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of applications:
	collection := connection.StatusBoard().V1().Applications()

	// Retrieve the list of applications using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	size := 10
	page := 1
	for {
		// Retrieve the page:
		response, err := collection.List().
			//Fullname("similitudinem-consentiat"). // Uncomment this to restrict results by fullname
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			return err
		}

		// Display the page:
		response.Items().Each(func(application *sb.Application) bool {
			fmt.Printf("%s - %s - %s\n", application.ID(), application.Name(), application.Owners())
			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if response.Size() < size {
			break
		}
		page++
	}

	// Get a list of services for a given application
	application_id := "92b631f1-b4a7-47a7-89e4-ea9b26fafebe" // Adjust as needed
	services_collection := connection.StatusBoard().V1().Applications().Application(application_id).Services()
	page = 1

	for {
		// Retrieve the page:
		response, err := services_collection.List().
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			return err
		}

		// Display the page:
		response.Items().Each(func(service *sb.Service) bool {
			fmt.Printf("%s - %s\n", service.ID(), service.Name())
			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if response.Size() < size {
			break
		}
		page++
	}

	return nil
}
