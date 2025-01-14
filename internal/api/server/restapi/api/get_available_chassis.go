// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetAvailableChassisHandlerFunc turns a function with the right signature into a get available chassis handler
type GetAvailableChassisHandlerFunc func(GetAvailableChassisParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAvailableChassisHandlerFunc) Handle(params GetAvailableChassisParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAvailableChassisHandler interface for that can handle valid get available chassis params
type GetAvailableChassisHandler interface {
	Handle(GetAvailableChassisParams, *models.Principal) middleware.Responder
}

// NewGetAvailableChassis creates a new http.Handler for the get available chassis operation
func NewGetAvailableChassis(ctx *middleware.Context, handler GetAvailableChassisHandler) *GetAvailableChassis {
	return &GetAvailableChassis{Context: ctx, Handler: handler}
}

/*
	GetAvailableChassis swagger:route GET /chassis/available_chassis Chassis getAvailableChassis

Get available chassis
*/
type GetAvailableChassis struct {
	Context *middleware.Context
	Handler GetAvailableChassisHandler
}

func (o *GetAvailableChassis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAvailableChassisParams()
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
