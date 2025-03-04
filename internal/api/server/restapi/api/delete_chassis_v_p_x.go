// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteChassisVPXHandlerFunc turns a function with the right signature into a delete chassis v p x handler
type DeleteChassisVPXHandlerFunc func(DeleteChassisVPXParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteChassisVPXHandlerFunc) Handle(params DeleteChassisVPXParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteChassisVPXHandler interface for that can handle valid delete chassis v p x params
type DeleteChassisVPXHandler interface {
	Handle(DeleteChassisVPXParams, *models.Principal) middleware.Responder
}

// NewDeleteChassisVPX creates a new http.Handler for the delete chassis v p x operation
func NewDeleteChassisVPX(ctx *middleware.Context, handler DeleteChassisVPXHandler) *DeleteChassisVPX {
	return &DeleteChassisVPX{Context: ctx, Handler: handler}
}

/*
	DeleteChassisVPX swagger:route DELETE /chassis_vpx/{id} ChassisVPX deleteChassisVPX

Delete chassis VPX by ID
*/
type DeleteChassisVPX struct {
	Context *middleware.Context
	Handler DeleteChassisVPXHandler
}

func (o *DeleteChassisVPX) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteChassisVPXParams()
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
