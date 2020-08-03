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

// SupportCaseKind is the name of the type used to represent objects
// of type 'support_case'.
const SupportCaseKind = "SupportCase"

// SupportCaseLinkKind is the name of the type used to represent links
// to objects of type 'support_case'.
const SupportCaseLinkKind = "SupportCaseLink"

// SupportCaseNilKind is the name of the type used to nil references
// to objects of type 'support_case'.
const SupportCaseNilKind = "SupportCaseNil"

// SupportCase represents the values of the 'support_case' type.
//
//
type SupportCase struct {
	id            *string
	href          *string
	link          bool
	clusterUuid   *string
	description   *string
	eventStreamId *string
	severity      *string
	summary       *string
}

// Kind returns the name of the type of the object.
func (o *SupportCase) Kind() string {
	if o == nil {
		return SupportCaseNilKind
	}
	if o.link {
		return SupportCaseLinkKind
	}
	return SupportCaseKind
}

// ID returns the identifier of the object.
func (o *SupportCase) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *SupportCase) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *SupportCase) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *SupportCase) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *SupportCase) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *SupportCase) Empty() bool {
	return o == nil || (o.id == nil &&
		o.clusterUuid == nil &&
		o.description == nil &&
		o.eventStreamId == nil &&
		o.severity == nil &&
		o.summary == nil &&
		true)
}

// ClusterUuid returns the value of the 'cluster_uuid' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// (optional) cluster uuid in case we want to create an osl log.
func (o *SupportCase) ClusterUuid() string {
	if o != nil && o.clusterUuid != nil {
		return *o.clusterUuid
	}
	return ""
}

// GetClusterUuid returns the value of the 'cluster_uuid' attribute and
// a flag indicating if the attribute has a value.
//
// (optional) cluster uuid in case we want to create an osl log.
func (o *SupportCase) GetClusterUuid() (value string, ok bool) {
	ok = o != nil && o.clusterUuid != nil
	if ok {
		value = *o.clusterUuid
	}
	return
}

// Description returns the value of the 'description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case desciption.
func (o *SupportCase) Description() string {
	if o != nil && o.description != nil {
		return *o.description
	}
	return ""
}

// GetDescription returns the value of the 'description' attribute and
// a flag indicating if the attribute has a value.
//
// Support case desciption.
func (o *SupportCase) GetDescription() (value string, ok bool) {
	ok = o != nil && o.description != nil
	if ok {
		value = *o.description
	}
	return
}

// EventStreamId returns the value of the 'event_stream_id' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// (optional) event stream id for the support case so we could track it.
func (o *SupportCase) EventStreamId() string {
	if o != nil && o.eventStreamId != nil {
		return *o.eventStreamId
	}
	return ""
}

// GetEventStreamId returns the value of the 'event_stream_id' attribute and
// a flag indicating if the attribute has a value.
//
// (optional) event stream id for the support case so we could track it.
func (o *SupportCase) GetEventStreamId() (value string, ok bool) {
	ok = o != nil && o.eventStreamId != nil
	if ok {
		value = *o.eventStreamId
	}
	return
}

// Severity returns the value of the 'severity' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case severity.
func (o *SupportCase) Severity() string {
	if o != nil && o.severity != nil {
		return *o.severity
	}
	return ""
}

// GetSeverity returns the value of the 'severity' attribute and
// a flag indicating if the attribute has a value.
//
// Support case severity.
func (o *SupportCase) GetSeverity() (value string, ok bool) {
	ok = o != nil && o.severity != nil
	if ok {
		value = *o.severity
	}
	return
}

// Summary returns the value of the 'summary' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case title.
func (o *SupportCase) Summary() string {
	if o != nil && o.summary != nil {
		return *o.summary
	}
	return ""
}

// GetSummary returns the value of the 'summary' attribute and
// a flag indicating if the attribute has a value.
//
// Support case title.
func (o *SupportCase) GetSummary() (value string, ok bool) {
	ok = o != nil && o.summary != nil
	if ok {
		value = *o.summary
	}
	return
}

// SupportCaseListKind is the name of the type used to represent list of objects of
// type 'support_case'.
const SupportCaseListKind = "SupportCaseList"

// SupportCaseListLinkKind is the name of the type used to represent links to list
// of objects of type 'support_case'.
const SupportCaseListLinkKind = "SupportCaseListLink"

// SupportCaseNilKind is the name of the type used to nil lists of objects of
// type 'support_case'.
const SupportCaseListNilKind = "SupportCaseListNil"

// SupportCaseList is a list of values of the 'support_case' type.
type SupportCaseList struct {
	href  *string
	link  bool
	items []*SupportCase
}

// Kind returns the name of the type of the object.
func (l *SupportCaseList) Kind() string {
	if l == nil {
		return SupportCaseListNilKind
	}
	if l.link {
		return SupportCaseListLinkKind
	}
	return SupportCaseListKind
}

// Link returns true iif this is a link.
func (l *SupportCaseList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *SupportCaseList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *SupportCaseList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *SupportCaseList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *SupportCaseList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *SupportCaseList) Get(i int) *SupportCase {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *SupportCaseList) Slice() []*SupportCase {
	var slice []*SupportCase
	if l == nil {
		slice = make([]*SupportCase, 0)
	} else {
		slice = make([]*SupportCase, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *SupportCaseList) Each(f func(item *SupportCase) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *SupportCaseList) Range(f func(index int, item *SupportCase) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
