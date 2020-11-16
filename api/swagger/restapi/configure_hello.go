// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"api-first-01/api/swagger/restapi/operations"
	"api-first-01/api/swagger/restapi/operations/health"
	"api-first-01/api/swagger/restapi/operations/hello"
)

//go:generate swagger generate server --target ../../swagger --name Hello --spec ../../../specification/swagger.json --exclude-main

func configureFlags(api *operations.HelloAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HelloAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.HelloV1CreateHelloHandler == nil {
		api.HelloV1CreateHelloHandler = hello.V1CreateHelloHandlerFunc(func(params hello.V1CreateHelloParams) middleware.Responder {
			return middleware.NotImplemented("operation hello.V1CreateHello has not yet been implemented")
		})
	}
	if api.HealthV1HealthCheckHandler == nil {
		api.HealthV1HealthCheckHandler = health.V1HealthCheckHandlerFunc(func(params health.V1HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation health.V1HealthCheck has not yet been implemented")
		})
	}
	if api.HelloV1ListHellosHandler == nil {
		api.HelloV1ListHellosHandler = hello.V1ListHellosHandlerFunc(func(params hello.V1ListHellosParams) middleware.Responder {
			return middleware.NotImplemented("operation hello.V1ListHellos has not yet been implemented")
		})
	}
	if api.HelloV1ReadHelloHandler == nil {
		api.HelloV1ReadHelloHandler = hello.V1ReadHelloHandlerFunc(func(params hello.V1ReadHelloParams) middleware.Responder {
			return middleware.NotImplemented("operation hello.V1ReadHello has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the api executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
