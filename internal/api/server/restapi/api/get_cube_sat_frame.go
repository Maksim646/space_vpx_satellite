// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// GetCubeSatFrameHandlerFunc turns a function with the right signature into a get cube sat frame handler
type GetCubeSatFrameHandlerFunc func(GetCubeSatFrameParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCubeSatFrameHandlerFunc) Handle(params GetCubeSatFrameParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetCubeSatFrameHandler interface for that can handle valid get cube sat frame params
type GetCubeSatFrameHandler interface {
	Handle(GetCubeSatFrameParams, *models.Principal) middleware.Responder
}

// NewGetCubeSatFrame creates a new http.Handler for the get cube sat frame operation
func NewGetCubeSatFrame(ctx *middleware.Context, handler GetCubeSatFrameHandler) *GetCubeSatFrame {
	return &GetCubeSatFrame{Context: ctx, Handler: handler}
}

/*
	GetCubeSatFrame swagger:route GET /cube_sat_frame/{id} CubeSateFrame getCubeSatFrame

Get Cube Sat Frame
*/
type GetCubeSatFrame struct {
	Context *middleware.Context
	Handler GetCubeSatFrameHandler
}

func (o *GetCubeSatFrame) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCubeSatFrameParams()
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
