// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VariableDTO variable d t o
//
// swagger:model VariableDTO
type VariableDTO struct {

	// A set of all components that will be affected if the value of this variable is changed
	// Unique: true
	AffectedComponents []*AffectedComponentEntity `json:"affectedComponents"`

	// The name of the variable
	Name string `json:"name,omitempty"`

	// The ID of the Process Group where this Variable is defined
	ProcessGroupID string `json:"processGroupId,omitempty"`

	// The value of the variable
	Value string `json:"value,omitempty"`
}

// Validate validates this variable d t o
func (m *VariableDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffectedComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableDTO) validateAffectedComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.AffectedComponents) { // not required
		return nil
	}

	if err := validate.UniqueItems("affectedComponents", "body", m.AffectedComponents); err != nil {
		return err
	}

	for i := 0; i < len(m.AffectedComponents); i++ {
		if swag.IsZero(m.AffectedComponents[i]) { // not required
			continue
		}

		if m.AffectedComponents[i] != nil {
			if err := m.AffectedComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affectedComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariableDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableDTO) UnmarshalBinary(b []byte) error {
	var res VariableDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
