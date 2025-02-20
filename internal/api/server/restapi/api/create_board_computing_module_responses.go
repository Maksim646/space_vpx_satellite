// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// CreateBoardComputingModuleCreatedCode is the HTTP code returned for type CreateBoardComputingModuleCreated
const CreateBoardComputingModuleCreatedCode int = 201

/*
CreateBoardComputingModuleCreated Successfully created Board Computing Module

swagger:response createBoardComputingModuleCreated
*/
type CreateBoardComputingModuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.BoardComputingModule `json:"body,omitempty"`
}

// NewCreateBoardComputingModuleCreated creates CreateBoardComputingModuleCreated with default headers values
func NewCreateBoardComputingModuleCreated() *CreateBoardComputingModuleCreated {

	return &CreateBoardComputingModuleCreated{}
}

// WithPayload adds the payload to the create board computing module created response
func (o *CreateBoardComputingModuleCreated) WithPayload(payload *models.BoardComputingModule) *CreateBoardComputingModuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create board computing module created response
func (o *CreateBoardComputingModuleCreated) SetPayload(payload *models.BoardComputingModule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBoardComputingModuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBoardComputingModuleBadRequestCode is the HTTP code returned for type CreateBoardComputingModuleBadRequest
const CreateBoardComputingModuleBadRequestCode int = 400

/*
CreateBoardComputingModuleBadRequest Bad request (e.g., invalid input)

swagger:response createBoardComputingModuleBadRequest
*/
type CreateBoardComputingModuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBoardComputingModuleBadRequest creates CreateBoardComputingModuleBadRequest with default headers values
func NewCreateBoardComputingModuleBadRequest() *CreateBoardComputingModuleBadRequest {

	return &CreateBoardComputingModuleBadRequest{}
}

// WithPayload adds the payload to the create board computing module bad request response
func (o *CreateBoardComputingModuleBadRequest) WithPayload(payload *models.Error) *CreateBoardComputingModuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create board computing module bad request response
func (o *CreateBoardComputingModuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBoardComputingModuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBoardComputingModuleUnauthorizedCode is the HTTP code returned for type CreateBoardComputingModuleUnauthorized
const CreateBoardComputingModuleUnauthorizedCode int = 401

/*
CreateBoardComputingModuleUnauthorized Unauthorized

swagger:response createBoardComputingModuleUnauthorized
*/
type CreateBoardComputingModuleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBoardComputingModuleUnauthorized creates CreateBoardComputingModuleUnauthorized with default headers values
func NewCreateBoardComputingModuleUnauthorized() *CreateBoardComputingModuleUnauthorized {

	return &CreateBoardComputingModuleUnauthorized{}
}

// WithPayload adds the payload to the create board computing module unauthorized response
func (o *CreateBoardComputingModuleUnauthorized) WithPayload(payload *models.Error) *CreateBoardComputingModuleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create board computing module unauthorized response
func (o *CreateBoardComputingModuleUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBoardComputingModuleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBoardComputingModuleForbiddenCode is the HTTP code returned for type CreateBoardComputingModuleForbidden
const CreateBoardComputingModuleForbiddenCode int = 403

/*
CreateBoardComputingModuleForbidden Forbidden

swagger:response createBoardComputingModuleForbidden
*/
type CreateBoardComputingModuleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBoardComputingModuleForbidden creates CreateBoardComputingModuleForbidden with default headers values
func NewCreateBoardComputingModuleForbidden() *CreateBoardComputingModuleForbidden {

	return &CreateBoardComputingModuleForbidden{}
}

// WithPayload adds the payload to the create board computing module forbidden response
func (o *CreateBoardComputingModuleForbidden) WithPayload(payload *models.Error) *CreateBoardComputingModuleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create board computing module forbidden response
func (o *CreateBoardComputingModuleForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBoardComputingModuleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBoardComputingModuleUnprocessableEntityCode is the HTTP code returned for type CreateBoardComputingModuleUnprocessableEntity
const CreateBoardComputingModuleUnprocessableEntityCode int = 422

/*
CreateBoardComputingModuleUnprocessableEntity Unprocessable Entity (e.g., validation errors)

swagger:response createBoardComputingModuleUnprocessableEntity
*/
type CreateBoardComputingModuleUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBoardComputingModuleUnprocessableEntity creates CreateBoardComputingModuleUnprocessableEntity with default headers values
func NewCreateBoardComputingModuleUnprocessableEntity() *CreateBoardComputingModuleUnprocessableEntity {

	return &CreateBoardComputingModuleUnprocessableEntity{}
}

// WithPayload adds the payload to the create board computing module unprocessable entity response
func (o *CreateBoardComputingModuleUnprocessableEntity) WithPayload(payload *models.Error) *CreateBoardComputingModuleUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create board computing module unprocessable entity response
func (o *CreateBoardComputingModuleUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBoardComputingModuleUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBoardComputingModuleInternalServerErrorCode is the HTTP code returned for type CreateBoardComputingModuleInternalServerError
const CreateBoardComputingModuleInternalServerErrorCode int = 500

/*
CreateBoardComputingModuleInternalServerError Internal server error

swagger:response createBoardComputingModuleInternalServerError
*/
type CreateBoardComputingModuleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBoardComputingModuleInternalServerError creates CreateBoardComputingModuleInternalServerError with default headers values
func NewCreateBoardComputingModuleInternalServerError() *CreateBoardComputingModuleInternalServerError {

	return &CreateBoardComputingModuleInternalServerError{}
}

// WithPayload adds the payload to the create board computing module internal server error response
func (o *CreateBoardComputingModuleInternalServerError) WithPayload(payload *models.Error) *CreateBoardComputingModuleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create board computing module internal server error response
func (o *CreateBoardComputingModuleInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBoardComputingModuleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
