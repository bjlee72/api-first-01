package get

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"api-first-01/api/swagger/models"
	operations "api-first-01/api/swagger/restapi/operations/hello"
	"api-first-01/storage"
	"api-first-01/utils/logger"
)

// Dependency handles POST requests to 'platforms' API.
type Dependency struct {
	ContextCreator ContextCreator
	Storage        HelloStorage
}

// NewHandler returns the request handler out of the given dependency.
func (d Dependency) NewHandler() operations.V1ReadHelloHandlerFunc {
	if d.ContextCreator == nil {
		logger.Fatalf(nil, "context creator not set for Dependency")
	}
	if d.Storage == nil {
		logger.Fatalf(nil, "storage not set for Dependency")
	}

	return func(params operations.V1ReadHelloParams) middleware.Responder {
		ctx := d.ContextCreator.Create(params.HTTPRequest)
		return d.handle(ctx, params)
	}
}

func (d Dependency) handle(ctx context.Context, params operations.V1ReadHelloParams) middleware.Responder {
	resp, err := d.Storage.ReadHello(ctx, &storage.ReadHelloRequest{
		ID: params.HelloID,
	})
	if err != nil {
		if storage.ErrRecordNotFound.Matches(err) {
			return errNotFound(logger.Errorf(ctx, "couldn't find the hello: %s", err.Error()))
		}
		return errInternalServer(logger.Errorf(ctx, "couldn't find the hello: %s", err.Error()))
	}

	return operations.NewV1ReadHelloOK().
		WithPayload(&models.ReadHelloResponse{
			Hello: &models.Hello{
				ID:      resp.Hello.ID,
				Message: swag.String(resp.Hello.Message),
			},
		})
}
