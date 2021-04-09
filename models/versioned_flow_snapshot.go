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

// VersionedFlowSnapshot versioned flow snapshot
// swagger:model VersionedFlowSnapshot
type VersionedFlowSnapshot struct {

	// The bucket where the flow is located
	// Read Only: true
	Bucket *Bucket `json:"bucket,omitempty"`

	// The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow.
	ExternalControllerServices map[string]ExternalControllerServiceReference `json:"externalControllerServices,omitempty"`

	// The flow this snapshot is for
	// Read Only: true
	Flow *VersionedFlow `json:"flow,omitempty"`

	// The contents of the versioned flow
	// Required: true
	FlowContents *VersionedProcessGroup `json:"flowContents"`

	// The optional encoding version of the flow contents.
	FlowEncodingVersion string `json:"flowEncodingVersion,omitempty"`

	// latest
	Latest bool `json:"latest,omitempty"`

	// The parameter contexts referenced by process groups in the flow contents. The mapping is from the name of the context to the context instance, and it is expected that any context in this map is referenced by at least one process group in this flow.
	ParameterContexts map[string]VersionedParameterContext `json:"parameterContexts,omitempty"`

	// The metadata for this snapshot
	// Required: true
	SnapshotMetadata *VersionedFlowSnapshotMetadata `json:"snapshotMetadata"`
}

// Validate validates this versioned flow snapshot
func (m *VersionedFlowSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalControllerServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterContexts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowSnapshot) validateBucket(formats strfmt.Registry) error {

	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *VersionedFlowSnapshot) validateExternalControllerServices(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalControllerServices) { // not required
		return nil
	}

	for k := range m.ExternalControllerServices {

		if err := validate.Required("externalControllerServices"+"."+k, "body", m.ExternalControllerServices[k]); err != nil {
			return err
		}
		if val, ok := m.ExternalControllerServices[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionedFlowSnapshot) validateFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.Flow) { // not required
		return nil
	}

	if m.Flow != nil {
		if err := m.Flow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *VersionedFlowSnapshot) validateFlowContents(formats strfmt.Registry) error {

	if err := validate.Required("flowContents", "body", m.FlowContents); err != nil {
		return err
	}

	if m.FlowContents != nil {
		if err := m.FlowContents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowContents")
			}
			return err
		}
	}

	return nil
}

func (m *VersionedFlowSnapshot) validateParameterContexts(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterContexts) { // not required
		return nil
	}

	for k := range m.ParameterContexts {

		if err := validate.Required("parameterContexts"+"."+k, "body", m.ParameterContexts[k]); err != nil {
			return err
		}
		if val, ok := m.ParameterContexts[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionedFlowSnapshot) validateSnapshotMetadata(formats strfmt.Registry) error {

	if err := validate.Required("snapshotMetadata", "body", m.SnapshotMetadata); err != nil {
		return err
	}

	if m.SnapshotMetadata != nil {
		if err := m.SnapshotMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedFlowSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowSnapshot) UnmarshalBinary(b []byte) error {
	var res VersionedFlowSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}