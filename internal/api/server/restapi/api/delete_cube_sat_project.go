// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteCubeSatProjectHandlerFunc turns a function with the right signature into a delete cube sat project handler
type DeleteCubeSatProjectHandlerFunc func(DeleteCubeSatProjectParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCubeSatProjectHandlerFunc) Handle(params DeleteCubeSatProjectParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteCubeSatProjectHandler interface for that can handle valid delete cube sat project params
type DeleteCubeSatProjectHandler interface {
	Handle(DeleteCubeSatProjectParams, *models.Principal) middleware.Responder
}

// NewDeleteCubeSatProject creates a new http.Handler for the delete cube sat project operation
func NewDeleteCubeSatProject(ctx *middleware.Context, handler DeleteCubeSatProjectHandler) *DeleteCubeSatProject {
	return &DeleteCubeSatProject{Context: ctx, Handler: handler}
}

/*
	DeleteCubeSatProject swagger:route DELETE /project/{id} CubeSatProject deleteCubeSatProject

Delete CubeSat project
*/
type DeleteCubeSatProject struct {
	Context *middleware.Context
	Handler DeleteCubeSatProjectHandler
}

func (o *DeleteCubeSatProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteCubeSatProjectParams()
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
