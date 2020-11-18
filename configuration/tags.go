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

// This file contains the code that is used to process tags like `!variable` and `!file` inside
// configuration files.

package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// tagRegistryEntry stores the description of one tag. When a tag is detected in a node the
// corresponding function will be called passing the node, so that the function can modify
// it as needed.
type tagRegistryEntry struct {
	name    string
	process func(name, replacement string, node *yaml.Node) error
}

// tagRegistry stores the collection of tags.
var tagRegistry []tagRegistryEntry

// registerTag adds a tag to the registry.
func registerTag(name string, process func(string, string, *yaml.Node) error) {
	tagRegistry = append(tagRegistry, tagRegistryEntry{
		name:    name,
		process: process,
	})
}

func init() {
	// Register tags:
	registerTag("variable", processVariableTag)
	registerTag("file", processFileTag)
}

// processTags process the tags present in the given YAML tree and returns the result. The following
// tags are supported:
//
//	!variable MYVARIABLE - Is replaced by the content of the environment variable `MYVARIABLE`.
//	!file file /my/file.txt - Is replaced by the content of the file `/my/file.txt`.
//
// The function will return an error if a tag references an environment variable or file that doesn't
// exist.
func processTags(node *yaml.Node) error {
	name, replacement := parseTag(node.Tag)
	if name != "" {
		for _, entry := range tagRegistry {
			if strings.HasPrefix(entry.name, name) {
				err := entry.process(name, replacement, node)
				if err != nil {
					return err
				}
			}
		}
	}
	for _, child := range node.Content {
		err := processTags(child)
		if err != nil {
			return err
		}
	}
	return nil
}

// processVariableTag is the implementation fo the `!variable` tag: replaces an environment variable
// reference with its value.
func processVariableTag(name, replacement string, node *yaml.Node) error {
	variable := strings.TrimSpace(node.Value)
	result, ok := os.LookupEnv(variable)
	if !ok {
		return fmt.Errorf("can't find environment variable '%s'", name)
	}
	node.SetString(result)
	node.Tag = "!!" + replacement
	return nil
}

// processFileTag is the implementation of the `!file` tag: replaces a file reference with the
// content of the file.
func processFileTag(name, replacement string, node *yaml.Node) error {
	file := strings.TrimSpace(node.Value)
	data, err := ioutil.ReadFile(file) // #nosec G304
	if err != nil {
		return err
	}
	result := strings.TrimSpace(string(data))
	node.SetString(result)
	node.Tag = "!!" + replacement
	return nil
}

// parseTag extract from the given tag the name and the replacement. For example, the tag `!file`
// has name `file` and no replacement and the tag `!file/int` has name `file` and
// replacement `int`.
func parseTag(tag string) (name, replacement string) {
	if !strings.HasPrefix(tag, "!") {
		return
	}
	slash := strings.Index(tag, "/")
	if slash == -1 {
		name = tag[1:]
		replacement = ""
	} else {
		name = tag[1:slash]
		replacement = tag[slash+1:]
	}
	return
}
