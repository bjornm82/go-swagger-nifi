// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ComponentStateDTO component state d t o
// swagger:model ComponentStateDTO
type ComponentStateDTO struct {

	// The cluster state for this component, or null if this NiFi is a standalone instance.
	ClusterState *StateMapDTO `json:"clusterState,omitempty"`

	// The component identifier.
	ComponentID string `json:"componentId,omitempty"`

	// The local state for this component.
	LocalState *StateMapDTO `json:"localState,omitempty"`

	// Description of the state this component persists.
	StateDescription string `json:"stateDescription,omitempty"`
}

// Validate validates this component state d t o
func (m *ComponentStateDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentStateDTO) validateClusterState(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterState) { // not required
		return nil
	}

	if m.ClusterState != nil {
		if err := m.ClusterState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterState")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentStateDTO) validateLocalState(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalState) { // not required
		return nil
	}

	if m.LocalState != nil {
		if err := m.LocalState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentStateDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentStateDTO) UnmarshalBinary(b []byte) error {
	var res ComponentStateDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
