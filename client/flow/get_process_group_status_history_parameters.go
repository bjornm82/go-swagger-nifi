// Code generated by go-swagger; DO NOT EDIT.

package flow

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

// NewGetProcessGroupStatusHistoryParams creates a new GetProcessGroupStatusHistoryParams object
// with the default values initialized.
func NewGetProcessGroupStatusHistoryParams() *GetProcessGroupStatusHistoryParams {
	var ()
	return &GetProcessGroupStatusHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessGroupStatusHistoryParamsWithTimeout creates a new GetProcessGroupStatusHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessGroupStatusHistoryParamsWithTimeout(timeout time.Duration) *GetProcessGroupStatusHistoryParams {
	var ()
	return &GetProcessGroupStatusHistoryParams{

		timeout: timeout,
	}
}

// NewGetProcessGroupStatusHistoryParamsWithContext creates a new GetProcessGroupStatusHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessGroupStatusHistoryParamsWithContext(ctx context.Context) *GetProcessGroupStatusHistoryParams {
	var ()
	return &GetProcessGroupStatusHistoryParams{

		Context: ctx,
	}
}

// NewGetProcessGroupStatusHistoryParamsWithHTTPClient creates a new GetProcessGroupStatusHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProcessGroupStatusHistoryParamsWithHTTPClient(client *http.Client) *GetProcessGroupStatusHistoryParams {
	var ()
	return &GetProcessGroupStatusHistoryParams{
		HTTPClient: client,
	}
}

/*GetProcessGroupStatusHistoryParams contains all the parameters to send to the API endpoint
for the get process group status history operation typically these are written to a http.Request
*/
type GetProcessGroupStatusHistoryParams struct {

	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) WithTimeout(timeout time.Duration) *GetProcessGroupStatusHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) WithContext(ctx context.Context) *GetProcessGroupStatusHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) WithHTTPClient(client *http.Client) *GetProcessGroupStatusHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) WithID(id string) *GetProcessGroupStatusHistoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get process group status history params
func (o *GetProcessGroupStatusHistoryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessGroupStatusHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
