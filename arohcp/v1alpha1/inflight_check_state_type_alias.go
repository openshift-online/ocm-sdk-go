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

// InflightCheckState represents the values of the 'inflight_check_state' enumerated type.
type InflightCheckState = api_v1alpha1.InflightCheckState

const (
	// The inflight check failed.
	InflightCheckStateFailed InflightCheckState = api_v1alpha1.InflightCheckStateFailed
	// The inflight check passed.
	InflightCheckStatePassed InflightCheckState = api_v1alpha1.InflightCheckStatePassed
	// The inflight check did not start running yet.
	InflightCheckStatePending InflightCheckState = api_v1alpha1.InflightCheckStatePending
	// The inflight check is currently running.
	InflightCheckStateRunning InflightCheckState = api_v1alpha1.InflightCheckStateRunning
)
