// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateCubeSatVhfAntennaSystemByProjectHandlerFunc turns a function with the right signature into a update cube sat vhf antenna system by project handler
type UpdateCubeSatVhfAntennaSystemByProjectHandlerFunc func(UpdateCubeSatVhfAntennaSystemByProjectParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCubeSatVhfAntennaSystemByProjectHandlerFunc) Handle(params UpdateCubeSatVhfAntennaSystemByProjectParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateCubeSatVhfAntennaSystemByProjectHandler interface for that can handle valid update cube sat vhf antenna system by project params
type UpdateCubeSatVhfAntennaSystemByProjectHandler interface {
	Handle(UpdateCubeSatVhfAntennaSystemByProjectParams, *models.Principal) middleware.Responder
}

// NewUpdateCubeSatVhfAntennaSystemByProject creates a new http.Handler for the update cube sat vhf antenna system by project operation
func NewUpdateCubeSatVhfAntennaSystemByProject(ctx *middleware.Context, handler UpdateCubeSatVhfAntennaSystemByProjectHandler) *UpdateCubeSatVhfAntennaSystemByProject {
	return &UpdateCubeSatVhfAntennaSystemByProject{Context: ctx, Handler: handler}
}

/*
	UpdateCubeSatVhfAntennaSystemByProject swagger:route PUT /projects/{project_id}/vhf_antenna_system CubeSatProject updateCubeSatVhfAntennaSystemByProject

Update the VHF Antenna System of a CubeSat Project
*/
type UpdateCubeSatVhfAntennaSystemByProject struct {
	Context *middleware.Context
	Handler UpdateCubeSatVhfAntennaSystemByProjectHandler
}

func (o *UpdateCubeSatVhfAntennaSystemByProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateCubeSatVhfAntennaSystemByProjectParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
