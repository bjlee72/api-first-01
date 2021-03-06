// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V1ReadHelloStatusHandlerFunc turns a function with the right signature into a v1 read hello status handler
type V1ReadHelloStatusHandlerFunc func(V1ReadHelloStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn V1ReadHelloStatusHandlerFunc) Handle(params V1ReadHelloStatusParams) middleware.Responder {
	return fn(params)
}

// V1ReadHelloStatusHandler interface for that can handle valid v1 read hello status params
type V1ReadHelloStatusHandler interface {
	Handle(V1ReadHelloStatusParams) middleware.Responder
}

// NewV1ReadHelloStatus creates a new http.Handler for the v1 read hello status operation
func NewV1ReadHelloStatus(ctx *middleware.Context, handler V1ReadHelloStatusHandler) *V1ReadHelloStatus {
	return &V1ReadHelloStatus{Context: ctx, Handler: handler}
}

/*V1ReadHelloStatus swagger:route GET /v1/hellos/{hello_id}/status Hello v1ReadHelloStatus

Get the status of a hello.

Get the current status of a specific hello.

*/
type V1ReadHelloStatus struct {
	Context *middleware.Context
	Handler V1ReadHelloStatusHandler
}

func (o *V1ReadHelloStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV1ReadHelloStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
