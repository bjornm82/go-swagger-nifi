// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FlowBreadcrumbDTO flow breadcrumb d t o
//
// swagger:model FlowBreadcrumbDTO
type FlowBreadcrumbDTO struct {

	// The id of the group.
	ID string `json:"id,omitempty"`

	// The id of the group.
	Name string `json:"name,omitempty"`

	// The process group version control information or null if not version controlled.
	VersionControlInformation *VersionControlInformationDTO `json:"versionControlInformation,omitempty"`
}

// Validate validates this flow breadcrumb d t o
func (m *FlowBreadcrumbDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionControlInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowBreadcrumbDTO) validateVersionControlInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionControlInformation) { // not required
		return nil
	}

	if m.VersionControlInformation != nil {
		if err := m.VersionControlInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionControlInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowBreadcrumbDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowBreadcrumbDTO) UnmarshalBinary(b []byte) error {
	var res FlowBreadcrumbDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
