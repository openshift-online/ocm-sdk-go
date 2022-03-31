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

package v1 // github.com/openshift-online/ocm-sdk-go/webrca/v1

import (
	"context"
	"net/http"
	time "time"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// EventsServer represents the interface the manages the 'events' resource.
type EventsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of events
	List(ctx context.Context, request *EventsListServerRequest, response *EventsListServerResponse) error

	// Event returns the target 'event' server for the given identifier.
	//
	//
	Event(id string) EventServer
}

// EventsListServerRequest is the request for the 'list' method.
type EventsListServerRequest struct {
	createdAfter     *time.Time
	createdBefore    *time.Time
	eventType        *string
	note             *string
	orderBy          *string
	page             *int
	showSystemEvents *bool
	size             *int
}

// CreatedAfter returns the value of the 'created_after' parameter.
//
//
func (r *EventsListServerRequest) CreatedAfter() time.Time {
	if r != nil && r.createdAfter != nil {
		return *r.createdAfter
	}
	return time.Time{}
}

// GetCreatedAfter returns the value of the 'created_after' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetCreatedAfter() (value time.Time, ok bool) {
	ok = r != nil && r.createdAfter != nil
	if ok {
		value = *r.createdAfter
	}
	return
}

// CreatedBefore returns the value of the 'created_before' parameter.
//
//
func (r *EventsListServerRequest) CreatedBefore() time.Time {
	if r != nil && r.createdBefore != nil {
		return *r.createdBefore
	}
	return time.Time{}
}

// GetCreatedBefore returns the value of the 'created_before' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetCreatedBefore() (value time.Time, ok bool) {
	ok = r != nil && r.createdBefore != nil
	if ok {
		value = *r.createdBefore
	}
	return
}

// EventType returns the value of the 'event_type' parameter.
//
//
func (r *EventsListServerRequest) EventType() string {
	if r != nil && r.eventType != nil {
		return *r.eventType
	}
	return ""
}

// GetEventType returns the value of the 'event_type' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetEventType() (value string, ok bool) {
	ok = r != nil && r.eventType != nil
	if ok {
		value = *r.eventType
	}
	return
}

// Note returns the value of the 'note' parameter.
//
//
func (r *EventsListServerRequest) Note() string {
	if r != nil && r.note != nil {
		return *r.note
	}
	return ""
}

// GetNote returns the value of the 'note' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetNote() (value string, ok bool) {
	ok = r != nil && r.note != nil
	if ok {
		value = *r.note
	}
	return
}

// OrderBy returns the value of the 'order_by' parameter.
//
//
func (r *EventsListServerRequest) OrderBy() string {
	if r != nil && r.orderBy != nil {
		return *r.orderBy
	}
	return ""
}

// GetOrderBy returns the value of the 'order_by' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetOrderBy() (value string, ok bool) {
	ok = r != nil && r.orderBy != nil
	if ok {
		value = *r.orderBy
	}
	return
}

// Page returns the value of the 'page' parameter.
//
//
func (r *EventsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// ShowSystemEvents returns the value of the 'show_system_events' parameter.
//
//
func (r *EventsListServerRequest) ShowSystemEvents() bool {
	if r != nil && r.showSystemEvents != nil {
		return *r.showSystemEvents
	}
	return false
}

// GetShowSystemEvents returns the value of the 'show_system_events' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetShowSystemEvents() (value bool, ok bool) {
	ok = r != nil && r.showSystemEvents != nil
	if ok {
		value = *r.showSystemEvents
	}
	return
}

// Size returns the value of the 'size' parameter.
//
//
func (r *EventsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EventsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// EventsListServerResponse is the response for the 'list' method.
type EventsListServerResponse struct {
	status int
	err    *errors.Error
	items  *EventList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
//
func (r *EventsListServerResponse) Items(value *EventList) *EventsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
//
func (r *EventsListServerResponse) Page(value int) *EventsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
//
func (r *EventsListServerResponse) Size(value int) *EventsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
//
func (r *EventsListServerResponse) Total(value int) *EventsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *EventsListServerResponse) Status(value int) *EventsListServerResponse {
	r.status = value
	return r
}

// dispatchEvents navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchEvents(w http.ResponseWriter, r *http.Request, server EventsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptEventsListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.Event(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchEvent(w, r, target, segments[1:])
	}
}

// adaptEventsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptEventsListRequest(w http.ResponseWriter, r *http.Request, server EventsServer) {
	request := &EventsListServerRequest{}
	err := readEventsListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &EventsListServerResponse{}
	response.status = 200
	err = server.List(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeEventsListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
