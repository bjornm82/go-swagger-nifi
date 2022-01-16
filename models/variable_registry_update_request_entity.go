// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VariableRegistryUpdateRequestEntity variable registry update request entity
//
// swagger:model VariableRegistryUpdateRequestEntity
type VariableRegistryUpdateRequestEntity struct {

	// The revision for the Process Group that owns this variable registry.
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`

	// The Variable Registry Update Request
	Request *VariableRegistryUpdateRequestDTO `json:"request,omitempty"`
}

// Validate validates this variable registry update request entity
func (m *VariableRegistryUpdateRequestEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableRegistryUpdateRequestEntity) validateProcessGroupRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessGroupRevision) { // not required
		return nil
	}

	if m.ProcessGroupRevision != nil {
		if err := m.ProcessGroupRevision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupRevision")
			}
			return err
		}
	}

	return nil
}

func (m *VariableRegistryUpdateRequestEntity) validateRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariableRegistryUpdateRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableRegistryUpdateRequestEntity) UnmarshalBinary(b []byte) error {
	var res VariableRegistryUpdateRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
