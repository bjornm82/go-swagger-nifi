// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ParameterContextUpdateRequestDTO parameter context update request d t o
// swagger:model ParameterContextUpdateRequestDTO
type ParameterContextUpdateRequestDTO struct {

	// Whether or not the request is completed
	// Read Only: true
	Complete *bool `json:"complete,omitempty"`

	// The reason for the request failing, or null if the request has not failed
	// Read Only: true
	FailureReason string `json:"failureReason,omitempty"`

	// The timestamp of when the request was last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// The Parameter Context that is being operated on. This may not be populated until the request has successfully completed.
	// Read Only: true
	ParameterContext *ParameterContextDTO `json:"parameterContext,omitempty"`

	// A value between 0 and 100 (inclusive) indicating how close the request is to completion
	// Read Only: true
	PercentCompleted int32 `json:"percentCompleted,omitempty"`

	// The components that are referenced by the update.
	// Read Only: true
	// Unique: true
	ReferencingComponents []*AffectedComponentEntity `json:"referencingComponents"`

	// The ID of the request
	// Read Only: true
	RequestID string `json:"requestId,omitempty"`

	// A description of the current state of the request
	// Read Only: true
	State string `json:"state,omitempty"`

	// The timestamp of when the request was submitted
	// Read Only: true
	// Format: date-time
	SubmissionTime strfmt.DateTime `json:"submissionTime,omitempty"`

	// The steps that are required in order to complete the request, along with the status of each
	// Read Only: true
	UpdateSteps []*ParameterContextUpdateStepDTO `json:"updateSteps"`

	// The URI for the request
	// Read Only: true
	URI string `json:"uri,omitempty"`
}

// Validate validates this parameter context update request d t o
func (m *ParameterContextUpdateRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferencingComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterContextUpdateRequestDTO) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ParameterContextUpdateRequestDTO) validateParameterContext(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterContext) { // not required
		return nil
	}

	if m.ParameterContext != nil {
		if err := m.ParameterContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterContext")
			}
			return err
		}
	}

	return nil
}

func (m *ParameterContextUpdateRequestDTO) validateReferencingComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferencingComponents) { // not required
		return nil
	}

	if err := validate.UniqueItems("referencingComponents", "body", m.ReferencingComponents); err != nil {
		return err
	}

	for i := 0; i < len(m.ReferencingComponents); i++ {
		if swag.IsZero(m.ReferencingComponents[i]) { // not required
			continue
		}

		if m.ReferencingComponents[i] != nil {
			if err := m.ReferencingComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterContextUpdateRequestDTO) validateSubmissionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("submissionTime", "body", "date-time", m.SubmissionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ParameterContextUpdateRequestDTO) validateUpdateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateSteps); i++ {
		if swag.IsZero(m.UpdateSteps[i]) { // not required
			continue
		}

		if m.UpdateSteps[i] != nil {
			if err := m.UpdateSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParameterContextUpdateRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterContextUpdateRequestDTO) UnmarshalBinary(b []byte) error {
	var res ParameterContextUpdateRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}