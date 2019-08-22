/*
Copyright (c) 2019 Red Hat, Inc.

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

package v1 // github.com/openshift-online/uhc-sdk-go/accountsmgmt/v1

import (
	time "time"
)

// SubscriptionBuilder contains the data and logic needed to build 'subscription' objects.
//
//
type SubscriptionBuilder struct {
	id                 *string
	href               *string
	link               bool
	plan               *PlanBuilder
	registryCredential *RegistryCredentialBuilder
	clusterID          *string
	externalClusterID  *string
	organizationID     *string
	lastTelemetryDate  *time.Time
	creator            *AccountBuilder
}

// NewSubscription creates a new builder of 'subscription' objects.
func NewSubscription() *SubscriptionBuilder {
	return new(SubscriptionBuilder)
}

// ID sets the identifier of the object.
func (b *SubscriptionBuilder) ID(value string) *SubscriptionBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *SubscriptionBuilder) HREF(value string) *SubscriptionBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *SubscriptionBuilder) Link(value bool) *SubscriptionBuilder {
	b.link = value
	return b
}

// Plan sets the value of the 'plan' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) Plan(value *PlanBuilder) *SubscriptionBuilder {
	b.plan = value
	return b
}

// RegistryCredential sets the value of the 'registry_credential' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) RegistryCredential(value *RegistryCredentialBuilder) *SubscriptionBuilder {
	b.registryCredential = value
	return b
}

// ClusterID sets the value of the 'cluster_ID' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) ClusterID(value string) *SubscriptionBuilder {
	b.clusterID = &value
	return b
}

// ExternalClusterID sets the value of the 'external_cluster_ID' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) ExternalClusterID(value string) *SubscriptionBuilder {
	b.externalClusterID = &value
	return b
}

// OrganizationID sets the value of the 'organization_ID' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) OrganizationID(value string) *SubscriptionBuilder {
	b.organizationID = &value
	return b
}

// LastTelemetryDate sets the value of the 'last_telemetry_date' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) LastTelemetryDate(value time.Time) *SubscriptionBuilder {
	b.lastTelemetryDate = &value
	return b
}

// Creator sets the value of the 'creator' attribute
// to the given value.
//
//
func (b *SubscriptionBuilder) Creator(value *AccountBuilder) *SubscriptionBuilder {
	b.creator = value
	return b
}

// Build creates a 'subscription' object using the configuration stored in the builder.
func (b *SubscriptionBuilder) Build() (object *Subscription, err error) {
	object = new(Subscription)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.plan != nil {
		object.plan, err = b.plan.Build()
		if err != nil {
			return
		}
	}
	if b.registryCredential != nil {
		object.registryCredential, err = b.registryCredential.Build()
		if err != nil {
			return
		}
	}
	if b.clusterID != nil {
		object.clusterID = b.clusterID
	}
	if b.externalClusterID != nil {
		object.externalClusterID = b.externalClusterID
	}
	if b.organizationID != nil {
		object.organizationID = b.organizationID
	}
	if b.lastTelemetryDate != nil {
		object.lastTelemetryDate = b.lastTelemetryDate
	}
	if b.creator != nil {
		object.creator, err = b.creator.Build()
		if err != nil {
			return
		}
	}
	return
}
