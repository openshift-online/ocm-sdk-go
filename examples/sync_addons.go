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

// This example shows how to update the collection of add-ons from an external source. To simplify
// things that external source is a YAML document embedded in this source file, but it could be an
// external file or an external collection of files.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ghodss/yaml"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := sdk.NewGoLoggerBuilder().
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
		URL("https://clusters-service.apps-crc.testing").
		Insecure(true).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the collection of add-ons:
	collection := connection.ClustersMgmt().V1().Addons()

	// Load the sets of add-ons from the YAML file and from the API:
	fileIndex, err := loadYAML(ctx, fileData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't load YAML data: %v\n", err)
		os.Exit(1)
	}
	apiIndex, err := loadAPI(ctx, collection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't load API data: %v\n", err)
		os.Exit(1)
	}

	// Add to the API the items that are in the file but not in the API:
	for id, fileItem := range fileIndex {
		_, ok := apiIndex[id]
		if !ok {
			apiItem, err := cmv1.NewAddOn().
				Copy(fileItem).
				Enabled(true).
				Build()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't build add-on: %v\n", err)
				os.Exit(1)
			}
			_, err = collection.Add().
				Body(apiItem).
				SendContext(ctx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't add add-on: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// Update in the API the items that are in the file and in the API:
	for id, fileItem := range fileIndex {
		_, ok := apiIndex[id]
		if ok {
			apiItem, err := cmv1.NewAddOn().
				Name(fileItem.Name()).
				Description(fileItem.Description()).
				Icon(fileItem.Icon()).
				Enabled(true).
				Build()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't build patch: %v\n", err)
				os.Exit(1)
			}
			_, err = collection.Addon(id).Update().
				Body(apiItem).
				SendContext(ctx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't update add-on: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// Disable in the API the items that are in the API but not in the file:
	for id := range apiIndex {
		_, ok := fileIndex[id]
		if !ok {
			apiItem, err := cmv1.NewAddOn().
				Enabled(false).
				Build()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't build patch: %v\n", err)
				os.Exit(1)
			}
			_, err = collection.Addon(id).Update().
				Body(apiItem).
				SendContext(ctx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't update add-on: %v\n", err)
				os.Exit(1)
			}
		}
	}
}

// loadYAML loads the add-ons from the YAML document and returns a map where the key is the
// identifier of the add-on and the value is the add-on object.
func loadYAML(ctx context.Context, data []byte) (result map[string]*cmv1.AddOn, err error) {
	// Load the list of add-ons from the API:
	data, err = yaml.YAMLToJSON(data)
	if err != nil {
		return
	}
	items, err := cmv1.UnmarshalAddOnList(data)
	if err != nil {
		return
	}

	// Populate the map:
	result = map[string]*cmv1.AddOn{}
	for _, item := range items {
		result[item.ID()] = item
	}

	return
}

// loadAPI loads the add-ons from the API and returns a map where the key is the identifier of the
// add-on and the value is the add-on object.
func loadAPI(ctx context.Context, collection *cmv1.AddOnsClient) (result map[string]*cmv1.AddOn,
	err error) {
	// Retrieve the list of add-ons using pages of ten items, till we get a page that has less
	// items than requests, as that marks the end of the collection:
	size := 10
	page := 1
	for {
		// Retrieve the page:
		var response *cmv1.AddOnsListResponse
		response, err = collection.List().
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			return
		}

		// Process the page:
		result = map[string]*cmv1.AddOn{}
		response.Items().Each(func(item *cmv1.AddOn) bool {
			result[item.ID()] = item
			return true
		})

		// Break the loop if the size of the page is less than requested, otherwise go to
		// the next page:
		if response.Size() < size {
			break
		}
		page++
	}

	return
}

var fileData = []byte(`
- id: black
  name: Black add-on
  description: Very black add-on
  icon: iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAAAAAA6fptVAAAACklEQVQI12NgAAAAAgAB4iG8MwAAAABJRU5ErkJggg==
- id: white
  name: White add-on
  description: Very white add-on
  icon: iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAAAAAA6fptVAAAACklEQVQI12NgAAAAAgAB4iG8MwAAAABJRU5ErkJggg==
`)
