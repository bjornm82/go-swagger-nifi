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

// NewProcessorsGetStateParams creates a new ProcessorsGetStateParams object
// with the default values initialized.
func NewProcessorsGetStateParams() *ProcessorsGetStateParams {
	var ()
	return &ProcessorsGetStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProcessorsGetStateParamsWithTimeout creates a new ProcessorsGetStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProcessorsGetStateParamsWithTimeout(timeout time.Duration) *ProcessorsGetStateParams {
	var ()
	return &ProcessorsGetStateParams{

		timeout: timeout,
	}
}

// NewProcessorsGetStateParamsWithContext creates a new ProcessorsGetStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProcessorsGetStateParamsWithContext(ctx context.Context) *ProcessorsGetStateParams {
	var ()
	return &ProcessorsGetStateParams{

		Context: ctx,
	}
}

// NewProcessorsGetStateParamsWithHTTPClient creates a new ProcessorsGetStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProcessorsGetStateParamsWithHTTPClient(client *http.Client) *ProcessorsGetStateParams {
	var ()
	return &ProcessorsGetStateParams{
		HTTPClient: client,
	}
}

/*ProcessorsGetStateParams contains all the parameters to send to the API endpoint
for the processors get state operation typically these are written to a http.Request
*/
type ProcessorsGetStateParams struct {

	/*ID
	  The processor id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the processors get state params
func (o *ProcessorsGetStateParams) WithTimeout(timeout time.Duration) *ProcessorsGetStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the processors get state params
func (o *ProcessorsGetStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the processors get state params
func (o *ProcessorsGetStateParams) WithContext(ctx context.Context) *ProcessorsGetStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the processors get state params
func (o *ProcessorsGetStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the processors get state params
func (o *ProcessorsGetStateParams) WithHTTPClient(client *http.Client) *ProcessorsGetStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the processors get state params
func (o *ProcessorsGetStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the processors get state params
func (o *ProcessorsGetStateParams) WithID(id string) *ProcessorsGetStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the processors get state params
func (o *ProcessorsGetStateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessorsGetStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
