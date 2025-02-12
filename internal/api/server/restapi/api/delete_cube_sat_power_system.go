// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteCubeSatPowerSystemHandlerFunc turns a function with the right signature into a delete cube sat power system handler
type DeleteCubeSatPowerSystemHandlerFunc func(DeleteCubeSatPowerSystemParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCubeSatPowerSystemHandlerFunc) Handle(params DeleteCubeSatPowerSystemParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteCubeSatPowerSystemHandler interface for that can handle valid delete cube sat power system params
type DeleteCubeSatPowerSystemHandler interface {
	Handle(DeleteCubeSatPowerSystemParams, *models.Principal) middleware.Responder
}

// NewDeleteCubeSatPowerSystem creates a new http.Handler for the delete cube sat power system operation
func NewDeleteCubeSatPowerSystem(ctx *middleware.Context, handler DeleteCubeSatPowerSystemHandler) *DeleteCubeSatPowerSystem {
	return &DeleteCubeSatPowerSystem{Context: ctx, Handler: handler}
}

/*
	DeleteCubeSatPowerSystem swagger:route DELETE /cube_sat_power_system/{id} PowerSystem deleteCubeSatPowerSystem

Delete Cube Sat Power System
*/
type DeleteCubeSatPowerSystem struct {
	Context *middleware.Context
	Handler DeleteCubeSatPowerSystemHandler
}

func (o *DeleteCubeSatPowerSystem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteCubeSatPowerSystemParams()
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
