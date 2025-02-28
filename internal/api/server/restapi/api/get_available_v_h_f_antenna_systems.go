// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetAvailableVHFAntennaSystemsHandlerFunc turns a function with the right signature into a get available v h f antenna systems handler
type GetAvailableVHFAntennaSystemsHandlerFunc func(GetAvailableVHFAntennaSystemsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAvailableVHFAntennaSystemsHandlerFunc) Handle(params GetAvailableVHFAntennaSystemsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAvailableVHFAntennaSystemsHandler interface for that can handle valid get available v h f antenna systems params
type GetAvailableVHFAntennaSystemsHandler interface {
	Handle(GetAvailableVHFAntennaSystemsParams, *models.Principal) middleware.Responder
}

// NewGetAvailableVHFAntennaSystems creates a new http.Handler for the get available v h f antenna systems operation
func NewGetAvailableVHFAntennaSystems(ctx *middleware.Context, handler GetAvailableVHFAntennaSystemsHandler) *GetAvailableVHFAntennaSystems {
	return &GetAvailableVHFAntennaSystems{Context: ctx, Handler: handler}
}

/*
	GetAvailableVHFAntennaSystems swagger:route GET /vhf_antenna_system/available_vhf_antenna_systems VHFAntennaSystem getAvailableVHFAntennaSystems

Get Available VHF Antenna Systems
*/
type GetAvailableVHFAntennaSystems struct {
	Context *middleware.Context
	Handler GetAvailableVHFAntennaSystemsHandler
}

func (o *GetAvailableVHFAntennaSystems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAvailableVHFAntennaSystemsParams()
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
