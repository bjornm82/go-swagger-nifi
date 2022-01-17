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

// NewGetStateProcessorsParams creates a new GetStateProcessorsParams object
// with the default values initialized.
func NewGetStateProcessorsParams() *GetStateProcessorsParams {
	var ()
	return &GetStateProcessorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStateProcessorsParamsWithTimeout creates a new GetStateProcessorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStateProcessorsParamsWithTimeout(timeout time.Duration) *GetStateProcessorsParams {
	var ()
	return &GetStateProcessorsParams{

		timeout: timeout,
	}
}

// NewGetStateProcessorsParamsWithContext creates a new GetStateProcessorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStateProcessorsParamsWithContext(ctx context.Context) *GetStateProcessorsParams {
	var ()
	return &GetStateProcessorsParams{

		Context: ctx,
	}
}

// NewGetStateProcessorsParamsWithHTTPClient creates a new GetStateProcessorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStateProcessorsParamsWithHTTPClient(client *http.Client) *GetStateProcessorsParams {
	var ()
	return &GetStateProcessorsParams{
		HTTPClient: client,
	}
}

/*GetStateProcessorsParams contains all the parameters to send to the API endpoint
for the get state processors operation typically these are written to a http.Request
*/
type GetStateProcessorsParams struct {

	/*ID
	  The processor id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get state processors params
func (o *GetStateProcessorsParams) WithTimeout(timeout time.Duration) *GetStateProcessorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get state processors params
func (o *GetStateProcessorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get state processors params
func (o *GetStateProcessorsParams) WithContext(ctx context.Context) *GetStateProcessorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get state processors params
func (o *GetStateProcessorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get state processors params
func (o *GetStateProcessorsParams) WithHTTPClient(client *http.Client) *GetStateProcessorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get state processors params
func (o *GetStateProcessorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get state processors params
func (o *GetStateProcessorsParams) WithID(id string) *GetStateProcessorsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get state processors params
func (o *GetStateProcessorsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStateProcessorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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