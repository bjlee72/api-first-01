// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateHelloResponse A response to the request to create a hello.
//
// swagger:model createHelloResponse
type CreateHelloResponse struct {

	// hello
	// Required: true
	Hello *Hello `json:"hello"`
}

// Validate validates this create hello response
func (m *CreateHelloResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHello(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateHelloResponse) validateHello(formats strfmt.Registry) error {

	if err := validate.Required("hello", "body", m.Hello); err != nil {
		return err
	}

	if m.Hello != nil {
		if err := m.Hello.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hello")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateHelloResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateHelloResponse) UnmarshalBinary(b []byte) error {
	var res CreateHelloResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
