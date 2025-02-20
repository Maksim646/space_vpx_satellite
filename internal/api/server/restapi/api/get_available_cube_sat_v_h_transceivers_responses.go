// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetAvailableCubeSatVHTransceiversOKCode is the HTTP code returned for type GetAvailableCubeSatVHTransceiversOK
const GetAvailableCubeSatVHTransceiversOKCode int = 200

/*
GetAvailableCubeSatVHTransceiversOK Get Available CubeSat VHF Transceivers Response

swagger:response getAvailableCubeSatVHTransceiversOK
*/
type GetAvailableCubeSatVHTransceiversOK struct {

	/*
	  In: Body
	*/
	Payload *models.CubeSatVHFTransceivers `json:"body,omitempty"`
}

// NewGetAvailableCubeSatVHTransceiversOK creates GetAvailableCubeSatVHTransceiversOK with default headers values
func NewGetAvailableCubeSatVHTransceiversOK() *GetAvailableCubeSatVHTransceiversOK {

	return &GetAvailableCubeSatVHTransceiversOK{}
}

// WithPayload adds the payload to the get available cube sat v h transceivers o k response
func (o *GetAvailableCubeSatVHTransceiversOK) WithPayload(payload *models.CubeSatVHFTransceivers) *GetAvailableCubeSatVHTransceiversOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available cube sat v h transceivers o k response
func (o *GetAvailableCubeSatVHTransceiversOK) SetPayload(payload *models.CubeSatVHFTransceivers) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableCubeSatVHTransceiversOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableCubeSatVHTransceiversBadRequestCode is the HTTP code returned for type GetAvailableCubeSatVHTransceiversBadRequest
const GetAvailableCubeSatVHTransceiversBadRequestCode int = 400

/*
GetAvailableCubeSatVHTransceiversBadRequest Bad request

swagger:response getAvailableCubeSatVHTransceiversBadRequest
*/
type GetAvailableCubeSatVHTransceiversBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableCubeSatVHTransceiversBadRequest creates GetAvailableCubeSatVHTransceiversBadRequest with default headers values
func NewGetAvailableCubeSatVHTransceiversBadRequest() *GetAvailableCubeSatVHTransceiversBadRequest {

	return &GetAvailableCubeSatVHTransceiversBadRequest{}
}

// WithPayload adds the payload to the get available cube sat v h transceivers bad request response
func (o *GetAvailableCubeSatVHTransceiversBadRequest) WithPayload(payload *models.Error) *GetAvailableCubeSatVHTransceiversBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available cube sat v h transceivers bad request response
func (o *GetAvailableCubeSatVHTransceiversBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableCubeSatVHTransceiversBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableCubeSatVHTransceiversForbiddenCode is the HTTP code returned for type GetAvailableCubeSatVHTransceiversForbidden
const GetAvailableCubeSatVHTransceiversForbiddenCode int = 403

/*
GetAvailableCubeSatVHTransceiversForbidden Forbidden

swagger:response getAvailableCubeSatVHTransceiversForbidden
*/
type GetAvailableCubeSatVHTransceiversForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableCubeSatVHTransceiversForbidden creates GetAvailableCubeSatVHTransceiversForbidden with default headers values
func NewGetAvailableCubeSatVHTransceiversForbidden() *GetAvailableCubeSatVHTransceiversForbidden {

	return &GetAvailableCubeSatVHTransceiversForbidden{}
}

// WithPayload adds the payload to the get available cube sat v h transceivers forbidden response
func (o *GetAvailableCubeSatVHTransceiversForbidden) WithPayload(payload *models.Error) *GetAvailableCubeSatVHTransceiversForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available cube sat v h transceivers forbidden response
func (o *GetAvailableCubeSatVHTransceiversForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableCubeSatVHTransceiversForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableCubeSatVHTransceiversUnprocessableEntityCode is the HTTP code returned for type GetAvailableCubeSatVHTransceiversUnprocessableEntity
const GetAvailableCubeSatVHTransceiversUnprocessableEntityCode int = 422

/*
GetAvailableCubeSatVHTransceiversUnprocessableEntity Unprocessable Entity

swagger:response getAvailableCubeSatVHTransceiversUnprocessableEntity
*/
type GetAvailableCubeSatVHTransceiversUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableCubeSatVHTransceiversUnprocessableEntity creates GetAvailableCubeSatVHTransceiversUnprocessableEntity with default headers values
func NewGetAvailableCubeSatVHTransceiversUnprocessableEntity() *GetAvailableCubeSatVHTransceiversUnprocessableEntity {

	return &GetAvailableCubeSatVHTransceiversUnprocessableEntity{}
}

// WithPayload adds the payload to the get available cube sat v h transceivers unprocessable entity response
func (o *GetAvailableCubeSatVHTransceiversUnprocessableEntity) WithPayload(payload *models.Error) *GetAvailableCubeSatVHTransceiversUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available cube sat v h transceivers unprocessable entity response
func (o *GetAvailableCubeSatVHTransceiversUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableCubeSatVHTransceiversUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvailableCubeSatVHTransceiversInternalServerErrorCode is the HTTP code returned for type GetAvailableCubeSatVHTransceiversInternalServerError
const GetAvailableCubeSatVHTransceiversInternalServerErrorCode int = 500

/*
GetAvailableCubeSatVHTransceiversInternalServerError Internal server error

swagger:response getAvailableCubeSatVHTransceiversInternalServerError
*/
type GetAvailableCubeSatVHTransceiversInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAvailableCubeSatVHTransceiversInternalServerError creates GetAvailableCubeSatVHTransceiversInternalServerError with default headers values
func NewGetAvailableCubeSatVHTransceiversInternalServerError() *GetAvailableCubeSatVHTransceiversInternalServerError {

	return &GetAvailableCubeSatVHTransceiversInternalServerError{}
}

// WithPayload adds the payload to the get available cube sat v h transceivers internal server error response
func (o *GetAvailableCubeSatVHTransceiversInternalServerError) WithPayload(payload *models.Error) *GetAvailableCubeSatVHTransceiversInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get available cube sat v h transceivers internal server error response
func (o *GetAvailableCubeSatVHTransceiversInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvailableCubeSatVHTransceiversInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
