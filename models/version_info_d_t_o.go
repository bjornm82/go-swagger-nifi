// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VersionInfoDTO version info d t o
// swagger:model VersionInfoDTO
type VersionInfoDTO struct {

	// Build branch
	BuildBranch string `json:"buildBranch,omitempty"`

	// Build revision or commit hash
	BuildRevision string `json:"buildRevision,omitempty"`

	// Build tag
	BuildTag string `json:"buildTag,omitempty"`

	// Build timestamp
	// Format: date-time
	BuildTimestamp strfmt.DateTime `json:"buildTimestamp,omitempty"`

	// Java JVM vendor
	JavaVendor string `json:"javaVendor,omitempty"`

	// Java version
	JavaVersion string `json:"javaVersion,omitempty"`

	// The version of this NiFi.
	NiFiVersion string `json:"niFiVersion,omitempty"`

	// Host operating system architecture
	OsArchitecture string `json:"osArchitecture,omitempty"`

	// Host operating system name
	OsName string `json:"osName,omitempty"`

	// Host operating system version
	OsVersion string `json:"osVersion,omitempty"`
}

// Validate validates this version info d t o
func (m *VersionInfoDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionInfoDTO) validateBuildTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("buildTimestamp", "body", "date-time", m.BuildTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionInfoDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionInfoDTO) UnmarshalBinary(b []byte) error {
	var res VersionInfoDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
