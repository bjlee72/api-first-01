package list

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"api-first-01/api/swagger/models"
	operations "api-first-01/api/swagger/restapi/operations/hello"
)

func errBadRequest(err error) middleware.Responder {
	return operations.NewV1ListHellosBadRequest().WithPayload(&models.Error{
		Message: swag.String(err.Error()),
		Code:    swag.Int64(http.StatusBadRequest),
	})
}

func errInternalServer(err error) middleware.Responder {
	return operations.NewV1ListHellosInternalServerError().WithPayload(&models.Error{
		Message: swag.String(err.Error()),
		Code:    swag.Int64(http.StatusInternalServerError),
	})
}
