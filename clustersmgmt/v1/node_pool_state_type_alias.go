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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	api_v1 "github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1"
)

// NodePoolStateKind is the name of the type used to represent objects
// of type 'node_pool_state'.
const NodePoolStateKind = api_v1.NodePoolStateKind

// NodePoolStateLinkKind is the name of the type used to represent links
// to objects of type 'node_pool_state'.
const NodePoolStateLinkKind = api_v1.NodePoolStateLinkKind

// NodePoolStateNilKind is the name of the type used to nil references
// to objects of type 'node_pool_state'.
const NodePoolStateNilKind = api_v1.NodePoolStateNilKind

// NodePoolState represents the values of the 'node_pool_state' type.
//
// Representation of the status of a node pool.
type NodePoolState = api_v1.NodePoolState
type NodePoolStateList = api_v1.NodePoolStateList
