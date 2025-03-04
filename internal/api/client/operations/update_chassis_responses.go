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

// UpdateChassisReader is a Reader for the UpdateChassis structure.
type UpdateChassisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateChassisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateChassisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateChassisBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateChassisForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateChassisUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateChassisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /chassis/{id}] UpdateChassis", response, response.Code())
	}
}

// NewUpdateChassisOK creates a UpdateChassisOK with default headers values
func NewUpdateChassisOK() *UpdateChassisOK {
	return &UpdateChassisOK{}
}

/*
UpdateChassisOK describes a response with status code 200, with default header values.

Update Chassis Response
*/
type UpdateChassisOK struct {
	Payload *models.Chassis
}

// IsSuccess returns true when this update chassis o k response has a 2xx status code
func (o *UpdateChassisOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update chassis o k response has a 3xx status code
func (o *UpdateChassisOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update chassis o k response has a 4xx status code
func (o *UpdateChassisOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update chassis o k response has a 5xx status code
func (o *UpdateChassisOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update chassis o k response a status code equal to that given
func (o *UpdateChassisOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update chassis o k response
func (o *UpdateChassisOK) Code() int {
	return 200
}

func (o *UpdateChassisOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisOK %s", 200, payload)
}

func (o *UpdateChassisOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisOK %s", 200, payload)
}

func (o *UpdateChassisOK) GetPayload() *models.Chassis {
	return o.Payload
}

func (o *UpdateChassisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Chassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChassisBadRequest creates a UpdateChassisBadRequest with default headers values
func NewUpdateChassisBadRequest() *UpdateChassisBadRequest {
	return &UpdateChassisBadRequest{}
}

/*
UpdateChassisBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateChassisBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update chassis bad request response has a 2xx status code
func (o *UpdateChassisBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update chassis bad request response has a 3xx status code
func (o *UpdateChassisBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update chassis bad request response has a 4xx status code
func (o *UpdateChassisBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update chassis bad request response has a 5xx status code
func (o *UpdateChassisBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update chassis bad request response a status code equal to that given
func (o *UpdateChassisBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update chassis bad request response
func (o *UpdateChassisBadRequest) Code() int {
	return 400
}

func (o *UpdateChassisBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisBadRequest %s", 400, payload)
}

func (o *UpdateChassisBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisBadRequest %s", 400, payload)
}

func (o *UpdateChassisBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateChassisBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChassisForbidden creates a UpdateChassisForbidden with default headers values
func NewUpdateChassisForbidden() *UpdateChassisForbidden {
	return &UpdateChassisForbidden{}
}

/*
UpdateChassisForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateChassisForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this update chassis forbidden response has a 2xx status code
func (o *UpdateChassisForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update chassis forbidden response has a 3xx status code
func (o *UpdateChassisForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update chassis forbidden response has a 4xx status code
func (o *UpdateChassisForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update chassis forbidden response has a 5xx status code
func (o *UpdateChassisForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update chassis forbidden response a status code equal to that given
func (o *UpdateChassisForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update chassis forbidden response
func (o *UpdateChassisForbidden) Code() int {
	return 403
}

func (o *UpdateChassisForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisForbidden %s", 403, payload)
}

func (o *UpdateChassisForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisForbidden %s", 403, payload)
}

func (o *UpdateChassisForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateChassisForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChassisUnprocessableEntity creates a UpdateChassisUnprocessableEntity with default headers values
func NewUpdateChassisUnprocessableEntity() *UpdateChassisUnprocessableEntity {
	return &UpdateChassisUnprocessableEntity{}
}

/*
UpdateChassisUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type UpdateChassisUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this update chassis unprocessable entity response has a 2xx status code
func (o *UpdateChassisUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update chassis unprocessable entity response has a 3xx status code
func (o *UpdateChassisUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update chassis unprocessable entity response has a 4xx status code
func (o *UpdateChassisUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update chassis unprocessable entity response has a 5xx status code
func (o *UpdateChassisUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update chassis unprocessable entity response a status code equal to that given
func (o *UpdateChassisUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update chassis unprocessable entity response
func (o *UpdateChassisUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateChassisUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisUnprocessableEntity %s", 422, payload)
}

func (o *UpdateChassisUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisUnprocessableEntity %s", 422, payload)
}

func (o *UpdateChassisUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateChassisUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChassisInternalServerError creates a UpdateChassisInternalServerError with default headers values
func NewUpdateChassisInternalServerError() *UpdateChassisInternalServerError {
	return &UpdateChassisInternalServerError{}
}

/*
UpdateChassisInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateChassisInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update chassis internal server error response has a 2xx status code
func (o *UpdateChassisInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update chassis internal server error response has a 3xx status code
func (o *UpdateChassisInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update chassis internal server error response has a 4xx status code
func (o *UpdateChassisInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update chassis internal server error response has a 5xx status code
func (o *UpdateChassisInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update chassis internal server error response a status code equal to that given
func (o *UpdateChassisInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update chassis internal server error response
func (o *UpdateChassisInternalServerError) Code() int {
	return 500
}

func (o *UpdateChassisInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisInternalServerError %s", 500, payload)
}

func (o *UpdateChassisInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /chassis/{id}][%d] updateChassisInternalServerError %s", 500, payload)
}

func (o *UpdateChassisInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateChassisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
