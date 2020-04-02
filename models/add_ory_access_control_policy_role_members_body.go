// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddOryAccessControlPolicyRoleMembersBody add ory access control policy role members body
// swagger:model addOryAccessControlPolicyRoleMembersBody
type AddOryAccessControlPolicyRoleMembersBody struct {

	// The members to be added.
	Members []string `json:"members"`
}

// Validate validates this add ory access control policy role members body
func (m *AddOryAccessControlPolicyRoleMembersBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddOryAccessControlPolicyRoleMembersBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOryAccessControlPolicyRoleMembersBody) UnmarshalBinary(b []byte) error {
	var res AddOryAccessControlPolicyRoleMembersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
