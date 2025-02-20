// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetAvailableVHFAntennaSystemsOKCode is the HTTP code returned for type GetAvailableVHFAntennaSystemsOK
const GetAvailableVHFAntennaSystemsOKCode int = 200

/*
GetAvailableVHFAntennaSystemsOK Get Available VHF Antenna Systems Response

swagger:response getAvailableVHFAntennaSystemsOK
*/
type GetAvailableVHFAntennaSystemsOK struct {

	/*
	  In: Body
	*/
	Payload *models.VHFAntennaSystems `json:"body,omitempty"`
}

// NewGetAvailableVHFAntennaSystemsOK creates GetAvailableVHFAntennaSystemsOK with default headers values
func NewGetAvailableVHFAntennaSystemsOK() *GetAvailableVHFAntennaSystemsOK {

	return &GetAvailableVHFAntennaSystemsOK{}
}

// WithPayload adds the payload to the get available v h f antenna systems o k response
func (o *GetAvailableVHFAntennaSystemsOK) WithPayload(payload *models.VHFAntennaSystems) *GetAvailableVHFAntennaSystemsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available v h f antenna systems o k response
func (o *GetAvailableVHFAntennaSystemsOK) SetPayload(payload *models.VHFAntennaSystems) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableVHFAntennaSystemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableVHFAntennaSystemsBadRequestCode is the HTTP code returned for type GetAvailableVHFAntennaSystemsBadRequest
const GetAvailableVHFAntennaSystemsBadRequestCode int = 400

/*
GetAvailableVHFAntennaSystemsBadRequest Bad request

swagger:response getAvailableVHFAntennaSystemsBadRequest
*/
type GetAvailableVHFAntennaSystemsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableVHFAntennaSystemsBadRequest creates GetAvailableVHFAntennaSystemsBadRequest with default headers values
func NewGetAvailableVHFAntennaSystemsBadRequest() *GetAvailableVHFAntennaSystemsBadRequest {

	return &GetAvailableVHFAntennaSystemsBadRequest{}
}

// WithPayload adds the payload to the get available v h f antenna systems bad request response
func (o *GetAvailableVHFAntennaSystemsBadRequest) WithPayload(payload *models.Error) *GetAvailableVHFAntennaSystemsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available v h f antenna systems bad request response
func (o *GetAvailableVHFAntennaSystemsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableVHFAntennaSystemsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableVHFAntennaSystemsForbiddenCode is the HTTP code returned for type GetAvailableVHFAntennaSystemsForbidden
const GetAvailableVHFAntennaSystemsForbiddenCode int = 403

/*
GetAvailableVHFAntennaSystemsForbidden Forbidden

swagger:response getAvailableVHFAntennaSystemsForbidden
*/
type GetAvailableVHFAntennaSystemsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableVHFAntennaSystemsForbidden creates GetAvailableVHFAntennaSystemsForbidden with default headers values
func NewGetAvailableVHFAntennaSystemsForbidden() *GetAvailableVHFAntennaSystemsForbidden {

	return &GetAvailableVHFAntennaSystemsForbidden{}
}

// WithPayload adds the payload to the get available v h f antenna systems forbidden response
func (o *GetAvailableVHFAntennaSystemsForbidden) WithPayload(payload *models.Error) *GetAvailableVHFAntennaSystemsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available v h f antenna systems forbidden response
func (o *GetAvailableVHFAntennaSystemsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableVHFAntennaSystemsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableVHFAntennaSystemsUnprocessableEntityCode is the HTTP code returned for type GetAvailableVHFAntennaSystemsUnprocessableEntity
const GetAvailableVHFAntennaSystemsUnprocessableEntityCode int = 422

/*
GetAvailableVHFAntennaSystemsUnprocessableEntity Unprocessable Entity

swagger:response getAvailableVHFAntennaSystemsUnprocessableEntity
*/
type GetAvailableVHFAntennaSystemsUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableVHFAntennaSystemsUnprocessableEntity creates GetAvailableVHFAntennaSystemsUnprocessableEntity with default headers values
func NewGetAvailableVHFAntennaSystemsUnprocessableEntity() *GetAvailableVHFAntennaSystemsUnprocessableEntity {

	return &GetAvailableVHFAntennaSystemsUnprocessableEntity{}
}

// WithPayload adds the payload to the get available v h f antenna systems unprocessable entity response
func (o *GetAvailableVHFAntennaSystemsUnprocessableEntity) WithPayload(payload *models.Error) *GetAvailableVHFAntennaSystemsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available v h f antenna systems unprocessable entity response
func (o *GetAvailableVHFAntennaSystemsUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableVHFAntennaSystemsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableVHFAntennaSystemsInternalServerErrorCode is the HTTP code returned for type GetAvailableVHFAntennaSystemsInternalServerError
const GetAvailableVHFAntennaSystemsInternalServerErrorCode int = 500

/*
GetAvailableVHFAntennaSystemsInternalServerError Internal server error

swagger:response getAvailableVHFAntennaSystemsInternalServerError
*/
type GetAvailableVHFAntennaSystemsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableVHFAntennaSystemsInternalServerError creates GetAvailableVHFAntennaSystemsInternalServerError with default headers values
func NewGetAvailableVHFAntennaSystemsInternalServerError() *GetAvailableVHFAntennaSystemsInternalServerError {

	return &GetAvailableVHFAntennaSystemsInternalServerError{}
}

// WithPayload adds the payload to the get available v h f antenna systems internal server error response
func (o *GetAvailableVHFAntennaSystemsInternalServerError) WithPayload(payload *models.Error) *GetAvailableVHFAntennaSystemsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available v h f antenna systems internal server error response
func (o *GetAvailableVHFAntennaSystemsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableVHFAntennaSystemsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
