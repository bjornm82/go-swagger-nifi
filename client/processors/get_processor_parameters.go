// Code generated by go-swagger; DO NOT EDIT.

package processors

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

// NewGetProcessorParams creates a new GetProcessorParams object
// with the default values initialized.
func NewGetProcessorParams() *GetProcessorParams {
	var ()
	return &GetProcessorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessorParamsWithTimeout creates a new GetProcessorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessorParamsWithTimeout(timeout time.Duration) *GetProcessorParams {
	var ()
	return &GetProcessorParams{

		timeout: timeout,
	}
}

// NewGetProcessorParamsWithContext creates a new GetProcessorParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessorParamsWithContext(ctx context.Context) *GetProcessorParams {
	var ()
	return &GetProcessorParams{

		Context: ctx,
	}
}

// NewGetProcessorParamsWithHTTPClient creates a new GetProcessorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProcessorParamsWithHTTPClient(client *http.Client) *GetProcessorParams {
	var ()
	return &GetProcessorParams{
		HTTPClient: client,
	}
}

/*GetProcessorParams contains all the parameters to send to the API endpoint
for the get processor operation typically these are written to a http.Request
*/
type GetProcessorParams struct {

	/*ID
	  The processor id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get processor params
func (o *GetProcessorParams) WithTimeout(timeout time.Duration) *GetProcessorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get processor params
func (o *GetProcessorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get processor params
func (o *GetProcessorParams) WithContext(ctx context.Context) *GetProcessorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get processor params
func (o *GetProcessorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get processor params
func (o *GetProcessorParams) WithHTTPClient(client *http.Client) *GetProcessorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get processor params
func (o *GetProcessorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get processor params
func (o *GetProcessorParams) WithID(id string) *GetProcessorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get processor params
func (o *GetProcessorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
