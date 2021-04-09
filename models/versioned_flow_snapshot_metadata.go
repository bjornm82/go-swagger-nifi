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

// VersionedFlowSnapshotMetadata versioned flow snapshot metadata
// swagger:model VersionedFlowSnapshotMetadata
type VersionedFlowSnapshotMetadata struct {

	// The user that created this snapshot of the flow.
	// Read Only: true
	Author string `json:"author,omitempty"`

	// The identifier of the bucket this snapshot belongs to.
	// Required: true
	BucketIdentifier *string `json:"bucketIdentifier"`

	// The comments provided by the user when creating the snapshot.
	Comments string `json:"comments,omitempty"`

	// The identifier of the flow this snapshot belongs to.
	// Required: true
	FlowIdentifier *string `json:"flowIdentifier"`

	// An WebLink to this entity.
	// Read Only: true
	Link *JaxbLink `json:"link,omitempty"`

	// The timestamp when the flow was saved, as milliseconds since epoch.
	// Read Only: true
	// Minimum: 1
	Timestamp int64 `json:"timestamp,omitempty"`

	// The version of this snapshot of the flow.
	// Required: true
	// Minimum: -1
	Version *int32 `json:"version"`
}

// Validate validates this versioned flow snapshot metadata
func (m *VersionedFlowSnapshotMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowSnapshotMetadata) validateBucketIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("bucketIdentifier", "body", m.BucketIdentifier); err != nil {
		return err
	}

	return nil
}

func (m *VersionedFlowSnapshotMetadata) validateFlowIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("flowIdentifier", "body", m.FlowIdentifier); err != nil {
		return err
	}

	return nil
}

func (m *VersionedFlowSnapshotMetadata) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *VersionedFlowSnapshotMetadata) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.MinimumInt("timestamp", "body", int64(m.Timestamp), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *VersionedFlowSnapshotMetadata) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), -1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedFlowSnapshotMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowSnapshotMetadata) UnmarshalBinary(b []byte) error {
	var res VersionedFlowSnapshotMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}