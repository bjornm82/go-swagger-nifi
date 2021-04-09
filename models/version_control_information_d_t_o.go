// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VersionControlInformationDTO version control information d t o
// swagger:model VersionControlInformationDTO
type VersionControlInformationDTO struct {

	// The ID of the bucket that the flow is stored in
	BucketID string `json:"bucketId,omitempty"`

	// The name of the bucket that the flow is stored in
	// Read Only: true
	BucketName string `json:"bucketName,omitempty"`

	// The description of the flow
	FlowDescription string `json:"flowDescription,omitempty"`

	// The ID of the flow
	FlowID string `json:"flowId,omitempty"`

	// The name of the flow
	FlowName string `json:"flowName,omitempty"`

	// The ID of the Process Group that is under version control
	GroupID string `json:"groupId,omitempty"`

	// The ID of the registry that the flow is stored in
	RegistryID string `json:"registryId,omitempty"`

	// The name of the registry that the flow is stored in
	// Read Only: true
	RegistryName string `json:"registryName,omitempty"`

	// The current state of the Process Group, as it relates to the Versioned Flow
	// Read Only: true
	// Enum: [LOCALLY_MODIFIED STALE LOCALLY_MODIFIED_AND_STALE UP_TO_DATE SYNC_FAILURE]
	State string `json:"state,omitempty"`

	// Explanation of why the group is in the specified state
	// Read Only: true
	StateExplanation string `json:"stateExplanation,omitempty"`

	// The version of the flow
	Version int32 `json:"version,omitempty"`
}

// Validate validates this version control information d t o
func (m *VersionControlInformationDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var versionControlInformationDTOTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOCALLY_MODIFIED","STALE","LOCALLY_MODIFIED_AND_STALE","UP_TO_DATE","SYNC_FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionControlInformationDTOTypeStatePropEnum = append(versionControlInformationDTOTypeStatePropEnum, v)
	}
}

const (

	// VersionControlInformationDTOStateLOCALLYMODIFIED captures enum value "LOCALLY_MODIFIED"
	VersionControlInformationDTOStateLOCALLYMODIFIED string = "LOCALLY_MODIFIED"

	// VersionControlInformationDTOStateSTALE captures enum value "STALE"
	VersionControlInformationDTOStateSTALE string = "STALE"

	// VersionControlInformationDTOStateLOCALLYMODIFIEDANDSTALE captures enum value "LOCALLY_MODIFIED_AND_STALE"
	VersionControlInformationDTOStateLOCALLYMODIFIEDANDSTALE string = "LOCALLY_MODIFIED_AND_STALE"

	// VersionControlInformationDTOStateUPTODATE captures enum value "UP_TO_DATE"
	VersionControlInformationDTOStateUPTODATE string = "UP_TO_DATE"

	// VersionControlInformationDTOStateSYNCFAILURE captures enum value "SYNC_FAILURE"
	VersionControlInformationDTOStateSYNCFAILURE string = "SYNC_FAILURE"
)

// prop value enum
func (m *VersionControlInformationDTO) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, versionControlInformationDTOTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VersionControlInformationDTO) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionControlInformationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionControlInformationDTO) UnmarshalBinary(b []byte) error {
	var res VersionControlInformationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
