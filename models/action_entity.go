// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ActionEntity action entity
// swagger:model ActionEntity
type ActionEntity struct {

	// action
	Action *ActionDTO `json:"action,omitempty"`

	// Indicates whether the user can read a given resource.
	// Read Only: true
	CanRead *bool `json:"canRead,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// source Id
	SourceID string `json:"sourceId,omitempty"`

	// The timestamp of the action.
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this action entity
func (m *ActionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionEntity) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionEntity) UnmarshalBinary(b []byte) error {
	var res ActionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
