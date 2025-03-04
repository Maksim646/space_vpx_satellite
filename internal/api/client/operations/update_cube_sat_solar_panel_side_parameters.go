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

// NewUpdateCubeSatSolarPanelSideParams creates a new UpdateCubeSatSolarPanelSideParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCubeSatSolarPanelSideParams() *UpdateCubeSatSolarPanelSideParams {
	return &UpdateCubeSatSolarPanelSideParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCubeSatSolarPanelSideParamsWithTimeout creates a new UpdateCubeSatSolarPanelSideParams object
// with the ability to set a timeout on a request.
func NewUpdateCubeSatSolarPanelSideParamsWithTimeout(timeout time.Duration) *UpdateCubeSatSolarPanelSideParams {
	return &UpdateCubeSatSolarPanelSideParams{
		timeout: timeout,
	}
}

// NewUpdateCubeSatSolarPanelSideParamsWithContext creates a new UpdateCubeSatSolarPanelSideParams object
// with the ability to set a context for a request.
func NewUpdateCubeSatSolarPanelSideParamsWithContext(ctx context.Context) *UpdateCubeSatSolarPanelSideParams {
	return &UpdateCubeSatSolarPanelSideParams{
		Context: ctx,
	}
}

// NewUpdateCubeSatSolarPanelSideParamsWithHTTPClient creates a new UpdateCubeSatSolarPanelSideParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCubeSatSolarPanelSideParamsWithHTTPClient(client *http.Client) *UpdateCubeSatSolarPanelSideParams {
	return &UpdateCubeSatSolarPanelSideParams{
		HTTPClient: client,
	}
}

/*
UpdateCubeSatSolarPanelSideParams contains all the parameters to send to the API endpoint

	for the update cube sat solar panel side operation.

	Typically these are written to a http.Request.
*/
type UpdateCubeSatSolarPanelSideParams struct {

	/* UpdateSolarPanelSideBody.

	   Update Solar Panel Side Body
	*/
	UpdateSolarPanelSideBody *models.UpdateSolarPanelSideBody

	/* ID.

	   The ID of cube sat solar panel side
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cube sat solar panel side params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCubeSatSolarPanelSideParams) WithDefaults() *UpdateCubeSatSolarPanelSideParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cube sat solar panel side params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCubeSatSolarPanelSideParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) WithTimeout(timeout time.Duration) *UpdateCubeSatSolarPanelSideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) WithContext(ctx context.Context) *UpdateCubeSatSolarPanelSideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) WithHTTPClient(client *http.Client) *UpdateCubeSatSolarPanelSideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateSolarPanelSideBody adds the updateSolarPanelSideBody to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) WithUpdateSolarPanelSideBody(updateSolarPanelSideBody *models.UpdateSolarPanelSideBody) *UpdateCubeSatSolarPanelSideParams {
	o.SetUpdateSolarPanelSideBody(updateSolarPanelSideBody)
	return o
}

// SetUpdateSolarPanelSideBody adds the updateSolarPanelSideBody to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) SetUpdateSolarPanelSideBody(updateSolarPanelSideBody *models.UpdateSolarPanelSideBody) {
	o.UpdateSolarPanelSideBody = updateSolarPanelSideBody
}

// WithID adds the id to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) WithID(id int64) *UpdateCubeSatSolarPanelSideParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update cube sat solar panel side params
func (o *UpdateCubeSatSolarPanelSideParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCubeSatSolarPanelSideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateSolarPanelSideBody != nil {
		if err := r.SetBodyParam(o.UpdateSolarPanelSideBody); err != nil {
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
