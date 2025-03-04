// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteVHFAntennaSystemOKCode is the HTTP code returned for type DeleteVHFAntennaSystemOK
const DeleteVHFAntennaSystemOKCode int = 200

/*
DeleteVHFAntennaSystemOK Delete VHF Antenna System Response

swagger:response deleteVHFAntennaSystemOK
*/
type DeleteVHFAntennaSystemOK struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVHFAntennaSystemOK creates DeleteVHFAntennaSystemOK with default headers values
func NewDeleteVHFAntennaSystemOK() *DeleteVHFAntennaSystemOK {

	return &DeleteVHFAntennaSystemOK{}
}

// WithPayload adds the payload to the delete v h f antenna system o k response
func (o *DeleteVHFAntennaSystemOK) WithPayload(payload *models.Error) *DeleteVHFAntennaSystemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v h f antenna system o k response
func (o *DeleteVHFAntennaSystemOK) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVHFAntennaSystemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVHFAntennaSystemBadRequestCode is the HTTP code returned for type DeleteVHFAntennaSystemBadRequest
const DeleteVHFAntennaSystemBadRequestCode int = 400

/*
DeleteVHFAntennaSystemBadRequest Bad request

swagger:response deleteVHFAntennaSystemBadRequest
*/
type DeleteVHFAntennaSystemBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVHFAntennaSystemBadRequest creates DeleteVHFAntennaSystemBadRequest with default headers values
func NewDeleteVHFAntennaSystemBadRequest() *DeleteVHFAntennaSystemBadRequest {

	return &DeleteVHFAntennaSystemBadRequest{}
}

// WithPayload adds the payload to the delete v h f antenna system bad request response
func (o *DeleteVHFAntennaSystemBadRequest) WithPayload(payload *models.Error) *DeleteVHFAntennaSystemBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v h f antenna system bad request response
func (o *DeleteVHFAntennaSystemBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVHFAntennaSystemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVHFAntennaSystemForbiddenCode is the HTTP code returned for type DeleteVHFAntennaSystemForbidden
const DeleteVHFAntennaSystemForbiddenCode int = 403

/*
DeleteVHFAntennaSystemForbidden Forbidden

swagger:response deleteVHFAntennaSystemForbidden
*/
type DeleteVHFAntennaSystemForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVHFAntennaSystemForbidden creates DeleteVHFAntennaSystemForbidden with default headers values
func NewDeleteVHFAntennaSystemForbidden() *DeleteVHFAntennaSystemForbidden {

	return &DeleteVHFAntennaSystemForbidden{}
}

// WithPayload adds the payload to the delete v h f antenna system forbidden response
func (o *DeleteVHFAntennaSystemForbidden) WithPayload(payload *models.Error) *DeleteVHFAntennaSystemForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v h f antenna system forbidden response
func (o *DeleteVHFAntennaSystemForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVHFAntennaSystemForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVHFAntennaSystemUnprocessableEntityCode is the HTTP code returned for type DeleteVHFAntennaSystemUnprocessableEntity
const DeleteVHFAntennaSystemUnprocessableEntityCode int = 422

/*
DeleteVHFAntennaSystemUnprocessableEntity Unprocessable Entity

swagger:response deleteVHFAntennaSystemUnprocessableEntity
*/
type DeleteVHFAntennaSystemUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVHFAntennaSystemUnprocessableEntity creates DeleteVHFAntennaSystemUnprocessableEntity with default headers values
func NewDeleteVHFAntennaSystemUnprocessableEntity() *DeleteVHFAntennaSystemUnprocessableEntity {

	return &DeleteVHFAntennaSystemUnprocessableEntity{}
}

// WithPayload adds the payload to the delete v h f antenna system unprocessable entity response
func (o *DeleteVHFAntennaSystemUnprocessableEntity) WithPayload(payload *models.Error) *DeleteVHFAntennaSystemUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v h f antenna system unprocessable entity response
func (o *DeleteVHFAntennaSystemUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVHFAntennaSystemUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVHFAntennaSystemInternalServerErrorCode is the HTTP code returned for type DeleteVHFAntennaSystemInternalServerError
const DeleteVHFAntennaSystemInternalServerErrorCode int = 500

/*
DeleteVHFAntennaSystemInternalServerError Internal server error

swagger:response deleteVHFAntennaSystemInternalServerError
*/
type DeleteVHFAntennaSystemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVHFAntennaSystemInternalServerError creates DeleteVHFAntennaSystemInternalServerError with default headers values
func NewDeleteVHFAntennaSystemInternalServerError() *DeleteVHFAntennaSystemInternalServerError {

	return &DeleteVHFAntennaSystemInternalServerError{}
}

// WithPayload adds the payload to the delete v h f antenna system internal server error response
func (o *DeleteVHFAntennaSystemInternalServerError) WithPayload(payload *models.Error) *DeleteVHFAntennaSystemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v h f antenna system internal server error response
func (o *DeleteVHFAntennaSystemInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVHFAntennaSystemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
