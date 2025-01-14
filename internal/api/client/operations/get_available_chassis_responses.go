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

// GetAvailableChassisReader is a Reader for the GetAvailableChassis structure.
type GetAvailableChassisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableChassisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAvailableChassisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAvailableChassisBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAvailableChassisForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAvailableChassisUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAvailableChassisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /chassis/available_chassis] GetAvailableChassis", response, response.Code())
	}
}

// NewGetAvailableChassisOK creates a GetAvailableChassisOK with default headers values
func NewGetAvailableChassisOK() *GetAvailableChassisOK {
	return &GetAvailableChassisOK{}
}

/*
GetAvailableChassisOK describes a response with status code 200, with default header values.

Get Available Chassis Response
*/
type GetAvailableChassisOK struct {
	Payload *models.Chassises
}

// IsSuccess returns true when this get available chassis o k response has a 2xx status code
func (o *GetAvailableChassisOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get available chassis o k response has a 3xx status code
func (o *GetAvailableChassisOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available chassis o k response has a 4xx status code
func (o *GetAvailableChassisOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available chassis o k response has a 5xx status code
func (o *GetAvailableChassisOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get available chassis o k response a status code equal to that given
func (o *GetAvailableChassisOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get available chassis o k response
func (o *GetAvailableChassisOK) Code() int {
	return 200
}

func (o *GetAvailableChassisOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisOK %s", 200, payload)
}

func (o *GetAvailableChassisOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisOK %s", 200, payload)
}

func (o *GetAvailableChassisOK) GetPayload() *models.Chassises {
	return o.Payload
}

func (o *GetAvailableChassisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Chassises)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableChassisBadRequest creates a GetAvailableChassisBadRequest with default headers values
func NewGetAvailableChassisBadRequest() *GetAvailableChassisBadRequest {
	return &GetAvailableChassisBadRequest{}
}

/*
GetAvailableChassisBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAvailableChassisBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get available chassis bad request response has a 2xx status code
func (o *GetAvailableChassisBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available chassis bad request response has a 3xx status code
func (o *GetAvailableChassisBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available chassis bad request response has a 4xx status code
func (o *GetAvailableChassisBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get available chassis bad request response has a 5xx status code
func (o *GetAvailableChassisBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get available chassis bad request response a status code equal to that given
func (o *GetAvailableChassisBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get available chassis bad request response
func (o *GetAvailableChassisBadRequest) Code() int {
	return 400
}

func (o *GetAvailableChassisBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisBadRequest %s", 400, payload)
}

func (o *GetAvailableChassisBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisBadRequest %s", 400, payload)
}

func (o *GetAvailableChassisBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAvailableChassisBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableChassisForbidden creates a GetAvailableChassisForbidden with default headers values
func NewGetAvailableChassisForbidden() *GetAvailableChassisForbidden {
	return &GetAvailableChassisForbidden{}
}

/*
GetAvailableChassisForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAvailableChassisForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get available chassis forbidden response has a 2xx status code
func (o *GetAvailableChassisForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available chassis forbidden response has a 3xx status code
func (o *GetAvailableChassisForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available chassis forbidden response has a 4xx status code
func (o *GetAvailableChassisForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get available chassis forbidden response has a 5xx status code
func (o *GetAvailableChassisForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get available chassis forbidden response a status code equal to that given
func (o *GetAvailableChassisForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get available chassis forbidden response
func (o *GetAvailableChassisForbidden) Code() int {
	return 403
}

func (o *GetAvailableChassisForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisForbidden %s", 403, payload)
}

func (o *GetAvailableChassisForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisForbidden %s", 403, payload)
}

func (o *GetAvailableChassisForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAvailableChassisForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableChassisUnprocessableEntity creates a GetAvailableChassisUnprocessableEntity with default headers values
func NewGetAvailableChassisUnprocessableEntity() *GetAvailableChassisUnprocessableEntity {
	return &GetAvailableChassisUnprocessableEntity{}
}

/*
GetAvailableChassisUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetAvailableChassisUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get available chassis unprocessable entity response has a 2xx status code
func (o *GetAvailableChassisUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available chassis unprocessable entity response has a 3xx status code
func (o *GetAvailableChassisUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available chassis unprocessable entity response has a 4xx status code
func (o *GetAvailableChassisUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get available chassis unprocessable entity response has a 5xx status code
func (o *GetAvailableChassisUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get available chassis unprocessable entity response a status code equal to that given
func (o *GetAvailableChassisUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get available chassis unprocessable entity response
func (o *GetAvailableChassisUnprocessableEntity) Code() int {
	return 422
}

func (o *GetAvailableChassisUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisUnprocessableEntity %s", 422, payload)
}

func (o *GetAvailableChassisUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisUnprocessableEntity %s", 422, payload)
}

func (o *GetAvailableChassisUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAvailableChassisUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableChassisInternalServerError creates a GetAvailableChassisInternalServerError with default headers values
func NewGetAvailableChassisInternalServerError() *GetAvailableChassisInternalServerError {
	return &GetAvailableChassisInternalServerError{}
}

/*
GetAvailableChassisInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAvailableChassisInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get available chassis internal server error response has a 2xx status code
func (o *GetAvailableChassisInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available chassis internal server error response has a 3xx status code
func (o *GetAvailableChassisInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available chassis internal server error response has a 4xx status code
func (o *GetAvailableChassisInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available chassis internal server error response has a 5xx status code
func (o *GetAvailableChassisInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get available chassis internal server error response a status code equal to that given
func (o *GetAvailableChassisInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get available chassis internal server error response
func (o *GetAvailableChassisInternalServerError) Code() int {
	return 500
}

func (o *GetAvailableChassisInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisInternalServerError %s", 500, payload)
}

func (o *GetAvailableChassisInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/available_chassis][%d] getAvailableChassisInternalServerError %s", 500, payload)
}

func (o *GetAvailableChassisInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAvailableChassisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
