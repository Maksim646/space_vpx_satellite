// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateCubeSatFrameHandlerFunc turns a function with the right signature into a update cube sat frame handler
type UpdateCubeSatFrameHandlerFunc func(UpdateCubeSatFrameParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCubeSatFrameHandlerFunc) Handle(params UpdateCubeSatFrameParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateCubeSatFrameHandler interface for that can handle valid update cube sat frame params
type UpdateCubeSatFrameHandler interface {
	Handle(UpdateCubeSatFrameParams, *models.Principal) middleware.Responder
}

// NewUpdateCubeSatFrame creates a new http.Handler for the update cube sat frame operation
func NewUpdateCubeSatFrame(ctx *middleware.Context, handler UpdateCubeSatFrameHandler) *UpdateCubeSatFrame {
	return &UpdateCubeSatFrame{Context: ctx, Handler: handler}
}

/*
	UpdateCubeSatFrame swagger:route PATCH /cube_sat_frame/{id} CubeSateFrame updateCubeSatFrame

Update Cube Sat Frame
*/
type UpdateCubeSatFrame struct {
	Context *middleware.Context
	Handler UpdateCubeSatFrameHandler
}

func (o *UpdateCubeSatFrame) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateCubeSatFrameParams()
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
