// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reporting tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reporting tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ClearStateReportingTasks(params *ClearStateReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*ClearStateReportingTasksOK, error)

	GetPropertyDescriptorReportingTasks(params *GetPropertyDescriptorReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertyDescriptorReportingTasksOK, error)

	GetReportingTask(params *GetReportingTaskParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportingTaskOK, error)

	GetStateReportingTasks(params *GetStateReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetStateReportingTasksOK, error)

	RemoveReportingTask(params *RemoveReportingTaskParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveReportingTaskOK, error)

	UpdateReportingTask(params *UpdateReportingTaskParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateReportingTaskOK, error)

	UpdateRunStatusReportingTasks(params *UpdateRunStatusReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRunStatusReportingTasksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ClearStateReportingTasks clears the state for a reporting task
*/
func (a *Client) ClearStateReportingTasks(params *ClearStateReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*ClearStateReportingTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClearStateReportingTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clearStateReportingTasks",
		Method:             "POST",
		PathPattern:        "/reporting-tasks/{id}/state/clear-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClearStateReportingTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClearStateReportingTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clearStateReportingTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPropertyDescriptorReportingTasks gets a reporting task property descriptor
*/
func (a *Client) GetPropertyDescriptorReportingTasks(params *GetPropertyDescriptorReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertyDescriptorReportingTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPropertyDescriptorReportingTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPropertyDescriptorReportingTasks",
		Method:             "GET",
		PathPattern:        "/reporting-tasks/{id}/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPropertyDescriptorReportingTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPropertyDescriptorReportingTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPropertyDescriptorReportingTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReportingTask gets a reporting task
*/
func (a *Client) GetReportingTask(params *GetReportingTaskParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportingTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportingTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReportingTask",
		Method:             "GET",
		PathPattern:        "/reporting-tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReportingTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReportingTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReportingTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStateReportingTasks gets the state for a reporting task
*/
func (a *Client) GetStateReportingTasks(params *GetStateReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*GetStateReportingTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStateReportingTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStateReportingTasks",
		Method:             "GET",
		PathPattern:        "/reporting-tasks/{id}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStateReportingTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStateReportingTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStateReportingTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveReportingTask deletes a reporting task
*/
func (a *Client) RemoveReportingTask(params *RemoveReportingTaskParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveReportingTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveReportingTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeReportingTask",
		Method:             "DELETE",
		PathPattern:        "/reporting-tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveReportingTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveReportingTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeReportingTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateReportingTask updates a reporting task
*/
func (a *Client) UpdateReportingTask(params *UpdateReportingTaskParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateReportingTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReportingTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateReportingTask",
		Method:             "PUT",
		PathPattern:        "/reporting-tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateReportingTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateReportingTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateReportingTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRunStatusReportingTasks updates run status of a reporting task
*/
func (a *Client) UpdateRunStatusReportingTasks(params *UpdateRunStatusReportingTasksParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRunStatusReportingTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRunStatusReportingTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRunStatusReportingTasks",
		Method:             "PUT",
		PathPattern:        "/reporting-tasks/{id}/run-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRunStatusReportingTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRunStatusReportingTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRunStatusReportingTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
