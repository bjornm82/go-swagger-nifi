// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProcessGroupParams creates a new GetProcessGroupParams object
// with the default values initialized.
func NewGetProcessGroupParams() *GetProcessGroupParams {
	var ()
	return &GetProcessGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessGroupParamsWithTimeout creates a new GetProcessGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessGroupParamsWithTimeout(timeout time.Duration) *GetProcessGroupParams {
	var ()
	return &GetProcessGroupParams{

		timeout: timeout,
	}
}

// NewGetProcessGroupParamsWithContext creates a new GetProcessGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessGroupParamsWithContext(ctx context.Context) *GetProcessGroupParams {
	var ()
	return &GetProcessGroupParams{

		Context: ctx,
	}
}

// NewGetProcessGroupParamsWithHTTPClient creates a new GetProcessGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProcessGroupParamsWithHTTPClient(client *http.Client) *GetProcessGroupParams {
	var ()
	return &GetProcessGroupParams{
		HTTPClient: client,
	}
}

/*GetProcessGroupParams contains all the parameters to send to the API endpoint
for the get process group operation typically these are written to a http.Request
*/
type GetProcessGroupParams struct {

	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get process group params
func (o *GetProcessGroupParams) WithTimeout(timeout time.Duration) *GetProcessGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get process group params
func (o *GetProcessGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get process group params
func (o *GetProcessGroupParams) WithContext(ctx context.Context) *GetProcessGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get process group params
func (o *GetProcessGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get process group params
func (o *GetProcessGroupParams) WithHTTPClient(client *http.Client) *GetProcessGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get process group params
func (o *GetProcessGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get process group params
func (o *GetProcessGroupParams) WithID(id string) *GetProcessGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get process group params
func (o *GetProcessGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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