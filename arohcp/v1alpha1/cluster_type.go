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

// ClusterKind is the name of the type used to represent objects
// of type 'cluster'.
const ClusterKind = "Cluster"

// ClusterLinkKind is the name of the type used to represent links
// to objects of type 'cluster'.
const ClusterLinkKind = "ClusterLink"

// ClusterNilKind is the name of the type used to nil references
// to objects of type 'cluster'.
const ClusterNilKind = "ClusterNil"

// Cluster represents the values of the 'cluster' type.
type Cluster struct {
	bitmap_ uint32
	id      string
	href    string
}

// Kind returns the name of the type of the object.
func (o *Cluster) Kind() string {
	if o == nil {
		return ClusterNilKind
	}
	if o.bitmap_&1 != 0 {
		return ClusterLinkKind
	}
	return ClusterKind
}

// Link returns true iif this is a link.
func (o *Cluster) Link() bool {
	return o != nil && o.bitmap_&1 != 0
}

// ID returns the identifier of the object.
func (o *Cluster) ID() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Cluster) GetID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *Cluster) HREF() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Cluster) GetHREF() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Cluster) Empty() bool {
	return o == nil || o.bitmap_&^1 == 0
}

// ClusterListKind is the name of the type used to represent list of objects of
// type 'cluster'.
const ClusterListKind = "ClusterList"

// ClusterListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster'.
const ClusterListLinkKind = "ClusterListLink"

// ClusterNilKind is the name of the type used to nil lists of objects of
// type 'cluster'.
const ClusterListNilKind = "ClusterListNil"

// ClusterList is a list of values of the 'cluster' type.
type ClusterList struct {
	href  string
	link  bool
	items []*Cluster
}

// Kind returns the name of the type of the object.
func (l *ClusterList) Kind() string {
	if l == nil {
		return ClusterListNilKind
	}
	if l.link {
		return ClusterListLinkKind
	}
	return ClusterListKind
}

// Link returns true iif this is a link.
func (l *ClusterList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ClusterList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ClusterList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *ClusterList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterList) Get(i int) *Cluster {
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
func (l *ClusterList) Slice() []*Cluster {
	var slice []*Cluster
	if l == nil {
		slice = make([]*Cluster, 0)
	} else {
		slice = make([]*Cluster, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterList) Each(f func(item *Cluster) bool) {
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
func (l *ClusterList) Range(f func(index int, item *Cluster) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
