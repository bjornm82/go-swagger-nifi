// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DropRequestDTO drop request d t o
//
// swagger:model DropRequestDTO
type DropRequestDTO struct {

	// The count and size of flow files currently queued.
	Current string `json:"current,omitempty"`

	// The number of flow files currently queued.
	CurrentCount int32 `json:"currentCount,omitempty"`

	// The size of flow files currently queued in bytes.
	CurrentSize int64 `json:"currentSize,omitempty"`

	// The count and size of flow files that have been dropped thus far.
	Dropped string `json:"dropped,omitempty"`

	// The number of flow files that have been dropped thus far.
	DroppedCount int32 `json:"droppedCount,omitempty"`

	// The size of flow files that have been dropped thus far in bytes.
	DroppedSize int64 `json:"droppedSize,omitempty"`

	// The reason, if any, that this drop request failed.
	FailureReason string `json:"failureReason,omitempty"`

	// Whether the query has finished.
	Finished bool `json:"finished,omitempty"`

	// The id for this drop request.
	ID string `json:"id,omitempty"`

	// The last time this drop request was updated.
	LastUpdated string `json:"lastUpdated,omitempty"`

	// The count and size of flow files to be dropped as a result of this request.
	Original string `json:"original,omitempty"`

	// The number of flow files to be dropped as a result of this request.
	OriginalCount int32 `json:"originalCount,omitempty"`

	// The size of flow files to be dropped as a result of this request in bytes.
	OriginalSize int64 `json:"originalSize,omitempty"`

	// The current percent complete.
	PercentCompleted int32 `json:"percentCompleted,omitempty"`

	// The current state of the drop request.
	State string `json:"state,omitempty"`

	// The timestamp when the query was submitted.
	SubmissionTime string `json:"submissionTime,omitempty"`

	// The URI for future requests to this drop request.
	URI string `json:"uri,omitempty"`
}

// Validate validates this drop request d t o
func (m *DropRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DropRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DropRequestDTO) UnmarshalBinary(b []byte) error {
	var res DropRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
