// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// UpdateBoardComputingModuleHandlerFunc turns a function with the right signature into a update board computing module handler
type UpdateBoardComputingModuleHandlerFunc func(UpdateBoardComputingModuleParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateBoardComputingModuleHandlerFunc) Handle(params UpdateBoardComputingModuleParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateBoardComputingModuleHandler interface for that can handle valid update board computing module params
type UpdateBoardComputingModuleHandler interface {
	Handle(UpdateBoardComputingModuleParams, *models.Principal) middleware.Responder
}

// NewUpdateBoardComputingModule creates a new http.Handler for the update board computing module operation
func NewUpdateBoardComputingModule(ctx *middleware.Context, handler UpdateBoardComputingModuleHandler) *UpdateBoardComputingModule {
	return &UpdateBoardComputingModule{Context: ctx, Handler: handler}
}

/*
	UpdateBoardComputingModule swagger:route PATCH /board_computing_module/{id} BoardComputingModule updateBoardComputingModule

Update a Board Computing Module by ID
*/
type UpdateBoardComputingModule struct {
	Context *middleware.Context
	Handler UpdateBoardComputingModuleHandler
}

func (o *UpdateBoardComputingModule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateBoardComputingModuleParams()
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
