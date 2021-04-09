// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reporting tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reporting tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
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
		ConsumesMediaTypes: []string{"application/json"},
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
	return result.(*GetReportingTaskOK), nil

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
		ConsumesMediaTypes: []string{"application/json"},
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
	return result.(*RemoveReportingTaskOK), nil

}

/*
ReportingTasksClearState clears the state for a reporting task
*/
func (a *Client) ReportingTasksClearState(params *ReportingTasksClearStateParams, authInfo runtime.ClientAuthInfoWriter) (*ReportingTasksClearStateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportingTasksClearStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportingTasksClearState",
		Method:             "POST",
		PathPattern:        "/reporting-tasks/{id}/state/clear-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportingTasksClearStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReportingTasksClearStateCreated), nil

}

/*
ReportingTasksGetPropertyDescriptor gets a reporting task property descriptor
*/
func (a *Client) ReportingTasksGetPropertyDescriptor(params *ReportingTasksGetPropertyDescriptorParams, authInfo runtime.ClientAuthInfoWriter) (*ReportingTasksGetPropertyDescriptorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportingTasksGetPropertyDescriptorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportingTasksGetPropertyDescriptor",
		Method:             "GET",
		PathPattern:        "/reporting-tasks/{id}/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportingTasksGetPropertyDescriptorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReportingTasksGetPropertyDescriptorOK), nil

}

/*
ReportingTasksGetState gets the state for a reporting task
*/
func (a *Client) ReportingTasksGetState(params *ReportingTasksGetStateParams, authInfo runtime.ClientAuthInfoWriter) (*ReportingTasksGetStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportingTasksGetStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportingTasksGetState",
		Method:             "GET",
		PathPattern:        "/reporting-tasks/{id}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportingTasksGetStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReportingTasksGetStateOK), nil

}

/*
ReportingTasksUpdateRunStatus updates run status of a reporting task
*/
func (a *Client) ReportingTasksUpdateRunStatus(params *ReportingTasksUpdateRunStatusParams, authInfo runtime.ClientAuthInfoWriter) (*ReportingTasksUpdateRunStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportingTasksUpdateRunStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportingTasksUpdateRunStatus",
		Method:             "PUT",
		PathPattern:        "/reporting-tasks/{id}/run-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportingTasksUpdateRunStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReportingTasksUpdateRunStatusOK), nil

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
	return result.(*UpdateReportingTaskOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
