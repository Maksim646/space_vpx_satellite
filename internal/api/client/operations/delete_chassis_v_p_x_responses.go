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

// DeleteChassisVPXReader is a Reader for the DeleteChassisVPX structure.
type DeleteChassisVPXReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteChassisVPXReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteChassisVPXOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteChassisVPXForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteChassisVPXNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteChassisVPXInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /chassis_vpx/{id}] DeleteChassisVPX", response, response.Code())
	}
}

// NewDeleteChassisVPXOK creates a DeleteChassisVPXOK with default headers values
func NewDeleteChassisVPXOK() *DeleteChassisVPXOK {
	return &DeleteChassisVPXOK{}
}

/*
DeleteChassisVPXOK describes a response with status code 200, with default header values.

Chassis VPX deleted successfully
*/
type DeleteChassisVPXOK struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete chassis v p x o k response has a 2xx status code
func (o *DeleteChassisVPXOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete chassis v p x o k response has a 3xx status code
func (o *DeleteChassisVPXOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete chassis v p x o k response has a 4xx status code
func (o *DeleteChassisVPXOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete chassis v p x o k response has a 5xx status code
func (o *DeleteChassisVPXOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete chassis v p x o k response a status code equal to that given
func (o *DeleteChassisVPXOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete chassis v p x o k response
func (o *DeleteChassisVPXOK) Code() int {
	return 200
}

func (o *DeleteChassisVPXOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXOK %s", 200, payload)
}

func (o *DeleteChassisVPXOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXOK %s", 200, payload)
}

func (o *DeleteChassisVPXOK) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteChassisVPXOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChassisVPXForbidden creates a DeleteChassisVPXForbidden with default headers values
func NewDeleteChassisVPXForbidden() *DeleteChassisVPXForbidden {
	return &DeleteChassisVPXForbidden{}
}

/*
DeleteChassisVPXForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteChassisVPXForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete chassis v p x forbidden response has a 2xx status code
func (o *DeleteChassisVPXForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete chassis v p x forbidden response has a 3xx status code
func (o *DeleteChassisVPXForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete chassis v p x forbidden response has a 4xx status code
func (o *DeleteChassisVPXForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete chassis v p x forbidden response has a 5xx status code
func (o *DeleteChassisVPXForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete chassis v p x forbidden response a status code equal to that given
func (o *DeleteChassisVPXForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete chassis v p x forbidden response
func (o *DeleteChassisVPXForbidden) Code() int {
	return 403
}

func (o *DeleteChassisVPXForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXForbidden %s", 403, payload)
}

func (o *DeleteChassisVPXForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXForbidden %s", 403, payload)
}

func (o *DeleteChassisVPXForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteChassisVPXForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChassisVPXNotFound creates a DeleteChassisVPXNotFound with default headers values
func NewDeleteChassisVPXNotFound() *DeleteChassisVPXNotFound {
	return &DeleteChassisVPXNotFound{}
}

/*
DeleteChassisVPXNotFound describes a response with status code 404, with default header values.

Chassis VPX not found
*/
type DeleteChassisVPXNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete chassis v p x not found response has a 2xx status code
func (o *DeleteChassisVPXNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete chassis v p x not found response has a 3xx status code
func (o *DeleteChassisVPXNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete chassis v p x not found response has a 4xx status code
func (o *DeleteChassisVPXNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete chassis v p x not found response has a 5xx status code
func (o *DeleteChassisVPXNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete chassis v p x not found response a status code equal to that given
func (o *DeleteChassisVPXNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete chassis v p x not found response
func (o *DeleteChassisVPXNotFound) Code() int {
	return 404
}

func (o *DeleteChassisVPXNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXNotFound %s", 404, payload)
}

func (o *DeleteChassisVPXNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXNotFound %s", 404, payload)
}

func (o *DeleteChassisVPXNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteChassisVPXNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChassisVPXInternalServerError creates a DeleteChassisVPXInternalServerError with default headers values
func NewDeleteChassisVPXInternalServerError() *DeleteChassisVPXInternalServerError {
	return &DeleteChassisVPXInternalServerError{}
}

/*
DeleteChassisVPXInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteChassisVPXInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete chassis v p x internal server error response has a 2xx status code
func (o *DeleteChassisVPXInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete chassis v p x internal server error response has a 3xx status code
func (o *DeleteChassisVPXInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete chassis v p x internal server error response has a 4xx status code
func (o *DeleteChassisVPXInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete chassis v p x internal server error response has a 5xx status code
func (o *DeleteChassisVPXInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete chassis v p x internal server error response a status code equal to that given
func (o *DeleteChassisVPXInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete chassis v p x internal server error response
func (o *DeleteChassisVPXInternalServerError) Code() int {
	return 500
}

func (o *DeleteChassisVPXInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXInternalServerError %s", 500, payload)
}

func (o *DeleteChassisVPXInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /chassis_vpx/{id}][%d] deleteChassisVPXInternalServerError %s", 500, payload)
}

func (o *DeleteChassisVPXInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteChassisVPXInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
