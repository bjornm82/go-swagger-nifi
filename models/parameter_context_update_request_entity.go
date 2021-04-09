// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ParameterContextUpdateRequestEntity parameter context update request entity
// swagger:model ParameterContextUpdateRequestEntity
type ParameterContextUpdateRequestEntity struct {

	// The Revision of the Parameter Context
	ParameterContextRevision *RevisionDTO `json:"parameterContextRevision,omitempty"`

	// The Update Request
	Request *ParameterContextUpdateRequestDTO `json:"request,omitempty"`
}

// Validate validates this parameter context update request entity
func (m *ParameterContextUpdateRequestEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameterContextRevision(formats); err != nil {
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

func (m *ParameterContextUpdateRequestEntity) validateParameterContextRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterContextRevision) { // not required
		return nil
	}

	if m.ParameterContextRevision != nil {
		if err := m.ParameterContextRevision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterContextRevision")
			}
			return err
		}
	}

	return nil
}

func (m *ParameterContextUpdateRequestEntity) validateRequest(formats strfmt.Registry) error {

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
func (m *ParameterContextUpdateRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterContextUpdateRequestEntity) UnmarshalBinary(b []byte) error {
	var res ParameterContextUpdateRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
