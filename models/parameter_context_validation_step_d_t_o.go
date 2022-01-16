// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParameterContextValidationStepDTO parameter context validation step d t o
//
// swagger:model ParameterContextValidationStepDTO
type ParameterContextValidationStepDTO struct {

	// Whether or not this step has completed
	Complete bool `json:"complete,omitempty"`

	// Explanation of what happens in this step
	Description string `json:"description,omitempty"`

	// An explanation of why this step failed, or null if this step did not fail
	FailureReason string `json:"failureReason,omitempty"`
}

// Validate validates this parameter context validation step d t o
func (m *ParameterContextValidationStepDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterContextValidationStepDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterContextValidationStepDTO) UnmarshalBinary(b []byte) error {
	var res ParameterContextValidationStepDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
