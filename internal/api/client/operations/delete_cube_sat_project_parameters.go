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

// NewDeleteCubeSatProjectParams creates a new DeleteCubeSatProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCubeSatProjectParams() *DeleteCubeSatProjectParams {
	return &DeleteCubeSatProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCubeSatProjectParamsWithTimeout creates a new DeleteCubeSatProjectParams object
// with the ability to set a timeout on a request.
func NewDeleteCubeSatProjectParamsWithTimeout(timeout time.Duration) *DeleteCubeSatProjectParams {
	return &DeleteCubeSatProjectParams{
		timeout: timeout,
	}
}

// NewDeleteCubeSatProjectParamsWithContext creates a new DeleteCubeSatProjectParams object
// with the ability to set a context for a request.
func NewDeleteCubeSatProjectParamsWithContext(ctx context.Context) *DeleteCubeSatProjectParams {
	return &DeleteCubeSatProjectParams{
		Context: ctx,
	}
}

// NewDeleteCubeSatProjectParamsWithHTTPClient creates a new DeleteCubeSatProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCubeSatProjectParamsWithHTTPClient(client *http.Client) *DeleteCubeSatProjectParams {
	return &DeleteCubeSatProjectParams{
		HTTPClient: client,
	}
}

/*
DeleteCubeSatProjectParams contains all the parameters to send to the API endpoint

	for the delete cube sat project operation.

	Typically these are written to a http.Request.
*/
type DeleteCubeSatProjectParams struct {

	/* ID.

	   The ID of the project to retrieve
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cube sat project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCubeSatProjectParams) WithDefaults() *DeleteCubeSatProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cube sat project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCubeSatProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) WithTimeout(timeout time.Duration) *DeleteCubeSatProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) WithContext(ctx context.Context) *DeleteCubeSatProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) WithHTTPClient(client *http.Client) *DeleteCubeSatProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) WithID(id string) *DeleteCubeSatProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete cube sat project params
func (o *DeleteCubeSatProjectParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCubeSatProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
