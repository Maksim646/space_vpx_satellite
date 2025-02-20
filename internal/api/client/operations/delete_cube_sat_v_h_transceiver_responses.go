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

// DeleteCubeSatVHTransceiverReader is a Reader for the DeleteCubeSatVHTransceiver structure.
type DeleteCubeSatVHTransceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCubeSatVHTransceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCubeSatVHTransceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCubeSatVHTransceiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCubeSatVHTransceiverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteCubeSatVHTransceiverUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCubeSatVHTransceiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cube_sat_vhf_transceiver/{id}] DeleteCubeSatVHTransceiver", response, response.Code())
	}
}

// NewDeleteCubeSatVHTransceiverOK creates a DeleteCubeSatVHTransceiverOK with default headers values
func NewDeleteCubeSatVHTransceiverOK() *DeleteCubeSatVHTransceiverOK {
	return &DeleteCubeSatVHTransceiverOK{}
}

/*
DeleteCubeSatVHTransceiverOK describes a response with status code 200, with default header values.

Delete CubeSat VHF Transceiver Response
*/
type DeleteCubeSatVHTransceiverOK struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat v h transceiver o k response has a 2xx status code
func (o *DeleteCubeSatVHTransceiverOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cube sat v h transceiver o k response has a 3xx status code
func (o *DeleteCubeSatVHTransceiverOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat v h transceiver o k response has a 4xx status code
func (o *DeleteCubeSatVHTransceiverOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cube sat v h transceiver o k response has a 5xx status code
func (o *DeleteCubeSatVHTransceiverOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat v h transceiver o k response a status code equal to that given
func (o *DeleteCubeSatVHTransceiverOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete cube sat v h transceiver o k response
func (o *DeleteCubeSatVHTransceiverOK) Code() int {
	return 200
}

func (o *DeleteCubeSatVHTransceiverOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverOK %s", 200, payload)
}

func (o *DeleteCubeSatVHTransceiverOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverOK %s", 200, payload)
}

func (o *DeleteCubeSatVHTransceiverOK) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatVHTransceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatVHTransceiverBadRequest creates a DeleteCubeSatVHTransceiverBadRequest with default headers values
func NewDeleteCubeSatVHTransceiverBadRequest() *DeleteCubeSatVHTransceiverBadRequest {
	return &DeleteCubeSatVHTransceiverBadRequest{}
}

/*
DeleteCubeSatVHTransceiverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteCubeSatVHTransceiverBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat v h transceiver bad request response has a 2xx status code
func (o *DeleteCubeSatVHTransceiverBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat v h transceiver bad request response has a 3xx status code
func (o *DeleteCubeSatVHTransceiverBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat v h transceiver bad request response has a 4xx status code
func (o *DeleteCubeSatVHTransceiverBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cube sat v h transceiver bad request response has a 5xx status code
func (o *DeleteCubeSatVHTransceiverBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat v h transceiver bad request response a status code equal to that given
func (o *DeleteCubeSatVHTransceiverBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete cube sat v h transceiver bad request response
func (o *DeleteCubeSatVHTransceiverBadRequest) Code() int {
	return 400
}

func (o *DeleteCubeSatVHTransceiverBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverBadRequest %s", 400, payload)
}

func (o *DeleteCubeSatVHTransceiverBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverBadRequest %s", 400, payload)
}

func (o *DeleteCubeSatVHTransceiverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatVHTransceiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatVHTransceiverForbidden creates a DeleteCubeSatVHTransceiverForbidden with default headers values
func NewDeleteCubeSatVHTransceiverForbidden() *DeleteCubeSatVHTransceiverForbidden {
	return &DeleteCubeSatVHTransceiverForbidden{}
}

/*
DeleteCubeSatVHTransceiverForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCubeSatVHTransceiverForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat v h transceiver forbidden response has a 2xx status code
func (o *DeleteCubeSatVHTransceiverForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat v h transceiver forbidden response has a 3xx status code
func (o *DeleteCubeSatVHTransceiverForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat v h transceiver forbidden response has a 4xx status code
func (o *DeleteCubeSatVHTransceiverForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cube sat v h transceiver forbidden response has a 5xx status code
func (o *DeleteCubeSatVHTransceiverForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat v h transceiver forbidden response a status code equal to that given
func (o *DeleteCubeSatVHTransceiverForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete cube sat v h transceiver forbidden response
func (o *DeleteCubeSatVHTransceiverForbidden) Code() int {
	return 403
}

func (o *DeleteCubeSatVHTransceiverForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverForbidden %s", 403, payload)
}

func (o *DeleteCubeSatVHTransceiverForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverForbidden %s", 403, payload)
}

func (o *DeleteCubeSatVHTransceiverForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatVHTransceiverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatVHTransceiverUnprocessableEntity creates a DeleteCubeSatVHTransceiverUnprocessableEntity with default headers values
func NewDeleteCubeSatVHTransceiverUnprocessableEntity() *DeleteCubeSatVHTransceiverUnprocessableEntity {
	return &DeleteCubeSatVHTransceiverUnprocessableEntity{}
}

/*
DeleteCubeSatVHTransceiverUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type DeleteCubeSatVHTransceiverUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat v h transceiver unprocessable entity response has a 2xx status code
func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat v h transceiver unprocessable entity response has a 3xx status code
func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat v h transceiver unprocessable entity response has a 4xx status code
func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cube sat v h transceiver unprocessable entity response has a 5xx status code
func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat v h transceiver unprocessable entity response a status code equal to that given
func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete cube sat v h transceiver unprocessable entity response
func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverUnprocessableEntity %s", 422, payload)
}

func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverUnprocessableEntity %s", 422, payload)
}

func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatVHTransceiverUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatVHTransceiverInternalServerError creates a DeleteCubeSatVHTransceiverInternalServerError with default headers values
func NewDeleteCubeSatVHTransceiverInternalServerError() *DeleteCubeSatVHTransceiverInternalServerError {
	return &DeleteCubeSatVHTransceiverInternalServerError{}
}

/*
DeleteCubeSatVHTransceiverInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteCubeSatVHTransceiverInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat v h transceiver internal server error response has a 2xx status code
func (o *DeleteCubeSatVHTransceiverInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat v h transceiver internal server error response has a 3xx status code
func (o *DeleteCubeSatVHTransceiverInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat v h transceiver internal server error response has a 4xx status code
func (o *DeleteCubeSatVHTransceiverInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cube sat v h transceiver internal server error response has a 5xx status code
func (o *DeleteCubeSatVHTransceiverInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete cube sat v h transceiver internal server error response a status code equal to that given
func (o *DeleteCubeSatVHTransceiverInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete cube sat v h transceiver internal server error response
func (o *DeleteCubeSatVHTransceiverInternalServerError) Code() int {
	return 500
}

func (o *DeleteCubeSatVHTransceiverInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverInternalServerError %s", 500, payload)
}

func (o *DeleteCubeSatVHTransceiverInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cube_sat_vhf_transceiver/{id}][%d] deleteCubeSatVHTransceiverInternalServerError %s", 500, payload)
}

func (o *DeleteCubeSatVHTransceiverInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatVHTransceiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
