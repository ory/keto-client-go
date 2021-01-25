// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthorizationResult AuthorizationResult is the result of an access control decision. It contains the decision outcome.
//
// swagger:model authorizationResult
type AuthorizationResult struct {

	// Allowed is true if the request should be allowed and false otherwise.
	// Required: true
	Allowed *bool `json:"allowed"`
}

// Validate validates this authorization result
func (m *AuthorizationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationResult) validateAllowed(formats strfmt.Registry) error {

	if err := validate.Required("allowed", "body", m.Allowed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this authorization result based on context it is used
func (m *AuthorizationResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationResult) UnmarshalBinary(b []byte) error {
	var res AuthorizationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
