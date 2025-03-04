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

// NewGetAvailableChassisParams creates a new GetAvailableChassisParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAvailableChassisParams() *GetAvailableChassisParams {
	return &GetAvailableChassisParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvailableChassisParamsWithTimeout creates a new GetAvailableChassisParams object
// with the ability to set a timeout on a request.
func NewGetAvailableChassisParamsWithTimeout(timeout time.Duration) *GetAvailableChassisParams {
	return &GetAvailableChassisParams{
		timeout: timeout,
	}
}

// NewGetAvailableChassisParamsWithContext creates a new GetAvailableChassisParams object
// with the ability to set a context for a request.
func NewGetAvailableChassisParamsWithContext(ctx context.Context) *GetAvailableChassisParams {
	return &GetAvailableChassisParams{
		Context: ctx,
	}
}

// NewGetAvailableChassisParamsWithHTTPClient creates a new GetAvailableChassisParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAvailableChassisParamsWithHTTPClient(client *http.Client) *GetAvailableChassisParams {
	return &GetAvailableChassisParams{
		HTTPClient: client,
	}
}

/*
GetAvailableChassisParams contains all the parameters to send to the API endpoint

	for the get available chassis operation.

	Typically these are written to a http.Request.
*/
type GetAvailableChassisParams struct {

	/* FilterChassisByMaxHeightFrom.

	   Filter By Max Height
	*/
	FilterChassisByMaxHeightFrom *float64

	/* FilterChassisByMaxLengthFrom.

	   Filter By Max Lenghth
	*/
	FilterChassisByMaxLengthFrom *float64

	/* FilterChassisByMaxPowerHandlingCapabilityPerBoardFrom.

	   Filter By Max Power Handling Capability Per Board
	*/
	FilterChassisByMaxPowerHandlingCapabilityPerBoardFrom *float64

	/* FilterChassisByMaxTemperaturePerBoardFrom.

	   Filter By Max Temperature Per Board
	*/
	FilterChassisByMaxTemperaturePerBoardFrom *float64

	/* FilterChassisByMaxWeightFrom.

	   Filter By Max Weight
	*/
	FilterChassisByMaxWeightFrom *float64

	/* FilterChassisByMaxWidthFrom.

	   Filter By Max Width
	*/
	FilterChassisByMaxWidthFrom *float64

	/* FilterChassisByMinHeightTo.

	   Filter By Min Height
	*/
	FilterChassisByMinHeightTo *float64

	/* FilterChassisByMinLengthTo.

	   Filter By Min Length
	*/
	FilterChassisByMinLengthTo *float64

	/* FilterChassisByMinPowerHandlingCapabilityPerBoardTo.

	   Filter By Min Power Handling Capability Per Board
	*/
	FilterChassisByMinPowerHandlingCapabilityPerBoardTo *float64

	/* FilterChassisByMinTemperaturePerBoardTo.

	   Filter By Min Temperature Per Board
	*/
	FilterChassisByMinTemperaturePerBoardTo *float64

	/* FilterChassisByMinWeightTo.

	   Filter By Min Weight
	*/
	FilterChassisByMinWeightTo *float64

	/* FilterChassisByMinWidthTo.

	   Filter By Min Width
	*/
	FilterChassisByMinWidthTo *float64

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

// WithDefaults hydrates default values in the get available chassis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAvailableChassisParams) WithDefaults() *GetAvailableChassisParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get available chassis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAvailableChassisParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get available chassis params
func (o *GetAvailableChassisParams) WithTimeout(timeout time.Duration) *GetAvailableChassisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get available chassis params
func (o *GetAvailableChassisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get available chassis params
func (o *GetAvailableChassisParams) WithContext(ctx context.Context) *GetAvailableChassisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get available chassis params
func (o *GetAvailableChassisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get available chassis params
func (o *GetAvailableChassisParams) WithHTTPClient(client *http.Client) *GetAvailableChassisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get available chassis params
func (o *GetAvailableChassisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterChassisByMaxHeightFrom adds the filterChassisByMaxHeightFrom to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMaxHeightFrom(filterChassisByMaxHeightFrom *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMaxHeightFrom(filterChassisByMaxHeightFrom)
	return o
}

// SetFilterChassisByMaxHeightFrom adds the filterChassisByMaxHeightFrom to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMaxHeightFrom(filterChassisByMaxHeightFrom *float64) {
	o.FilterChassisByMaxHeightFrom = filterChassisByMaxHeightFrom
}

// WithFilterChassisByMaxLengthFrom adds the filterChassisByMaxLengthFrom to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMaxLengthFrom(filterChassisByMaxLengthFrom *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMaxLengthFrom(filterChassisByMaxLengthFrom)
	return o
}

// SetFilterChassisByMaxLengthFrom adds the filterChassisByMaxLengthFrom to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMaxLengthFrom(filterChassisByMaxLengthFrom *float64) {
	o.FilterChassisByMaxLengthFrom = filterChassisByMaxLengthFrom
}

// WithFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom adds the filterChassisByMaxPowerHandlingCapabilityPerBoardFrom to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom(filterChassisByMaxPowerHandlingCapabilityPerBoardFrom *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom(filterChassisByMaxPowerHandlingCapabilityPerBoardFrom)
	return o
}

// SetFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom adds the filterChassisByMaxPowerHandlingCapabilityPerBoardFrom to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom(filterChassisByMaxPowerHandlingCapabilityPerBoardFrom *float64) {
	o.FilterChassisByMaxPowerHandlingCapabilityPerBoardFrom = filterChassisByMaxPowerHandlingCapabilityPerBoardFrom
}

// WithFilterChassisByMaxTemperaturePerBoardFrom adds the filterChassisByMaxTemperaturePerBoardFrom to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMaxTemperaturePerBoardFrom(filterChassisByMaxTemperaturePerBoardFrom *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMaxTemperaturePerBoardFrom(filterChassisByMaxTemperaturePerBoardFrom)
	return o
}

// SetFilterChassisByMaxTemperaturePerBoardFrom adds the filterChassisByMaxTemperaturePerBoardFrom to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMaxTemperaturePerBoardFrom(filterChassisByMaxTemperaturePerBoardFrom *float64) {
	o.FilterChassisByMaxTemperaturePerBoardFrom = filterChassisByMaxTemperaturePerBoardFrom
}

// WithFilterChassisByMaxWeightFrom adds the filterChassisByMaxWeightFrom to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMaxWeightFrom(filterChassisByMaxWeightFrom *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMaxWeightFrom(filterChassisByMaxWeightFrom)
	return o
}

// SetFilterChassisByMaxWeightFrom adds the filterChassisByMaxWeightFrom to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMaxWeightFrom(filterChassisByMaxWeightFrom *float64) {
	o.FilterChassisByMaxWeightFrom = filterChassisByMaxWeightFrom
}

// WithFilterChassisByMaxWidthFrom adds the filterChassisByMaxWidthFrom to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMaxWidthFrom(filterChassisByMaxWidthFrom *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMaxWidthFrom(filterChassisByMaxWidthFrom)
	return o
}

// SetFilterChassisByMaxWidthFrom adds the filterChassisByMaxWidthFrom to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMaxWidthFrom(filterChassisByMaxWidthFrom *float64) {
	o.FilterChassisByMaxWidthFrom = filterChassisByMaxWidthFrom
}

// WithFilterChassisByMinHeightTo adds the filterChassisByMinHeightTo to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMinHeightTo(filterChassisByMinHeightTo *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMinHeightTo(filterChassisByMinHeightTo)
	return o
}

// SetFilterChassisByMinHeightTo adds the filterChassisByMinHeightTo to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMinHeightTo(filterChassisByMinHeightTo *float64) {
	o.FilterChassisByMinHeightTo = filterChassisByMinHeightTo
}

// WithFilterChassisByMinLengthTo adds the filterChassisByMinLengthTo to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMinLengthTo(filterChassisByMinLengthTo *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMinLengthTo(filterChassisByMinLengthTo)
	return o
}

// SetFilterChassisByMinLengthTo adds the filterChassisByMinLengthTo to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMinLengthTo(filterChassisByMinLengthTo *float64) {
	o.FilterChassisByMinLengthTo = filterChassisByMinLengthTo
}

// WithFilterChassisByMinPowerHandlingCapabilityPerBoardTo adds the filterChassisByMinPowerHandlingCapabilityPerBoardTo to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMinPowerHandlingCapabilityPerBoardTo(filterChassisByMinPowerHandlingCapabilityPerBoardTo *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMinPowerHandlingCapabilityPerBoardTo(filterChassisByMinPowerHandlingCapabilityPerBoardTo)
	return o
}

// SetFilterChassisByMinPowerHandlingCapabilityPerBoardTo adds the filterChassisByMinPowerHandlingCapabilityPerBoardTo to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMinPowerHandlingCapabilityPerBoardTo(filterChassisByMinPowerHandlingCapabilityPerBoardTo *float64) {
	o.FilterChassisByMinPowerHandlingCapabilityPerBoardTo = filterChassisByMinPowerHandlingCapabilityPerBoardTo
}

// WithFilterChassisByMinTemperaturePerBoardTo adds the filterChassisByMinTemperaturePerBoardTo to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMinTemperaturePerBoardTo(filterChassisByMinTemperaturePerBoardTo *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMinTemperaturePerBoardTo(filterChassisByMinTemperaturePerBoardTo)
	return o
}

// SetFilterChassisByMinTemperaturePerBoardTo adds the filterChassisByMinTemperaturePerBoardTo to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMinTemperaturePerBoardTo(filterChassisByMinTemperaturePerBoardTo *float64) {
	o.FilterChassisByMinTemperaturePerBoardTo = filterChassisByMinTemperaturePerBoardTo
}

// WithFilterChassisByMinWeightTo adds the filterChassisByMinWeightTo to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMinWeightTo(filterChassisByMinWeightTo *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMinWeightTo(filterChassisByMinWeightTo)
	return o
}

// SetFilterChassisByMinWeightTo adds the filterChassisByMinWeightTo to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMinWeightTo(filterChassisByMinWeightTo *float64) {
	o.FilterChassisByMinWeightTo = filterChassisByMinWeightTo
}

// WithFilterChassisByMinWidthTo adds the filterChassisByMinWidthTo to the get available chassis params
func (o *GetAvailableChassisParams) WithFilterChassisByMinWidthTo(filterChassisByMinWidthTo *float64) *GetAvailableChassisParams {
	o.SetFilterChassisByMinWidthTo(filterChassisByMinWidthTo)
	return o
}

// SetFilterChassisByMinWidthTo adds the filterChassisByMinWidthTo to the get available chassis params
func (o *GetAvailableChassisParams) SetFilterChassisByMinWidthTo(filterChassisByMinWidthTo *float64) {
	o.FilterChassisByMinWidthTo = filterChassisByMinWidthTo
}

// WithLimit adds the limit to the get available chassis params
func (o *GetAvailableChassisParams) WithLimit(limit int64) *GetAvailableChassisParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get available chassis params
func (o *GetAvailableChassisParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get available chassis params
func (o *GetAvailableChassisParams) WithOffset(offset int64) *GetAvailableChassisParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get available chassis params
func (o *GetAvailableChassisParams) SetOffset(offset int64) {
	o.Offset = offset
}

// WithSortField adds the sortField to the get available chassis params
func (o *GetAvailableChassisParams) WithSortField(sortField *string) *GetAvailableChassisParams {
	o.SetSortField(sortField)
	return o
}

// SetSortField adds the sortField to the get available chassis params
func (o *GetAvailableChassisParams) SetSortField(sortField *string) {
	o.SortField = sortField
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvailableChassisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterChassisByMaxHeightFrom != nil {

		// query param FilterChassisByMaxHeight[from]
		var qrFilterChassisByMaxHeightFrom float64

		if o.FilterChassisByMaxHeightFrom != nil {
			qrFilterChassisByMaxHeightFrom = *o.FilterChassisByMaxHeightFrom
		}
		qFilterChassisByMaxHeightFrom := swag.FormatFloat64(qrFilterChassisByMaxHeightFrom)
		if qFilterChassisByMaxHeightFrom != "" {

			if err := r.SetQueryParam("FilterChassisByMaxHeight[from]", qFilterChassisByMaxHeightFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMaxLengthFrom != nil {

		// query param FilterChassisByMaxLength[from]
		var qrFilterChassisByMaxLengthFrom float64

		if o.FilterChassisByMaxLengthFrom != nil {
			qrFilterChassisByMaxLengthFrom = *o.FilterChassisByMaxLengthFrom
		}
		qFilterChassisByMaxLengthFrom := swag.FormatFloat64(qrFilterChassisByMaxLengthFrom)
		if qFilterChassisByMaxLengthFrom != "" {

			if err := r.SetQueryParam("FilterChassisByMaxLength[from]", qFilterChassisByMaxLengthFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMaxPowerHandlingCapabilityPerBoardFrom != nil {

		// query param FilterChassisByMaxPowerHandlingCapabilityPerBoard[from]
		var qrFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom float64

		if o.FilterChassisByMaxPowerHandlingCapabilityPerBoardFrom != nil {
			qrFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom = *o.FilterChassisByMaxPowerHandlingCapabilityPerBoardFrom
		}
		qFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom := swag.FormatFloat64(qrFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom)
		if qFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom != "" {

			if err := r.SetQueryParam("FilterChassisByMaxPowerHandlingCapabilityPerBoard[from]", qFilterChassisByMaxPowerHandlingCapabilityPerBoardFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMaxTemperaturePerBoardFrom != nil {

		// query param FilterChassisByMaxTemperaturePerBoard[from]
		var qrFilterChassisByMaxTemperaturePerBoardFrom float64

		if o.FilterChassisByMaxTemperaturePerBoardFrom != nil {
			qrFilterChassisByMaxTemperaturePerBoardFrom = *o.FilterChassisByMaxTemperaturePerBoardFrom
		}
		qFilterChassisByMaxTemperaturePerBoardFrom := swag.FormatFloat64(qrFilterChassisByMaxTemperaturePerBoardFrom)
		if qFilterChassisByMaxTemperaturePerBoardFrom != "" {

			if err := r.SetQueryParam("FilterChassisByMaxTemperaturePerBoard[from]", qFilterChassisByMaxTemperaturePerBoardFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMaxWeightFrom != nil {

		// query param FilterChassisByMaxWeight[from]
		var qrFilterChassisByMaxWeightFrom float64

		if o.FilterChassisByMaxWeightFrom != nil {
			qrFilterChassisByMaxWeightFrom = *o.FilterChassisByMaxWeightFrom
		}
		qFilterChassisByMaxWeightFrom := swag.FormatFloat64(qrFilterChassisByMaxWeightFrom)
		if qFilterChassisByMaxWeightFrom != "" {

			if err := r.SetQueryParam("FilterChassisByMaxWeight[from]", qFilterChassisByMaxWeightFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMaxWidthFrom != nil {

		// query param FilterChassisByMaxWidth[from]
		var qrFilterChassisByMaxWidthFrom float64

		if o.FilterChassisByMaxWidthFrom != nil {
			qrFilterChassisByMaxWidthFrom = *o.FilterChassisByMaxWidthFrom
		}
		qFilterChassisByMaxWidthFrom := swag.FormatFloat64(qrFilterChassisByMaxWidthFrom)
		if qFilterChassisByMaxWidthFrom != "" {

			if err := r.SetQueryParam("FilterChassisByMaxWidth[from]", qFilterChassisByMaxWidthFrom); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMinHeightTo != nil {

		// query param FilterChassisByMinHeight[to]
		var qrFilterChassisByMinHeightTo float64

		if o.FilterChassisByMinHeightTo != nil {
			qrFilterChassisByMinHeightTo = *o.FilterChassisByMinHeightTo
		}
		qFilterChassisByMinHeightTo := swag.FormatFloat64(qrFilterChassisByMinHeightTo)
		if qFilterChassisByMinHeightTo != "" {

			if err := r.SetQueryParam("FilterChassisByMinHeight[to]", qFilterChassisByMinHeightTo); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMinLengthTo != nil {

		// query param FilterChassisByMinLength[to]
		var qrFilterChassisByMinLengthTo float64

		if o.FilterChassisByMinLengthTo != nil {
			qrFilterChassisByMinLengthTo = *o.FilterChassisByMinLengthTo
		}
		qFilterChassisByMinLengthTo := swag.FormatFloat64(qrFilterChassisByMinLengthTo)
		if qFilterChassisByMinLengthTo != "" {

			if err := r.SetQueryParam("FilterChassisByMinLength[to]", qFilterChassisByMinLengthTo); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMinPowerHandlingCapabilityPerBoardTo != nil {

		// query param FilterChassisByMinPowerHandlingCapabilityPerBoard[to]
		var qrFilterChassisByMinPowerHandlingCapabilityPerBoardTo float64

		if o.FilterChassisByMinPowerHandlingCapabilityPerBoardTo != nil {
			qrFilterChassisByMinPowerHandlingCapabilityPerBoardTo = *o.FilterChassisByMinPowerHandlingCapabilityPerBoardTo
		}
		qFilterChassisByMinPowerHandlingCapabilityPerBoardTo := swag.FormatFloat64(qrFilterChassisByMinPowerHandlingCapabilityPerBoardTo)
		if qFilterChassisByMinPowerHandlingCapabilityPerBoardTo != "" {

			if err := r.SetQueryParam("FilterChassisByMinPowerHandlingCapabilityPerBoard[to]", qFilterChassisByMinPowerHandlingCapabilityPerBoardTo); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMinTemperaturePerBoardTo != nil {

		// query param FilterChassisByMinTemperaturePerBoard[to]
		var qrFilterChassisByMinTemperaturePerBoardTo float64

		if o.FilterChassisByMinTemperaturePerBoardTo != nil {
			qrFilterChassisByMinTemperaturePerBoardTo = *o.FilterChassisByMinTemperaturePerBoardTo
		}
		qFilterChassisByMinTemperaturePerBoardTo := swag.FormatFloat64(qrFilterChassisByMinTemperaturePerBoardTo)
		if qFilterChassisByMinTemperaturePerBoardTo != "" {

			if err := r.SetQueryParam("FilterChassisByMinTemperaturePerBoard[to]", qFilterChassisByMinTemperaturePerBoardTo); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMinWeightTo != nil {

		// query param FilterChassisByMinWeight[to]
		var qrFilterChassisByMinWeightTo float64

		if o.FilterChassisByMinWeightTo != nil {
			qrFilterChassisByMinWeightTo = *o.FilterChassisByMinWeightTo
		}
		qFilterChassisByMinWeightTo := swag.FormatFloat64(qrFilterChassisByMinWeightTo)
		if qFilterChassisByMinWeightTo != "" {

			if err := r.SetQueryParam("FilterChassisByMinWeight[to]", qFilterChassisByMinWeightTo); err != nil {
				return err
			}
		}
	}

	if o.FilterChassisByMinWidthTo != nil {

		// query param FilterChassisByMinWidth[to]
		var qrFilterChassisByMinWidthTo float64

		if o.FilterChassisByMinWidthTo != nil {
			qrFilterChassisByMinWidthTo = *o.FilterChassisByMinWidthTo
		}
		qFilterChassisByMinWidthTo := swag.FormatFloat64(qrFilterChassisByMinWidthTo)
		if qFilterChassisByMinWidthTo != "" {

			if err := r.SetQueryParam("FilterChassisByMinWidth[to]", qFilterChassisByMinWidthTo); err != nil {
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
