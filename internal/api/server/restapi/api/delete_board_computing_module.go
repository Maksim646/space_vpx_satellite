// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	models "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
)

// DeleteBoardComputingModuleHandlerFunc turns a function with the right signature into a delete board computing module handler
type DeleteBoardComputingModuleHandlerFunc func(DeleteBoardComputingModuleParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBoardComputingModuleHandlerFunc) Handle(params DeleteBoardComputingModuleParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteBoardComputingModuleHandler interface for that can handle valid delete board computing module params
type DeleteBoardComputingModuleHandler interface {
	Handle(DeleteBoardComputingModuleParams, *models.Principal) middleware.Responder
}

// NewDeleteBoardComputingModule creates a new http.Handler for the delete board computing module operation
func NewDeleteBoardComputingModule(ctx *middleware.Context, handler DeleteBoardComputingModuleHandler) *DeleteBoardComputingModule {
	return &DeleteBoardComputingModule{Context: ctx, Handler: handler}
}

/*
	DeleteBoardComputingModule swagger:route DELETE /board_computing_module/{id} BoardComputingModule deleteBoardComputingModule

Delete a Board Computing Module by ID
*/
type DeleteBoardComputingModule struct {
	Context *middleware.Context
	Handler DeleteBoardComputingModuleHandler
}

func (o *DeleteBoardComputingModule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteBoardComputingModuleParams()
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
