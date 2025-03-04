// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// NewUpdateCubeSatFrameByProjectParams creates a new UpdateCubeSatFrameByProjectParams object
//
// There are no default values defined in the spec.
func NewUpdateCubeSatFrameByProjectParams() UpdateCubeSatFrameByProjectParams {

	return UpdateCubeSatFrameByProjectParams{}
}

// UpdateCubeSatFrameByProjectParams contains all the bound params for the update cube sat frame by project operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateCubeSatFrameByProject
type UpdateCubeSatFrameByProjectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Frame object to update
	  Required: true
	  In: body
	*/
	AddCubeSatFrame *models.AddCubeSatFrame
	/*The ID of the project to update
	  Required: true
	  In: path
	*/
	ProjectID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateCubeSatFrameByProjectParams() beforehand.
func (o *UpdateCubeSatFrameByProjectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AddCubeSatFrame
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("addCubeSatFrame", "body", ""))
			} else {
				res = append(res, errors.NewParseError("addCubeSatFrame", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.AddCubeSatFrame = &body
			}
		}
	} else {
		res = append(res, errors.Required("addCubeSatFrame", "body", ""))
	}

	rProjectID, rhkProjectID, _ := route.Params.GetOK("project_id")
	if err := o.bindProjectID(rProjectID, rhkProjectID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProjectID binds and validates parameter ProjectID from path.
func (o *UpdateCubeSatFrameByProjectParams) bindProjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ProjectID = raw

	return nil
}
