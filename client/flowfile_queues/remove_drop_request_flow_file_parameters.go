// Code generated by go-swagger; DO NOT EDIT.

package flowfile_queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRemoveDropRequestFlowFileParams creates a new RemoveDropRequestFlowFileParams object
// with the default values initialized.
func NewRemoveDropRequestFlowFileParams() *RemoveDropRequestFlowFileParams {
	var ()
	return &RemoveDropRequestFlowFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveDropRequestFlowFileParamsWithTimeout creates a new RemoveDropRequestFlowFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveDropRequestFlowFileParamsWithTimeout(timeout time.Duration) *RemoveDropRequestFlowFileParams {
	var ()
	return &RemoveDropRequestFlowFileParams{

		timeout: timeout,
	}
}

// NewRemoveDropRequestFlowFileParamsWithContext creates a new RemoveDropRequestFlowFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveDropRequestFlowFileParamsWithContext(ctx context.Context) *RemoveDropRequestFlowFileParams {
	var ()
	return &RemoveDropRequestFlowFileParams{

		Context: ctx,
	}
}

// NewRemoveDropRequestFlowFileParamsWithHTTPClient creates a new RemoveDropRequestFlowFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveDropRequestFlowFileParamsWithHTTPClient(client *http.Client) *RemoveDropRequestFlowFileParams {
	var ()
	return &RemoveDropRequestFlowFileParams{
		HTTPClient: client,
	}
}

/*RemoveDropRequestFlowFileParams contains all the parameters to send to the API endpoint
for the remove drop request flow file operation typically these are written to a http.Request
*/
type RemoveDropRequestFlowFileParams struct {

	/*DropRequestID
	  The drop request id.

	*/
	DropRequestID string
	/*ID
	  The connection id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) WithTimeout(timeout time.Duration) *RemoveDropRequestFlowFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) WithContext(ctx context.Context) *RemoveDropRequestFlowFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) WithHTTPClient(client *http.Client) *RemoveDropRequestFlowFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDropRequestID adds the dropRequestID to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) WithDropRequestID(dropRequestID string) *RemoveDropRequestFlowFileParams {
	o.SetDropRequestID(dropRequestID)
	return o
}

// SetDropRequestID adds the dropRequestId to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) SetDropRequestID(dropRequestID string) {
	o.DropRequestID = dropRequestID
}

// WithID adds the id to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) WithID(id string) *RemoveDropRequestFlowFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove drop request flow file params
func (o *RemoveDropRequestFlowFileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveDropRequestFlowFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param drop-request-id
	if err := r.SetPathParam("drop-request-id", o.DropRequestID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
