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

// ConnectionDTO connection d t o
// swagger:model ConnectionDTO
type ConnectionDTO struct {

	// The relationships that the source of the connection currently supports.
	// Read Only: true
	// Unique: true
	AvailableRelationships []string `json:"availableRelationships"`

	// The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureDataSizeThreshold string `json:"backPressureDataSizeThreshold,omitempty"`

	// The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureObjectThreshold int64 `json:"backPressureObjectThreshold,omitempty"`

	// The bend points on the connection.
	Bends []*PositionDTO `json:"bends"`

	// The destination of the connection.
	Destination *ConnectableDTO `json:"destination,omitempty"`

	// The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it.
	FlowFileExpiration string `json:"flowFileExpiration,omitempty"`

	// The z index of the connection.
	GetzIndex int64 `json:"getzIndex,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The index of the bend point where to place the connection label.
	LabelIndex int32 `json:"labelIndex,omitempty"`

	// Whether or not data should be compressed when being transferred between nodes in the cluster.
	// Enum: [DO_NOT_COMPRESS COMPRESS_ATTRIBUTES_ONLY COMPRESS_ATTRIBUTES_AND_CONTENT]
	LoadBalanceCompression string `json:"loadBalanceCompression,omitempty"`

	// The FlowFile Attribute to use for determining which node a FlowFile will go to if the Load Balancing Strategy is set to PARTITION_BY_ATTRIBUTE
	LoadBalancePartitionAttribute string `json:"loadBalancePartitionAttribute,omitempty"`

	// The current status of the Connection's Load Balancing Activities. Status can indicate that Load Balancing is not configured for the connection, that Load Balancing is configured but inactive (not currently transferring data to another node), or that Load Balancing is configured and actively transferring data to another node.
	// Read Only: true
	// Enum: [LOAD_BALANCE_NOT_CONFIGURED LOAD_BALANCE_INACTIVE LOAD_BALANCE_ACTIVE]
	LoadBalanceStatus string `json:"loadBalanceStatus,omitempty"`

	// How to load balance the data in this Connection across the nodes in the cluster.
	// Enum: [DO_NOT_LOAD_BALANCE PARTITION_BY_ATTRIBUTE ROUND_ROBIN SINGLE_NODE]
	LoadBalanceStrategy string `json:"loadBalanceStrategy,omitempty"`

	// The name of the connection.
	Name string `json:"name,omitempty"`

	// The id of parent process group of this component if applicable.
	ParentGroupID string `json:"parentGroupId,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// The comparators used to prioritize the queue.
	Prioritizers []string `json:"prioritizers"`

	// The selected relationship that comprise the connection.
	// Unique: true
	SelectedRelationships []string `json:"selectedRelationships"`

	// The source of the connection.
	Source *ConnectableDTO `json:"source,omitempty"`

	// The ID of the corresponding component that is under version control
	VersionedComponentID string `json:"versionedComponentId,omitempty"`
}

// Validate validates this connection d t o
func (m *ConnectionDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBends(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalanceCompression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalanceStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionDTO) validateAvailableRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableRelationships) { // not required
		return nil
	}

	if err := validate.UniqueItems("availableRelationships", "body", m.AvailableRelationships); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionDTO) validateBends(formats strfmt.Registry) error {

	if swag.IsZero(m.Bends) { // not required
		return nil
	}

	for i := 0; i < len(m.Bends); i++ {
		if swag.IsZero(m.Bends[i]) { // not required
			continue
		}

		if m.Bends[i] != nil {
			if err := m.Bends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConnectionDTO) validateDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

var connectionDTOTypeLoadBalanceCompressionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DO_NOT_COMPRESS","COMPRESS_ATTRIBUTES_ONLY","COMPRESS_ATTRIBUTES_AND_CONTENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionDTOTypeLoadBalanceCompressionPropEnum = append(connectionDTOTypeLoadBalanceCompressionPropEnum, v)
	}
}

const (

	// ConnectionDTOLoadBalanceCompressionDONOTCOMPRESS captures enum value "DO_NOT_COMPRESS"
	ConnectionDTOLoadBalanceCompressionDONOTCOMPRESS string = "DO_NOT_COMPRESS"

	// ConnectionDTOLoadBalanceCompressionCOMPRESSATTRIBUTESONLY captures enum value "COMPRESS_ATTRIBUTES_ONLY"
	ConnectionDTOLoadBalanceCompressionCOMPRESSATTRIBUTESONLY string = "COMPRESS_ATTRIBUTES_ONLY"

	// ConnectionDTOLoadBalanceCompressionCOMPRESSATTRIBUTESANDCONTENT captures enum value "COMPRESS_ATTRIBUTES_AND_CONTENT"
	ConnectionDTOLoadBalanceCompressionCOMPRESSATTRIBUTESANDCONTENT string = "COMPRESS_ATTRIBUTES_AND_CONTENT"
)

// prop value enum
func (m *ConnectionDTO) validateLoadBalanceCompressionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, connectionDTOTypeLoadBalanceCompressionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConnectionDTO) validateLoadBalanceCompression(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalanceCompression) { // not required
		return nil
	}

	// value enum
	if err := m.validateLoadBalanceCompressionEnum("loadBalanceCompression", "body", m.LoadBalanceCompression); err != nil {
		return err
	}

	return nil
}

var connectionDTOTypeLoadBalanceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOAD_BALANCE_NOT_CONFIGURED","LOAD_BALANCE_INACTIVE","LOAD_BALANCE_ACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionDTOTypeLoadBalanceStatusPropEnum = append(connectionDTOTypeLoadBalanceStatusPropEnum, v)
	}
}

const (

	// ConnectionDTOLoadBalanceStatusLOADBALANCENOTCONFIGURED captures enum value "LOAD_BALANCE_NOT_CONFIGURED"
	ConnectionDTOLoadBalanceStatusLOADBALANCENOTCONFIGURED string = "LOAD_BALANCE_NOT_CONFIGURED"

	// ConnectionDTOLoadBalanceStatusLOADBALANCEINACTIVE captures enum value "LOAD_BALANCE_INACTIVE"
	ConnectionDTOLoadBalanceStatusLOADBALANCEINACTIVE string = "LOAD_BALANCE_INACTIVE"

	// ConnectionDTOLoadBalanceStatusLOADBALANCEACTIVE captures enum value "LOAD_BALANCE_ACTIVE"
	ConnectionDTOLoadBalanceStatusLOADBALANCEACTIVE string = "LOAD_BALANCE_ACTIVE"
)

// prop value enum
func (m *ConnectionDTO) validateLoadBalanceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, connectionDTOTypeLoadBalanceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConnectionDTO) validateLoadBalanceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalanceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLoadBalanceStatusEnum("loadBalanceStatus", "body", m.LoadBalanceStatus); err != nil {
		return err
	}

	return nil
}

var connectionDTOTypeLoadBalanceStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DO_NOT_LOAD_BALANCE","PARTITION_BY_ATTRIBUTE","ROUND_ROBIN","SINGLE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionDTOTypeLoadBalanceStrategyPropEnum = append(connectionDTOTypeLoadBalanceStrategyPropEnum, v)
	}
}

const (

	// ConnectionDTOLoadBalanceStrategyDONOTLOADBALANCE captures enum value "DO_NOT_LOAD_BALANCE"
	ConnectionDTOLoadBalanceStrategyDONOTLOADBALANCE string = "DO_NOT_LOAD_BALANCE"

	// ConnectionDTOLoadBalanceStrategyPARTITIONBYATTRIBUTE captures enum value "PARTITION_BY_ATTRIBUTE"
	ConnectionDTOLoadBalanceStrategyPARTITIONBYATTRIBUTE string = "PARTITION_BY_ATTRIBUTE"

	// ConnectionDTOLoadBalanceStrategyROUNDROBIN captures enum value "ROUND_ROBIN"
	ConnectionDTOLoadBalanceStrategyROUNDROBIN string = "ROUND_ROBIN"

	// ConnectionDTOLoadBalanceStrategySINGLENODE captures enum value "SINGLE_NODE"
	ConnectionDTOLoadBalanceStrategySINGLENODE string = "SINGLE_NODE"
)

// prop value enum
func (m *ConnectionDTO) validateLoadBalanceStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, connectionDTOTypeLoadBalanceStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConnectionDTO) validateLoadBalanceStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalanceStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateLoadBalanceStrategyEnum("loadBalanceStrategy", "body", m.LoadBalanceStrategy); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionDTO) validatePosition(formats strfmt.Registry) error {

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

func (m *ConnectionDTO) validateSelectedRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedRelationships) { // not required
		return nil
	}

	if err := validate.UniqueItems("selectedRelationships", "body", m.SelectedRelationships); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionDTO) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionDTO) UnmarshalBinary(b []byte) error {
	var res ConnectionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
