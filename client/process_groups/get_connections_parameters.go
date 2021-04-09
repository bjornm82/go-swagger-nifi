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

// NewGetConnectionsParams creates a new GetConnectionsParams object
// with the default values initialized.
func NewGetConnectionsParams() *GetConnectionsParams {
	var ()
	return &GetConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConnectionsParamsWithTimeout creates a new GetConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConnectionsParamsWithTimeout(timeout time.Duration) *GetConnectionsParams {
	var ()
	return &GetConnectionsParams{

		timeout: timeout,
	}
}

// NewGetConnectionsParamsWithContext creates a new GetConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConnectionsParamsWithContext(ctx context.Context) *GetConnectionsParams {
	var ()
	return &GetConnectionsParams{

		Context: ctx,
	}
}

// NewGetConnectionsParamsWithHTTPClient creates a new GetConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConnectionsParamsWithHTTPClient(client *http.Client) *GetConnectionsParams {
	var ()
	return &GetConnectionsParams{
		HTTPClient: client,
	}
}

/*GetConnectionsParams contains all the parameters to send to the API endpoint
for the get connections operation typically these are written to a http.Request
*/
type GetConnectionsParams struct {

	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get connections params
func (o *GetConnectionsParams) WithTimeout(timeout time.Duration) *GetConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get connections params
func (o *GetConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get connections params
func (o *GetConnectionsParams) WithContext(ctx context.Context) *GetConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get connections params
func (o *GetConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get connections params
func (o *GetConnectionsParams) WithHTTPClient(client *http.Client) *GetConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get connections params
func (o *GetConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get connections params
func (o *GetConnectionsParams) WithID(id string) *GetConnectionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get connections params
func (o *GetConnectionsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
