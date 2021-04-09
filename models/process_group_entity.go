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

// ProcessGroupEntity process group entity
// swagger:model ProcessGroupEntity
type ProcessGroupEntity struct {

	// The number of active remote ports in the process group.
	ActiveRemotePortCount int32 `json:"activeRemotePortCount,omitempty"`

	// The bulletins for this component.
	Bulletins []*BulletinEntity `json:"bulletins"`

	// component
	Component *ProcessGroupDTO `json:"component,omitempty"`

	// The number of disabled components in the process group.
	DisabledCount int32 `json:"disabledCount,omitempty"`

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The number of inactive remote ports in the process group.
	InactiveRemotePortCount int32 `json:"inactiveRemotePortCount,omitempty"`

	// The number of input ports in the process group.
	// Read Only: true
	InputPortCount int32 `json:"inputPortCount,omitempty"`

	// The number of invalid components in the process group.
	InvalidCount int32 `json:"invalidCount,omitempty"`

	// The number of local input ports in the process group.
	LocalInputPortCount int32 `json:"localInputPortCount,omitempty"`

	// The number of local output ports in the process group.
	LocalOutputPortCount int32 `json:"localOutputPortCount,omitempty"`

	// The number of locally modified and stale versioned process groups in the process group.
	LocallyModifiedAndStaleCount int32 `json:"locallyModifiedAndStaleCount,omitempty"`

	// The number of locally modified versioned process groups in the process group.
	LocallyModifiedCount int32 `json:"locallyModifiedCount,omitempty"`

	// The number of output ports in the process group.
	// Read Only: true
	OutputPortCount int32 `json:"outputPortCount,omitempty"`

	// The Parameter Context, or null if no Parameter Context has been bound to the Process Group
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`

	// The permissions for this component.
	Permissions *PermissionsDTO `json:"permissions,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// The number of public input ports in the process group.
	PublicInputPortCount int32 `json:"publicInputPortCount,omitempty"`

	// The number of public output ports in the process group.
	PublicOutputPortCount int32 `json:"publicOutputPortCount,omitempty"`

	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// The number of running components in this process group.
	RunningCount int32 `json:"runningCount,omitempty"`

	// The number of stale versioned process groups in the process group.
	StaleCount int32 `json:"staleCount,omitempty"`

	// The status of the process group.
	Status *ProcessGroupStatusDTO `json:"status,omitempty"`

	// The number of stopped components in the process group.
	StoppedCount int32 `json:"stoppedCount,omitempty"`

	// The number of versioned process groups in the process group that are unable to sync to a registry.
	SyncFailureCount int32 `json:"syncFailureCount,omitempty"`

	// The number of up to date versioned process groups in the process group.
	UpToDateCount int32 `json:"upToDateCount,omitempty"`

	// The URI for futures requests to the component.
	URI string `json:"uri,omitempty"`

	// Returns the Versioned Flow that describes the contents of the Versioned Flow to be imported
	// Read Only: true
	VersionedFlowSnapshot *VersionedFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`

	// The current state of the Process Group, as it relates to the Versioned Flow
	// Read Only: true
	// Enum: [LOCALLY_MODIFIED STALE LOCALLY_MODIFIED_AND_STALE UP_TO_DATE SYNC_FAILURE]
	VersionedFlowState string `json:"versionedFlowState,omitempty"`
}

// Validate validates this process group entity
func (m *ProcessGroupEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionedFlowSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionedFlowState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupEntity) validateBulletins(formats strfmt.Registry) error {

	if swag.IsZero(m.Bulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.Bulletins); i++ {
		if swag.IsZero(m.Bulletins[i]) { // not required
			continue
		}

		if m.Bulletins[i] != nil {
			if err := m.Bulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProcessGroupEntity) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupEntity) validateParameterContext(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterContext) { // not required
		return nil
	}

	if m.ParameterContext != nil {
		if err := m.ParameterContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterContext")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupEntity) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupEntity) validatePosition(formats strfmt.Registry) error {

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

func (m *ProcessGroupEntity) validateRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupEntity) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupEntity) validateVersionedFlowSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionedFlowSnapshot) { // not required
		return nil
	}

	if m.VersionedFlowSnapshot != nil {
		if err := m.VersionedFlowSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlowSnapshot")
			}
			return err
		}
	}

	return nil
}

var processGroupEntityTypeVersionedFlowStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOCALLY_MODIFIED","STALE","LOCALLY_MODIFIED_AND_STALE","UP_TO_DATE","SYNC_FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		processGroupEntityTypeVersionedFlowStatePropEnum = append(processGroupEntityTypeVersionedFlowStatePropEnum, v)
	}
}

const (

	// ProcessGroupEntityVersionedFlowStateLOCALLYMODIFIED captures enum value "LOCALLY_MODIFIED"
	ProcessGroupEntityVersionedFlowStateLOCALLYMODIFIED string = "LOCALLY_MODIFIED"

	// ProcessGroupEntityVersionedFlowStateSTALE captures enum value "STALE"
	ProcessGroupEntityVersionedFlowStateSTALE string = "STALE"

	// ProcessGroupEntityVersionedFlowStateLOCALLYMODIFIEDANDSTALE captures enum value "LOCALLY_MODIFIED_AND_STALE"
	ProcessGroupEntityVersionedFlowStateLOCALLYMODIFIEDANDSTALE string = "LOCALLY_MODIFIED_AND_STALE"

	// ProcessGroupEntityVersionedFlowStateUPTODATE captures enum value "UP_TO_DATE"
	ProcessGroupEntityVersionedFlowStateUPTODATE string = "UP_TO_DATE"

	// ProcessGroupEntityVersionedFlowStateSYNCFAILURE captures enum value "SYNC_FAILURE"
	ProcessGroupEntityVersionedFlowStateSYNCFAILURE string = "SYNC_FAILURE"
)

// prop value enum
func (m *ProcessGroupEntity) validateVersionedFlowStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, processGroupEntityTypeVersionedFlowStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProcessGroupEntity) validateVersionedFlowState(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionedFlowState) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionedFlowStateEnum("versionedFlowState", "body", m.VersionedFlowState); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupEntity) UnmarshalBinary(b []byte) error {
	var res ProcessGroupEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}