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

// CreateSolarPanelTopReader is a Reader for the CreateSolarPanelTop structure.
type CreateSolarPanelTopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSolarPanelTopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSolarPanelTopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSolarPanelTopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSolarPanelTopForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateSolarPanelTopUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSolarPanelTopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /solar_panel_top] CreateSolarPanelTop", response, response.Code())
	}
}

// NewCreateSolarPanelTopOK creates a CreateSolarPanelTopOK with default headers values
func NewCreateSolarPanelTopOK() *CreateSolarPanelTopOK {
	return &CreateSolarPanelTopOK{}
}

/*
CreateSolarPanelTopOK describes a response with status code 200, with default header values.

Create Solar Panel Top Response
*/
type CreateSolarPanelTopOK struct {
	Payload *models.SolarPanelTop
}

// IsSuccess returns true when this create solar panel top o k response has a 2xx status code
func (o *CreateSolarPanelTopOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create solar panel top o k response has a 3xx status code
func (o *CreateSolarPanelTopOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel top o k response has a 4xx status code
func (o *CreateSolarPanelTopOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create solar panel top o k response has a 5xx status code
func (o *CreateSolarPanelTopOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel top o k response a status code equal to that given
func (o *CreateSolarPanelTopOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create solar panel top o k response
func (o *CreateSolarPanelTopOK) Code() int {
	return 200
}

func (o *CreateSolarPanelTopOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopOK %s", 200, payload)
}

func (o *CreateSolarPanelTopOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopOK %s", 200, payload)
}

func (o *CreateSolarPanelTopOK) GetPayload() *models.SolarPanelTop {
	return o.Payload
}

func (o *CreateSolarPanelTopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SolarPanelTop)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelTopBadRequest creates a CreateSolarPanelTopBadRequest with default headers values
func NewCreateSolarPanelTopBadRequest() *CreateSolarPanelTopBadRequest {
	return &CreateSolarPanelTopBadRequest{}
}

/*
CreateSolarPanelTopBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSolarPanelTopBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel top bad request response has a 2xx status code
func (o *CreateSolarPanelTopBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel top bad request response has a 3xx status code
func (o *CreateSolarPanelTopBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel top bad request response has a 4xx status code
func (o *CreateSolarPanelTopBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create solar panel top bad request response has a 5xx status code
func (o *CreateSolarPanelTopBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel top bad request response a status code equal to that given
func (o *CreateSolarPanelTopBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create solar panel top bad request response
func (o *CreateSolarPanelTopBadRequest) Code() int {
	return 400
}

func (o *CreateSolarPanelTopBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopBadRequest %s", 400, payload)
}

func (o *CreateSolarPanelTopBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopBadRequest %s", 400, payload)
}

func (o *CreateSolarPanelTopBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelTopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelTopForbidden creates a CreateSolarPanelTopForbidden with default headers values
func NewCreateSolarPanelTopForbidden() *CreateSolarPanelTopForbidden {
	return &CreateSolarPanelTopForbidden{}
}

/*
CreateSolarPanelTopForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateSolarPanelTopForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel top forbidden response has a 2xx status code
func (o *CreateSolarPanelTopForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel top forbidden response has a 3xx status code
func (o *CreateSolarPanelTopForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel top forbidden response has a 4xx status code
func (o *CreateSolarPanelTopForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create solar panel top forbidden response has a 5xx status code
func (o *CreateSolarPanelTopForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel top forbidden response a status code equal to that given
func (o *CreateSolarPanelTopForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create solar panel top forbidden response
func (o *CreateSolarPanelTopForbidden) Code() int {
	return 403
}

func (o *CreateSolarPanelTopForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopForbidden %s", 403, payload)
}

func (o *CreateSolarPanelTopForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopForbidden %s", 403, payload)
}

func (o *CreateSolarPanelTopForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelTopForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelTopUnprocessableEntity creates a CreateSolarPanelTopUnprocessableEntity with default headers values
func NewCreateSolarPanelTopUnprocessableEntity() *CreateSolarPanelTopUnprocessableEntity {
	return &CreateSolarPanelTopUnprocessableEntity{}
}

/*
CreateSolarPanelTopUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CreateSolarPanelTopUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel top unprocessable entity response has a 2xx status code
func (o *CreateSolarPanelTopUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel top unprocessable entity response has a 3xx status code
func (o *CreateSolarPanelTopUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel top unprocessable entity response has a 4xx status code
func (o *CreateSolarPanelTopUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create solar panel top unprocessable entity response has a 5xx status code
func (o *CreateSolarPanelTopUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel top unprocessable entity response a status code equal to that given
func (o *CreateSolarPanelTopUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create solar panel top unprocessable entity response
func (o *CreateSolarPanelTopUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateSolarPanelTopUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopUnprocessableEntity %s", 422, payload)
}

func (o *CreateSolarPanelTopUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopUnprocessableEntity %s", 422, payload)
}

func (o *CreateSolarPanelTopUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelTopUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelTopInternalServerError creates a CreateSolarPanelTopInternalServerError with default headers values
func NewCreateSolarPanelTopInternalServerError() *CreateSolarPanelTopInternalServerError {
	return &CreateSolarPanelTopInternalServerError{}
}

/*
CreateSolarPanelTopInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateSolarPanelTopInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel top internal server error response has a 2xx status code
func (o *CreateSolarPanelTopInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel top internal server error response has a 3xx status code
func (o *CreateSolarPanelTopInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel top internal server error response has a 4xx status code
func (o *CreateSolarPanelTopInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create solar panel top internal server error response has a 5xx status code
func (o *CreateSolarPanelTopInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create solar panel top internal server error response a status code equal to that given
func (o *CreateSolarPanelTopInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create solar panel top internal server error response
func (o *CreateSolarPanelTopInternalServerError) Code() int {
	return 500
}

func (o *CreateSolarPanelTopInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopInternalServerError %s", 500, payload)
}

func (o *CreateSolarPanelTopInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel_top][%d] createSolarPanelTopInternalServerError %s", 500, payload)
}

func (o *CreateSolarPanelTopInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelTopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
