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

// This example shows how to load a configuration and then dump the resulting YAML.

package main

import (
	"context"
	"os"

	"github.com/openshift-online/ocm-sdk-go/v2/configuration"
)

func dumpConfig(ctx context.Context, args []string) error {
	// Create the configuration and load all the files given in the command line:
	builder := configuration.New()
	if len(args) > 0 {
		for _, arg := range args {
			builder.Load(arg)
		}
	} else {
		builder.Load(os.Stdin)
	}
	object, err := builder.Build()
	if err != nil {
		return err
	}

	// Dump the resulting YAML file:
	effective, err := object.Effective()
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(effective)
	if err != nil {
		return err
	}

	return nil
}
