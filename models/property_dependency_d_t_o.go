// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PropertyDependencyDTO property dependency d t o
//
// swagger:model PropertyDependencyDTO
type PropertyDependencyDTO struct {

	// The values for the property that satisfies the dependency, or null if the dependency is satisfied by the presence of any value for the associated property name
	// Unique: true
	DependentValues []string `json:"dependentValues"`

	// The name of the property that is being depended upon
	PropertyName string `json:"propertyName,omitempty"`
}

// Validate validates this property dependency d t o
func (m *PropertyDependencyDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependentValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDependencyDTO) validateDependentValues(formats strfmt.Registry) error {

	if swag.IsZero(m.DependentValues) { // not required
		return nil
	}

	if err := validate.UniqueItems("dependentValues", "body", m.DependentValues); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyDependencyDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyDependencyDTO) UnmarshalBinary(b []byte) error {
	var res PropertyDependencyDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
