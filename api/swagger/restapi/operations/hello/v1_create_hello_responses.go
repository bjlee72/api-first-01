// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"api-first-01/api/swagger/models"
)

// V1CreateHelloOKCode is the HTTP code returned for type V1CreateHelloOK
const V1CreateHelloOKCode int = 200

/*V1CreateHelloOK The hello is successfully created.

swagger:response v1CreateHelloOK
*/
type V1CreateHelloOK struct {

	/*
	  In: Body
	*/
	Payload *models.CreateHelloResponse `json:"body,omitempty"`
}

// NewV1CreateHelloOK creates V1CreateHelloOK with default headers values
func NewV1CreateHelloOK() *V1CreateHelloOK {

	return &V1CreateHelloOK{}
}

// WithPayload adds the payload to the v1 create hello o k response
func (o *V1CreateHelloOK) WithPayload(payload *models.CreateHelloResponse) *V1CreateHelloOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v1 create hello o k response
func (o *V1CreateHelloOK) SetPayload(payload *models.CreateHelloResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V1CreateHelloOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V1CreateHelloBadRequestCode is the HTTP code returned for type V1CreateHelloBadRequest
const V1CreateHelloBadRequestCode int = 400

/*V1CreateHelloBadRequest Bad Request

swagger:response v1CreateHelloBadRequest
*/
type V1CreateHelloBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV1CreateHelloBadRequest creates V1CreateHelloBadRequest with default headers values
func NewV1CreateHelloBadRequest() *V1CreateHelloBadRequest {

	return &V1CreateHelloBadRequest{}
}

// WithPayload adds the payload to the v1 create hello bad request response
func (o *V1CreateHelloBadRequest) WithPayload(payload *models.Error) *V1CreateHelloBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v1 create hello bad request response
func (o *V1CreateHelloBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V1CreateHelloBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V1CreateHelloInternalServerErrorCode is the HTTP code returned for type V1CreateHelloInternalServerError
const V1CreateHelloInternalServerErrorCode int = 500

/*V1CreateHelloInternalServerError Internal Server Error

swagger:response v1CreateHelloInternalServerError
*/
type V1CreateHelloInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV1CreateHelloInternalServerError creates V1CreateHelloInternalServerError with default headers values
func NewV1CreateHelloInternalServerError() *V1CreateHelloInternalServerError {

	return &V1CreateHelloInternalServerError{}
}

// WithPayload adds the payload to the v1 create hello internal server error response
func (o *V1CreateHelloInternalServerError) WithPayload(payload *models.Error) *V1CreateHelloInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v1 create hello internal server error response
func (o *V1CreateHelloInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V1CreateHelloInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}