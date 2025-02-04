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
)

// NewGetCubeSatFramesParams creates a new GetCubeSatFramesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCubeSatFramesParams() *GetCubeSatFramesParams {
	return &GetCubeSatFramesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCubeSatFramesParamsWithTimeout creates a new GetCubeSatFramesParams object
// with the ability to set a timeout on a request.
func NewGetCubeSatFramesParamsWithTimeout(timeout time.Duration) *GetCubeSatFramesParams {
	return &GetCubeSatFramesParams{
		timeout: timeout,
	}
}

// NewGetCubeSatFramesParamsWithContext creates a new GetCubeSatFramesParams object
// with the ability to set a context for a request.
func NewGetCubeSatFramesParamsWithContext(ctx context.Context) *GetCubeSatFramesParams {
	return &GetCubeSatFramesParams{
		Context: ctx,
	}
}

// NewGetCubeSatFramesParamsWithHTTPClient creates a new GetCubeSatFramesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCubeSatFramesParamsWithHTTPClient(client *http.Client) *GetCubeSatFramesParams {
	return &GetCubeSatFramesParams{
		HTTPClient: client,
	}
}

/*
GetCubeSatFramesParams contains all the parameters to send to the API endpoint

	for the get cube sat frames operation.

	Typically these are written to a http.Request.
*/
type GetCubeSatFramesParams struct {

	/* FilterCubeSatFrameByMinLengthMax.

	   Filter By Max Length
	*/
	FilterCubeSatFrameByMinLengthMax *float64

	/* FilterCubeSatFrameByMinLengthMin.

	   Filter By Min Length
	*/
	FilterCubeSatFrameByMinLengthMin *float64

	/* Limit.

	   Offset Configs
	*/
	Limit int64

	/* Offset.

	   Offset Configs
	*/
	Offset int64

	/* SortField.

	   sort parametrs
	*/
	SortField *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cube sat frames params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCubeSatFramesParams) WithDefaults() *GetCubeSatFramesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cube sat frames params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCubeSatFramesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithTimeout(timeout time.Duration) *GetCubeSatFramesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithContext(ctx context.Context) *GetCubeSatFramesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithHTTPClient(client *http.Client) *GetCubeSatFramesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterCubeSatFrameByMinLengthMax adds the filterCubeSatFrameByMinLengthMax to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithFilterCubeSatFrameByMinLengthMax(filterCubeSatFrameByMinLengthMax *float64) *GetCubeSatFramesParams {
	o.SetFilterCubeSatFrameByMinLengthMax(filterCubeSatFrameByMinLengthMax)
	return o
}

// SetFilterCubeSatFrameByMinLengthMax adds the filterCubeSatFrameByMinLengthMax to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetFilterCubeSatFrameByMinLengthMax(filterCubeSatFrameByMinLengthMax *float64) {
	o.FilterCubeSatFrameByMinLengthMax = filterCubeSatFrameByMinLengthMax
}

// WithFilterCubeSatFrameByMinLengthMin adds the filterCubeSatFrameByMinLengthMin to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithFilterCubeSatFrameByMinLengthMin(filterCubeSatFrameByMinLengthMin *float64) *GetCubeSatFramesParams {
	o.SetFilterCubeSatFrameByMinLengthMin(filterCubeSatFrameByMinLengthMin)
	return o
}

// SetFilterCubeSatFrameByMinLengthMin adds the filterCubeSatFrameByMinLengthMin to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetFilterCubeSatFrameByMinLengthMin(filterCubeSatFrameByMinLengthMin *float64) {
	o.FilterCubeSatFrameByMinLengthMin = filterCubeSatFrameByMinLengthMin
}

// WithLimit adds the limit to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithLimit(limit int64) *GetCubeSatFramesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithOffset(offset int64) *GetCubeSatFramesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetOffset(offset int64) {
	o.Offset = offset
}

// WithSortField adds the sortField to the get cube sat frames params
func (o *GetCubeSatFramesParams) WithSortField(sortField *string) *GetCubeSatFramesParams {
	o.SetSortField(sortField)
	return o
}

// SetSortField adds the sortField to the get cube sat frames params
func (o *GetCubeSatFramesParams) SetSortField(sortField *string) {
	o.SortField = sortField
}

// WriteToRequest writes these params to a swagger request
func (o *GetCubeSatFramesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCubeSatFrameByMinLengthMax != nil {

		// query param FilterCubeSatFrameByMinLength[max]
		var qrFilterCubeSatFrameByMinLengthMax float64

		if o.FilterCubeSatFrameByMinLengthMax != nil {
			qrFilterCubeSatFrameByMinLengthMax = *o.FilterCubeSatFrameByMinLengthMax
		}
		qFilterCubeSatFrameByMinLengthMax := swag.FormatFloat64(qrFilterCubeSatFrameByMinLengthMax)
		if qFilterCubeSatFrameByMinLengthMax != "" {

			if err := r.SetQueryParam("FilterCubeSatFrameByMinLength[max]", qFilterCubeSatFrameByMinLengthMax); err != nil {
				return err
			}
		}
	}

	if o.FilterCubeSatFrameByMinLengthMin != nil {

		// query param FilterCubeSatFrameByMinLength[min]
		var qrFilterCubeSatFrameByMinLengthMin float64

		if o.FilterCubeSatFrameByMinLengthMin != nil {
			qrFilterCubeSatFrameByMinLengthMin = *o.FilterCubeSatFrameByMinLengthMin
		}
		qFilterCubeSatFrameByMinLengthMin := swag.FormatFloat64(qrFilterCubeSatFrameByMinLengthMin)
		if qFilterCubeSatFrameByMinLengthMin != "" {

			if err := r.SetQueryParam("FilterCubeSatFrameByMinLength[min]", qFilterCubeSatFrameByMinLengthMin); err != nil {
				return err
			}
		}
	}

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)
	if qLimit != "" {

		if err := r.SetQueryParam("limit", qLimit); err != nil {
			return err
		}
	}

	// query param offset
	qrOffset := o.Offset
	qOffset := swag.FormatInt64(qrOffset)
	if qOffset != "" {

		if err := r.SetQueryParam("offset", qOffset); err != nil {
			return err
		}
	}

	if o.SortField != nil {

		// query param sort[field]
		var qrSortField string

		if o.SortField != nil {
			qrSortField = *o.SortField
		}
		qSortField := qrSortField
		if qSortField != "" {

			if err := r.SetQueryParam("sort[field]", qSortField); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
