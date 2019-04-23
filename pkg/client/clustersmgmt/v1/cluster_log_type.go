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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// ClusterLogKind is the name of the type used to represent objects
// of type 'cluster_log'.
const ClusterLogKind = "ClusterLog"

// ClusterLogLinkKind is the name of the type used to represent links
// to objects of type 'cluster_log'.
const ClusterLogLinkKind = "ClusterLogLink"

// ClusterLogNilKind is the name of the type used to nil references
// to objects of type 'cluster_log'.
const ClusterLogNilKind = "ClusterLogNil"

// ClusterLog represents the values of the 'cluster_log' type.
//
// Logs of the cluster.
type ClusterLog struct {
	id      *string
	href    *string
	link    bool
	content *string
}

// Kind returns the name of the type of the object.
func (o *ClusterLog) Kind() string {
	if o == nil {
		return ClusterLogNilKind
	}
	if o.link {
		return ClusterLogLinkKind
	}
	return ClusterLogKind
}

// ID returns the identifier of the object.
func (o *ClusterLog) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ClusterLog) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *ClusterLog) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *ClusterLog) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ClusterLog) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Content returns the value of the 'content' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Content of the log.
func (o *ClusterLog) Content() string {
	if o != nil && o.content != nil {
		return *o.content
	}
	return ""
}

// GetContent returns the value of the 'content' attribute and
// a flag indicating if the attribute has a value.
//
// Content of the log.
func (o *ClusterLog) GetContent() (value string, ok bool) {
	ok = o != nil && o.content != nil
	if ok {
		value = *o.content
	}
	return
}

// ClusterLogList is a list of values of the 'cluster_log' type.
type ClusterLogList struct {
	items []*ClusterLog
}

// Len returns the length of the list.
func (l *ClusterLogList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterLogList) Slice() []*ClusterLog {
	var slice []*ClusterLog
	if l == nil {
		slice = make([]*ClusterLog, 0)
	} else {
		slice = make([]*ClusterLog, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterLogList) Each(f func(item *ClusterLog) bool) {
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
func (l *ClusterLogList) Range(f func(index int, item *ClusterLog) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
