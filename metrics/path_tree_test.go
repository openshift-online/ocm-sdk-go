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

// This file contains tests for the URL path tree.

package metrics

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2/dsl/table" // nolint
	. "github.com/onsi/gomega"              // nolint
)

var _ = DescribeTable(
	"Add",
	func(original string, paths []string, expected string) {
		var tree *pathTree
		err := json.Unmarshal([]byte(original), &tree)
		Expect(err).ToNot(HaveOccurred())
		for _, path := range paths {
			tree.add(path)
		}
		actual, err := json.Marshal(tree)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(MatchJSON(expected))
	},
	Entry(
		"Empty path",
		`{}`,
		[]string{
			``,
		},
		`{}`,
	),
	Entry(
		"Non existing path with one segment",
		`{}`,
		[]string{
			`/api`,
		},
		`{
			"api": null
		}`,
	),
	Entry(
		"Non existing path with two segments",
		`{}`,
		[]string{
			`/api/clusters_mgmt`,
		},
		`{
			"api": {
				"clusters_mgmt": null
			}
		}`,
	),
	Entry(
		"Non existing path with three segments",
		`{}`,
		[]string{
			`/api/clusters_mgmt/v1`,
		},
		`{
			"api": {
				"clusters_mgmt": {
					"v1": null
				}
			}
		}`,
	),
	Entry(
		"Existing path with one segment",
		`{
			"api": null
		}`,
		[]string{
			`/api`,
		},
		`{
			"api": null
		}`,
	),
	Entry(
		"Existing path with two segments",
		`{
			"api": {
				"clusters_mgmt": null
			}
		}`,
		[]string{
			`/api/clusters_mgmt`,
		},
		`{
			"api": {
				"clusters_mgmt": null
			}
		}`,
	),
	Entry(
		"Existing path with three segments",
		`{
			"api": {
				"clusters_mgmt": {
					"v1": null
				}
			}
		}`,
		[]string{
			`/api/clusters_mgmt/v1`,
		},
		`{
			"api": {
				"clusters_mgmt": {
					"v1": null
				}
			}
		}`,
	),
	Entry(
		"Appends to partially existing path",
		`{
			"api": null
		}`,
		[]string{
			`/api/clusters_mgmt`,
		},
		`{
			"api": {
				"clusters_mgmt": null
			}
		}`,
	),
	Entry(
		"Adds default token URL",
		`{
			"api": {
				"clusters_mgmt": null
			}
		}`,
		[]string{
			`/auth/realms/redhat-external/protocol/openid-connect/token`,
		},
		`{
			"api": {
				"clusters_mgmt": null
			},
			"auth": {
				"realms": {
					"redhat-external": {
						"protocol": {
							"openid-connect": {
								"token": null
							}
						}
					}
				}
			}
		}`,
	),
	Entry(
		"Merges prefix",
		`{
			"api": {
				"clusters_mgmt": null
			}
		}`,
		[]string{
			`/api/clusters_mgmt`,
			`/api/accounts_mgmt`,
			`/api/service_logs`,
		},
		`{
			"api": {
				"clusters_mgmt": null,
				"accounts_mgmt": null,
				"service_logs": null
			}
		}`,
	),
)
