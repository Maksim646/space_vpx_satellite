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

// NewGetCubeSatPowerSystemsParams creates a new GetCubeSatPowerSystemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCubeSatPowerSystemsParams() *GetCubeSatPowerSystemsParams {
	return &GetCubeSatPowerSystemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCubeSatPowerSystemsParamsWithTimeout creates a new GetCubeSatPowerSystemsParams object
// with the ability to set a timeout on a request.
func NewGetCubeSatPowerSystemsParamsWithTimeout(timeout time.Duration) *GetCubeSatPowerSystemsParams {
	return &GetCubeSatPowerSystemsParams{
		timeout: timeout,
	}
}

// NewGetCubeSatPowerSystemsParamsWithContext creates a new GetCubeSatPowerSystemsParams object
// with the ability to set a context for a request.
func NewGetCubeSatPowerSystemsParamsWithContext(ctx context.Context) *GetCubeSatPowerSystemsParams {
	return &GetCubeSatPowerSystemsParams{
		Context: ctx,
	}
}

// NewGetCubeSatPowerSystemsParamsWithHTTPClient creates a new GetCubeSatPowerSystemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCubeSatPowerSystemsParamsWithHTTPClient(client *http.Client) *GetCubeSatPowerSystemsParams {
	return &GetCubeSatPowerSystemsParams{
		HTTPClient: client,
	}
}

/*
GetCubeSatPowerSystemsParams contains all the parameters to send to the API endpoint

	for the get cube sat power systems operation.

	Typically these are written to a http.Request.
*/
type GetCubeSatPowerSystemsParams struct {

	/* FilterCubeSatPowerSystemByLengthMax.

	   Filter By Max Length
	*/
	FilterCubeSatPowerSystemByLengthMax *float64

	/* FilterCubeSatPowerSystemByLengthMin.

	   Filter By Min Length
	*/
	FilterCubeSatPowerSystemByLengthMin *float64

	/* Limit.

	   Offset Configs
	*/
	Limit int64

	/* Offset.

	   Offset Configs
	*/
	Offset int64

	/* SortField.

	   sort parameters
	*/
	SortField *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cube sat power systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCubeSatPowerSystemsParams) WithDefaults() *GetCubeSatPowerSystemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cube sat power systems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCubeSatPowerSystemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithTimeout(timeout time.Duration) *GetCubeSatPowerSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithContext(ctx context.Context) *GetCubeSatPowerSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithHTTPClient(client *http.Client) *GetCubeSatPowerSystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterCubeSatPowerSystemByLengthMax adds the filterCubeSatPowerSystemByLengthMax to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithFilterCubeSatPowerSystemByLengthMax(filterCubeSatPowerSystemByLengthMax *float64) *GetCubeSatPowerSystemsParams {
	o.SetFilterCubeSatPowerSystemByLengthMax(filterCubeSatPowerSystemByLengthMax)
	return o
}

// SetFilterCubeSatPowerSystemByLengthMax adds the filterCubeSatPowerSystemByLengthMax to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetFilterCubeSatPowerSystemByLengthMax(filterCubeSatPowerSystemByLengthMax *float64) {
	o.FilterCubeSatPowerSystemByLengthMax = filterCubeSatPowerSystemByLengthMax
}

// WithFilterCubeSatPowerSystemByLengthMin adds the filterCubeSatPowerSystemByLengthMin to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithFilterCubeSatPowerSystemByLengthMin(filterCubeSatPowerSystemByLengthMin *float64) *GetCubeSatPowerSystemsParams {
	o.SetFilterCubeSatPowerSystemByLengthMin(filterCubeSatPowerSystemByLengthMin)
	return o
}

// SetFilterCubeSatPowerSystemByLengthMin adds the filterCubeSatPowerSystemByLengthMin to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetFilterCubeSatPowerSystemByLengthMin(filterCubeSatPowerSystemByLengthMin *float64) {
	o.FilterCubeSatPowerSystemByLengthMin = filterCubeSatPowerSystemByLengthMin
}

// WithLimit adds the limit to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithLimit(limit int64) *GetCubeSatPowerSystemsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithOffset(offset int64) *GetCubeSatPowerSystemsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetOffset(offset int64) {
	o.Offset = offset
}

// WithSortField adds the sortField to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) WithSortField(sortField *string) *GetCubeSatPowerSystemsParams {
	o.SetSortField(sortField)
	return o
}

// SetSortField adds the sortField to the get cube sat power systems params
func (o *GetCubeSatPowerSystemsParams) SetSortField(sortField *string) {
	o.SortField = sortField
}

// WriteToRequest writes these params to a swagger request
func (o *GetCubeSatPowerSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCubeSatPowerSystemByLengthMax != nil {

		// query param FilterCubeSatPowerSystemByLength[max]
		var qrFilterCubeSatPowerSystemByLengthMax float64

		if o.FilterCubeSatPowerSystemByLengthMax != nil {
			qrFilterCubeSatPowerSystemByLengthMax = *o.FilterCubeSatPowerSystemByLengthMax
		}
		qFilterCubeSatPowerSystemByLengthMax := swag.FormatFloat64(qrFilterCubeSatPowerSystemByLengthMax)
		if qFilterCubeSatPowerSystemByLengthMax != "" {

			if err := r.SetQueryParam("FilterCubeSatPowerSystemByLength[max]", qFilterCubeSatPowerSystemByLengthMax); err != nil {
				return err
			}
		}
	}

	if o.FilterCubeSatPowerSystemByLengthMin != nil {

		// query param FilterCubeSatPowerSystemByLength[min]
		var qrFilterCubeSatPowerSystemByLengthMin float64

		if o.FilterCubeSatPowerSystemByLengthMin != nil {
			qrFilterCubeSatPowerSystemByLengthMin = *o.FilterCubeSatPowerSystemByLengthMin
		}
		qFilterCubeSatPowerSystemByLengthMin := swag.FormatFloat64(qrFilterCubeSatPowerSystemByLengthMin)
		if qFilterCubeSatPowerSystemByLengthMin != "" {

			if err := r.SetQueryParam("FilterCubeSatPowerSystemByLength[min]", qFilterCubeSatPowerSystemByLengthMin); err != nil {
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
