// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProcessGroupFlowEntity process group flow entity
// swagger:model ProcessGroupFlowEntity
type ProcessGroupFlowEntity struct {

	// The access policy for this process group.
	Permissions *PermissionsDTO `json:"permissions,omitempty"`

	// process group flow
	ProcessGroupFlow *ProcessGroupFlowDTO `json:"processGroupFlow,omitempty"`
}

// Validate validates this process group flow entity
func (m *ProcessGroupFlowEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessGroupFlow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupFlowEntity) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupFlowEntity) validateProcessGroupFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessGroupFlow) { // not required
		return nil
	}

	if m.ProcessGroupFlow != nil {
		if err := m.ProcessGroupFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupFlow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupFlowEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupFlowEntity) UnmarshalBinary(b []byte) error {
	var res ProcessGroupFlowEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}