// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeConnectionStatisticsSnapshotDTO node connection statistics snapshot d t o
//
// swagger:model NodeConnectionStatisticsSnapshotDTO
type NodeConnectionStatisticsSnapshotDTO struct {

	// The API address of the node
	Address string `json:"address,omitempty"`

	// The API port used to communicate with the node
	APIPort int32 `json:"apiPort,omitempty"`

	// The unique ID that identifies the node
	NodeID string `json:"nodeId,omitempty"`

	// The connection status snapshot from the node.
	StatisticsSnapshot *ConnectionStatisticsSnapshotDTO `json:"statisticsSnapshot,omitempty"`
}

// Validate validates this node connection statistics snapshot d t o
func (m *NodeConnectionStatisticsSnapshotDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatisticsSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeConnectionStatisticsSnapshotDTO) validateStatisticsSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.StatisticsSnapshot) { // not required
		return nil
	}

	if m.StatisticsSnapshot != nil {
		if err := m.StatisticsSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statisticsSnapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeConnectionStatisticsSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeConnectionStatisticsSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res NodeConnectionStatisticsSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
