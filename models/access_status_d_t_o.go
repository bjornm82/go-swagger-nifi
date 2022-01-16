// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessStatusDTO access status d t o
//
// swagger:model AccessStatusDTO
type AccessStatusDTO struct {

	// The user identity.
	Identity string `json:"identity,omitempty"`

	// Additional details about the user access status.
	Message string `json:"message,omitempty"`

	// The user access status.
	Status string `json:"status,omitempty"`
}

// Validate validates this access status d t o
func (m *AccessStatusDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessStatusDTO) UnmarshalBinary(b []byte) error {
	var res AccessStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
