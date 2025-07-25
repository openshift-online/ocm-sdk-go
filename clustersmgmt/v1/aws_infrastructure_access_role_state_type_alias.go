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

// AWSInfrastructureAccessRoleState represents the values of the 'AWS_infrastructure_access_role_state' enumerated type.
type AWSInfrastructureAccessRoleState = api_v1.AWSInfrastructureAccessRoleState

const (
	// Role definition is invalid. Role can't be used.
	AWSInfrastructureAccessRoleStateInvalid AWSInfrastructureAccessRoleState = api_v1.AWSInfrastructureAccessRoleStateInvalid
	// This is a special state intended for the user know
	// that the access role has been removed by SRE,
	// but there are still grants referencing it.
	// Role can't be used in a new grant.
	AWSInfrastructureAccessRoleStateRemoved AWSInfrastructureAccessRoleState = api_v1.AWSInfrastructureAccessRoleStateRemoved
	// Access role is valid an can be used.
	// Only valid roles can be used in a role grant.
	AWSInfrastructureAccessRoleStateValid AWSInfrastructureAccessRoleState = api_v1.AWSInfrastructureAccessRoleStateValid
)
