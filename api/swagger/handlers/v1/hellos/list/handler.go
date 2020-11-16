package list

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
func (d Dependency) NewHandler() operations.V1ListHellosHandlerFunc {
	if d.ContextCreator == nil {
		logger.Fatalf(nil, "context creator not set for Dependency")
	}
	if d.Storage == nil {
		logger.Fatalf(nil, "storage not set for Dependency")
	}

	return func(params operations.V1ListHellosParams) middleware.Responder {
		ctx := d.ContextCreator.Create(params.HTTPRequest)
		return d.handle(ctx, params)
	}
}

func (d Dependency) handle(ctx context.Context, params operations.V1ListHellosParams) middleware.Responder {
	resp, err := d.Storage.ListHellos(ctx, &storage.ListHellosRequest{
		Count: params.Count,
	})
	if err != nil {
		return errInternalServer(logger.Errorf(ctx, "couldn't find the hello: %s", err.Error()))
	}

	hellos := make([]*models.Hello, 0, len(resp.Hellos))
	for _, h := range resp.Hellos {
		hellos = append(hellos, &models.Hello{
			ID:      h.ID,
			Message: swag.String(h.Message),
		})
	}

	return operations.NewV1ListHellosOK().
		WithPayload(&models.ListHellosResponse{
			Hellos: hellos,
		})
}
