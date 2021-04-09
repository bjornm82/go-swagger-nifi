// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClusterSummaryDTO cluster summary d t o
// swagger:model ClusterSummaryDTO
type ClusterSummaryDTO struct {

	// Whether this NiFi instance is clustered.
	Clustered bool `json:"clustered,omitempty"`

	// The number of nodes that are currently connected to the cluster
	ConnectedNodeCount int32 `json:"connectedNodeCount,omitempty"`

	// When clustered, reports the number of nodes connected vs the number of nodes in the cluster.
	ConnectedNodes string `json:"connectedNodes,omitempty"`

	// Whether this NiFi instance is connected to a cluster.
	ConnectedToCluster bool `json:"connectedToCluster,omitempty"`

	// The number of nodes in the cluster, regardless of whether or not they are connected
	TotalNodeCount int32 `json:"totalNodeCount,omitempty"`
}

// Validate validates this cluster summary d t o
func (m *ClusterSummaryDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSummaryDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSummaryDTO) UnmarshalBinary(b []byte) error {
	var res ClusterSummaryDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
