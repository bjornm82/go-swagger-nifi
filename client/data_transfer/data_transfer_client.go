// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new data transfer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data transfer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CommitInputPortTransaction commits or cancel the specified transaction
*/
func (a *Client) CommitInputPortTransaction(params *CommitInputPortTransactionParams, authInfo runtime.ClientAuthInfoWriter) (*CommitInputPortTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitInputPortTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "commitInputPortTransaction",
		Method:             "DELETE",
		PathPattern:        "/data-transfer/input-ports/{portId}/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CommitInputPortTransactionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CommitInputPortTransactionOK), nil

}

/*
CommitOutputPortTransaction commits or cancel the specified transaction
*/
func (a *Client) CommitOutputPortTransaction(params *CommitOutputPortTransactionParams, authInfo runtime.ClientAuthInfoWriter) (*CommitOutputPortTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitOutputPortTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "commitOutputPortTransaction",
		Method:             "DELETE",
		PathPattern:        "/data-transfer/output-ports/{portId}/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CommitOutputPortTransactionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CommitOutputPortTransactionOK), nil

}

/*
CreatePortTransaction creates a transaction to the specified output port or input port
*/
func (a *Client) CreatePortTransaction(params *CreatePortTransactionParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePortTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePortTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPortTransaction",
		Method:             "POST",
		PathPattern:        "/data-transfer/{portType}/{portId}/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePortTransactionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePortTransactionOK), nil

}

/*
ExtendInputPortTransactionTTL extends transaction TTL
*/
func (a *Client) ExtendInputPortTransactionTTL(params *ExtendInputPortTransactionTTLParams, authInfo runtime.ClientAuthInfoWriter) (*ExtendInputPortTransactionTTLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtendInputPortTransactionTTLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "extendInputPortTransactionTTL",
		Method:             "PUT",
		PathPattern:        "/data-transfer/input-ports/{portId}/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExtendInputPortTransactionTTLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExtendInputPortTransactionTTLOK), nil

}

/*
ExtendOutputPortTransactionTTL extends transaction TTL
*/
func (a *Client) ExtendOutputPortTransactionTTL(params *ExtendOutputPortTransactionTTLParams, authInfo runtime.ClientAuthInfoWriter) (*ExtendOutputPortTransactionTTLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExtendOutputPortTransactionTTLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "extendOutputPortTransactionTTL",
		Method:             "PUT",
		PathPattern:        "/data-transfer/output-ports/{portId}/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExtendOutputPortTransactionTTLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExtendOutputPortTransactionTTLOK), nil

}

/*
ReceiveFlowFiles transfers flow files to the input port
*/
func (a *Client) ReceiveFlowFiles(params *ReceiveFlowFilesParams, authInfo runtime.ClientAuthInfoWriter) (*ReceiveFlowFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReceiveFlowFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "receiveFlowFiles",
		Method:             "POST",
		PathPattern:        "/data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReceiveFlowFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReceiveFlowFilesOK), nil

}

/*
TransferFlowFiles transfers flow files from the output port
*/
func (a *Client) TransferFlowFiles(params *TransferFlowFilesParams, authInfo runtime.ClientAuthInfoWriter) (*TransferFlowFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferFlowFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "transferFlowFiles",
		Method:             "GET",
		PathPattern:        "/data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TransferFlowFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TransferFlowFilesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}