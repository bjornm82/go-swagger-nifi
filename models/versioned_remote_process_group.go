// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VersionedRemoteProcessGroup versioned remote process group
// swagger:model VersionedRemoteProcessGroup
type VersionedRemoteProcessGroup struct {

	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`

	// The time period used for the timeout when communicating with the target.
	CommunicationsTimeout string `json:"communicationsTimeout,omitempty"`

	// component type
	// Enum: [CONNECTION PROCESSOR PROCESS_GROUP REMOTE_PROCESS_GROUP INPUT_PORT OUTPUT_PORT REMOTE_INPUT_PORT REMOTE_OUTPUT_PORT FUNNEL LABEL CONTROLLER_SERVICE]
	ComponentType string `json:"componentType,omitempty"`

	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`

	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`

	// A Set of Input Ports that can be connected to, in order to send data to the remote NiFi instance
	// Unique: true
	InputPorts []*VersionedRemoteGroupPort `json:"inputPorts"`

	// The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier.
	LocalNetworkInterface string `json:"localNetworkInterface,omitempty"`

	// The component's name
	Name string `json:"name,omitempty"`

	// A Set of Output Ports that can be connected to, in order to pull data from the remote NiFi instance
	// Unique: true
	OutputPorts []*VersionedRemoteGroupPort `json:"outputPorts"`

	// The component's position on the graph
	Position *Position `json:"position,omitempty"`

	// proxy host
	ProxyHost string `json:"proxyHost,omitempty"`

	// proxy port
	ProxyPort int32 `json:"proxyPort,omitempty"`

	// proxy user
	ProxyUser string `json:"proxyUser,omitempty"`

	// [DEPRECATED] The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first uri in the uris. If neither target uri nor uris are set, then returns null.
	TargetURI string `json:"targetUri,omitempty"`

	// The target URIs of the remote process group. If target uris is not set but target uri is set, then returns the single target uri. If neither target uris nor target uri is set, then returns null.
	TargetUris string `json:"targetUris,omitempty"`

	// The Transport Protocol that is used for Site-to-Site communications
	// Enum: [RAW HTTP]
	TransportProtocol string `json:"transportProtocol,omitempty"`

	// When yielding, this amount of time must elapse before the remote process group is scheduled again.
	YieldDuration string `json:"yieldDuration,omitempty"`
}

// Validate validates this versioned remote process group
func (m *VersionedRemoteProcessGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var versionedRemoteProcessGroupTypeComponentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONNECTION","PROCESSOR","PROCESS_GROUP","REMOTE_PROCESS_GROUP","INPUT_PORT","OUTPUT_PORT","REMOTE_INPUT_PORT","REMOTE_OUTPUT_PORT","FUNNEL","LABEL","CONTROLLER_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedRemoteProcessGroupTypeComponentTypePropEnum = append(versionedRemoteProcessGroupTypeComponentTypePropEnum, v)
	}
}

const (

	// VersionedRemoteProcessGroupComponentTypeCONNECTION captures enum value "CONNECTION"
	VersionedRemoteProcessGroupComponentTypeCONNECTION string = "CONNECTION"

	// VersionedRemoteProcessGroupComponentTypePROCESSOR captures enum value "PROCESSOR"
	VersionedRemoteProcessGroupComponentTypePROCESSOR string = "PROCESSOR"

	// VersionedRemoteProcessGroupComponentTypePROCESSGROUP captures enum value "PROCESS_GROUP"
	VersionedRemoteProcessGroupComponentTypePROCESSGROUP string = "PROCESS_GROUP"

	// VersionedRemoteProcessGroupComponentTypeREMOTEPROCESSGROUP captures enum value "REMOTE_PROCESS_GROUP"
	VersionedRemoteProcessGroupComponentTypeREMOTEPROCESSGROUP string = "REMOTE_PROCESS_GROUP"

	// VersionedRemoteProcessGroupComponentTypeINPUTPORT captures enum value "INPUT_PORT"
	VersionedRemoteProcessGroupComponentTypeINPUTPORT string = "INPUT_PORT"

	// VersionedRemoteProcessGroupComponentTypeOUTPUTPORT captures enum value "OUTPUT_PORT"
	VersionedRemoteProcessGroupComponentTypeOUTPUTPORT string = "OUTPUT_PORT"

	// VersionedRemoteProcessGroupComponentTypeREMOTEINPUTPORT captures enum value "REMOTE_INPUT_PORT"
	VersionedRemoteProcessGroupComponentTypeREMOTEINPUTPORT string = "REMOTE_INPUT_PORT"

	// VersionedRemoteProcessGroupComponentTypeREMOTEOUTPUTPORT captures enum value "REMOTE_OUTPUT_PORT"
	VersionedRemoteProcessGroupComponentTypeREMOTEOUTPUTPORT string = "REMOTE_OUTPUT_PORT"

	// VersionedRemoteProcessGroupComponentTypeFUNNEL captures enum value "FUNNEL"
	VersionedRemoteProcessGroupComponentTypeFUNNEL string = "FUNNEL"

	// VersionedRemoteProcessGroupComponentTypeLABEL captures enum value "LABEL"
	VersionedRemoteProcessGroupComponentTypeLABEL string = "LABEL"

	// VersionedRemoteProcessGroupComponentTypeCONTROLLERSERVICE captures enum value "CONTROLLER_SERVICE"
	VersionedRemoteProcessGroupComponentTypeCONTROLLERSERVICE string = "CONTROLLER_SERVICE"
)

// prop value enum
func (m *VersionedRemoteProcessGroup) validateComponentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, versionedRemoteProcessGroupTypeComponentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VersionedRemoteProcessGroup) validateComponentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ComponentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentTypeEnum("componentType", "body", m.ComponentType); err != nil {
		return err
	}

	return nil
}

func (m *VersionedRemoteProcessGroup) validateInputPorts(formats strfmt.Registry) error {

	if swag.IsZero(m.InputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("inputPorts", "body", m.InputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.InputPorts); i++ {
		if swag.IsZero(m.InputPorts[i]) { // not required
			continue
		}

		if m.InputPorts[i] != nil {
			if err := m.InputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VersionedRemoteProcessGroup) validateOutputPorts(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("outputPorts", "body", m.OutputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.OutputPorts); i++ {
		if swag.IsZero(m.OutputPorts[i]) { // not required
			continue
		}

		if m.OutputPorts[i] != nil {
			if err := m.OutputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VersionedRemoteProcessGroup) validatePosition(formats strfmt.Registry) error {

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

var versionedRemoteProcessGroupTypeTransportProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RAW","HTTP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedRemoteProcessGroupTypeTransportProtocolPropEnum = append(versionedRemoteProcessGroupTypeTransportProtocolPropEnum, v)
	}
}

const (

	// VersionedRemoteProcessGroupTransportProtocolRAW captures enum value "RAW"
	VersionedRemoteProcessGroupTransportProtocolRAW string = "RAW"

	// VersionedRemoteProcessGroupTransportProtocolHTTP captures enum value "HTTP"
	VersionedRemoteProcessGroupTransportProtocolHTTP string = "HTTP"
)

// prop value enum
func (m *VersionedRemoteProcessGroup) validateTransportProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, versionedRemoteProcessGroupTypeTransportProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VersionedRemoteProcessGroup) validateTransportProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.TransportProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransportProtocolEnum("transportProtocol", "body", m.TransportProtocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedRemoteProcessGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedRemoteProcessGroup) UnmarshalBinary(b []byte) error {
	var res VersionedRemoteProcessGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
