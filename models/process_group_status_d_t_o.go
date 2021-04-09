// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProcessGroupStatusDTO process group status d t o
// swagger:model ProcessGroupStatusDTO
type ProcessGroupStatusDTO struct {

	// The aggregate status of all nodes in the cluster
	AggregateSnapshot *ProcessGroupStatusSnapshotDTO `json:"aggregateSnapshot,omitempty"`

	// The ID of the Process Group
	ID string `json:"id,omitempty"`

	// The name of the Process Group
	Name string `json:"name,omitempty"`

	// The status reported by each node in the cluster. If the NiFi instance is a standalone instance, rather than a clustered instance, this value may be null.
	NodeSnapshots []*NodeProcessGroupStatusSnapshotDTO `json:"nodeSnapshots"`

	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
}

// Validate validates this process group status d t o
func (m *ProcessGroupStatusDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupStatusDTO) validateAggregateSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.AggregateSnapshot) { // not required
		return nil
	}

	if m.AggregateSnapshot != nil {
		if err := m.AggregateSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregateSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupStatusDTO) validateNodeSnapshots(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeSnapshots); i++ {
		if swag.IsZero(m.NodeSnapshots[i]) { // not required
			continue
		}

		if m.NodeSnapshots[i] != nil {
			if err := m.NodeSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupStatusDTO) UnmarshalBinary(b []byte) error {
	var res ProcessGroupStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
