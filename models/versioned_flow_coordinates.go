// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VersionedFlowCoordinates versioned flow coordinates
//
// swagger:model VersionedFlowCoordinates
type VersionedFlowCoordinates struct {

	// The UUID of the bucket that the flow resides in
	BucketID string `json:"bucketId,omitempty"`

	// The UUID of the flow
	FlowID string `json:"flowId,omitempty"`

	// Whether or not these coordinates point to the latest version of the flow
	Latest bool `json:"latest,omitempty"`

	// The URL of the Flow Registry that contains the flow
	RegistryURL string `json:"registryUrl,omitempty"`

	// The version of the flow
	Version int32 `json:"version,omitempty"`
}

// Validate validates this versioned flow coordinates
func (m *VersionedFlowCoordinates) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VersionedFlowCoordinates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowCoordinates) UnmarshalBinary(b []byte) error {
	var res VersionedFlowCoordinates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
