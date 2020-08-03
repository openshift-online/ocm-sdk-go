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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// SupportCaseBuilder contains the data and logic needed to build 'support_case' objects.
//
//
type SupportCaseBuilder struct {
	id            *string
	href          *string
	link          bool
	clusterUuid   *string
	description   *string
	eventStreamId *string
	severity      *string
	summary       *string
}

// NewSupportCase creates a new builder of 'support_case' objects.
func NewSupportCase() *SupportCaseBuilder {
	return new(SupportCaseBuilder)
}

// ID sets the identifier of the object.
func (b *SupportCaseBuilder) ID(value string) *SupportCaseBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *SupportCaseBuilder) HREF(value string) *SupportCaseBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *SupportCaseBuilder) Link(value bool) *SupportCaseBuilder {
	b.link = value
	return b
}

// ClusterUuid sets the value of the 'cluster_uuid' attribute to the given value.
//
//
func (b *SupportCaseBuilder) ClusterUuid(value string) *SupportCaseBuilder {
	b.clusterUuid = &value
	return b
}

// Description sets the value of the 'description' attribute to the given value.
//
//
func (b *SupportCaseBuilder) Description(value string) *SupportCaseBuilder {
	b.description = &value
	return b
}

// EventStreamId sets the value of the 'event_stream_id' attribute to the given value.
//
//
func (b *SupportCaseBuilder) EventStreamId(value string) *SupportCaseBuilder {
	b.eventStreamId = &value
	return b
}

// Severity sets the value of the 'severity' attribute to the given value.
//
//
func (b *SupportCaseBuilder) Severity(value string) *SupportCaseBuilder {
	b.severity = &value
	return b
}

// Summary sets the value of the 'summary' attribute to the given value.
//
//
func (b *SupportCaseBuilder) Summary(value string) *SupportCaseBuilder {
	b.summary = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *SupportCaseBuilder) Copy(object *SupportCase) *SupportCaseBuilder {
	if object == nil {
		return b
	}
	b.id = object.id
	b.href = object.href
	b.link = object.link
	b.clusterUuid = object.clusterUuid
	b.description = object.description
	b.eventStreamId = object.eventStreamId
	b.severity = object.severity
	b.summary = object.summary
	return b
}

// Build creates a 'support_case' object using the configuration stored in the builder.
func (b *SupportCaseBuilder) Build() (object *SupportCase, err error) {
	object = new(SupportCase)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	object.clusterUuid = b.clusterUuid
	object.description = b.description
	object.eventStreamId = b.eventStreamId
	object.severity = b.severity
	object.summary = b.summary
	return
}
