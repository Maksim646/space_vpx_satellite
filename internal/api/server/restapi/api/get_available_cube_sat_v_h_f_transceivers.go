// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetAvailableCubeSatVHFTransceiversHandlerFunc turns a function with the right signature into a get available cube sat v h f transceivers handler
type GetAvailableCubeSatVHFTransceiversHandlerFunc func(GetAvailableCubeSatVHFTransceiversParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAvailableCubeSatVHFTransceiversHandlerFunc) Handle(params GetAvailableCubeSatVHFTransceiversParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAvailableCubeSatVHFTransceiversHandler interface for that can handle valid get available cube sat v h f transceivers params
type GetAvailableCubeSatVHFTransceiversHandler interface {
	Handle(GetAvailableCubeSatVHFTransceiversParams, *models.Principal) middleware.Responder
}

// NewGetAvailableCubeSatVHFTransceivers creates a new http.Handler for the get available cube sat v h f transceivers operation
func NewGetAvailableCubeSatVHFTransceivers(ctx *middleware.Context, handler GetAvailableCubeSatVHFTransceiversHandler) *GetAvailableCubeSatVHFTransceivers {
	return &GetAvailableCubeSatVHFTransceivers{Context: ctx, Handler: handler}
}

/*
	GetAvailableCubeSatVHFTransceivers swagger:route GET /cube_sat_vhf_transceiver/available CubeSatVHFTransceiver getAvailableCubeSatVHFTransceivers

Get Available CubeSat VHF Transceivers
*/
type GetAvailableCubeSatVHFTransceivers struct {
	Context *middleware.Context
	Handler GetAvailableCubeSatVHFTransceiversHandler
}

func (o *GetAvailableCubeSatVHFTransceivers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAvailableCubeSatVHFTransceiversParams()
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
