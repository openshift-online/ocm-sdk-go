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

package v1 // github.com/openshift-online/ocm-sdk-go/servicemgmt/v1

import (
	api_v1 "github.com/openshift-online/ocm-api-model/clientapi/servicemgmt/v1"
)

// ManagedServiceKind is the name of the type used to represent objects
// of type 'managed_service'.
const ManagedServiceKind = api_v1.ManagedServiceKind

// ManagedServiceLinkKind is the name of the type used to represent links
// to objects of type 'managed_service'.
const ManagedServiceLinkKind = api_v1.ManagedServiceLinkKind

// ManagedServiceNilKind is the name of the type used to nil references
// to objects of type 'managed_service'.
const ManagedServiceNilKind = api_v1.ManagedServiceNilKind

// ManagedService represents the values of the 'managed_service' type.
//
// Represents data about a running Managed Service.
type ManagedService = api_v1.ManagedService
type ManagedServiceList = api_v1.ManagedServiceList
