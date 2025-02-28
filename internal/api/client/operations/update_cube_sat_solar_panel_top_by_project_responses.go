// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateCubeSatSolarPanelTopByProjectReader is a Reader for the UpdateCubeSatSolarPanelTopByProject structure.
type UpdateCubeSatSolarPanelTopByProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCubeSatSolarPanelTopByProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCubeSatSolarPanelTopByProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCubeSatSolarPanelTopByProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCubeSatSolarPanelTopByProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCubeSatSolarPanelTopByProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCubeSatSolarPanelTopByProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateCubeSatSolarPanelTopByProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCubeSatSolarPanelTopByProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /projects/{project_id}/solar_panel_top] UpdateCubeSatSolarPanelTopByProject", response, response.Code())
	}
}

// NewUpdateCubeSatSolarPanelTopByProjectOK creates a UpdateCubeSatSolarPanelTopByProjectOK with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectOK() *UpdateCubeSatSolarPanelTopByProjectOK {
	return &UpdateCubeSatSolarPanelTopByProjectOK{}
}

/*
UpdateCubeSatSolarPanelTopByProjectOK describes a response with status code 200, with default header values.

Successfully updated Solar Panel Top
*/
type UpdateCubeSatSolarPanelTopByProjectOK struct {
	Payload *models.CubeSatProject
}

// IsSuccess returns true when this update cube sat solar panel top by project o k response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cube sat solar panel top by project o k response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project o k response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cube sat solar panel top by project o k response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update cube sat solar panel top by project o k response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update cube sat solar panel top by project o k response
func (o *UpdateCubeSatSolarPanelTopByProjectOK) Code() int {
	return 200
}

func (o *UpdateCubeSatSolarPanelTopByProjectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectOK %s", 200, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectOK %s", 200, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectOK) GetPayload() *models.CubeSatProject {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CubeSatProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCubeSatSolarPanelTopByProjectBadRequest creates a UpdateCubeSatSolarPanelTopByProjectBadRequest with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectBadRequest() *UpdateCubeSatSolarPanelTopByProjectBadRequest {
	return &UpdateCubeSatSolarPanelTopByProjectBadRequest{}
}

/*
UpdateCubeSatSolarPanelTopByProjectBadRequest describes a response with status code 400, with default header values.

Bad request (e.g., invalid input)
*/
type UpdateCubeSatSolarPanelTopByProjectBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cube sat solar panel top by project bad request response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cube sat solar panel top by project bad request response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project bad request response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cube sat solar panel top by project bad request response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update cube sat solar panel top by project bad request response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update cube sat solar panel top by project bad request response
func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) Code() int {
	return 400
}

func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectBadRequest %s", 400, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectBadRequest %s", 400, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCubeSatSolarPanelTopByProjectUnauthorized creates a UpdateCubeSatSolarPanelTopByProjectUnauthorized with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectUnauthorized() *UpdateCubeSatSolarPanelTopByProjectUnauthorized {
	return &UpdateCubeSatSolarPanelTopByProjectUnauthorized{}
}

/*
UpdateCubeSatSolarPanelTopByProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateCubeSatSolarPanelTopByProjectUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cube sat solar panel top by project unauthorized response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cube sat solar panel top by project unauthorized response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project unauthorized response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cube sat solar panel top by project unauthorized response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update cube sat solar panel top by project unauthorized response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update cube sat solar panel top by project unauthorized response
func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) Code() int {
	return 401
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectUnauthorized %s", 401, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectUnauthorized %s", 401, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCubeSatSolarPanelTopByProjectForbidden creates a UpdateCubeSatSolarPanelTopByProjectForbidden with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectForbidden() *UpdateCubeSatSolarPanelTopByProjectForbidden {
	return &UpdateCubeSatSolarPanelTopByProjectForbidden{}
}

/*
UpdateCubeSatSolarPanelTopByProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCubeSatSolarPanelTopByProjectForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cube sat solar panel top by project forbidden response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cube sat solar panel top by project forbidden response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project forbidden response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cube sat solar panel top by project forbidden response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update cube sat solar panel top by project forbidden response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update cube sat solar panel top by project forbidden response
func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) Code() int {
	return 403
}

func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectForbidden %s", 403, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectForbidden %s", 403, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCubeSatSolarPanelTopByProjectNotFound creates a UpdateCubeSatSolarPanelTopByProjectNotFound with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectNotFound() *UpdateCubeSatSolarPanelTopByProjectNotFound {
	return &UpdateCubeSatSolarPanelTopByProjectNotFound{}
}

/*
UpdateCubeSatSolarPanelTopByProjectNotFound describes a response with status code 404, with default header values.

Project not found
*/
type UpdateCubeSatSolarPanelTopByProjectNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cube sat solar panel top by project not found response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cube sat solar panel top by project not found response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project not found response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cube sat solar panel top by project not found response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update cube sat solar panel top by project not found response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update cube sat solar panel top by project not found response
func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) Code() int {
	return 404
}

func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectNotFound %s", 404, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectNotFound %s", 404, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCubeSatSolarPanelTopByProjectUnprocessableEntity creates a UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectUnprocessableEntity() *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity {
	return &UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity{}
}

/*
UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity (e.g., validation errors)
*/
type UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cube sat solar panel top by project unprocessable entity response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cube sat solar panel top by project unprocessable entity response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project unprocessable entity response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cube sat solar panel top by project unprocessable entity response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update cube sat solar panel top by project unprocessable entity response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update cube sat solar panel top by project unprocessable entity response
func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectUnprocessableEntity %s", 422, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectUnprocessableEntity %s", 422, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCubeSatSolarPanelTopByProjectInternalServerError creates a UpdateCubeSatSolarPanelTopByProjectInternalServerError with default headers values
func NewUpdateCubeSatSolarPanelTopByProjectInternalServerError() *UpdateCubeSatSolarPanelTopByProjectInternalServerError {
	return &UpdateCubeSatSolarPanelTopByProjectInternalServerError{}
}

/*
UpdateCubeSatSolarPanelTopByProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateCubeSatSolarPanelTopByProjectInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cube sat solar panel top by project internal server error response has a 2xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cube sat solar panel top by project internal server error response has a 3xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cube sat solar panel top by project internal server error response has a 4xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cube sat solar panel top by project internal server error response has a 5xx status code
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update cube sat solar panel top by project internal server error response a status code equal to that given
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update cube sat solar panel top by project internal server error response
func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) Code() int {
	return 500
}

func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectInternalServerError %s", 500, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /projects/{project_id}/solar_panel_top][%d] updateCubeSatSolarPanelTopByProjectInternalServerError %s", 500, payload)
}

func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCubeSatSolarPanelTopByProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
