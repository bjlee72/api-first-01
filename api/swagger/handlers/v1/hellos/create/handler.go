package create

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rs/xid"

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
func (d Dependency) NewHandler() operations.V1CreateHelloHandlerFunc {
	if d.ContextCreator == nil {
		logger.Fatalf(nil, "context creator not set for Dependency")
	}
	if d.Storage == nil {
		logger.Fatalf(nil, "storage not set for Dependency")
	}

	return func(params operations.V1CreateHelloParams) middleware.Responder {
		ctx := d.ContextCreator.Create(params.HTTPRequest)
		return d.handle(ctx, params)
	}
}

func (d Dependency) handle(ctx context.Context, params operations.V1CreateHelloParams) middleware.Responder {
	if params.Body.Hello.ID != "" {
		// bad request error handling example
		return errBadRequest(logger.Errorf(ctx, "id cannot be specified in the request: %s", params.Body.Hello.ID))
	}

	resp, err := d.Storage.CreateHello(ctx, &storage.CreateHelloRequest{
		Hello: &storage.Hello{
			ID:      xid.New().String(),
			Message: "hello creation message",
		},
	})
	if err != nil {
		if storage.ErrBadRequest.Matches(err) {
			return errBadRequest(logger.Errorf(ctx, "bad request to storage: %s", err.Error()))
		}
		return errInternalServer(logger.Errorf(ctx, "couldn't create the hello: %s", err.Error()))
	}

	return operations.NewV1CreateHelloOK().
		WithPayload(&models.CreateHelloResponse{
			Hello: &models.Hello{
				ID:      resp.Hello.ID,
				Message: swag.String(resp.Hello.Message),
			},
		})
}
