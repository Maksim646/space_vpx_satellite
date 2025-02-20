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

// GetCubeSatVHTransceiverReader is a Reader for the GetCubeSatVHTransceiver structure.
type GetCubeSatVHTransceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCubeSatVHTransceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCubeSatVHTransceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCubeSatVHTransceiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCubeSatVHTransceiverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCubeSatVHTransceiverUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCubeSatVHTransceiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cube_sat_vhf_transceiver/{id}] GetCubeSatVHTransceiver", response, response.Code())
	}
}

// NewGetCubeSatVHTransceiverOK creates a GetCubeSatVHTransceiverOK with default headers values
func NewGetCubeSatVHTransceiverOK() *GetCubeSatVHTransceiverOK {
	return &GetCubeSatVHTransceiverOK{}
}

/*
GetCubeSatVHTransceiverOK describes a response with status code 200, with default header values.

Get CubeSat VHF Transceiver Response
*/
type GetCubeSatVHTransceiverOK struct {
	Payload *models.VHFTransceiver
}

// IsSuccess returns true when this get cube sat v h transceiver o k response has a 2xx status code
func (o *GetCubeSatVHTransceiverOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cube sat v h transceiver o k response has a 3xx status code
func (o *GetCubeSatVHTransceiverOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat v h transceiver o k response has a 4xx status code
func (o *GetCubeSatVHTransceiverOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cube sat v h transceiver o k response has a 5xx status code
func (o *GetCubeSatVHTransceiverOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat v h transceiver o k response a status code equal to that given
func (o *GetCubeSatVHTransceiverOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cube sat v h transceiver o k response
func (o *GetCubeSatVHTransceiverOK) Code() int {
	return 200
}

func (o *GetCubeSatVHTransceiverOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverOK %s", 200, payload)
}

func (o *GetCubeSatVHTransceiverOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverOK %s", 200, payload)
}

func (o *GetCubeSatVHTransceiverOK) GetPayload() *models.VHFTransceiver {
	return o.Payload
}

func (o *GetCubeSatVHTransceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VHFTransceiver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatVHTransceiverBadRequest creates a GetCubeSatVHTransceiverBadRequest with default headers values
func NewGetCubeSatVHTransceiverBadRequest() *GetCubeSatVHTransceiverBadRequest {
	return &GetCubeSatVHTransceiverBadRequest{}
}

/*
GetCubeSatVHTransceiverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCubeSatVHTransceiverBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat v h transceiver bad request response has a 2xx status code
func (o *GetCubeSatVHTransceiverBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat v h transceiver bad request response has a 3xx status code
func (o *GetCubeSatVHTransceiverBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat v h transceiver bad request response has a 4xx status code
func (o *GetCubeSatVHTransceiverBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cube sat v h transceiver bad request response has a 5xx status code
func (o *GetCubeSatVHTransceiverBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat v h transceiver bad request response a status code equal to that given
func (o *GetCubeSatVHTransceiverBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cube sat v h transceiver bad request response
func (o *GetCubeSatVHTransceiverBadRequest) Code() int {
	return 400
}

func (o *GetCubeSatVHTransceiverBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverBadRequest %s", 400, payload)
}

func (o *GetCubeSatVHTransceiverBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverBadRequest %s", 400, payload)
}

func (o *GetCubeSatVHTransceiverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatVHTransceiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatVHTransceiverForbidden creates a GetCubeSatVHTransceiverForbidden with default headers values
func NewGetCubeSatVHTransceiverForbidden() *GetCubeSatVHTransceiverForbidden {
	return &GetCubeSatVHTransceiverForbidden{}
}

/*
GetCubeSatVHTransceiverForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCubeSatVHTransceiverForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat v h transceiver forbidden response has a 2xx status code
func (o *GetCubeSatVHTransceiverForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat v h transceiver forbidden response has a 3xx status code
func (o *GetCubeSatVHTransceiverForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat v h transceiver forbidden response has a 4xx status code
func (o *GetCubeSatVHTransceiverForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cube sat v h transceiver forbidden response has a 5xx status code
func (o *GetCubeSatVHTransceiverForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat v h transceiver forbidden response a status code equal to that given
func (o *GetCubeSatVHTransceiverForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cube sat v h transceiver forbidden response
func (o *GetCubeSatVHTransceiverForbidden) Code() int {
	return 403
}

func (o *GetCubeSatVHTransceiverForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverForbidden %s", 403, payload)
}

func (o *GetCubeSatVHTransceiverForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverForbidden %s", 403, payload)
}

func (o *GetCubeSatVHTransceiverForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatVHTransceiverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatVHTransceiverUnprocessableEntity creates a GetCubeSatVHTransceiverUnprocessableEntity with default headers values
func NewGetCubeSatVHTransceiverUnprocessableEntity() *GetCubeSatVHTransceiverUnprocessableEntity {
	return &GetCubeSatVHTransceiverUnprocessableEntity{}
}

/*
GetCubeSatVHTransceiverUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetCubeSatVHTransceiverUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat v h transceiver unprocessable entity response has a 2xx status code
func (o *GetCubeSatVHTransceiverUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat v h transceiver unprocessable entity response has a 3xx status code
func (o *GetCubeSatVHTransceiverUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat v h transceiver unprocessable entity response has a 4xx status code
func (o *GetCubeSatVHTransceiverUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cube sat v h transceiver unprocessable entity response has a 5xx status code
func (o *GetCubeSatVHTransceiverUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get cube sat v h transceiver unprocessable entity response a status code equal to that given
func (o *GetCubeSatVHTransceiverUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get cube sat v h transceiver unprocessable entity response
func (o *GetCubeSatVHTransceiverUnprocessableEntity) Code() int {
	return 422
}

func (o *GetCubeSatVHTransceiverUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverUnprocessableEntity %s", 422, payload)
}

func (o *GetCubeSatVHTransceiverUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverUnprocessableEntity %s", 422, payload)
}

func (o *GetCubeSatVHTransceiverUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatVHTransceiverUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCubeSatVHTransceiverInternalServerError creates a GetCubeSatVHTransceiverInternalServerError with default headers values
func NewGetCubeSatVHTransceiverInternalServerError() *GetCubeSatVHTransceiverInternalServerError {
	return &GetCubeSatVHTransceiverInternalServerError{}
}

/*
GetCubeSatVHTransceiverInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCubeSatVHTransceiverInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cube sat v h transceiver internal server error response has a 2xx status code
func (o *GetCubeSatVHTransceiverInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cube sat v h transceiver internal server error response has a 3xx status code
func (o *GetCubeSatVHTransceiverInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cube sat v h transceiver internal server error response has a 4xx status code
func (o *GetCubeSatVHTransceiverInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cube sat v h transceiver internal server error response has a 5xx status code
func (o *GetCubeSatVHTransceiverInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cube sat v h transceiver internal server error response a status code equal to that given
func (o *GetCubeSatVHTransceiverInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cube sat v h transceiver internal server error response
func (o *GetCubeSatVHTransceiverInternalServerError) Code() int {
	return 500
}

func (o *GetCubeSatVHTransceiverInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverInternalServerError %s", 500, payload)
}

func (o *GetCubeSatVHTransceiverInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cube_sat_vhf_transceiver/{id}][%d] getCubeSatVHTransceiverInternalServerError %s", 500, payload)
}

func (o *GetCubeSatVHTransceiverInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCubeSatVHTransceiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
