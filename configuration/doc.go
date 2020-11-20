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

// Package configuration provides a mechanism to load configuration from JSON or YAML files. The
// typical use will be to create a configuration object and then load one or more configuration
// sources:
//
//	// Load the configuration from a file:
//	cfg, err := configuration.New().
//		Load("myconfig.yaml").
//		Build()
//	if err != nil {
//		...
//	}
//
// Once the configuration is loaded it can be copied into an object containing the same tags
// used by the YAML library:
//
//	// Copy the configuration into our object:
//	type MyConfig struct {
//		MyKey string `yaml:"mykey"`
//		YouKey int `yaml:"yourkey"`
//	}
//	var myCfg MyConfig
//	err = cfg.Populate(&myCfg)
//	if err != nil {
//		...
//	}
//
// The advantage of using this configuration instead of using plain YAML is that configuration
// sources can use the the `!variable` and `!file` tags to reference environment variables
// or files. For example:
//
//	mykey: !variable MYVARIABLE
//	yourkey: !file /my/file.txt
//
// The following tags are supported:
//
//	!variable MYVARIABLE - Is replaced by the content of the environment variable `MYVARIABLE`.
//	!file /my/file.txt - Is replaced by the content of the file `/my/file.txt`.
//	!shell myscript - Is replaced by the result of executing the `myscript` shell script.
//
// Tag names can be abbreviated. For example these are all valid tags:
//
//	!var MYVARIABLE - Replaced by the content of the environment variablel `MYVARIABLE`.
//	!v MYVARIABLE - Replaced by the content of the environment variablel `MYVARIABLE`.
//	!f /my/file.txt - Replaced by the content of the `/my.file.txt` file.
//	!sh myscript - Replaced by the result of execution the `myscript` shell script.
//
// The `file` tag trims all leading and traling white space from the content of the file.
//
// By default the tags replace the node they are applied to with a string. This will not work for
// fields that are declared of other types in the configuration struct. In those cases it is
// possible to add a suffix to the tag to indicate the type of the replacmenet.  For example:
//
//  # A configuration with an integer loaded from an environment variable
//  # and a boolean loaded from a file:
//  myid: !variable/int MYID
//  myenabled: !file/bool /my/enabled.txt
//
// This can be used with the following Go code:
//
//	type MyConfig struct {
// 		MyId      int  `yaml:"myid"`
//		MyEnabled bool `yaml:"myenabled"`
//	}
//	var myCfg MyConfig
//	err = cfg.Populate(&myCfg)
//	if err != nil {
//		...
//	}
//
// When multiple sources are configured (calling the Load method multiple times) they will all
// be merged, and sources loaded later sources will override sources loaded earlier.
package configuration
