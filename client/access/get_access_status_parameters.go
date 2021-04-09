// Code generated by go-swagger; DO NOT EDIT.

package access

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

// NewGetAccessStatusParams creates a new GetAccessStatusParams object
// with the default values initialized.
func NewGetAccessStatusParams() *GetAccessStatusParams {

	return &GetAccessStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessStatusParamsWithTimeout creates a new GetAccessStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccessStatusParamsWithTimeout(timeout time.Duration) *GetAccessStatusParams {

	return &GetAccessStatusParams{

		timeout: timeout,
	}
}

// NewGetAccessStatusParamsWithContext creates a new GetAccessStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccessStatusParamsWithContext(ctx context.Context) *GetAccessStatusParams {

	return &GetAccessStatusParams{

		Context: ctx,
	}
}

// NewGetAccessStatusParamsWithHTTPClient creates a new GetAccessStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccessStatusParamsWithHTTPClient(client *http.Client) *GetAccessStatusParams {

	return &GetAccessStatusParams{
		HTTPClient: client,
	}
}

/*GetAccessStatusParams contains all the parameters to send to the API endpoint
for the get access status operation typically these are written to a http.Request
*/
type GetAccessStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get access status params
func (o *GetAccessStatusParams) WithTimeout(timeout time.Duration) *GetAccessStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access status params
func (o *GetAccessStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access status params
func (o *GetAccessStatusParams) WithContext(ctx context.Context) *GetAccessStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access status params
func (o *GetAccessStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access status params
func (o *GetAccessStatusParams) WithHTTPClient(client *http.Client) *GetAccessStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access status params
func (o *GetAccessStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
