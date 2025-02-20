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

// CreateCubeSatVHTransceiverReader is a Reader for the CreateCubeSatVHTransceiver structure.
type CreateCubeSatVHTransceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCubeSatVHTransceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCubeSatVHTransceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCubeSatVHTransceiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCubeSatVHTransceiverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCubeSatVHTransceiverUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCubeSatVHTransceiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cube_sat_vhf_transceiver] CreateCubeSatVHTransceiver", response, response.Code())
	}
}

// NewCreateCubeSatVHTransceiverOK creates a CreateCubeSatVHTransceiverOK with default headers values
func NewCreateCubeSatVHTransceiverOK() *CreateCubeSatVHTransceiverOK {
	return &CreateCubeSatVHTransceiverOK{}
}

/*
CreateCubeSatVHTransceiverOK describes a response with status code 200, with default header values.

Create CubeSat VHF Transceiver Response
*/
type CreateCubeSatVHTransceiverOK struct {
	Payload *models.VHFTransceiver
}

// IsSuccess returns true when this create cube sat v h transceiver o k response has a 2xx status code
func (o *CreateCubeSatVHTransceiverOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cube sat v h transceiver o k response has a 3xx status code
func (o *CreateCubeSatVHTransceiverOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cube sat v h transceiver o k response has a 4xx status code
func (o *CreateCubeSatVHTransceiverOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cube sat v h transceiver o k response has a 5xx status code
func (o *CreateCubeSatVHTransceiverOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create cube sat v h transceiver o k response a status code equal to that given
func (o *CreateCubeSatVHTransceiverOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create cube sat v h transceiver o k response
func (o *CreateCubeSatVHTransceiverOK) Code() int {
	return 200
}

func (o *CreateCubeSatVHTransceiverOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverOK %s", 200, payload)
}

func (o *CreateCubeSatVHTransceiverOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverOK %s", 200, payload)
}

func (o *CreateCubeSatVHTransceiverOK) GetPayload() *models.VHFTransceiver {
	return o.Payload
}

func (o *CreateCubeSatVHTransceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VHFTransceiver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCubeSatVHTransceiverBadRequest creates a CreateCubeSatVHTransceiverBadRequest with default headers values
func NewCreateCubeSatVHTransceiverBadRequest() *CreateCubeSatVHTransceiverBadRequest {
	return &CreateCubeSatVHTransceiverBadRequest{}
}

/*
CreateCubeSatVHTransceiverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateCubeSatVHTransceiverBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create cube sat v h transceiver bad request response has a 2xx status code
func (o *CreateCubeSatVHTransceiverBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cube sat v h transceiver bad request response has a 3xx status code
func (o *CreateCubeSatVHTransceiverBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cube sat v h transceiver bad request response has a 4xx status code
func (o *CreateCubeSatVHTransceiverBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cube sat v h transceiver bad request response has a 5xx status code
func (o *CreateCubeSatVHTransceiverBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create cube sat v h transceiver bad request response a status code equal to that given
func (o *CreateCubeSatVHTransceiverBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create cube sat v h transceiver bad request response
func (o *CreateCubeSatVHTransceiverBadRequest) Code() int {
	return 400
}

func (o *CreateCubeSatVHTransceiverBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverBadRequest %s", 400, payload)
}

func (o *CreateCubeSatVHTransceiverBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverBadRequest %s", 400, payload)
}

func (o *CreateCubeSatVHTransceiverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCubeSatVHTransceiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCubeSatVHTransceiverForbidden creates a CreateCubeSatVHTransceiverForbidden with default headers values
func NewCreateCubeSatVHTransceiverForbidden() *CreateCubeSatVHTransceiverForbidden {
	return &CreateCubeSatVHTransceiverForbidden{}
}

/*
CreateCubeSatVHTransceiverForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateCubeSatVHTransceiverForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this create cube sat v h transceiver forbidden response has a 2xx status code
func (o *CreateCubeSatVHTransceiverForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cube sat v h transceiver forbidden response has a 3xx status code
func (o *CreateCubeSatVHTransceiverForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cube sat v h transceiver forbidden response has a 4xx status code
func (o *CreateCubeSatVHTransceiverForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cube sat v h transceiver forbidden response has a 5xx status code
func (o *CreateCubeSatVHTransceiverForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create cube sat v h transceiver forbidden response a status code equal to that given
func (o *CreateCubeSatVHTransceiverForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create cube sat v h transceiver forbidden response
func (o *CreateCubeSatVHTransceiverForbidden) Code() int {
	return 403
}

func (o *CreateCubeSatVHTransceiverForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverForbidden %s", 403, payload)
}

func (o *CreateCubeSatVHTransceiverForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverForbidden %s", 403, payload)
}

func (o *CreateCubeSatVHTransceiverForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCubeSatVHTransceiverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCubeSatVHTransceiverUnprocessableEntity creates a CreateCubeSatVHTransceiverUnprocessableEntity with default headers values
func NewCreateCubeSatVHTransceiverUnprocessableEntity() *CreateCubeSatVHTransceiverUnprocessableEntity {
	return &CreateCubeSatVHTransceiverUnprocessableEntity{}
}

/*
CreateCubeSatVHTransceiverUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CreateCubeSatVHTransceiverUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this create cube sat v h transceiver unprocessable entity response has a 2xx status code
func (o *CreateCubeSatVHTransceiverUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cube sat v h transceiver unprocessable entity response has a 3xx status code
func (o *CreateCubeSatVHTransceiverUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cube sat v h transceiver unprocessable entity response has a 4xx status code
func (o *CreateCubeSatVHTransceiverUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cube sat v h transceiver unprocessable entity response has a 5xx status code
func (o *CreateCubeSatVHTransceiverUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create cube sat v h transceiver unprocessable entity response a status code equal to that given
func (o *CreateCubeSatVHTransceiverUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create cube sat v h transceiver unprocessable entity response
func (o *CreateCubeSatVHTransceiverUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateCubeSatVHTransceiverUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverUnprocessableEntity %s", 422, payload)
}

func (o *CreateCubeSatVHTransceiverUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverUnprocessableEntity %s", 422, payload)
}

func (o *CreateCubeSatVHTransceiverUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCubeSatVHTransceiverUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCubeSatVHTransceiverInternalServerError creates a CreateCubeSatVHTransceiverInternalServerError with default headers values
func NewCreateCubeSatVHTransceiverInternalServerError() *CreateCubeSatVHTransceiverInternalServerError {
	return &CreateCubeSatVHTransceiverInternalServerError{}
}

/*
CreateCubeSatVHTransceiverInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateCubeSatVHTransceiverInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create cube sat v h transceiver internal server error response has a 2xx status code
func (o *CreateCubeSatVHTransceiverInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cube sat v h transceiver internal server error response has a 3xx status code
func (o *CreateCubeSatVHTransceiverInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cube sat v h transceiver internal server error response has a 4xx status code
func (o *CreateCubeSatVHTransceiverInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cube sat v h transceiver internal server error response has a 5xx status code
func (o *CreateCubeSatVHTransceiverInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create cube sat v h transceiver internal server error response a status code equal to that given
func (o *CreateCubeSatVHTransceiverInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create cube sat v h transceiver internal server error response
func (o *CreateCubeSatVHTransceiverInternalServerError) Code() int {
	return 500
}

func (o *CreateCubeSatVHTransceiverInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverInternalServerError %s", 500, payload)
}

func (o *CreateCubeSatVHTransceiverInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cube_sat_vhf_transceiver][%d] createCubeSatVHTransceiverInternalServerError %s", 500, payload)
}

func (o *CreateCubeSatVHTransceiverInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCubeSatVHTransceiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
