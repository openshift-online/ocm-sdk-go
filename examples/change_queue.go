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

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/openshift-online/ocm-sdk-go/database"
	"github.com/openshift-online/ocm-sdk-go/logging"
)

// This example shows how to create a queue that will call a function for each change that is done
// in a database table. To use adjust the `URL` parameter of the constructor of the queue and create
// a `my_data` table, for example:
//
//	create table my_data (
//		id integer not null primary key,
//		name text not null
//	);
//
// Then run this example in one window and in another window do some operations with that table, for
// example insert a new record:
//
//	insert into my_data (id, name) values (123, 'my_name');
//
// If everything works correctly you will see something like this in the output:
//
//	Serial: 57
//	Timestamp: 2021-09-02T19:14:23+02:00
//	Source: my_data
//	Operation: insert
//	Old:
//	New: {
//	  "id": 124,
//	  "name": "my_name"
//	}
//
// Try some other operations on the table to see more results.

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

	// Create the queue:
	queue, err := database.NewChangeQueue().
		Logger(logger).
		URL("postgres://service:service123@localhost/service?sslmode=disable").
		Name("my_queue").
		Table("my_data").
		Install(true).
		Callback(print).
		Build(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create queue: %v\n", err)
		os.Exit(1)
	}

	// Wait a while and then close the queue:
	time.Sleep(5 * time.Minute)
	err = queue.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't close queue: %v\n", err)
		os.Exit(1)
	}
}

func print(ctx context.Context, item *database.ChangeQueueItem) {
	fmt.Printf("Serial: %d\n", item.Serial)
	fmt.Printf("Timestamp: %s\n", item.Timestamp.Format(time.RFC3339))
	fmt.Printf("Source: %s\n", item.Source)
	fmt.Printf("Operation: %s\n", item.Operation)
	fmt.Printf("Old: %s\n", render(item.Old))
	fmt.Printf("New: %s\n", render(item.New))
	fmt.Printf("\n")
}

func render(data []byte) string {
	var object map[string]interface{}
	err := json.Unmarshal(data, &object)
	if err != nil {
		return string(data)
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(object)
	if err != nil {
		return string(data)
	}
	return strings.TrimSpace(buffer.String())
}
