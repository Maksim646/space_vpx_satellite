// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetCubeSatSolarPanelsTopURL generates an URL for the get cube sat solar panels top operation
type GetCubeSatSolarPanelsTopURL struct {
	FilterCubeSatSolarPanelByLengthMax *float64
	FilterCubeSatSolarPanelByLengthMin *float64
	Limit                              int64
	Offset                             int64
	SortField                          *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCubeSatSolarPanelsTopURL) WithBasePath(bp string) *GetCubeSatSolarPanelsTopURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCubeSatSolarPanelsTopURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetCubeSatSolarPanelsTopURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/solar_panel_top/available_solar_panel_top"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var filterCubeSatSolarPanelByLengthMaxQ string
	if o.FilterCubeSatSolarPanelByLengthMax != nil {
		filterCubeSatSolarPanelByLengthMaxQ = swag.FormatFloat64(*o.FilterCubeSatSolarPanelByLengthMax)
	}
	if filterCubeSatSolarPanelByLengthMaxQ != "" {
		qs.Set("FilterCubeSatSolarPanelByLength[max]", filterCubeSatSolarPanelByLengthMaxQ)
	}

	var filterCubeSatSolarPanelByLengthMinQ string
	if o.FilterCubeSatSolarPanelByLengthMin != nil {
		filterCubeSatSolarPanelByLengthMinQ = swag.FormatFloat64(*o.FilterCubeSatSolarPanelByLengthMin)
	}
	if filterCubeSatSolarPanelByLengthMinQ != "" {
		qs.Set("FilterCubeSatSolarPanelByLength[min]", filterCubeSatSolarPanelByLengthMinQ)
	}

	limitQ := swag.FormatInt64(o.Limit)
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	offsetQ := swag.FormatInt64(o.Offset)
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var sortFieldQ string
	if o.SortField != nil {
		sortFieldQ = *o.SortField
	}
	if sortFieldQ != "" {
		qs.Set("sort[field]", sortFieldQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetCubeSatSolarPanelsTopURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetCubeSatSolarPanelsTopURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetCubeSatSolarPanelsTopURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetCubeSatSolarPanelsTopURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetCubeSatSolarPanelsTopURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetCubeSatSolarPanelsTopURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
