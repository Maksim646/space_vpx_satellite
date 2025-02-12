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

// DeleteCubeSatSolarPanelTopReader is a Reader for the DeleteCubeSatSolarPanelTop structure.
type DeleteCubeSatSolarPanelTopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCubeSatSolarPanelTopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCubeSatSolarPanelTopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCubeSatSolarPanelTopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCubeSatSolarPanelTopForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteCubeSatSolarPanelTopUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCubeSatSolarPanelTopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /solar_panel_top/{id}] DeleteCubeSatSolarPanelTop", response, response.Code())
	}
}

// NewDeleteCubeSatSolarPanelTopOK creates a DeleteCubeSatSolarPanelTopOK with default headers values
func NewDeleteCubeSatSolarPanelTopOK() *DeleteCubeSatSolarPanelTopOK {
	return &DeleteCubeSatSolarPanelTopOK{}
}

/*
DeleteCubeSatSolarPanelTopOK describes a response with status code 200, with default header values.

Delete Cube Sat Solar Panel Top Response
*/
type DeleteCubeSatSolarPanelTopOK struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat solar panel top o k response has a 2xx status code
func (o *DeleteCubeSatSolarPanelTopOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cube sat solar panel top o k response has a 3xx status code
func (o *DeleteCubeSatSolarPanelTopOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat solar panel top o k response has a 4xx status code
func (o *DeleteCubeSatSolarPanelTopOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cube sat solar panel top o k response has a 5xx status code
func (o *DeleteCubeSatSolarPanelTopOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat solar panel top o k response a status code equal to that given
func (o *DeleteCubeSatSolarPanelTopOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete cube sat solar panel top o k response
func (o *DeleteCubeSatSolarPanelTopOK) Code() int {
	return 200
}

func (o *DeleteCubeSatSolarPanelTopOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopOK %s", 200, payload)
}

func (o *DeleteCubeSatSolarPanelTopOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopOK %s", 200, payload)
}

func (o *DeleteCubeSatSolarPanelTopOK) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatSolarPanelTopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatSolarPanelTopBadRequest creates a DeleteCubeSatSolarPanelTopBadRequest with default headers values
func NewDeleteCubeSatSolarPanelTopBadRequest() *DeleteCubeSatSolarPanelTopBadRequest {
	return &DeleteCubeSatSolarPanelTopBadRequest{}
}

/*
DeleteCubeSatSolarPanelTopBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteCubeSatSolarPanelTopBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat solar panel top bad request response has a 2xx status code
func (o *DeleteCubeSatSolarPanelTopBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat solar panel top bad request response has a 3xx status code
func (o *DeleteCubeSatSolarPanelTopBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat solar panel top bad request response has a 4xx status code
func (o *DeleteCubeSatSolarPanelTopBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cube sat solar panel top bad request response has a 5xx status code
func (o *DeleteCubeSatSolarPanelTopBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat solar panel top bad request response a status code equal to that given
func (o *DeleteCubeSatSolarPanelTopBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete cube sat solar panel top bad request response
func (o *DeleteCubeSatSolarPanelTopBadRequest) Code() int {
	return 400
}

func (o *DeleteCubeSatSolarPanelTopBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopBadRequest %s", 400, payload)
}

func (o *DeleteCubeSatSolarPanelTopBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopBadRequest %s", 400, payload)
}

func (o *DeleteCubeSatSolarPanelTopBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatSolarPanelTopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatSolarPanelTopForbidden creates a DeleteCubeSatSolarPanelTopForbidden with default headers values
func NewDeleteCubeSatSolarPanelTopForbidden() *DeleteCubeSatSolarPanelTopForbidden {
	return &DeleteCubeSatSolarPanelTopForbidden{}
}

/*
DeleteCubeSatSolarPanelTopForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCubeSatSolarPanelTopForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat solar panel top forbidden response has a 2xx status code
func (o *DeleteCubeSatSolarPanelTopForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat solar panel top forbidden response has a 3xx status code
func (o *DeleteCubeSatSolarPanelTopForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat solar panel top forbidden response has a 4xx status code
func (o *DeleteCubeSatSolarPanelTopForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cube sat solar panel top forbidden response has a 5xx status code
func (o *DeleteCubeSatSolarPanelTopForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat solar panel top forbidden response a status code equal to that given
func (o *DeleteCubeSatSolarPanelTopForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete cube sat solar panel top forbidden response
func (o *DeleteCubeSatSolarPanelTopForbidden) Code() int {
	return 403
}

func (o *DeleteCubeSatSolarPanelTopForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopForbidden %s", 403, payload)
}

func (o *DeleteCubeSatSolarPanelTopForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopForbidden %s", 403, payload)
}

func (o *DeleteCubeSatSolarPanelTopForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatSolarPanelTopForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatSolarPanelTopUnprocessableEntity creates a DeleteCubeSatSolarPanelTopUnprocessableEntity with default headers values
func NewDeleteCubeSatSolarPanelTopUnprocessableEntity() *DeleteCubeSatSolarPanelTopUnprocessableEntity {
	return &DeleteCubeSatSolarPanelTopUnprocessableEntity{}
}

/*
DeleteCubeSatSolarPanelTopUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type DeleteCubeSatSolarPanelTopUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat solar panel top unprocessable entity response has a 2xx status code
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat solar panel top unprocessable entity response has a 3xx status code
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat solar panel top unprocessable entity response has a 4xx status code
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cube sat solar panel top unprocessable entity response has a 5xx status code
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cube sat solar panel top unprocessable entity response a status code equal to that given
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete cube sat solar panel top unprocessable entity response
func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopUnprocessableEntity %s", 422, payload)
}

func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopUnprocessableEntity %s", 422, payload)
}

func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatSolarPanelTopUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCubeSatSolarPanelTopInternalServerError creates a DeleteCubeSatSolarPanelTopInternalServerError with default headers values
func NewDeleteCubeSatSolarPanelTopInternalServerError() *DeleteCubeSatSolarPanelTopInternalServerError {
	return &DeleteCubeSatSolarPanelTopInternalServerError{}
}

/*
DeleteCubeSatSolarPanelTopInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteCubeSatSolarPanelTopInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete cube sat solar panel top internal server error response has a 2xx status code
func (o *DeleteCubeSatSolarPanelTopInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cube sat solar panel top internal server error response has a 3xx status code
func (o *DeleteCubeSatSolarPanelTopInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cube sat solar panel top internal server error response has a 4xx status code
func (o *DeleteCubeSatSolarPanelTopInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cube sat solar panel top internal server error response has a 5xx status code
func (o *DeleteCubeSatSolarPanelTopInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete cube sat solar panel top internal server error response a status code equal to that given
func (o *DeleteCubeSatSolarPanelTopInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete cube sat solar panel top internal server error response
func (o *DeleteCubeSatSolarPanelTopInternalServerError) Code() int {
	return 500
}

func (o *DeleteCubeSatSolarPanelTopInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopInternalServerError %s", 500, payload)
}

func (o *DeleteCubeSatSolarPanelTopInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /solar_panel_top/{id}][%d] deleteCubeSatSolarPanelTopInternalServerError %s", 500, payload)
}

func (o *DeleteCubeSatSolarPanelTopInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCubeSatSolarPanelTopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
