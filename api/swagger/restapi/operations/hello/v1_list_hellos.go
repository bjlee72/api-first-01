// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V1ListHellosHandlerFunc turns a function with the right signature into a v1 list hellos handler
type V1ListHellosHandlerFunc func(V1ListHellosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn V1ListHellosHandlerFunc) Handle(params V1ListHellosParams) middleware.Responder {
	return fn(params)
}

// V1ListHellosHandler interface for that can handle valid v1 list hellos params
type V1ListHellosHandler interface {
	Handle(V1ListHellosParams) middleware.Responder
}

// NewV1ListHellos creates a new http.Handler for the v1 list hellos operation
func NewV1ListHellos(ctx *middleware.Context, handler V1ListHellosHandler) *V1ListHellos {
	return &V1ListHellos{Context: ctx, Handler: handler}
}

/*V1ListHellos swagger:route GET /v1/hellos Hello v1ListHellos

List Hellos.

List Hellos which are in the system.

*/
type V1ListHellos struct {
	Context *middleware.Context
	Handler V1ListHellosHandler
}

func (o *V1ListHellos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV1ListHellosParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}