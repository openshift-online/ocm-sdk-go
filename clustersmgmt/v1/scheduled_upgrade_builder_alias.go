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

// ScheduledUpgradeBuilder contains the data and logic needed to build 'scheduled_upgrade' objects.
//
// Representation of one upgrade scheduled somewhere in the fleet, regardless of
// which kind of upgrade policy scheduled it. Read-only; the authoritative
// resource is linked from 'policy'.
type ScheduledUpgradeBuilder = api_v1.ScheduledUpgradeBuilder

// NewScheduledUpgrade creates a new builder of 'scheduled_upgrade' objects.
var NewScheduledUpgrade = api_v1.NewScheduledUpgrade
