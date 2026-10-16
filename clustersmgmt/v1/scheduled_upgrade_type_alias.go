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

// ScheduledUpgradeKind is the name of the type used to represent objects
// of type 'scheduled_upgrade'.
const ScheduledUpgradeKind = api_v1.ScheduledUpgradeKind

// ScheduledUpgradeLinkKind is the name of the type used to represent links
// to objects of type 'scheduled_upgrade'.
const ScheduledUpgradeLinkKind = api_v1.ScheduledUpgradeLinkKind

// ScheduledUpgradeNilKind is the name of the type used to nil references
// to objects of type 'scheduled_upgrade'.
const ScheduledUpgradeNilKind = api_v1.ScheduledUpgradeNilKind

// ScheduledUpgrade represents the values of the 'scheduled_upgrade' type.
//
// Representation of one upgrade scheduled somewhere in the fleet, regardless of
// which kind of upgrade policy scheduled it. Read-only; the authoritative
// resource is linked from 'policy'.
type ScheduledUpgrade = api_v1.ScheduledUpgrade

// ScheduledUpgradeListKind is the name of the type used to represent list of objects of
// type 'scheduled_upgrade'.
const ScheduledUpgradeListKind = api_v1.ScheduledUpgradeListKind

// ScheduledUpgradeListLinkKind is the name of the type used to represent links to list
// of objects of type 'scheduled_upgrade'.
const ScheduledUpgradeListLinkKind = api_v1.ScheduledUpgradeListLinkKind

// ScheduledUpgradeNilKind is the name of the type used to nil lists of objects of
// type 'scheduled_upgrade'.
const ScheduledUpgradeListNilKind = api_v1.ScheduledUpgradeListNilKind

type ScheduledUpgradeList = api_v1.ScheduledUpgradeList
