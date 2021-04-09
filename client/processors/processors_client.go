// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new processors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for processors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteProcessor deletes a processor
*/
func (a *Client) DeleteProcessor(params *DeleteProcessorParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProcessorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProcessorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProcessor",
		Method:             "DELETE",
		PathPattern:        "/processors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProcessorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProcessorOK), nil

}

/*
GetProcessor gets a processor
*/
func (a *Client) GetProcessor(params *GetProcessorParams, authInfo runtime.ClientAuthInfoWriter) (*GetProcessorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProcessor",
		Method:             "GET",
		PathPattern:        "/processors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProcessorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProcessorOK), nil

}

/*
GetProcessorDiagnostics gets diagnostics information about a processor

Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) GetProcessorDiagnostics(params *GetProcessorDiagnosticsParams, authInfo runtime.ClientAuthInfoWriter) (*GetProcessorDiagnosticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessorDiagnosticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProcessorDiagnostics",
		Method:             "GET",
		PathPattern:        "/processors/{id}/diagnostics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProcessorDiagnosticsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProcessorDiagnosticsOK), nil

}

/*
GetProcessorRunStatusDetails submits a query to retrieve the run status details of all processors that are in the given list of processor ids
*/
func (a *Client) GetProcessorRunStatusDetails(params *GetProcessorRunStatusDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetProcessorRunStatusDetailsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessorRunStatusDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProcessorRunStatusDetails",
		Method:             "POST",
		PathPattern:        "/processors/run-status-details/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProcessorRunStatusDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProcessorRunStatusDetailsCreated), nil

}

/*
ProcessorsClearState clears the state for a processor
*/
func (a *Client) ProcessorsClearState(params *ProcessorsClearStateParams, authInfo runtime.ClientAuthInfoWriter) (*ProcessorsClearStateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProcessorsClearStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "processorsClearState",
		Method:             "POST",
		PathPattern:        "/processors/{id}/state/clear-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProcessorsClearStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProcessorsClearStateCreated), nil

}

/*
ProcessorsGetPropertyDescriptor gets the descriptor for a processor property
*/
func (a *Client) ProcessorsGetPropertyDescriptor(params *ProcessorsGetPropertyDescriptorParams, authInfo runtime.ClientAuthInfoWriter) (*ProcessorsGetPropertyDescriptorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProcessorsGetPropertyDescriptorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "processorsGetPropertyDescriptor",
		Method:             "GET",
		PathPattern:        "/processors/{id}/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProcessorsGetPropertyDescriptorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProcessorsGetPropertyDescriptorOK), nil

}

/*
ProcessorsGetState gets the state for a processor
*/
func (a *Client) ProcessorsGetState(params *ProcessorsGetStateParams, authInfo runtime.ClientAuthInfoWriter) (*ProcessorsGetStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProcessorsGetStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "processorsGetState",
		Method:             "GET",
		PathPattern:        "/processors/{id}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProcessorsGetStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProcessorsGetStateOK), nil

}

/*
ProcessorsUpdateRunStatus updates run status of a processor
*/
func (a *Client) ProcessorsUpdateRunStatus(params *ProcessorsUpdateRunStatusParams, authInfo runtime.ClientAuthInfoWriter) (*ProcessorsUpdateRunStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProcessorsUpdateRunStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "processorsUpdateRunStatus",
		Method:             "PUT",
		PathPattern:        "/processors/{id}/run-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProcessorsUpdateRunStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProcessorsUpdateRunStatusOK), nil

}

/*
TerminateProcessor terminates a processor essentially deleting its threads and any active tasks
*/
func (a *Client) TerminateProcessor(params *TerminateProcessorParams, authInfo runtime.ClientAuthInfoWriter) (*TerminateProcessorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateProcessorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "terminateProcessor",
		Method:             "DELETE",
		PathPattern:        "/processors/{id}/threads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TerminateProcessorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TerminateProcessorOK), nil

}

/*
UpdateProcessor updates a processor
*/
func (a *Client) UpdateProcessor(params *UpdateProcessorParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProcessorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProcessorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProcessor",
		Method:             "PUT",
		PathPattern:        "/processors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProcessorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProcessorOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
