// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ParameterContextUpdateStepDTO parameter context update step d t o
// swagger:model ParameterContextUpdateStepDTO
type ParameterContextUpdateStepDTO struct {

	// Whether or not this step has completed
	// Read Only: true
	Complete *bool `json:"complete,omitempty"`

	// Explanation of what happens in this step
	// Read Only: true
	Description string `json:"description,omitempty"`

	// An explanation of why this step failed, or null if this step did not fail
	// Read Only: true
	FailureReason string `json:"failureReason,omitempty"`
}

// Validate validates this parameter context update step d t o
func (m *ParameterContextUpdateStepDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterContextUpdateStepDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterContextUpdateStepDTO) UnmarshalBinary(b []byte) error {
	var res ParameterContextUpdateStepDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
