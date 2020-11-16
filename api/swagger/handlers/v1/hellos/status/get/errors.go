package get

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"api-first-01/api/swagger/models"
	operations "api-first-01/api/swagger/restapi/operations/hello"
)

func errBadRequest(err error) middleware.Responder {
	return operations.NewV1ReadHelloStatusBadRequest().WithPayload(&models.Error{
		Message: swag.String(err.Error()),
		Code:    swag.Int64(http.StatusBadRequest),
	})
}

func errNotFound(err error) middleware.Responder {
	return operations.NewV1ReadHelloStatusNotFound().WithPayload(&models.Error{
		Message: swag.String(err.Error()),
		Code:    swag.Int64(http.StatusNotFound),
	})
}

func errInternalServer(err error) middleware.Responder {
	return operations.NewV1ReadHelloStatusInternalServerError().WithPayload(&models.Error{
		Message: swag.String(err.Error()),
		Code:    swag.Int64(http.StatusInternalServerError),
	})
}
