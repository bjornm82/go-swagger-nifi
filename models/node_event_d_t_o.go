// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NodeEventDTO node event d t o
// swagger:model NodeEventDTO
type NodeEventDTO struct {

	// The category of the node event.
	Category string `json:"category,omitempty"`

	// The message in the node event.
	Message string `json:"message,omitempty"`

	// The timestamp of the node event.
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this node event d t o
func (m *NodeEventDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeEventDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeEventDTO) UnmarshalBinary(b []byte) error {
	var res NodeEventDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
