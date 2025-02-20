// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateVHFAntennaSystemHandlerFunc turns a function with the right signature into a update v h f antenna system handler
type UpdateVHFAntennaSystemHandlerFunc func(UpdateVHFAntennaSystemParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVHFAntennaSystemHandlerFunc) Handle(params UpdateVHFAntennaSystemParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateVHFAntennaSystemHandler interface for that can handle valid update v h f antenna system params
type UpdateVHFAntennaSystemHandler interface {
	Handle(UpdateVHFAntennaSystemParams, *models.Principal) middleware.Responder
}

// NewUpdateVHFAntennaSystem creates a new http.Handler for the update v h f antenna system operation
func NewUpdateVHFAntennaSystem(ctx *middleware.Context, handler UpdateVHFAntennaSystemHandler) *UpdateVHFAntennaSystem {
	return &UpdateVHFAntennaSystem{Context: ctx, Handler: handler}
}

/*
	UpdateVHFAntennaSystem swagger:route PATCH /vhf_antenna_system/{id} VHFAntennaSystem updateVHFAntennaSystem

Update VHF Antenna System
*/
type UpdateVHFAntennaSystem struct {
	Context *middleware.Context
	Handler UpdateVHFAntennaSystemHandler
}

func (o *UpdateVHFAntennaSystem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateVHFAntennaSystemParams()
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
