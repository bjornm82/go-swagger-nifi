// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProcessGroupPortDTO remote process group port d t o
//
// swagger:model RemoteProcessGroupPortDTO
type RemoteProcessGroupPortDTO struct {

	// The batch settings for data transmission.
	BatchSettings *BatchSettingsDTO `json:"batchSettings,omitempty"`

	// The comments as configured on the target port.
	Comments string `json:"comments,omitempty"`

	// The number of task that may transmit flowfiles to the target port concurrently.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`

	// Whether the port has either an incoming or outgoing connection.
	Connected bool `json:"connected,omitempty"`

	// Whether the target port exists.
	Exists bool `json:"exists,omitempty"`

	// The id of the remote process group that the port resides in.
	GroupID string `json:"groupId,omitempty"`

	// The id of the port.
	ID string `json:"id,omitempty"`

	// The name of the target port.
	Name string `json:"name,omitempty"`

	// The id of the target port.
	TargetID string `json:"targetId,omitempty"`

	// Whether the target port is running.
	TargetRunning bool `json:"targetRunning,omitempty"`

	// Whether the remote port is configured for transmission.
	Transmitting bool `json:"transmitting,omitempty"`

	// Whether the flowfiles are compressed when sent to the target port.
	UseCompression bool `json:"useCompression,omitempty"`

	// The ID of the corresponding component that is under version control
	VersionedComponentID string `json:"versionedComponentId,omitempty"`
}

// Validate validates this remote process group port d t o
func (m *RemoteProcessGroupPortDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBatchSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProcessGroupPortDTO) validateBatchSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.BatchSettings) { // not required
		return nil
	}

	if m.BatchSettings != nil {
		if err := m.BatchSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("batchSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProcessGroupPortDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProcessGroupPortDTO) UnmarshalBinary(b []byte) error {
	var res RemoteProcessGroupPortDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
