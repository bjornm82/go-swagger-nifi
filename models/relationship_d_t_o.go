// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RelationshipDTO relationship d t o
//
// swagger:model RelationshipDTO
type RelationshipDTO struct {

	// Whether or not flowfiles sent to this relationship should auto terminate.
	AutoTerminate bool `json:"autoTerminate,omitempty"`

	// The relationship description.
	Description string `json:"description,omitempty"`

	// The relationship name.
	Name string `json:"name,omitempty"`
}

// Validate validates this relationship d t o
func (m *RelationshipDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipDTO) UnmarshalBinary(b []byte) error {
	var res RelationshipDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
