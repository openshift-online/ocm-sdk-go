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

package v1alpha1 // github.com/openshift-online/ocm-sdk-go/arohcp/v1alpha1

import (
	api_v1alpha1 "github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1"
)

// GCPImageOverrideKind is the name of the type used to represent objects
// of type 'GCP_image_override'.
const GCPImageOverrideKind = api_v1alpha1.GCPImageOverrideKind

// GCPImageOverrideLinkKind is the name of the type used to represent links
// to objects of type 'GCP_image_override'.
const GCPImageOverrideLinkKind = api_v1alpha1.GCPImageOverrideLinkKind

// GCPImageOverrideNilKind is the name of the type used to nil references
// to objects of type 'GCP_image_override'.
const GCPImageOverrideNilKind = api_v1alpha1.GCPImageOverrideNilKind

// GCPImageOverride represents the values of the 'GCP_image_override' type.
//
// GcpImageOverride specifies what a GCP VM Image should be used for a particular product and billing model
type GCPImageOverride = api_v1alpha1.GCPImageOverride
type GCPImageOverrideList = api_v1alpha1.GCPImageOverrideList
