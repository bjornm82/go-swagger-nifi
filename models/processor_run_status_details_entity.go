// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProcessorRunStatusDetailsEntity processor run status details entity
// swagger:model ProcessorRunStatusDetailsEntity
type ProcessorRunStatusDetailsEntity struct {

	// The permissions for the Processor.
	Permissions *PermissionsDTO `json:"permissions,omitempty"`

	// The revision for the Processor.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// The details of a Processor's run status
	RunStatusDetails *ProcessorRunStatusDetailsDTO `json:"runStatusDetails,omitempty"`
}

// Validate validates this processor run status details entity
func (m *ProcessorRunStatusDetailsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunStatusDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorRunStatusDetailsEntity) validatePermissions(formats strfmt.Registry) error {

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

func (m *ProcessorRunStatusDetailsEntity) validateRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessorRunStatusDetailsEntity) validateRunStatusDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.RunStatusDetails) { // not required
		return nil
	}

	if m.RunStatusDetails != nil {
		if err := m.RunStatusDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runStatusDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessorRunStatusDetailsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorRunStatusDetailsEntity) UnmarshalBinary(b []byte) error {
	var res ProcessorRunStatusDetailsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
