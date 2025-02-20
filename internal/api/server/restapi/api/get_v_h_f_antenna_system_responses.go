// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetVHFAntennaSystemOKCode is the HTTP code returned for type GetVHFAntennaSystemOK
const GetVHFAntennaSystemOKCode int = 200

/*
GetVHFAntennaSystemOK Get VHF Antenna System Response

swagger:response getVHFAntennaSystemOK
*/
type GetVHFAntennaSystemOK struct {

	/*
	  In: Body
	*/
	Payload *models.VHFAntennaSystem `json:"body,omitempty"`
}

// NewGetVHFAntennaSystemOK creates GetVHFAntennaSystemOK with default headers values
func NewGetVHFAntennaSystemOK() *GetVHFAntennaSystemOK {

	return &GetVHFAntennaSystemOK{}
}

// WithPayload adds the payload to the get v h f antenna system o k response
func (o *GetVHFAntennaSystemOK) WithPayload(payload *models.VHFAntennaSystem) *GetVHFAntennaSystemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v h f antenna system o k response
func (o *GetVHFAntennaSystemOK) SetPayload(payload *models.VHFAntennaSystem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVHFAntennaSystemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVHFAntennaSystemBadRequestCode is the HTTP code returned for type GetVHFAntennaSystemBadRequest
const GetVHFAntennaSystemBadRequestCode int = 400

/*
GetVHFAntennaSystemBadRequest Bad request

swagger:response getVHFAntennaSystemBadRequest
*/
type GetVHFAntennaSystemBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVHFAntennaSystemBadRequest creates GetVHFAntennaSystemBadRequest with default headers values
func NewGetVHFAntennaSystemBadRequest() *GetVHFAntennaSystemBadRequest {

	return &GetVHFAntennaSystemBadRequest{}
}

// WithPayload adds the payload to the get v h f antenna system bad request response
func (o *GetVHFAntennaSystemBadRequest) WithPayload(payload *models.Error) *GetVHFAntennaSystemBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v h f antenna system bad request response
func (o *GetVHFAntennaSystemBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVHFAntennaSystemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVHFAntennaSystemForbiddenCode is the HTTP code returned for type GetVHFAntennaSystemForbidden
const GetVHFAntennaSystemForbiddenCode int = 403

/*
GetVHFAntennaSystemForbidden Forbidden

swagger:response getVHFAntennaSystemForbidden
*/
type GetVHFAntennaSystemForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVHFAntennaSystemForbidden creates GetVHFAntennaSystemForbidden with default headers values
func NewGetVHFAntennaSystemForbidden() *GetVHFAntennaSystemForbidden {

	return &GetVHFAntennaSystemForbidden{}
}

// WithPayload adds the payload to the get v h f antenna system forbidden response
func (o *GetVHFAntennaSystemForbidden) WithPayload(payload *models.Error) *GetVHFAntennaSystemForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v h f antenna system forbidden response
func (o *GetVHFAntennaSystemForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVHFAntennaSystemForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVHFAntennaSystemUnprocessableEntityCode is the HTTP code returned for type GetVHFAntennaSystemUnprocessableEntity
const GetVHFAntennaSystemUnprocessableEntityCode int = 422

/*
GetVHFAntennaSystemUnprocessableEntity Unprocessable Entity

swagger:response getVHFAntennaSystemUnprocessableEntity
*/
type GetVHFAntennaSystemUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVHFAntennaSystemUnprocessableEntity creates GetVHFAntennaSystemUnprocessableEntity with default headers values
func NewGetVHFAntennaSystemUnprocessableEntity() *GetVHFAntennaSystemUnprocessableEntity {

	return &GetVHFAntennaSystemUnprocessableEntity{}
}

// WithPayload adds the payload to the get v h f antenna system unprocessable entity response
func (o *GetVHFAntennaSystemUnprocessableEntity) WithPayload(payload *models.Error) *GetVHFAntennaSystemUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v h f antenna system unprocessable entity response
func (o *GetVHFAntennaSystemUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVHFAntennaSystemUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVHFAntennaSystemInternalServerErrorCode is the HTTP code returned for type GetVHFAntennaSystemInternalServerError
const GetVHFAntennaSystemInternalServerErrorCode int = 500

/*
GetVHFAntennaSystemInternalServerError Internal server error

swagger:response getVHFAntennaSystemInternalServerError
*/
type GetVHFAntennaSystemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVHFAntennaSystemInternalServerError creates GetVHFAntennaSystemInternalServerError with default headers values
func NewGetVHFAntennaSystemInternalServerError() *GetVHFAntennaSystemInternalServerError {

	return &GetVHFAntennaSystemInternalServerError{}
}

// WithPayload adds the payload to the get v h f antenna system internal server error response
func (o *GetVHFAntennaSystemInternalServerError) WithPayload(payload *models.Error) *GetVHFAntennaSystemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v h f antenna system internal server error response
func (o *GetVHFAntennaSystemInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVHFAntennaSystemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
