// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NodeProcessorStatusSnapshotDTO node processor status snapshot d t o
// swagger:model NodeProcessorStatusSnapshotDTO
type NodeProcessorStatusSnapshotDTO struct {

	// The API address of the node
	Address string `json:"address,omitempty"`

	// The API port used to communicate with the node
	APIPort int32 `json:"apiPort,omitempty"`

	// The unique ID that identifies the node
	NodeID string `json:"nodeId,omitempty"`

	// The processor status snapshot from the node.
	StatusSnapshot *ProcessorStatusSnapshotDTO `json:"statusSnapshot,omitempty"`
}

// Validate validates this node processor status snapshot d t o
func (m *NodeProcessorStatusSnapshotDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeProcessorStatusSnapshotDTO) validateStatusSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusSnapshot) { // not required
		return nil
	}

	if m.StatusSnapshot != nil {
		if err := m.StatusSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusSnapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeProcessorStatusSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeProcessorStatusSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res NodeProcessorStatusSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
