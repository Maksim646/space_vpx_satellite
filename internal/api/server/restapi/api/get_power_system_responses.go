// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetPowerSystemOKCode is the HTTP code returned for type GetPowerSystemOK
const GetPowerSystemOKCode int = 200

/*
GetPowerSystemOK Get Power System Response

swagger:response getPowerSystemOK
*/
type GetPowerSystemOK struct {

	/*
	  In: Body
	*/
	Payload *models.CubeSatPowerSystem `json:"body,omitempty"`
}

// NewGetPowerSystemOK creates GetPowerSystemOK with default headers values
func NewGetPowerSystemOK() *GetPowerSystemOK {

	return &GetPowerSystemOK{}
}

// WithPayload adds the payload to the get power system o k response
func (o *GetPowerSystemOK) WithPayload(payload *models.CubeSatPowerSystem) *GetPowerSystemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get power system o k response
func (o *GetPowerSystemOK) SetPayload(payload *models.CubeSatPowerSystem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPowerSystemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPowerSystemBadRequestCode is the HTTP code returned for type GetPowerSystemBadRequest
const GetPowerSystemBadRequestCode int = 400

/*
GetPowerSystemBadRequest Bad request

swagger:response getPowerSystemBadRequest
*/
type GetPowerSystemBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPowerSystemBadRequest creates GetPowerSystemBadRequest with default headers values
func NewGetPowerSystemBadRequest() *GetPowerSystemBadRequest {

	return &GetPowerSystemBadRequest{}
}

// WithPayload adds the payload to the get power system bad request response
func (o *GetPowerSystemBadRequest) WithPayload(payload *models.Error) *GetPowerSystemBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get power system bad request response
func (o *GetPowerSystemBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPowerSystemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPowerSystemForbiddenCode is the HTTP code returned for type GetPowerSystemForbidden
const GetPowerSystemForbiddenCode int = 403

/*
GetPowerSystemForbidden Forbidden

swagger:response getPowerSystemForbidden
*/
type GetPowerSystemForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPowerSystemForbidden creates GetPowerSystemForbidden with default headers values
func NewGetPowerSystemForbidden() *GetPowerSystemForbidden {

	return &GetPowerSystemForbidden{}
}

// WithPayload adds the payload to the get power system forbidden response
func (o *GetPowerSystemForbidden) WithPayload(payload *models.Error) *GetPowerSystemForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get power system forbidden response
func (o *GetPowerSystemForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPowerSystemForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPowerSystemUnprocessableEntityCode is the HTTP code returned for type GetPowerSystemUnprocessableEntity
const GetPowerSystemUnprocessableEntityCode int = 422

/*
GetPowerSystemUnprocessableEntity Unprocessable Entity

swagger:response getPowerSystemUnprocessableEntity
*/
type GetPowerSystemUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPowerSystemUnprocessableEntity creates GetPowerSystemUnprocessableEntity with default headers values
func NewGetPowerSystemUnprocessableEntity() *GetPowerSystemUnprocessableEntity {

	return &GetPowerSystemUnprocessableEntity{}
}

// WithPayload adds the payload to the get power system unprocessable entity response
func (o *GetPowerSystemUnprocessableEntity) WithPayload(payload *models.Error) *GetPowerSystemUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get power system unprocessable entity response
func (o *GetPowerSystemUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPowerSystemUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPowerSystemInternalServerErrorCode is the HTTP code returned for type GetPowerSystemInternalServerError
const GetPowerSystemInternalServerErrorCode int = 500

/*
GetPowerSystemInternalServerError Internal server error

swagger:response getPowerSystemInternalServerError
*/
type GetPowerSystemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPowerSystemInternalServerError creates GetPowerSystemInternalServerError with default headers values
func NewGetPowerSystemInternalServerError() *GetPowerSystemInternalServerError {

	return &GetPowerSystemInternalServerError{}
}

// WithPayload adds the payload to the get power system internal server error response
func (o *GetPowerSystemInternalServerError) WithPayload(payload *models.Error) *GetPowerSystemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get power system internal server error response
func (o *GetPowerSystemInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPowerSystemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
