// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// CreatePowerSystemHandlerFunc turns a function with the right signature into a create power system handler
type CreatePowerSystemHandlerFunc func(CreatePowerSystemParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePowerSystemHandlerFunc) Handle(params CreatePowerSystemParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreatePowerSystemHandler interface for that can handle valid create power system params
type CreatePowerSystemHandler interface {
	Handle(CreatePowerSystemParams, *models.Principal) middleware.Responder
}

// NewCreatePowerSystem creates a new http.Handler for the create power system operation
func NewCreatePowerSystem(ctx *middleware.Context, handler CreatePowerSystemHandler) *CreatePowerSystem {
	return &CreatePowerSystem{Context: ctx, Handler: handler}
}

/*
	CreatePowerSystem swagger:route POST /cube_sat_power_system PowerSystem createPowerSystem

Create Power System
*/
type CreatePowerSystem struct {
	Context *middleware.Context
	Handler CreatePowerSystemHandler
}

func (o *CreatePowerSystem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreatePowerSystemParams()
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
