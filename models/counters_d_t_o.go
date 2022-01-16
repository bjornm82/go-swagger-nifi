// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CountersDTO counters d t o
//
// swagger:model CountersDTO
type CountersDTO struct {

	// A Counters snapshot that represents the aggregate values of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	AggregateSnapshot *CountersSnapshotDTO `json:"aggregateSnapshot,omitempty"`

	// A Counters snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []*NodeCountersSnapshotDTO `json:"nodeSnapshots"`
}

// Validate validates this counters d t o
func (m *CountersDTO) Validate(formats strfmt.Registry) error {
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

func (m *CountersDTO) validateAggregateSnapshot(formats strfmt.Registry) error {

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

func (m *CountersDTO) validateNodeSnapshots(formats strfmt.Registry) error {

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
func (m *CountersDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountersDTO) UnmarshalBinary(b []byte) error {
	var res CountersDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
