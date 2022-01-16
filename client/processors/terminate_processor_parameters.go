// Code generated by go-swagger; DO NOT EDIT.

package processors

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

// NewTerminateProcessorParams creates a new TerminateProcessorParams object
// with the default values initialized.
func NewTerminateProcessorParams() *TerminateProcessorParams {
	var ()
	return &TerminateProcessorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTerminateProcessorParamsWithTimeout creates a new TerminateProcessorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTerminateProcessorParamsWithTimeout(timeout time.Duration) *TerminateProcessorParams {
	var ()
	return &TerminateProcessorParams{

		timeout: timeout,
	}
}

// NewTerminateProcessorParamsWithContext creates a new TerminateProcessorParams object
// with the default values initialized, and the ability to set a context for a request
func NewTerminateProcessorParamsWithContext(ctx context.Context) *TerminateProcessorParams {
	var ()
	return &TerminateProcessorParams{

		Context: ctx,
	}
}

// NewTerminateProcessorParamsWithHTTPClient creates a new TerminateProcessorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTerminateProcessorParamsWithHTTPClient(client *http.Client) *TerminateProcessorParams {
	var ()
	return &TerminateProcessorParams{
		HTTPClient: client,
	}
}

/*TerminateProcessorParams contains all the parameters to send to the API endpoint
for the terminate processor operation typically these are written to a http.Request
*/
type TerminateProcessorParams struct {

	/*ID
	  The processor id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the terminate processor params
func (o *TerminateProcessorParams) WithTimeout(timeout time.Duration) *TerminateProcessorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the terminate processor params
func (o *TerminateProcessorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the terminate processor params
func (o *TerminateProcessorParams) WithContext(ctx context.Context) *TerminateProcessorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the terminate processor params
func (o *TerminateProcessorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the terminate processor params
func (o *TerminateProcessorParams) WithHTTPClient(client *http.Client) *TerminateProcessorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the terminate processor params
func (o *TerminateProcessorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the terminate processor params
func (o *TerminateProcessorParams) WithID(id string) *TerminateProcessorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the terminate processor params
func (o *TerminateProcessorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TerminateProcessorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
