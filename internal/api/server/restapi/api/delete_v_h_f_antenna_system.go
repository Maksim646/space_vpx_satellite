// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteVHFAntennaSystemHandlerFunc turns a function with the right signature into a delete v h f antenna system handler
type DeleteVHFAntennaSystemHandlerFunc func(DeleteVHFAntennaSystemParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteVHFAntennaSystemHandlerFunc) Handle(params DeleteVHFAntennaSystemParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteVHFAntennaSystemHandler interface for that can handle valid delete v h f antenna system params
type DeleteVHFAntennaSystemHandler interface {
	Handle(DeleteVHFAntennaSystemParams, *models.Principal) middleware.Responder
}

// NewDeleteVHFAntennaSystem creates a new http.Handler for the delete v h f antenna system operation
func NewDeleteVHFAntennaSystem(ctx *middleware.Context, handler DeleteVHFAntennaSystemHandler) *DeleteVHFAntennaSystem {
	return &DeleteVHFAntennaSystem{Context: ctx, Handler: handler}
}

/*
	DeleteVHFAntennaSystem swagger:route DELETE /vhf_antenna_system/{id} VHFAntennaSystem deleteVHFAntennaSystem

Delete VHF Antenna System
*/
type DeleteVHFAntennaSystem struct {
	Context *middleware.Context
	Handler DeleteVHFAntennaSystemHandler
}

func (o *DeleteVHFAntennaSystem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteVHFAntennaSystemParams()
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
