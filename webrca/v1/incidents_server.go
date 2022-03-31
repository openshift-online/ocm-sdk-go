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

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// IncidentsServer represents the interface the manages the 'incidents' resource.
type IncidentsServer interface {

	// Add handles a request for the 'add' method.
	//
	//
	Add(ctx context.Context, request *IncidentsAddServerRequest, response *IncidentsAddServerResponse) error

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of incidents.
	List(ctx context.Context, request *IncidentsListServerRequest, response *IncidentsListServerResponse) error

	// Incident returns the target 'incident' server for the given identifier.
	//
	//
	Incident(id string) IncidentServer
}

// IncidentsAddServerRequest is the request for the 'add' method.
type IncidentsAddServerRequest struct {
	body *Incident
}

// Body returns the value of the 'body' parameter.
//
//
func (r *IncidentsAddServerRequest) Body() *Incident {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsAddServerRequest) GetBody() (value *Incident, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// IncidentsAddServerResponse is the response for the 'add' method.
type IncidentsAddServerResponse struct {
	status int
	err    *errors.Error
	body   *Incident
}

// Body sets the value of the 'body' parameter.
//
//
func (r *IncidentsAddServerResponse) Body(value *Incident) *IncidentsAddServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *IncidentsAddServerResponse) Status(value int) *IncidentsAddServerResponse {
	r.status = value
	return r
}

// IncidentsListServerRequest is the request for the 'list' method.
type IncidentsListServerRequest struct {
	creatorId            *string
	incidentCommanderId  *string
	incidentName         *string
	mine                 *bool
	onCallId             *string
	orderBy              *string
	page                 *int
	participantId        *string
	productId            *string
	publicId             *string
	responsibleManagerId *string
	size                 *int
}

// CreatorId returns the value of the 'creator_id' parameter.
//
//
func (r *IncidentsListServerRequest) CreatorId() string {
	if r != nil && r.creatorId != nil {
		return *r.creatorId
	}
	return ""
}

// GetCreatorId returns the value of the 'creator_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetCreatorId() (value string, ok bool) {
	ok = r != nil && r.creatorId != nil
	if ok {
		value = *r.creatorId
	}
	return
}

// IncidentCommanderId returns the value of the 'incident_commander_id' parameter.
//
//
func (r *IncidentsListServerRequest) IncidentCommanderId() string {
	if r != nil && r.incidentCommanderId != nil {
		return *r.incidentCommanderId
	}
	return ""
}

// GetIncidentCommanderId returns the value of the 'incident_commander_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetIncidentCommanderId() (value string, ok bool) {
	ok = r != nil && r.incidentCommanderId != nil
	if ok {
		value = *r.incidentCommanderId
	}
	return
}

// IncidentName returns the value of the 'incident_name' parameter.
//
// in Status String - TODO: This causes problems, figure out why
func (r *IncidentsListServerRequest) IncidentName() string {
	if r != nil && r.incidentName != nil {
		return *r.incidentName
	}
	return ""
}

// GetIncidentName returns the value of the 'incident_name' parameter and
// a flag indicating if the parameter has a value.
//
// in Status String - TODO: This causes problems, figure out why
func (r *IncidentsListServerRequest) GetIncidentName() (value string, ok bool) {
	ok = r != nil && r.incidentName != nil
	if ok {
		value = *r.incidentName
	}
	return
}

// Mine returns the value of the 'mine' parameter.
//
//
func (r *IncidentsListServerRequest) Mine() bool {
	if r != nil && r.mine != nil {
		return *r.mine
	}
	return false
}

// GetMine returns the value of the 'mine' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetMine() (value bool, ok bool) {
	ok = r != nil && r.mine != nil
	if ok {
		value = *r.mine
	}
	return
}

// OnCallId returns the value of the 'on_call_id' parameter.
//
//
func (r *IncidentsListServerRequest) OnCallId() string {
	if r != nil && r.onCallId != nil {
		return *r.onCallId
	}
	return ""
}

// GetOnCallId returns the value of the 'on_call_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetOnCallId() (value string, ok bool) {
	ok = r != nil && r.onCallId != nil
	if ok {
		value = *r.onCallId
	}
	return
}

// OrderBy returns the value of the 'order_by' parameter.
//
//
func (r *IncidentsListServerRequest) OrderBy() string {
	if r != nil && r.orderBy != nil {
		return *r.orderBy
	}
	return ""
}

// GetOrderBy returns the value of the 'order_by' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetOrderBy() (value string, ok bool) {
	ok = r != nil && r.orderBy != nil
	if ok {
		value = *r.orderBy
	}
	return
}

// Page returns the value of the 'page' parameter.
//
//
func (r *IncidentsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// ParticipantId returns the value of the 'participant_id' parameter.
//
//
func (r *IncidentsListServerRequest) ParticipantId() string {
	if r != nil && r.participantId != nil {
		return *r.participantId
	}
	return ""
}

// GetParticipantId returns the value of the 'participant_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetParticipantId() (value string, ok bool) {
	ok = r != nil && r.participantId != nil
	if ok {
		value = *r.participantId
	}
	return
}

// ProductId returns the value of the 'product_id' parameter.
//
//
func (r *IncidentsListServerRequest) ProductId() string {
	if r != nil && r.productId != nil {
		return *r.productId
	}
	return ""
}

// GetProductId returns the value of the 'product_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetProductId() (value string, ok bool) {
	ok = r != nil && r.productId != nil
	if ok {
		value = *r.productId
	}
	return
}

// PublicId returns the value of the 'public_id' parameter.
//
//
func (r *IncidentsListServerRequest) PublicId() string {
	if r != nil && r.publicId != nil {
		return *r.publicId
	}
	return ""
}

// GetPublicId returns the value of the 'public_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetPublicId() (value string, ok bool) {
	ok = r != nil && r.publicId != nil
	if ok {
		value = *r.publicId
	}
	return
}

// ResponsibleManagerId returns the value of the 'responsible_manager_id' parameter.
//
//
func (r *IncidentsListServerRequest) ResponsibleManagerId() string {
	if r != nil && r.responsibleManagerId != nil {
		return *r.responsibleManagerId
	}
	return ""
}

// GetResponsibleManagerId returns the value of the 'responsible_manager_id' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetResponsibleManagerId() (value string, ok bool) {
	ok = r != nil && r.responsibleManagerId != nil
	if ok {
		value = *r.responsibleManagerId
	}
	return
}

// Size returns the value of the 'size' parameter.
//
//
func (r *IncidentsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IncidentsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// IncidentsListServerResponse is the response for the 'list' method.
type IncidentsListServerResponse struct {
	status int
	err    *errors.Error
	items  *IncidentList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
//
func (r *IncidentsListServerResponse) Items(value *IncidentList) *IncidentsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
//
func (r *IncidentsListServerResponse) Page(value int) *IncidentsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
//
func (r *IncidentsListServerResponse) Size(value int) *IncidentsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
//
func (r *IncidentsListServerResponse) Total(value int) *IncidentsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *IncidentsListServerResponse) Status(value int) *IncidentsListServerResponse {
	r.status = value
	return r
}

// dispatchIncidents navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchIncidents(w http.ResponseWriter, r *http.Request, server IncidentsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "POST":
			adaptIncidentsAddRequest(w, r, server)
			return
		case "GET":
			adaptIncidentsListRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		target := server.Incident(segments[0])
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchIncident(w, r, target, segments[1:])
	}
}

// adaptIncidentsAddRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptIncidentsAddRequest(w http.ResponseWriter, r *http.Request, server IncidentsServer) {
	request := &IncidentsAddServerRequest{}
	err := readIncidentsAddRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &IncidentsAddServerResponse{}
	response.status = 201
	err = server.Add(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeIncidentsAddResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptIncidentsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptIncidentsListRequest(w http.ResponseWriter, r *http.Request, server IncidentsServer) {
	request := &IncidentsListServerRequest{}
	err := readIncidentsListRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &IncidentsListServerResponse{}
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
	err = writeIncidentsListResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
