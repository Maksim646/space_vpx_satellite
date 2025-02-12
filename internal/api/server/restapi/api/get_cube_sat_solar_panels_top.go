// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetCubeSatSolarPanelsTopHandlerFunc turns a function with the right signature into a get cube sat solar panels top handler
type GetCubeSatSolarPanelsTopHandlerFunc func(GetCubeSatSolarPanelsTopParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCubeSatSolarPanelsTopHandlerFunc) Handle(params GetCubeSatSolarPanelsTopParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetCubeSatSolarPanelsTopHandler interface for that can handle valid get cube sat solar panels top params
type GetCubeSatSolarPanelsTopHandler interface {
	Handle(GetCubeSatSolarPanelsTopParams, *models.Principal) middleware.Responder
}

// NewGetCubeSatSolarPanelsTop creates a new http.Handler for the get cube sat solar panels top operation
func NewGetCubeSatSolarPanelsTop(ctx *middleware.Context, handler GetCubeSatSolarPanelsTopHandler) *GetCubeSatSolarPanelsTop {
	return &GetCubeSatSolarPanelsTop{Context: ctx, Handler: handler}
}

/*
	GetCubeSatSolarPanelsTop swagger:route GET /solar_panel_top/available_solar_panel_top SolarPanelTop getCubeSatSolarPanelsTop

Get Cube Sat Solar Panels Top
*/
type GetCubeSatSolarPanelsTop struct {
	Context *middleware.Context
	Handler GetCubeSatSolarPanelsTopHandler
}

func (o *GetCubeSatSolarPanelsTop) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCubeSatSolarPanelsTopParams()
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
