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
)

// NewValidateCubeSatProjectByProjectParams creates a new ValidateCubeSatProjectByProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateCubeSatProjectByProjectParams() *ValidateCubeSatProjectByProjectParams {
	return &ValidateCubeSatProjectByProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCubeSatProjectByProjectParamsWithTimeout creates a new ValidateCubeSatProjectByProjectParams object
// with the ability to set a timeout on a request.
func NewValidateCubeSatProjectByProjectParamsWithTimeout(timeout time.Duration) *ValidateCubeSatProjectByProjectParams {
	return &ValidateCubeSatProjectByProjectParams{
		timeout: timeout,
	}
}

// NewValidateCubeSatProjectByProjectParamsWithContext creates a new ValidateCubeSatProjectByProjectParams object
// with the ability to set a context for a request.
func NewValidateCubeSatProjectByProjectParamsWithContext(ctx context.Context) *ValidateCubeSatProjectByProjectParams {
	return &ValidateCubeSatProjectByProjectParams{
		Context: ctx,
	}
}

// NewValidateCubeSatProjectByProjectParamsWithHTTPClient creates a new ValidateCubeSatProjectByProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateCubeSatProjectByProjectParamsWithHTTPClient(client *http.Client) *ValidateCubeSatProjectByProjectParams {
	return &ValidateCubeSatProjectByProjectParams{
		HTTPClient: client,
	}
}

/*
ValidateCubeSatProjectByProjectParams contains all the parameters to send to the API endpoint

	for the validate cube sat project by project operation.

	Typically these are written to a http.Request.
*/
type ValidateCubeSatProjectByProjectParams struct {

	/* ProjectID.

	   The ID of the project to validate
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate cube sat project by project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateCubeSatProjectByProjectParams) WithDefaults() *ValidateCubeSatProjectByProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate cube sat project by project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateCubeSatProjectByProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) WithTimeout(timeout time.Duration) *ValidateCubeSatProjectByProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) WithContext(ctx context.Context) *ValidateCubeSatProjectByProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) WithHTTPClient(client *http.Client) *ValidateCubeSatProjectByProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) WithProjectID(projectID string) *ValidateCubeSatProjectByProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the validate cube sat project by project params
func (o *ValidateCubeSatProjectByProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCubeSatProjectByProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
