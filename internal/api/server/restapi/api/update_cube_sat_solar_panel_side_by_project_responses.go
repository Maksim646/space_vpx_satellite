// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateCubeSatSolarPanelSideByProjectOKCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectOK
const UpdateCubeSatSolarPanelSideByProjectOKCode int = 200

/*
UpdateCubeSatSolarPanelSideByProjectOK Successfully updated Solar Panel Side

swagger:response updateCubeSatSolarPanelSideByProjectOK
*/
type UpdateCubeSatSolarPanelSideByProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.CubeSatProject `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectOK creates UpdateCubeSatSolarPanelSideByProjectOK with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectOK() *UpdateCubeSatSolarPanelSideByProjectOK {

	return &UpdateCubeSatSolarPanelSideByProjectOK{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project o k response
func (o *UpdateCubeSatSolarPanelSideByProjectOK) WithPayload(payload *models.CubeSatProject) *UpdateCubeSatSolarPanelSideByProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project o k response
func (o *UpdateCubeSatSolarPanelSideByProjectOK) SetPayload(payload *models.CubeSatProject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelSideByProjectBadRequestCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectBadRequest
const UpdateCubeSatSolarPanelSideByProjectBadRequestCode int = 400

/*
UpdateCubeSatSolarPanelSideByProjectBadRequest Bad request (e.g., invalid input)

swagger:response updateCubeSatSolarPanelSideByProjectBadRequest
*/
type UpdateCubeSatSolarPanelSideByProjectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectBadRequest creates UpdateCubeSatSolarPanelSideByProjectBadRequest with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectBadRequest() *UpdateCubeSatSolarPanelSideByProjectBadRequest {

	return &UpdateCubeSatSolarPanelSideByProjectBadRequest{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project bad request response
func (o *UpdateCubeSatSolarPanelSideByProjectBadRequest) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelSideByProjectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project bad request response
func (o *UpdateCubeSatSolarPanelSideByProjectBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelSideByProjectUnauthorizedCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectUnauthorized
const UpdateCubeSatSolarPanelSideByProjectUnauthorizedCode int = 401

/*
UpdateCubeSatSolarPanelSideByProjectUnauthorized Unauthorized

swagger:response updateCubeSatSolarPanelSideByProjectUnauthorized
*/
type UpdateCubeSatSolarPanelSideByProjectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectUnauthorized creates UpdateCubeSatSolarPanelSideByProjectUnauthorized with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectUnauthorized() *UpdateCubeSatSolarPanelSideByProjectUnauthorized {

	return &UpdateCubeSatSolarPanelSideByProjectUnauthorized{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project unauthorized response
func (o *UpdateCubeSatSolarPanelSideByProjectUnauthorized) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelSideByProjectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project unauthorized response
func (o *UpdateCubeSatSolarPanelSideByProjectUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelSideByProjectForbiddenCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectForbidden
const UpdateCubeSatSolarPanelSideByProjectForbiddenCode int = 403

/*
UpdateCubeSatSolarPanelSideByProjectForbidden Forbidden

swagger:response updateCubeSatSolarPanelSideByProjectForbidden
*/
type UpdateCubeSatSolarPanelSideByProjectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectForbidden creates UpdateCubeSatSolarPanelSideByProjectForbidden with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectForbidden() *UpdateCubeSatSolarPanelSideByProjectForbidden {

	return &UpdateCubeSatSolarPanelSideByProjectForbidden{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project forbidden response
func (o *UpdateCubeSatSolarPanelSideByProjectForbidden) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelSideByProjectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project forbidden response
func (o *UpdateCubeSatSolarPanelSideByProjectForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelSideByProjectNotFoundCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectNotFound
const UpdateCubeSatSolarPanelSideByProjectNotFoundCode int = 404

/*
UpdateCubeSatSolarPanelSideByProjectNotFound Project not found

swagger:response updateCubeSatSolarPanelSideByProjectNotFound
*/
type UpdateCubeSatSolarPanelSideByProjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectNotFound creates UpdateCubeSatSolarPanelSideByProjectNotFound with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectNotFound() *UpdateCubeSatSolarPanelSideByProjectNotFound {

	return &UpdateCubeSatSolarPanelSideByProjectNotFound{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project not found response
func (o *UpdateCubeSatSolarPanelSideByProjectNotFound) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelSideByProjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project not found response
func (o *UpdateCubeSatSolarPanelSideByProjectNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelSideByProjectUnprocessableEntityCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity
const UpdateCubeSatSolarPanelSideByProjectUnprocessableEntityCode int = 422

/*
UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity Unprocessable Entity (e.g., validation errors)

swagger:response updateCubeSatSolarPanelSideByProjectUnprocessableEntity
*/
type UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectUnprocessableEntity creates UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectUnprocessableEntity() *UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity {

	return &UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project unprocessable entity response
func (o *UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project unprocessable entity response
func (o *UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatSolarPanelSideByProjectInternalServerErrorCode is the HTTP code returned for type UpdateCubeSatSolarPanelSideByProjectInternalServerError
const UpdateCubeSatSolarPanelSideByProjectInternalServerErrorCode int = 500

/*
UpdateCubeSatSolarPanelSideByProjectInternalServerError Internal server error

swagger:response updateCubeSatSolarPanelSideByProjectInternalServerError
*/
type UpdateCubeSatSolarPanelSideByProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatSolarPanelSideByProjectInternalServerError creates UpdateCubeSatSolarPanelSideByProjectInternalServerError with default headers values
func NewUpdateCubeSatSolarPanelSideByProjectInternalServerError() *UpdateCubeSatSolarPanelSideByProjectInternalServerError {

	return &UpdateCubeSatSolarPanelSideByProjectInternalServerError{}
}

// WithPayload adds the payload to the update cube sat solar panel side by project internal server error response
func (o *UpdateCubeSatSolarPanelSideByProjectInternalServerError) WithPayload(payload *models.Error) *UpdateCubeSatSolarPanelSideByProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat solar panel side by project internal server error response
func (o *UpdateCubeSatSolarPanelSideByProjectInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatSolarPanelSideByProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
