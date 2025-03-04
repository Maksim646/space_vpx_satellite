// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateCubeSatSolarPanelTopByProjectOKCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectOK
const UpdateCubeSatSolarPanelTopByProjectOKCode int = 200

/*
UpdateCubeSatSolarPanelTopByProjectOK Successfully updated Solar Panel Top

swagger:response updateCubeSatSolarPanelTopByProjectOK
*/
type UpdateCubeSatSolarPanelTopByProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.CubeSatProject `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectOK creates UpdateCubeSatSolarPanelTopByProjectOK with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectOK() *UpdateCubeSatSolarPanelTopByProjectOK {

	return &UpdateCubeSatSolarPanelTopByProjectOK{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project o k response
func (o *UpdateCubeSatSolarPanelTopByProjectOK) WithPayload(payload *models.CubeSatProject) *UpdateCubeSatSolarPanelTopByProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project o k response
func (o *UpdateCubeSatSolarPanelTopByProjectOK) SetPayload(payload *models.CubeSatProject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelTopByProjectBadRequestCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectBadRequest
const UpdateCubeSatSolarPanelTopByProjectBadRequestCode int = 400

/*
UpdateCubeSatSolarPanelTopByProjectBadRequest Bad request (e.g., invalid input)

swagger:response updateCubeSatSolarPanelTopByProjectBadRequest
*/
type UpdateCubeSatSolarPanelTopByProjectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectBadRequest creates UpdateCubeSatSolarPanelTopByProjectBadRequest with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectBadRequest() *UpdateCubeSatSolarPanelTopByProjectBadRequest {

	return &UpdateCubeSatSolarPanelTopByProjectBadRequest{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project bad request response
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelTopByProjectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project bad request response
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelTopByProjectUnauthorizedCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectUnauthorized
const UpdateCubeSatSolarPanelTopByProjectUnauthorizedCode int = 401

/*
UpdateCubeSatSolarPanelTopByProjectUnauthorized Unauthorized

swagger:response updateCubeSatSolarPanelTopByProjectUnauthorized
*/
type UpdateCubeSatSolarPanelTopByProjectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectUnauthorized creates UpdateCubeSatSolarPanelTopByProjectUnauthorized with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectUnauthorized() *UpdateCubeSatSolarPanelTopByProjectUnauthorized {

	return &UpdateCubeSatSolarPanelTopByProjectUnauthorized{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project unauthorized response
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelTopByProjectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project unauthorized response
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelTopByProjectForbiddenCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectForbidden
const UpdateCubeSatSolarPanelTopByProjectForbiddenCode int = 403

/*
UpdateCubeSatSolarPanelTopByProjectForbidden Forbidden

swagger:response updateCubeSatSolarPanelTopByProjectForbidden
*/
type UpdateCubeSatSolarPanelTopByProjectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectForbidden creates UpdateCubeSatSolarPanelTopByProjectForbidden with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectForbidden() *UpdateCubeSatSolarPanelTopByProjectForbidden {

	return &UpdateCubeSatSolarPanelTopByProjectForbidden{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project forbidden response
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelTopByProjectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project forbidden response
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelTopByProjectNotFoundCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectNotFound
const UpdateCubeSatSolarPanelTopByProjectNotFoundCode int = 404

/*
UpdateCubeSatSolarPanelTopByProjectNotFound Project not found

swagger:response updateCubeSatSolarPanelTopByProjectNotFound
*/
type UpdateCubeSatSolarPanelTopByProjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectNotFound creates UpdateCubeSatSolarPanelTopByProjectNotFound with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectNotFound() *UpdateCubeSatSolarPanelTopByProjectNotFound {

	return &UpdateCubeSatSolarPanelTopByProjectNotFound{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project not found response
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelTopByProjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project not found response
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelTopByProjectUnprocessableEntityCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity
const UpdateCubeSatSolarPanelTopByProjectUnprocessableEntityCode int = 422

/*
UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity Unprocessable Entity (e.g., validation errors)

swagger:response updateCubeSatSolarPanelTopByProjectUnprocessableEntity
*/
type UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectUnprocessableEntity creates UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectUnprocessableEntity() *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity {

	return &UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project unprocessable entity response
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project unprocessable entity response
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelTopByProjectInternalServerErrorCode is the HTTP code returned for type UpdateCubeSatSolarPanelTopByProjectInternalServerError
const UpdateCubeSatSolarPanelTopByProjectInternalServerErrorCode int = 500

/*
UpdateCubeSatSolarPanelTopByProjectInternalServerError Internal server error

swagger:response updateCubeSatSolarPanelTopByProjectInternalServerError
*/
type UpdateCubeSatSolarPanelTopByProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelTopByProjectInternalServerError creates UpdateCubeSatSolarPanelTopByProjectInternalServerError with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectInternalServerError() *UpdateCubeSatSolarPanelTopByProjectInternalServerError {

	return &UpdateCubeSatSolarPanelTopByProjectInternalServerError{}
}

// WithPayload adds the payload to the update cube sat solar panel top by project internal server error response
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelTopByProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel top by project internal server error response
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
