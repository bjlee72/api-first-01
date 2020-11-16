package get

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"api-first-01/api/swagger/models"
	operations "api-first-01/api/swagger/restapi/operations/health"
	"api-first-01/utils/logger"
)

// Dependency expresses the downstream dependencies for this handler.
type Dependency struct {
	ContextCreator
}

// NewHandler converts dependencies into a handler that works on top of them.
func (d Dependency) NewHandler() operations.V1HealthCheckHandlerFunc {
	if d.ContextCreator == nil {
		logger.Fatalf(nil, "context creator not set for HandlerBuilder")
	}

	return func(params operations.V1HealthCheckParams) middleware.Responder {
		ctx := d.ContextCreator.Create(params.HTTPRequest)
		return d.handle(ctx, params)
	}
}

// The actual business logic implementation.
func (d Dependency) handle(_ context.Context, _ operations.V1HealthCheckParams) middleware.Responder {
	return operations.NewV1HealthCheckOK().
		WithPayload(&models.HealthCheckResponse{Healthy: swag.Bool(true)})
}
