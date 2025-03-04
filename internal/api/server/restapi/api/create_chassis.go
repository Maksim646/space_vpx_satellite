// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// CreateChassisHandlerFunc turns a function with the right signature into a create chassis handler
type CreateChassisHandlerFunc func(CreateChassisParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateChassisHandlerFunc) Handle(params CreateChassisParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateChassisHandler interface for that can handle valid create chassis params
type CreateChassisHandler interface {
	Handle(CreateChassisParams, *models.Principal) middleware.Responder
}

// NewCreateChassis creates a new http.Handler for the create chassis operation
func NewCreateChassis(ctx *middleware.Context, handler CreateChassisHandler) *CreateChassis {
	return &CreateChassis{Context: ctx, Handler: handler}
}

/*
	CreateChassis swagger:route POST /chassis Chassis createChassis

Create chasis
*/
type CreateChassis struct {
	Context *middleware.Context
	Handler CreateChassisHandler
}

func (o *CreateChassis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateChassisParams()
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
