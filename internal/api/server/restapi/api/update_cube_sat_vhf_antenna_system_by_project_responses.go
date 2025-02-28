// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateCubeSatVhfAntennaSystemByProjectOKCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectOK
const UpdateCubeSatVhfAntennaSystemByProjectOKCode int = 200

/*
UpdateCubeSatVhfAntennaSystemByProjectOK Successfully updated VHF Antenna System

swagger:response updateCubeSatVhfAntennaSystemByProjectOK
*/
type UpdateCubeSatVhfAntennaSystemByProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.CubeSatProject `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectOK creates UpdateCubeSatVhfAntennaSystemByProjectOK with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectOK() *UpdateCubeSatVhfAntennaSystemByProjectOK {

	return &UpdateCubeSatVhfAntennaSystemByProjectOK{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project o k response
func (o *UpdateCubeSatVhfAntennaSystemByProjectOK) WithPayload(payload *models.CubeSatProject) *UpdateCubeSatVhfAntennaSystemByProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project o k response
func (o *UpdateCubeSatVhfAntennaSystemByProjectOK) SetPayload(payload *models.CubeSatProject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatVhfAntennaSystemByProjectBadRequestCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectBadRequest
const UpdateCubeSatVhfAntennaSystemByProjectBadRequestCode int = 400

/*
UpdateCubeSatVhfAntennaSystemByProjectBadRequest Bad request (e.g., invalid input)

swagger:response updateCubeSatVhfAntennaSystemByProjectBadRequest
*/
type UpdateCubeSatVhfAntennaSystemByProjectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectBadRequest creates UpdateCubeSatVhfAntennaSystemByProjectBadRequest with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectBadRequest() *UpdateCubeSatVhfAntennaSystemByProjectBadRequest {

	return &UpdateCubeSatVhfAntennaSystemByProjectBadRequest{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project bad request response
func (o *UpdateCubeSatVhfAntennaSystemByProjectBadRequest) WithPayload(payload *models.Error) *UpdateCubeSatVhfAntennaSystemByProjectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project bad request response
func (o *UpdateCubeSatVhfAntennaSystemByProjectBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatVhfAntennaSystemByProjectUnauthorizedCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectUnauthorized
const UpdateCubeSatVhfAntennaSystemByProjectUnauthorizedCode int = 401

/*
UpdateCubeSatVhfAntennaSystemByProjectUnauthorized Unauthorized

swagger:response updateCubeSatVhfAntennaSystemByProjectUnauthorized
*/
type UpdateCubeSatVhfAntennaSystemByProjectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectUnauthorized creates UpdateCubeSatVhfAntennaSystemByProjectUnauthorized with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectUnauthorized() *UpdateCubeSatVhfAntennaSystemByProjectUnauthorized {

	return &UpdateCubeSatVhfAntennaSystemByProjectUnauthorized{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project unauthorized response
func (o *UpdateCubeSatVhfAntennaSystemByProjectUnauthorized) WithPayload(payload *models.Error) *UpdateCubeSatVhfAntennaSystemByProjectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project unauthorized response
func (o *UpdateCubeSatVhfAntennaSystemByProjectUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatVhfAntennaSystemByProjectForbiddenCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectForbidden
const UpdateCubeSatVhfAntennaSystemByProjectForbiddenCode int = 403

/*
UpdateCubeSatVhfAntennaSystemByProjectForbidden Forbidden

swagger:response updateCubeSatVhfAntennaSystemByProjectForbidden
*/
type UpdateCubeSatVhfAntennaSystemByProjectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectForbidden creates UpdateCubeSatVhfAntennaSystemByProjectForbidden with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectForbidden() *UpdateCubeSatVhfAntennaSystemByProjectForbidden {

	return &UpdateCubeSatVhfAntennaSystemByProjectForbidden{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project forbidden response
func (o *UpdateCubeSatVhfAntennaSystemByProjectForbidden) WithPayload(payload *models.Error) *UpdateCubeSatVhfAntennaSystemByProjectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project forbidden response
func (o *UpdateCubeSatVhfAntennaSystemByProjectForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatVhfAntennaSystemByProjectNotFoundCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectNotFound
const UpdateCubeSatVhfAntennaSystemByProjectNotFoundCode int = 404

/*
UpdateCubeSatVhfAntennaSystemByProjectNotFound Project not found

swagger:response updateCubeSatVhfAntennaSystemByProjectNotFound
*/
type UpdateCubeSatVhfAntennaSystemByProjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectNotFound creates UpdateCubeSatVhfAntennaSystemByProjectNotFound with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectNotFound() *UpdateCubeSatVhfAntennaSystemByProjectNotFound {

	return &UpdateCubeSatVhfAntennaSystemByProjectNotFound{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project not found response
func (o *UpdateCubeSatVhfAntennaSystemByProjectNotFound) WithPayload(payload *models.Error) *UpdateCubeSatVhfAntennaSystemByProjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project not found response
func (o *UpdateCubeSatVhfAntennaSystemByProjectNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntityCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity
const UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntityCode int = 422

/*
UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity Unprocessable Entity (e.g., validation errors)

swagger:response updateCubeSatVhfAntennaSystemByProjectUnprocessableEntity
*/
type UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity creates UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity() *UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity {

	return &UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project unprocessable entity response
func (o *UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity) WithPayload(payload *models.Error) *UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project unprocessable entity response
func (o *UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateCubeSatVhfAntennaSystemByProjectInternalServerErrorCode is the HTTP code returned for type UpdateCubeSatVhfAntennaSystemByProjectInternalServerError
const UpdateCubeSatVhfAntennaSystemByProjectInternalServerErrorCode int = 500

/*
UpdateCubeSatVhfAntennaSystemByProjectInternalServerError Internal server error

swagger:response updateCubeSatVhfAntennaSystemByProjectInternalServerError
*/
type UpdateCubeSatVhfAntennaSystemByProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCubeSatVhfAntennaSystemByProjectInternalServerError creates UpdateCubeSatVhfAntennaSystemByProjectInternalServerError with default headers values
func NewUpdateCubeSatVhfAntennaSystemByProjectInternalServerError() *UpdateCubeSatVhfAntennaSystemByProjectInternalServerError {

	return &UpdateCubeSatVhfAntennaSystemByProjectInternalServerError{}
}

// WithPayload adds the payload to the update cube sat vhf antenna system by project internal server error response
func (o *UpdateCubeSatVhfAntennaSystemByProjectInternalServerError) WithPayload(payload *models.Error) *UpdateCubeSatVhfAntennaSystemByProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cube sat vhf antenna system by project internal server error response
func (o *UpdateCubeSatVhfAntennaSystemByProjectInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCubeSatVhfAntennaSystemByProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
