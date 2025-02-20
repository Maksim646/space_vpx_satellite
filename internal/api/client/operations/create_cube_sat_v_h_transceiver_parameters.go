// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// NewCreateCubeSatVHTransceiverParams creates a new CreateCubeSatVHTransceiverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCubeSatVHTransceiverParams() *CreateCubeSatVHTransceiverParams {
	return &CreateCubeSatVHTransceiverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCubeSatVHTransceiverParamsWithTimeout creates a new CreateCubeSatVHTransceiverParams object
// with the ability to set a timeout on a request.
func NewCreateCubeSatVHTransceiverParamsWithTimeout(timeout time.Duration) *CreateCubeSatVHTransceiverParams {
	return &CreateCubeSatVHTransceiverParams{
		timeout: timeout,
	}
}

// NewCreateCubeSatVHTransceiverParamsWithContext creates a new CreateCubeSatVHTransceiverParams object
// with the ability to set a context for a request.
func NewCreateCubeSatVHTransceiverParamsWithContext(ctx context.Context) *CreateCubeSatVHTransceiverParams {
	return &CreateCubeSatVHTransceiverParams{
		Context: ctx,
	}
}

// NewCreateCubeSatVHTransceiverParamsWithHTTPClient creates a new CreateCubeSatVHTransceiverParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCubeSatVHTransceiverParamsWithHTTPClient(client *http.Client) *CreateCubeSatVHTransceiverParams {
	return &CreateCubeSatVHTransceiverParams{
		HTTPClient: client,
	}
}

/*
CreateCubeSatVHTransceiverParams contains all the parameters to send to the API endpoint

	for the create cube sat v h transceiver operation.

	Typically these are written to a http.Request.
*/
type CreateCubeSatVHTransceiverParams struct {

	/* CreateVHFTransceiverBody.

	   Create VHF Transceiver Body
	*/
	CreateVHFTransceiverBody *models.CreateVHFTransceiverBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cube sat v h transceiver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCubeSatVHTransceiverParams) WithDefaults() *CreateCubeSatVHTransceiverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cube sat v h transceiver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCubeSatVHTransceiverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) WithTimeout(timeout time.Duration) *CreateCubeSatVHTransceiverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) WithContext(ctx context.Context) *CreateCubeSatVHTransceiverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) WithHTTPClient(client *http.Client) *CreateCubeSatVHTransceiverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateVHFTransceiverBody adds the createVHFTransceiverBody to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) WithCreateVHFTransceiverBody(createVHFTransceiverBody *models.CreateVHFTransceiverBody) *CreateCubeSatVHTransceiverParams {
	o.SetCreateVHFTransceiverBody(createVHFTransceiverBody)
	return o
}

// SetCreateVHFTransceiverBody adds the createVHFTransceiverBody to the create cube sat v h transceiver params
func (o *CreateCubeSatVHTransceiverParams) SetCreateVHFTransceiverBody(createVHFTransceiverBody *models.CreateVHFTransceiverBody) {
	o.CreateVHFTransceiverBody = createVHFTransceiverBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCubeSatVHTransceiverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateVHFTransceiverBody != nil {
		if err := r.SetBodyParam(o.CreateVHFTransceiverBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
