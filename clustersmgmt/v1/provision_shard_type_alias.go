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

// ProvisionShardKind is the name of the type used to represent objects
// of type 'provision_shard'.
const ProvisionShardKind = api_v1.ProvisionShardKind

// ProvisionShardLinkKind is the name of the type used to represent links
// to objects of type 'provision_shard'.
const ProvisionShardLinkKind = api_v1.ProvisionShardLinkKind

// ProvisionShardNilKind is the name of the type used to nil references
// to objects of type 'provision_shard'.
const ProvisionShardNilKind = api_v1.ProvisionShardNilKind

// ProvisionShard represents the values of the 'provision_shard' type.
//
// Contains the properties of the provision shard, including AWS and GCP related configurations
type ProvisionShard = api_v1.ProvisionShard
type ProvisionShardList = api_v1.ProvisionShardList
