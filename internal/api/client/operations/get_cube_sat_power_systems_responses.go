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

// GetCubeSatPowerSystemsReader is a Reader for the GetCubeSatPowerSystems structure.
type GetCubeSatPowerSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCubeSatPowerSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCubeSatPowerSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCubeSatPowerSystemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCubeSatPowerSystemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCubeSatPowerSystemsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCubeSatPowerSystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cube_sat_power_system/available_power_systems] GetCubeSatPowerSystems", response, response.Code())
	}
}

// NewGetCubeSatPowerSystemsOK creates a GetCubeSatPowerSystemsOK with default headers values
func NewGetCubeSatPowerSystemsOK() *GetCubeSatPowerSystemsOK {
	return &GetCubeSatPowerSystemsOK{}
}

/*
GetCubeSatPowerSystemsOK describes a response with status code 200, with default header values.

Get CubeSat Power Systems Response
*/
type GetCubeSatPowerSystemsOK struct {
	Payload *models.CubeSatPowerSystems
}

// IsSuccess returns true when this get cube sat power systems o k response has a 2xx status code
func (o *GetCubeSatPowerSystemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cube sat power systems o k response has a 3xx status code
func (o *GetCubeSatPowerSystemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat power systems o k response has a 4xx status code
func (o *GetCubeSatPowerSystemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cube sat power systems o k response has a 5xx status code
func (o *GetCubeSatPowerSystemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat power systems o k response a status code equal to that given
func (o *GetCubeSatPowerSystemsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cube sat power systems o k response
func (o *GetCubeSatPowerSystemsOK) Code() int {
	return 200
}

func (o *GetCubeSatPowerSystemsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsOK %s", 200, payload)
}

func (o *GetCubeSatPowerSystemsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsOK %s", 200, payload)
}

func (o *GetCubeSatPowerSystemsOK) GetPayload() *models.CubeSatPowerSystems {
	return o.Payload
}

func (o *GetCubeSatPowerSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CubeSatPowerSystems)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatPowerSystemsBadRequest creates a GetCubeSatPowerSystemsBadRequest with default headers values
func NewGetCubeSatPowerSystemsBadRequest() *GetCubeSatPowerSystemsBadRequest {
	return &GetCubeSatPowerSystemsBadRequest{}
}

/*
GetCubeSatPowerSystemsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCubeSatPowerSystemsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat power systems bad request response has a 2xx status code
func (o *GetCubeSatPowerSystemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat power systems bad request response has a 3xx status code
func (o *GetCubeSatPowerSystemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat power systems bad request response has a 4xx status code
func (o *GetCubeSatPowerSystemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cube sat power systems bad request response has a 5xx status code
func (o *GetCubeSatPowerSystemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat power systems bad request response a status code equal to that given
func (o *GetCubeSatPowerSystemsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cube sat power systems bad request response
func (o *GetCubeSatPowerSystemsBadRequest) Code() int {
	return 400
}

func (o *GetCubeSatPowerSystemsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsBadRequest %s", 400, payload)
}

func (o *GetCubeSatPowerSystemsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsBadRequest %s", 400, payload)
}

func (o *GetCubeSatPowerSystemsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatPowerSystemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatPowerSystemsForbidden creates a GetCubeSatPowerSystemsForbidden with default headers values
func NewGetCubeSatPowerSystemsForbidden() *GetCubeSatPowerSystemsForbidden {
	return &GetCubeSatPowerSystemsForbidden{}
}

/*
GetCubeSatPowerSystemsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCubeSatPowerSystemsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat power systems forbidden response has a 2xx status code
func (o *GetCubeSatPowerSystemsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat power systems forbidden response has a 3xx status code
func (o *GetCubeSatPowerSystemsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat power systems forbidden response has a 4xx status code
func (o *GetCubeSatPowerSystemsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cube sat power systems forbidden response has a 5xx status code
func (o *GetCubeSatPowerSystemsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat power systems forbidden response a status code equal to that given
func (o *GetCubeSatPowerSystemsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cube sat power systems forbidden response
func (o *GetCubeSatPowerSystemsForbidden) Code() int {
	return 403
}

func (o *GetCubeSatPowerSystemsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsForbidden %s", 403, payload)
}

func (o *GetCubeSatPowerSystemsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsForbidden %s", 403, payload)
}

func (o *GetCubeSatPowerSystemsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatPowerSystemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatPowerSystemsUnprocessableEntity creates a GetCubeSatPowerSystemsUnprocessableEntity with default headers values
func NewGetCubeSatPowerSystemsUnprocessableEntity() *GetCubeSatPowerSystemsUnprocessableEntity {
	return &GetCubeSatPowerSystemsUnprocessableEntity{}
}

/*
GetCubeSatPowerSystemsUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetCubeSatPowerSystemsUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat power systems unprocessable entity response has a 2xx status code
func (o *GetCubeSatPowerSystemsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat power systems unprocessable entity response has a 3xx status code
func (o *GetCubeSatPowerSystemsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat power systems unprocessable entity response has a 4xx status code
func (o *GetCubeSatPowerSystemsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cube sat power systems unprocessable entity response has a 5xx status code
func (o *GetCubeSatPowerSystemsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat power systems unprocessable entity response a status code equal to that given
func (o *GetCubeSatPowerSystemsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get cube sat power systems unprocessable entity response
func (o *GetCubeSatPowerSystemsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetCubeSatPowerSystemsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsUnprocessableEntity %s", 422, payload)
}

func (o *GetCubeSatPowerSystemsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsUnprocessableEntity %s", 422, payload)
}

func (o *GetCubeSatPowerSystemsUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatPowerSystemsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatPowerSystemsInternalServerError creates a GetCubeSatPowerSystemsInternalServerError with default headers values
func NewGetCubeSatPowerSystemsInternalServerError() *GetCubeSatPowerSystemsInternalServerError {
	return &GetCubeSatPowerSystemsInternalServerError{}
}

/*
GetCubeSatPowerSystemsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCubeSatPowerSystemsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat power systems internal server error response has a 2xx status code
func (o *GetCubeSatPowerSystemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat power systems internal server error response has a 3xx status code
func (o *GetCubeSatPowerSystemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat power systems internal server error response has a 4xx status code
func (o *GetCubeSatPowerSystemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cube sat power systems internal server error response has a 5xx status code
func (o *GetCubeSatPowerSystemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cube sat power systems internal server error response a status code equal to that given
func (o *GetCubeSatPowerSystemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cube sat power systems internal server error response
func (o *GetCubeSatPowerSystemsInternalServerError) Code() int {
	return 500
}

func (o *GetCubeSatPowerSystemsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsInternalServerError %s", 500, payload)
}

func (o *GetCubeSatPowerSystemsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_power_system/available_power_systems][%d] getCubeSatPowerSystemsInternalServerError %s", 500, payload)
}

func (o *GetCubeSatPowerSystemsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatPowerSystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
