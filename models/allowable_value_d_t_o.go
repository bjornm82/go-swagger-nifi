// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AllowableValueDTO allowable value d t o
//
// swagger:model AllowableValueDTO
type AllowableValueDTO struct {

	// A description for this allowable value.
	Description string `json:"description,omitempty"`

	// A human readable value that is allowed for the property descriptor.
	DisplayName string `json:"displayName,omitempty"`

	// A value that is allowed for the property descriptor.
	Value string `json:"value,omitempty"`
}

// Validate validates this allowable value d t o
func (m *AllowableValueDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllowableValueDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowableValueDTO) UnmarshalBinary(b []byte) error {
	var res AllowableValueDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
