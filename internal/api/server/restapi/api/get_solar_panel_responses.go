// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetSolarPanelOKCode is the HTTP code returned for type GetSolarPanelOK
const GetSolarPanelOKCode int = 200

/*
GetSolarPanelOK Get SolarPanel Response

swagger:response getSolarPanelOK
*/
type GetSolarPanelOK struct {

	/*
	  In: Body
	*/
	Payload *models.SolarPanel `json:"body,omitempty"`
}

// NewGetSolarPanelOK creates GetSolarPanelOK with default headers values
func NewGetSolarPanelOK() *GetSolarPanelOK {

	return &GetSolarPanelOK{}
}

// WithPayload adds the payload to the get solar panel o k response
func (o *GetSolarPanelOK) WithPayload(payload *models.SolarPanel) *GetSolarPanelOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solar panel o k response
func (o *GetSolarPanelOK) SetPayload(payload *models.SolarPanel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolarPanelOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolarPanelBadRequestCode is the HTTP code returned for type GetSolarPanelBadRequest
const GetSolarPanelBadRequestCode int = 400

/*
GetSolarPanelBadRequest Bad request

swagger:response getSolarPanelBadRequest
*/
type GetSolarPanelBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolarPanelBadRequest creates GetSolarPanelBadRequest with default headers values
func NewGetSolarPanelBadRequest() *GetSolarPanelBadRequest {

	return &GetSolarPanelBadRequest{}
}

// WithPayload adds the payload to the get solar panel bad request response
func (o *GetSolarPanelBadRequest) WithPayload(payload *models.Error) *GetSolarPanelBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solar panel bad request response
func (o *GetSolarPanelBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolarPanelBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolarPanelForbiddenCode is the HTTP code returned for type GetSolarPanelForbidden
const GetSolarPanelForbiddenCode int = 403

/*
GetSolarPanelForbidden Forbidden

swagger:response getSolarPanelForbidden
*/
type GetSolarPanelForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolarPanelForbidden creates GetSolarPanelForbidden with default headers values
func NewGetSolarPanelForbidden() *GetSolarPanelForbidden {

	return &GetSolarPanelForbidden{}
}

// WithPayload adds the payload to the get solar panel forbidden response
func (o *GetSolarPanelForbidden) WithPayload(payload *models.Error) *GetSolarPanelForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solar panel forbidden response
func (o *GetSolarPanelForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolarPanelForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolarPanelUnprocessableEntityCode is the HTTP code returned for type GetSolarPanelUnprocessableEntity
const GetSolarPanelUnprocessableEntityCode int = 422

/*
GetSolarPanelUnprocessableEntity Unprocessable Entity

swagger:response getSolarPanelUnprocessableEntity
*/
type GetSolarPanelUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolarPanelUnprocessableEntity creates GetSolarPanelUnprocessableEntity with default headers values
func NewGetSolarPanelUnprocessableEntity() *GetSolarPanelUnprocessableEntity {

	return &GetSolarPanelUnprocessableEntity{}
}

// WithPayload adds the payload to the get solar panel unprocessable entity response
func (o *GetSolarPanelUnprocessableEntity) WithPayload(payload *models.Error) *GetSolarPanelUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solar panel unprocessable entity response
func (o *GetSolarPanelUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolarPanelUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolarPanelInternalServerErrorCode is the HTTP code returned for type GetSolarPanelInternalServerError
const GetSolarPanelInternalServerErrorCode int = 500

/*
GetSolarPanelInternalServerError Internal server error

swagger:response getSolarPanelInternalServerError
*/
type GetSolarPanelInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolarPanelInternalServerError creates GetSolarPanelInternalServerError with default headers values
func NewGetSolarPanelInternalServerError() *GetSolarPanelInternalServerError {

	return &GetSolarPanelInternalServerError{}
}

// WithPayload adds the payload to the get solar panel internal server error response
func (o *GetSolarPanelInternalServerError) WithPayload(payload *models.Error) *GetSolarPanelInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solar panel internal server error response
func (o *GetSolarPanelInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolarPanelInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
