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
	"github.com/go-openapi/swag"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// NewUpdateCubeSatSolarPanelTopParams creates a new UpdateCubeSatSolarPanelTopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCubeSatSolarPanelTopParams() *UpdateCubeSatSolarPanelTopParams {
	return &UpdateCubeSatSolarPanelTopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCubeSatSolarPanelTopParamsWithTimeout creates a new UpdateCubeSatSolarPanelTopParams object
// with the ability to set a timeout on a request.
func NewUpdateCubeSatSolarPanelTopParamsWithTimeout(timeout time.Duration) *UpdateCubeSatSolarPanelTopParams {
	return &UpdateCubeSatSolarPanelTopParams{
		timeout: timeout,
	}
}

// NewUpdateCubeSatSolarPanelTopParamsWithContext creates a new UpdateCubeSatSolarPanelTopParams object
// with the ability to set a context for a request.
func NewUpdateCubeSatSolarPanelTopParamsWithContext(ctx context.Context) *UpdateCubeSatSolarPanelTopParams {
	return &UpdateCubeSatSolarPanelTopParams{
		Context: ctx,
	}
}

// NewUpdateCubeSatSolarPanelTopParamsWithHTTPClient creates a new UpdateCubeSatSolarPanelTopParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCubeSatSolarPanelTopParamsWithHTTPClient(client *http.Client) *UpdateCubeSatSolarPanelTopParams {
	return &UpdateCubeSatSolarPanelTopParams{
		HTTPClient: client,
	}
}

/*
UpdateCubeSatSolarPanelTopParams contains all the parameters to send to the API endpoint

	for the update cube sat solar panel top operation.

	Typically these are written to a http.Request.
*/
type UpdateCubeSatSolarPanelTopParams struct {

	/* UpdateSolarPanelTopBody.

	   Update Solar Panel Top Body
	*/
	UpdateSolarPanelTopBody *models.UpdateSolarPanelTopBody

	/* ID.

	   The ID of cube sat solar panel top
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cube sat solar panel top params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCubeSatSolarPanelTopParams) WithDefaults() *UpdateCubeSatSolarPanelTopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cube sat solar panel top params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCubeSatSolarPanelTopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) WithTimeout(timeout time.Duration) *UpdateCubeSatSolarPanelTopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) WithContext(ctx context.Context) *UpdateCubeSatSolarPanelTopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) WithHTTPClient(client *http.Client) *UpdateCubeSatSolarPanelTopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateSolarPanelTopBody adds the updateSolarPanelTopBody to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) WithUpdateSolarPanelTopBody(updateSolarPanelTopBody *models.UpdateSolarPanelTopBody) *UpdateCubeSatSolarPanelTopParams {
	o.SetUpdateSolarPanelTopBody(updateSolarPanelTopBody)
	return o
}

// SetUpdateSolarPanelTopBody adds the updateSolarPanelTopBody to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) SetUpdateSolarPanelTopBody(updateSolarPanelTopBody *models.UpdateSolarPanelTopBody) {
	o.UpdateSolarPanelTopBody = updateSolarPanelTopBody
}

// WithID adds the id to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) WithID(id int64) *UpdateCubeSatSolarPanelTopParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update cube sat solar panel top params
func (o *UpdateCubeSatSolarPanelTopParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCubeSatSolarPanelTopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateSolarPanelTopBody != nil {
		if err := r.SetBodyParam(o.UpdateSolarPanelTopBody); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
