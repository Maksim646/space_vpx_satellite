// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteCubeSatSolarPanelTopOKCode is the HTTP code returned for type DeleteCubeSatSolarPanelTopOK
const DeleteCubeSatSolarPanelTopOKCode int = 200

/*
DeleteCubeSatSolarPanelTopOK Delete Cube Sat Solar Panel Top Response

swagger:response deleteCubeSatSolarPanelTopOK
*/
type DeleteCubeSatSolarPanelTopOK struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCubeSatSolarPanelTopOK creates DeleteCubeSatSolarPanelTopOK with default headers values
func NewDeleteCubeSatSolarPanelTopOK() *DeleteCubeSatSolarPanelTopOK {

	return &DeleteCubeSatSolarPanelTopOK{}
}

// WithPayload adds the payload to the delete cube sat solar panel top o k response
func (o *DeleteCubeSatSolarPanelTopOK) WithPayload(payload *models.Error) *DeleteCubeSatSolarPanelTopOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cube sat solar panel top o k response
func (o *DeleteCubeSatSolarPanelTopOK) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCubeSatSolarPanelTopOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCubeSatSolarPanelTopBadRequestCode is the HTTP code returned for type DeleteCubeSatSolarPanelTopBadRequest
const DeleteCubeSatSolarPanelTopBadRequestCode int = 400

/*
DeleteCubeSatSolarPanelTopBadRequest Bad request

swagger:response deleteCubeSatSolarPanelTopBadRequest
*/
type DeleteCubeSatSolarPanelTopBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCubeSatSolarPanelTopBadRequest creates DeleteCubeSatSolarPanelTopBadRequest with default headers values
func NewDeleteCubeSatSolarPanelTopBadRequest() *DeleteCubeSatSolarPanelTopBadRequest {

	return &DeleteCubeSatSolarPanelTopBadRequest{}
}

// WithPayload adds the payload to the delete cube sat solar panel top bad request response
func (o *DeleteCubeSatSolarPanelTopBadRequest) WithPayload(payload *models.Error) *DeleteCubeSatSolarPanelTopBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cube sat solar panel top bad request response
func (o *DeleteCubeSatSolarPanelTopBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCubeSatSolarPanelTopBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCubeSatSolarPanelTopForbiddenCode is the HTTP code returned for type DeleteCubeSatSolarPanelTopForbidden
const DeleteCubeSatSolarPanelTopForbiddenCode int = 403

/*
DeleteCubeSatSolarPanelTopForbidden Forbidden

swagger:response deleteCubeSatSolarPanelTopForbidden
*/
type DeleteCubeSatSolarPanelTopForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCubeSatSolarPanelTopForbidden creates DeleteCubeSatSolarPanelTopForbidden with default headers values
func NewDeleteCubeSatSolarPanelTopForbidden() *DeleteCubeSatSolarPanelTopForbidden {

	return &DeleteCubeSatSolarPanelTopForbidden{}
}

// WithPayload adds the payload to the delete cube sat solar panel top forbidden response
func (o *DeleteCubeSatSolarPanelTopForbidden) WithPayload(payload *models.Error) *DeleteCubeSatSolarPanelTopForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cube sat solar panel top forbidden response
func (o *DeleteCubeSatSolarPanelTopForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCubeSatSolarPanelTopForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCubeSatSolarPanelTopUnprocessableEntityCode is the HTTP code returned for type DeleteCubeSatSolarPanelTopUnprocessableEntity
const DeleteCubeSatSolarPanelTopUnprocessableEntityCode int = 422

/*
DeleteCubeSatSolarPanelTopUnprocessableEntity Unprocessable Entity

swagger:response deleteCubeSatSolarPanelTopUnprocessableEntity
*/
type DeleteCubeSatSolarPanelTopUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCubeSatSolarPanelTopUnprocessableEntity creates DeleteCubeSatSolarPanelTopUnprocessableEntity with default headers values
func NewDeleteCubeSatSolarPanelTopUnprocessableEntity() *DeleteCubeSatSolarPanelTopUnprocessableEntity {

	return &DeleteCubeSatSolarPanelTopUnprocessableEntity{}
}

// WithPayload adds the payload to the delete cube sat solar panel top unprocessable entity response
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) WithPayload(payload *models.Error) *DeleteCubeSatSolarPanelTopUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cube sat solar panel top unprocessable entity response
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCubeSatSolarPanelTopInternalServerErrorCode is the HTTP code returned for type DeleteCubeSatSolarPanelTopInternalServerError
const DeleteCubeSatSolarPanelTopInternalServerErrorCode int = 500

/*
DeleteCubeSatSolarPanelTopInternalServerError Internal server error

swagger:response deleteCubeSatSolarPanelTopInternalServerError
*/
type DeleteCubeSatSolarPanelTopInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCubeSatSolarPanelTopInternalServerError creates DeleteCubeSatSolarPanelTopInternalServerError with default headers values
func NewDeleteCubeSatSolarPanelTopInternalServerError() *DeleteCubeSatSolarPanelTopInternalServerError {

	return &DeleteCubeSatSolarPanelTopInternalServerError{}
}

// WithPayload adds the payload to the delete cube sat solar panel top internal server error response
func (o *DeleteCubeSatSolarPanelTopInternalServerError) WithPayload(payload *models.Error) *DeleteCubeSatSolarPanelTopInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cube sat solar panel top internal server error response
func (o *DeleteCubeSatSolarPanelTopInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCubeSatSolarPanelTopInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
