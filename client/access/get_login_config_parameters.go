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

// NewGetLoginConfigParams creates a new GetLoginConfigParams object
// with the default values initialized.
func NewGetLoginConfigParams() *GetLoginConfigParams {

	return &GetLoginConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginConfigParamsWithTimeout creates a new GetLoginConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoginConfigParamsWithTimeout(timeout time.Duration) *GetLoginConfigParams {

	return &GetLoginConfigParams{

		timeout: timeout,
	}
}

// NewGetLoginConfigParamsWithContext creates a new GetLoginConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoginConfigParamsWithContext(ctx context.Context) *GetLoginConfigParams {

	return &GetLoginConfigParams{

		Context: ctx,
	}
}

// NewGetLoginConfigParamsWithHTTPClient creates a new GetLoginConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoginConfigParamsWithHTTPClient(client *http.Client) *GetLoginConfigParams {

	return &GetLoginConfigParams{
		HTTPClient: client,
	}
}

/*GetLoginConfigParams contains all the parameters to send to the API endpoint
for the get login config operation typically these are written to a http.Request
*/
type GetLoginConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get login config params
func (o *GetLoginConfigParams) WithTimeout(timeout time.Duration) *GetLoginConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get login config params
func (o *GetLoginConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get login config params
func (o *GetLoginConfigParams) WithContext(ctx context.Context) *GetLoginConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get login config params
func (o *GetLoginConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get login config params
func (o *GetLoginConfigParams) WithHTTPClient(client *http.Client) *GetLoginConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get login config params
func (o *GetLoginConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}