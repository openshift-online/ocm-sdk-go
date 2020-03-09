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

package v1 // github.com/openshift-online/ocm-sdk-go/servicelogs/v1

import (
	time "time"
)

// LogEntryBuilder contains the data and logic needed to build 'log_entry' objects.
//
//
type LogEntryBuilder struct {
	id           *string
	href         *string
	link         bool
	clusterUUID  *string
	description  *string
	internalOnly *bool
	serviceName  *string
	serviceName  *string
	severity     *Severity
	summary      *string
	timestamp    *time.Time
}

// NewLogEntry creates a new builder of 'log_entry' objects.
func NewLogEntry() *LogEntryBuilder {
	return new(LogEntryBuilder)
}

// ID sets the identifier of the object.
func (b *LogEntryBuilder) ID(value string) *LogEntryBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *LogEntryBuilder) HREF(value string) *LogEntryBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *LogEntryBuilder) Link(value bool) *LogEntryBuilder {
	b.link = value
	return b
}

// ClusterUUID sets the value of the 'cluster_UUID' attribute to the given value.
//
//
func (b *LogEntryBuilder) ClusterUUID(value string) *LogEntryBuilder {
	b.clusterUUID = &value
	return b
}

// Description sets the value of the 'description' attribute to the given value.
//
//
func (b *LogEntryBuilder) Description(value string) *LogEntryBuilder {
	b.description = &value
	return b
}

// InternalOnly sets the value of the 'internal_only' attribute to the given value.
//
//
func (b *LogEntryBuilder) InternalOnly(value bool) *LogEntryBuilder {
	b.internalOnly = &value
	return b
}

// ServiceName sets the value of the 'service_name' attribute to the given value.
//
//
func (b *LogEntryBuilder) ServiceName(value string) *LogEntryBuilder {
	b.serviceName = &value
	return b
}

// ServiceName sets the value of the 'service_name' attribute to the given value.
//
//
func (b *LogEntryBuilder) ServiceName(value string) *LogEntryBuilder {
	b.serviceName = &value
	return b
}

// Severity sets the value of the 'severity' attribute to the given value.
//
//
func (b *LogEntryBuilder) Severity(value Severity) *LogEntryBuilder {
	b.severity = &value
	return b
}

// Summary sets the value of the 'summary' attribute to the given value.
//
//
func (b *LogEntryBuilder) Summary(value string) *LogEntryBuilder {
	b.summary = &value
	return b
}

// Timestamp sets the value of the 'timestamp' attribute to the given value.
//
//
func (b *LogEntryBuilder) Timestamp(value time.Time) *LogEntryBuilder {
	b.timestamp = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *LogEntryBuilder) Copy(object *LogEntry) *LogEntryBuilder {
	if object == nil {
		return b
	}
	b.id = object.id
	b.href = object.href
	b.link = object.link
	b.clusterUUID = object.clusterUUID
	b.description = object.description
	b.internalOnly = object.internalOnly
	b.serviceName = object.serviceName
	b.serviceName = object.serviceName
	b.severity = object.severity
	b.summary = object.summary
	b.timestamp = object.timestamp
	return b
}

// Build creates a 'log_entry' object using the configuration stored in the builder.
func (b *LogEntryBuilder) Build() (object *LogEntry, err error) {
	object = new(LogEntry)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	object.clusterUUID = b.clusterUUID
	object.description = b.description
	object.internalOnly = b.internalOnly
	object.serviceName = b.serviceName
	object.serviceName = b.serviceName
	object.severity = b.severity
	object.summary = b.summary
	object.timestamp = b.timestamp
	return
}
