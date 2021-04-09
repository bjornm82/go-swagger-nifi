// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StateEntryDTO state entry d t o
// swagger:model StateEntryDTO
type StateEntryDTO struct {

	// The label for the node where the state originated.
	ClusterNodeAddress string `json:"clusterNodeAddress,omitempty"`

	// The identifier for the node where the state originated.
	ClusterNodeID string `json:"clusterNodeId,omitempty"`

	// The key for this state.
	Key string `json:"key,omitempty"`

	// The value for this state.
	Value string `json:"value,omitempty"`
}

// Validate validates this state entry d t o
func (m *StateEntryDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateEntryDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateEntryDTO) UnmarshalBinary(b []byte) error {
	var res StateEntryDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
