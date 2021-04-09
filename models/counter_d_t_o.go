// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CounterDTO counter d t o
// swagger:model CounterDTO
type CounterDTO struct {

	// The context of the counter.
	Context string `json:"context,omitempty"`

	// The id of the counter.
	ID string `json:"id,omitempty"`

	// The name of the counter.
	Name string `json:"name,omitempty"`

	// The value of the counter.
	Value string `json:"value,omitempty"`

	// The value count.
	ValueCount int64 `json:"valueCount,omitempty"`
}

// Validate validates this counter d t o
func (m *CounterDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CounterDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterDTO) UnmarshalBinary(b []byte) error {
	var res CounterDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}