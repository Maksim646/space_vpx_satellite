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

// DeleteBoardComputingModuleReader is a Reader for the DeleteBoardComputingModule structure.
type DeleteBoardComputingModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBoardComputingModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBoardComputingModuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteBoardComputingModuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteBoardComputingModuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteBoardComputingModuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteBoardComputingModuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteBoardComputingModuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /board_computing_module/{id}] DeleteBoardComputingModule", response, response.Code())
	}
}

// NewDeleteBoardComputingModuleNoContent creates a DeleteBoardComputingModuleNoContent with default headers values
func NewDeleteBoardComputingModuleNoContent() *DeleteBoardComputingModuleNoContent {
	return &DeleteBoardComputingModuleNoContent{}
}

/*
DeleteBoardComputingModuleNoContent describes a response with status code 204, with default header values.

Successfully deleted Board Computing Module
*/
type DeleteBoardComputingModuleNoContent struct {
}

// IsSuccess returns true when this delete board computing module no content response has a 2xx status code
func (o *DeleteBoardComputingModuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete board computing module no content response has a 3xx status code
func (o *DeleteBoardComputingModuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete board computing module no content response has a 4xx status code
func (o *DeleteBoardComputingModuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete board computing module no content response has a 5xx status code
func (o *DeleteBoardComputingModuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete board computing module no content response a status code equal to that given
func (o *DeleteBoardComputingModuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete board computing module no content response
func (o *DeleteBoardComputingModuleNoContent) Code() int {
	return 204
}

func (o *DeleteBoardComputingModuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleNoContent", 204)
}

func (o *DeleteBoardComputingModuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleNoContent", 204)
}

func (o *DeleteBoardComputingModuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBoardComputingModuleBadRequest creates a DeleteBoardComputingModuleBadRequest with default headers values
func NewDeleteBoardComputingModuleBadRequest() *DeleteBoardComputingModuleBadRequest {
	return &DeleteBoardComputingModuleBadRequest{}
}

/*
DeleteBoardComputingModuleBadRequest describes a response with status code 400, with default header values.

Bad request (e.g., invalid input)
*/
type DeleteBoardComputingModuleBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete board computing module bad request response has a 2xx status code
func (o *DeleteBoardComputingModuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete board computing module bad request response has a 3xx status code
func (o *DeleteBoardComputingModuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete board computing module bad request response has a 4xx status code
func (o *DeleteBoardComputingModuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete board computing module bad request response has a 5xx status code
func (o *DeleteBoardComputingModuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete board computing module bad request response a status code equal to that given
func (o *DeleteBoardComputingModuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete board computing module bad request response
func (o *DeleteBoardComputingModuleBadRequest) Code() int {
	return 400
}

func (o *DeleteBoardComputingModuleBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleBadRequest %s", 400, payload)
}

func (o *DeleteBoardComputingModuleBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleBadRequest %s", 400, payload)
}

func (o *DeleteBoardComputingModuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBoardComputingModuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBoardComputingModuleUnauthorized creates a DeleteBoardComputingModuleUnauthorized with default headers values
func NewDeleteBoardComputingModuleUnauthorized() *DeleteBoardComputingModuleUnauthorized {
	return &DeleteBoardComputingModuleUnauthorized{}
}

/*
DeleteBoardComputingModuleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteBoardComputingModuleUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete board computing module unauthorized response has a 2xx status code
func (o *DeleteBoardComputingModuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete board computing module unauthorized response has a 3xx status code
func (o *DeleteBoardComputingModuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete board computing module unauthorized response has a 4xx status code
func (o *DeleteBoardComputingModuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete board computing module unauthorized response has a 5xx status code
func (o *DeleteBoardComputingModuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete board computing module unauthorized response a status code equal to that given
func (o *DeleteBoardComputingModuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete board computing module unauthorized response
func (o *DeleteBoardComputingModuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteBoardComputingModuleUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleUnauthorized %s", 401, payload)
}

func (o *DeleteBoardComputingModuleUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleUnauthorized %s", 401, payload)
}

func (o *DeleteBoardComputingModuleUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBoardComputingModuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBoardComputingModuleForbidden creates a DeleteBoardComputingModuleForbidden with default headers values
func NewDeleteBoardComputingModuleForbidden() *DeleteBoardComputingModuleForbidden {
	return &DeleteBoardComputingModuleForbidden{}
}

/*
DeleteBoardComputingModuleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteBoardComputingModuleForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete board computing module forbidden response has a 2xx status code
func (o *DeleteBoardComputingModuleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete board computing module forbidden response has a 3xx status code
func (o *DeleteBoardComputingModuleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete board computing module forbidden response has a 4xx status code
func (o *DeleteBoardComputingModuleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete board computing module forbidden response has a 5xx status code
func (o *DeleteBoardComputingModuleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete board computing module forbidden response a status code equal to that given
func (o *DeleteBoardComputingModuleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete board computing module forbidden response
func (o *DeleteBoardComputingModuleForbidden) Code() int {
	return 403
}

func (o *DeleteBoardComputingModuleForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleForbidden %s", 403, payload)
}

func (o *DeleteBoardComputingModuleForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleForbidden %s", 403, payload)
}

func (o *DeleteBoardComputingModuleForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBoardComputingModuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBoardComputingModuleNotFound creates a DeleteBoardComputingModuleNotFound with default headers values
func NewDeleteBoardComputingModuleNotFound() *DeleteBoardComputingModuleNotFound {
	return &DeleteBoardComputingModuleNotFound{}
}

/*
DeleteBoardComputingModuleNotFound describes a response with status code 404, with default header values.

Board Computing Module not found
*/
type DeleteBoardComputingModuleNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete board computing module not found response has a 2xx status code
func (o *DeleteBoardComputingModuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete board computing module not found response has a 3xx status code
func (o *DeleteBoardComputingModuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete board computing module not found response has a 4xx status code
func (o *DeleteBoardComputingModuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete board computing module not found response has a 5xx status code
func (o *DeleteBoardComputingModuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete board computing module not found response a status code equal to that given
func (o *DeleteBoardComputingModuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete board computing module not found response
func (o *DeleteBoardComputingModuleNotFound) Code() int {
	return 404
}

func (o *DeleteBoardComputingModuleNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleNotFound %s", 404, payload)
}

func (o *DeleteBoardComputingModuleNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleNotFound %s", 404, payload)
}

func (o *DeleteBoardComputingModuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBoardComputingModuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBoardComputingModuleInternalServerError creates a DeleteBoardComputingModuleInternalServerError with default headers values
func NewDeleteBoardComputingModuleInternalServerError() *DeleteBoardComputingModuleInternalServerError {
	return &DeleteBoardComputingModuleInternalServerError{}
}

/*
DeleteBoardComputingModuleInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteBoardComputingModuleInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete board computing module internal server error response has a 2xx status code
func (o *DeleteBoardComputingModuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete board computing module internal server error response has a 3xx status code
func (o *DeleteBoardComputingModuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete board computing module internal server error response has a 4xx status code
func (o *DeleteBoardComputingModuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete board computing module internal server error response has a 5xx status code
func (o *DeleteBoardComputingModuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete board computing module internal server error response a status code equal to that given
func (o *DeleteBoardComputingModuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete board computing module internal server error response
func (o *DeleteBoardComputingModuleInternalServerError) Code() int {
	return 500
}

func (o *DeleteBoardComputingModuleInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleInternalServerError %s", 500, payload)
}

func (o *DeleteBoardComputingModuleInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /board_computing_module/{id}][%d] deleteBoardComputingModuleInternalServerError %s", 500, payload)
}

func (o *DeleteBoardComputingModuleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBoardComputingModuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
