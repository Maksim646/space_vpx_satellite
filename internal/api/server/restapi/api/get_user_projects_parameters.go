// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetUserProjectsParams creates a new GetUserProjectsParams object
//
// There are no default values defined in the spec.
func NewGetUserProjectsParams() GetUserProjectsParams {

	return GetUserProjectsParams{}
}

// GetUserProjectsParams contains all the bound params for the get user projects operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetUserProjects
type GetUserProjectsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Offset Configs
	  Required: true
	  Minimum: 0
	  In: query
	*/
	Limit int64
	/*Offset Configs
	  Required: true
	  Minimum: 0
	  In: query
	*/
	Offset int64
	/*sort parametrs
	  In: query
	*/
	SortField *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUserProjectsParams() beforehand.
func (o *GetUserProjectsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortField, qhkSortField, _ := qs.GetOK("sort[field]")
	if err := o.bindSortField(qSortField, qhkSortField, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetUserProjectsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("limit", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("limit", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetUserProjectsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", o.Limit, 0, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetUserProjectsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("offset", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("offset", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetUserProjectsParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", o.Offset, 0, false); err != nil {
		return err
	}

	return nil
}

// bindSortField binds and validates parameter SortField from query.
func (o *GetUserProjectsParams) bindSortField(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.SortField = &raw

	if err := o.validateSortField(formats); err != nil {
		return err
	}

	return nil
}

// validateSortField carries on validations for parameter SortField
func (o *GetUserProjectsParams) validateSortField(formats strfmt.Registry) error {

	if err := validate.EnumCase("sort[field]", "query", *o.SortField, []interface{}{"created_at", "updated_at"}, true); err != nil {
		return err
	}

	return nil
}
