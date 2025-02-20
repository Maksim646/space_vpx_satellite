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

// GetVHFAntennaSystemReader is a Reader for the GetVHFAntennaSystem structure.
type GetVHFAntennaSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVHFAntennaSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVHFAntennaSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVHFAntennaSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVHFAntennaSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetVHFAntennaSystemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVHFAntennaSystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /vhf_antenna_system/{id}] GetVHFAntennaSystem", response, response.Code())
	}
}

// NewGetVHFAntennaSystemOK creates a GetVHFAntennaSystemOK with default headers values
func NewGetVHFAntennaSystemOK() *GetVHFAntennaSystemOK {
	return &GetVHFAntennaSystemOK{}
}

/*
GetVHFAntennaSystemOK describes a response with status code 200, with default header values.

Get VHF Antenna System Response
*/
type GetVHFAntennaSystemOK struct {
	Payload *models.VHFAntennaSystem
}

// IsSuccess returns true when this get v h f antenna system o k response has a 2xx status code
func (o *GetVHFAntennaSystemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v h f antenna system o k response has a 3xx status code
func (o *GetVHFAntennaSystemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v h f antenna system o k response has a 4xx status code
func (o *GetVHFAntennaSystemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v h f antenna system o k response has a 5xx status code
func (o *GetVHFAntennaSystemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v h f antenna system o k response a status code equal to that given
func (o *GetVHFAntennaSystemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v h f antenna system o k response
func (o *GetVHFAntennaSystemOK) Code() int {
	return 200
}

func (o *GetVHFAntennaSystemOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemOK %s", 200, payload)
}

func (o *GetVHFAntennaSystemOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemOK %s", 200, payload)
}

func (o *GetVHFAntennaSystemOK) GetPayload() *models.VHFAntennaSystem {
	return o.Payload
}

func (o *GetVHFAntennaSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VHFAntennaSystem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVHFAntennaSystemBadRequest creates a GetVHFAntennaSystemBadRequest with default headers values
func NewGetVHFAntennaSystemBadRequest() *GetVHFAntennaSystemBadRequest {
	return &GetVHFAntennaSystemBadRequest{}
}

/*
GetVHFAntennaSystemBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVHFAntennaSystemBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get v h f antenna system bad request response has a 2xx status code
func (o *GetVHFAntennaSystemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v h f antenna system bad request response has a 3xx status code
func (o *GetVHFAntennaSystemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v h f antenna system bad request response has a 4xx status code
func (o *GetVHFAntennaSystemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v h f antenna system bad request response has a 5xx status code
func (o *GetVHFAntennaSystemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v h f antenna system bad request response a status code equal to that given
func (o *GetVHFAntennaSystemBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v h f antenna system bad request response
func (o *GetVHFAntennaSystemBadRequest) Code() int {
	return 400
}

func (o *GetVHFAntennaSystemBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemBadRequest %s", 400, payload)
}

func (o *GetVHFAntennaSystemBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemBadRequest %s", 400, payload)
}

func (o *GetVHFAntennaSystemBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVHFAntennaSystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVHFAntennaSystemForbidden creates a GetVHFAntennaSystemForbidden with default headers values
func NewGetVHFAntennaSystemForbidden() *GetVHFAntennaSystemForbidden {
	return &GetVHFAntennaSystemForbidden{}
}

/*
GetVHFAntennaSystemForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVHFAntennaSystemForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get v h f antenna system forbidden response has a 2xx status code
func (o *GetVHFAntennaSystemForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v h f antenna system forbidden response has a 3xx status code
func (o *GetVHFAntennaSystemForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v h f antenna system forbidden response has a 4xx status code
func (o *GetVHFAntennaSystemForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v h f antenna system forbidden response has a 5xx status code
func (o *GetVHFAntennaSystemForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v h f antenna system forbidden response a status code equal to that given
func (o *GetVHFAntennaSystemForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v h f antenna system forbidden response
func (o *GetVHFAntennaSystemForbidden) Code() int {
	return 403
}

func (o *GetVHFAntennaSystemForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemForbidden %s", 403, payload)
}

func (o *GetVHFAntennaSystemForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemForbidden %s", 403, payload)
}

func (o *GetVHFAntennaSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVHFAntennaSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVHFAntennaSystemUnprocessableEntity creates a GetVHFAntennaSystemUnprocessableEntity with default headers values
func NewGetVHFAntennaSystemUnprocessableEntity() *GetVHFAntennaSystemUnprocessableEntity {
	return &GetVHFAntennaSystemUnprocessableEntity{}
}

/*
GetVHFAntennaSystemUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetVHFAntennaSystemUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get v h f antenna system unprocessable entity response has a 2xx status code
func (o *GetVHFAntennaSystemUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v h f antenna system unprocessable entity response has a 3xx status code
func (o *GetVHFAntennaSystemUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v h f antenna system unprocessable entity response has a 4xx status code
func (o *GetVHFAntennaSystemUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v h f antenna system unprocessable entity response has a 5xx status code
func (o *GetVHFAntennaSystemUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get v h f antenna system unprocessable entity response a status code equal to that given
func (o *GetVHFAntennaSystemUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get v h f antenna system unprocessable entity response
func (o *GetVHFAntennaSystemUnprocessableEntity) Code() int {
	return 422
}

func (o *GetVHFAntennaSystemUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemUnprocessableEntity %s", 422, payload)
}

func (o *GetVHFAntennaSystemUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemUnprocessableEntity %s", 422, payload)
}

func (o *GetVHFAntennaSystemUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVHFAntennaSystemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVHFAntennaSystemInternalServerError creates a GetVHFAntennaSystemInternalServerError with default headers values
func NewGetVHFAntennaSystemInternalServerError() *GetVHFAntennaSystemInternalServerError {
	return &GetVHFAntennaSystemInternalServerError{}
}

/*
GetVHFAntennaSystemInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetVHFAntennaSystemInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get v h f antenna system internal server error response has a 2xx status code
func (o *GetVHFAntennaSystemInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v h f antenna system internal server error response has a 3xx status code
func (o *GetVHFAntennaSystemInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v h f antenna system internal server error response has a 4xx status code
func (o *GetVHFAntennaSystemInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v h f antenna system internal server error response has a 5xx status code
func (o *GetVHFAntennaSystemInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v h f antenna system internal server error response a status code equal to that given
func (o *GetVHFAntennaSystemInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v h f antenna system internal server error response
func (o *GetVHFAntennaSystemInternalServerError) Code() int {
	return 500
}

func (o *GetVHFAntennaSystemInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemInternalServerError %s", 500, payload)
}

func (o *GetVHFAntennaSystemInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vhf_antenna_system/{id}][%d] getVHFAntennaSystemInternalServerError %s", 500, payload)
}

func (o *GetVHFAntennaSystemInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVHFAntennaSystemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
