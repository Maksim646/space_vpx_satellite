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

// CreateSolarPanelReader is a Reader for the CreateSolarPanel structure.
type CreateSolarPanelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSolarPanelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSolarPanelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSolarPanelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSolarPanelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateSolarPanelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSolarPanelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /solar_panel] CreateSolarPanel", response, response.Code())
	}
}

// NewCreateSolarPanelOK creates a CreateSolarPanelOK with default headers values
func NewCreateSolarPanelOK() *CreateSolarPanelOK {
	return &CreateSolarPanelOK{}
}

/*
CreateSolarPanelOK describes a response with status code 200, with default header values.

Create Solar Panel Response
*/
type CreateSolarPanelOK struct {
	Payload *models.SolarPanel
}

// IsSuccess returns true when this create solar panel o k response has a 2xx status code
func (o *CreateSolarPanelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create solar panel o k response has a 3xx status code
func (o *CreateSolarPanelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel o k response has a 4xx status code
func (o *CreateSolarPanelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create solar panel o k response has a 5xx status code
func (o *CreateSolarPanelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel o k response a status code equal to that given
func (o *CreateSolarPanelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create solar panel o k response
func (o *CreateSolarPanelOK) Code() int {
	return 200
}

func (o *CreateSolarPanelOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelOK %s", 200, payload)
}

func (o *CreateSolarPanelOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelOK %s", 200, payload)
}

func (o *CreateSolarPanelOK) GetPayload() *models.SolarPanel {
	return o.Payload
}

func (o *CreateSolarPanelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SolarPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelBadRequest creates a CreateSolarPanelBadRequest with default headers values
func NewCreateSolarPanelBadRequest() *CreateSolarPanelBadRequest {
	return &CreateSolarPanelBadRequest{}
}

/*
CreateSolarPanelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSolarPanelBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel bad request response has a 2xx status code
func (o *CreateSolarPanelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel bad request response has a 3xx status code
func (o *CreateSolarPanelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel bad request response has a 4xx status code
func (o *CreateSolarPanelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create solar panel bad request response has a 5xx status code
func (o *CreateSolarPanelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel bad request response a status code equal to that given
func (o *CreateSolarPanelBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create solar panel bad request response
func (o *CreateSolarPanelBadRequest) Code() int {
	return 400
}

func (o *CreateSolarPanelBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelBadRequest %s", 400, payload)
}

func (o *CreateSolarPanelBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelBadRequest %s", 400, payload)
}

func (o *CreateSolarPanelBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelForbidden creates a CreateSolarPanelForbidden with default headers values
func NewCreateSolarPanelForbidden() *CreateSolarPanelForbidden {
	return &CreateSolarPanelForbidden{}
}

/*
CreateSolarPanelForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateSolarPanelForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel forbidden response has a 2xx status code
func (o *CreateSolarPanelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel forbidden response has a 3xx status code
func (o *CreateSolarPanelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel forbidden response has a 4xx status code
func (o *CreateSolarPanelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create solar panel forbidden response has a 5xx status code
func (o *CreateSolarPanelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel forbidden response a status code equal to that given
func (o *CreateSolarPanelForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create solar panel forbidden response
func (o *CreateSolarPanelForbidden) Code() int {
	return 403
}

func (o *CreateSolarPanelForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelForbidden %s", 403, payload)
}

func (o *CreateSolarPanelForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelForbidden %s", 403, payload)
}

func (o *CreateSolarPanelForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelUnprocessableEntity creates a CreateSolarPanelUnprocessableEntity with default headers values
func NewCreateSolarPanelUnprocessableEntity() *CreateSolarPanelUnprocessableEntity {
	return &CreateSolarPanelUnprocessableEntity{}
}

/*
CreateSolarPanelUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CreateSolarPanelUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel unprocessable entity response has a 2xx status code
func (o *CreateSolarPanelUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel unprocessable entity response has a 3xx status code
func (o *CreateSolarPanelUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel unprocessable entity response has a 4xx status code
func (o *CreateSolarPanelUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create solar panel unprocessable entity response has a 5xx status code
func (o *CreateSolarPanelUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create solar panel unprocessable entity response a status code equal to that given
func (o *CreateSolarPanelUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create solar panel unprocessable entity response
func (o *CreateSolarPanelUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateSolarPanelUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelUnprocessableEntity %s", 422, payload)
}

func (o *CreateSolarPanelUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelUnprocessableEntity %s", 422, payload)
}

func (o *CreateSolarPanelUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSolarPanelInternalServerError creates a CreateSolarPanelInternalServerError with default headers values
func NewCreateSolarPanelInternalServerError() *CreateSolarPanelInternalServerError {
	return &CreateSolarPanelInternalServerError{}
}

/*
CreateSolarPanelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateSolarPanelInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create solar panel internal server error response has a 2xx status code
func (o *CreateSolarPanelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create solar panel internal server error response has a 3xx status code
func (o *CreateSolarPanelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create solar panel internal server error response has a 4xx status code
func (o *CreateSolarPanelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create solar panel internal server error response has a 5xx status code
func (o *CreateSolarPanelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create solar panel internal server error response a status code equal to that given
func (o *CreateSolarPanelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create solar panel internal server error response
func (o *CreateSolarPanelInternalServerError) Code() int {
	return 500
}

func (o *CreateSolarPanelInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelInternalServerError %s", 500, payload)
}

func (o *CreateSolarPanelInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /solar_panel][%d] createSolarPanelInternalServerError %s", 500, payload)
}

func (o *CreateSolarPanelInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSolarPanelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
