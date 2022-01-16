// Code generated by go-swagger; DO NOT EDIT.

package output_ports

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

// NewGetOutputPortParams creates a new GetOutputPortParams object
// with the default values initialized.
func NewGetOutputPortParams() *GetOutputPortParams {
	var ()
	return &GetOutputPortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutputPortParamsWithTimeout creates a new GetOutputPortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutputPortParamsWithTimeout(timeout time.Duration) *GetOutputPortParams {
	var ()
	return &GetOutputPortParams{

		timeout: timeout,
	}
}

// NewGetOutputPortParamsWithContext creates a new GetOutputPortParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutputPortParamsWithContext(ctx context.Context) *GetOutputPortParams {
	var ()
	return &GetOutputPortParams{

		Context: ctx,
	}
}

// NewGetOutputPortParamsWithHTTPClient creates a new GetOutputPortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutputPortParamsWithHTTPClient(client *http.Client) *GetOutputPortParams {
	var ()
	return &GetOutputPortParams{
		HTTPClient: client,
	}
}

/*GetOutputPortParams contains all the parameters to send to the API endpoint
for the get output port operation typically these are written to a http.Request
*/
type GetOutputPortParams struct {

	/*ID
	  The output port id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get output port params
func (o *GetOutputPortParams) WithTimeout(timeout time.Duration) *GetOutputPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get output port params
func (o *GetOutputPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get output port params
func (o *GetOutputPortParams) WithContext(ctx context.Context) *GetOutputPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get output port params
func (o *GetOutputPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get output port params
func (o *GetOutputPortParams) WithHTTPClient(client *http.Client) *GetOutputPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get output port params
func (o *GetOutputPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get output port params
func (o *GetOutputPortParams) WithID(id string) *GetOutputPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get output port params
func (o *GetOutputPortParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutputPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
