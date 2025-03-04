// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetUserCubeSatProjectsOKCode is the HTTP code returned for type GetUserCubeSatProjectsOK
const GetUserCubeSatProjectsOKCode int = 200

/*
GetUserCubeSatProjectsOK Get User CubeSat Project Response

swagger:response getUserCubeSatProjectsOK
*/
type GetUserCubeSatProjectsOK struct {

	/*
	  In: Body
	*/
	Payload *models.CubeSatProjects `json:"body,omitempty"`
}

// NewGetUserCubeSatProjectsOK creates GetUserCubeSatProjectsOK with default headers values
func NewGetUserCubeSatProjectsOK() *GetUserCubeSatProjectsOK {

	return &GetUserCubeSatProjectsOK{}
}

// WithPayload adds the payload to the get user cube sat projects o k response
func (o *GetUserCubeSatProjectsOK) WithPayload(payload *models.CubeSatProjects) *GetUserCubeSatProjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user cube sat projects o k response
func (o *GetUserCubeSatProjectsOK) SetPayload(payload *models.CubeSatProjects) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserCubeSatProjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserCubeSatProjectsBadRequestCode is the HTTP code returned for type GetUserCubeSatProjectsBadRequest
const GetUserCubeSatProjectsBadRequestCode int = 400

/*
GetUserCubeSatProjectsBadRequest Bad request

swagger:response getUserCubeSatProjectsBadRequest
*/
type GetUserCubeSatProjectsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserCubeSatProjectsBadRequest creates GetUserCubeSatProjectsBadRequest with default headers values
func NewGetUserCubeSatProjectsBadRequest() *GetUserCubeSatProjectsBadRequest {

	return &GetUserCubeSatProjectsBadRequest{}
}

// WithPayload adds the payload to the get user cube sat projects bad request response
func (o *GetUserCubeSatProjectsBadRequest) WithPayload(payload *models.Error) *GetUserCubeSatProjectsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user cube sat projects bad request response
func (o *GetUserCubeSatProjectsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserCubeSatProjectsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserCubeSatProjectsForbiddenCode is the HTTP code returned for type GetUserCubeSatProjectsForbidden
const GetUserCubeSatProjectsForbiddenCode int = 403

/*
GetUserCubeSatProjectsForbidden Forbidden

swagger:response getUserCubeSatProjectsForbidden
*/
type GetUserCubeSatProjectsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserCubeSatProjectsForbidden creates GetUserCubeSatProjectsForbidden with default headers values
func NewGetUserCubeSatProjectsForbidden() *GetUserCubeSatProjectsForbidden {

	return &GetUserCubeSatProjectsForbidden{}
}

// WithPayload adds the payload to the get user cube sat projects forbidden response
func (o *GetUserCubeSatProjectsForbidden) WithPayload(payload *models.Error) *GetUserCubeSatProjectsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user cube sat projects forbidden response
func (o *GetUserCubeSatProjectsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserCubeSatProjectsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserCubeSatProjectsUnprocessableEntityCode is the HTTP code returned for type GetUserCubeSatProjectsUnprocessableEntity
const GetUserCubeSatProjectsUnprocessableEntityCode int = 422

/*
GetUserCubeSatProjectsUnprocessableEntity Unprocessable Entity

swagger:response getUserCubeSatProjectsUnprocessableEntity
*/
type GetUserCubeSatProjectsUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserCubeSatProjectsUnprocessableEntity creates GetUserCubeSatProjectsUnprocessableEntity with default headers values
func NewGetUserCubeSatProjectsUnprocessableEntity() *GetUserCubeSatProjectsUnprocessableEntity {

	return &GetUserCubeSatProjectsUnprocessableEntity{}
}

// WithPayload adds the payload to the get user cube sat projects unprocessable entity response
func (o *GetUserCubeSatProjectsUnprocessableEntity) WithPayload(payload *models.Error) *GetUserCubeSatProjectsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user cube sat projects unprocessable entity response
func (o *GetUserCubeSatProjectsUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserCubeSatProjectsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserCubeSatProjectsInternalServerErrorCode is the HTTP code returned for type GetUserCubeSatProjectsInternalServerError
const GetUserCubeSatProjectsInternalServerErrorCode int = 500

/*
GetUserCubeSatProjectsInternalServerError Internal server error

swagger:response getUserCubeSatProjectsInternalServerError
*/
type GetUserCubeSatProjectsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserCubeSatProjectsInternalServerError creates GetUserCubeSatProjectsInternalServerError with default headers values
func NewGetUserCubeSatProjectsInternalServerError() *GetUserCubeSatProjectsInternalServerError {

	return &GetUserCubeSatProjectsInternalServerError{}
}

// WithPayload adds the payload to the get user cube sat projects internal server error response
func (o *GetUserCubeSatProjectsInternalServerError) WithPayload(payload *models.Error) *GetUserCubeSatProjectsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user cube sat projects internal server error response
func (o *GetUserCubeSatProjectsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserCubeSatProjectsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
