// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProcessGroupDTO remote process group d t o
//
// swagger:model RemoteProcessGroupDTO
type RemoteProcessGroupDTO struct {

	// The number of active remote input ports.
	ActiveRemoteInputPortCount int32 `json:"activeRemoteInputPortCount,omitempty"`

	// The number of active remote output ports.
	ActiveRemoteOutputPortCount int32 `json:"activeRemoteOutputPortCount,omitempty"`

	// Any remote authorization issues for the remote process group.
	AuthorizationIssues []string `json:"authorizationIssues"`

	// The comments for the remote process group.
	Comments string `json:"comments,omitempty"`

	// The time period used for the timeout when communicating with the target.
	CommunicationsTimeout string `json:"communicationsTimeout,omitempty"`

	// The contents of the remote process group. Will contain available input/output ports.
	Contents *RemoteProcessGroupContentsDTO `json:"contents,omitempty"`

	// The timestamp when this remote process group was last refreshed.
	FlowRefreshed string `json:"flowRefreshed,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The number of inactive remote input ports.
	InactiveRemoteInputPortCount int32 `json:"inactiveRemoteInputPortCount,omitempty"`

	// The number of inactive remote output ports.
	InactiveRemoteOutputPortCount int32 `json:"inactiveRemoteOutputPortCount,omitempty"`

	// The number of remote input ports currently available on the target.
	InputPortCount int32 `json:"inputPortCount,omitempty"`

	// The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier.
	LocalNetworkInterface string `json:"localNetworkInterface,omitempty"`

	// The name of the remote process group.
	Name string `json:"name,omitempty"`

	// The number of remote output ports currently available on the target.
	OutputPortCount int32 `json:"outputPortCount,omitempty"`

	// The id of parent process group of this component if applicable.
	ParentGroupID string `json:"parentGroupId,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// proxy host
	ProxyHost string `json:"proxyHost,omitempty"`

	// proxy password
	ProxyPassword string `json:"proxyPassword,omitempty"`

	// proxy port
	ProxyPort int32 `json:"proxyPort,omitempty"`

	// proxy user
	ProxyUser string `json:"proxyUser,omitempty"`

	// Whether the target is running securely.
	TargetSecure bool `json:"targetSecure,omitempty"`

	// The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first url in the urls. If neither target uri nor uris are set, then returns null.
	TargetURI string `json:"targetUri,omitempty"`

	// The target URI of the remote process group. If target uris is not set but target uri is set, then returns a collection containing the single target uri. If neither target uris nor uris are set, then returns null.
	TargetUris string `json:"targetUris,omitempty"`

	// Whether the remote process group is actively transmitting.
	Transmitting bool `json:"transmitting,omitempty"`

	// transport protocol
	TransportProtocol string `json:"transportProtocol,omitempty"`

	// The validation errors for the remote process group. These validation errors represent the problems with the remote process group that must be resolved before it can transmit.
	ValidationErrors []string `json:"validationErrors"`

	// The ID of the corresponding component that is under version control
	VersionedComponentID string `json:"versionedComponentId,omitempty"`

	// When yielding, this amount of time must elapse before the remote process group is scheduled again.
	YieldDuration string `json:"yieldDuration,omitempty"`
}

// Validate validates this remote process group d t o
func (m *RemoteProcessGroupDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProcessGroupDTO) validateContents(formats strfmt.Registry) error {

	if swag.IsZero(m.Contents) { // not required
		return nil
	}

	if m.Contents != nil {
		if err := m.Contents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contents")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteProcessGroupDTO) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProcessGroupDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProcessGroupDTO) UnmarshalBinary(b []byte) error {
	var res RemoteProcessGroupDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
