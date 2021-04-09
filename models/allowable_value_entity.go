// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AllowableValueEntity allowable value entity
// swagger:model AllowableValueEntity
type AllowableValueEntity struct {

	// allowable value
	AllowableValue *AllowableValueDTO `json:"allowableValue,omitempty"`

	// Indicates whether the user can read a given resource.
	// Read Only: true
	CanRead *bool `json:"canRead,omitempty"`
}

// Validate validates this allowable value entity
func (m *AllowableValueEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowableValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllowableValueEntity) validateAllowableValue(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowableValue) { // not required
		return nil
	}

	if m.AllowableValue != nil {
		if err := m.AllowableValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowableValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllowableValueEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowableValueEntity) UnmarshalBinary(b []byte) error {
	var res AllowableValueEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}