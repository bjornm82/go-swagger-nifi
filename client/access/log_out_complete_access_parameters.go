// Code generated by go-swagger; DO NOT EDIT.

package access

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

// NewLogOutCompleteAccessParams creates a new LogOutCompleteAccessParams object
// with the default values initialized.
func NewLogOutCompleteAccessParams() *LogOutCompleteAccessParams {

	return &LogOutCompleteAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogOutCompleteAccessParamsWithTimeout creates a new LogOutCompleteAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogOutCompleteAccessParamsWithTimeout(timeout time.Duration) *LogOutCompleteAccessParams {

	return &LogOutCompleteAccessParams{

		timeout: timeout,
	}
}

// NewLogOutCompleteAccessParamsWithContext creates a new LogOutCompleteAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogOutCompleteAccessParamsWithContext(ctx context.Context) *LogOutCompleteAccessParams {

	return &LogOutCompleteAccessParams{

		Context: ctx,
	}
}

// NewLogOutCompleteAccessParamsWithHTTPClient creates a new LogOutCompleteAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogOutCompleteAccessParamsWithHTTPClient(client *http.Client) *LogOutCompleteAccessParams {

	return &LogOutCompleteAccessParams{
		HTTPClient: client,
	}
}

/*LogOutCompleteAccessParams contains all the parameters to send to the API endpoint
for the log out complete access operation typically these are written to a http.Request
*/
type LogOutCompleteAccessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the log out complete access params
func (o *LogOutCompleteAccessParams) WithTimeout(timeout time.Duration) *LogOutCompleteAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log out complete access params
func (o *LogOutCompleteAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log out complete access params
func (o *LogOutCompleteAccessParams) WithContext(ctx context.Context) *LogOutCompleteAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log out complete access params
func (o *LogOutCompleteAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log out complete access params
func (o *LogOutCompleteAccessParams) WithHTTPClient(client *http.Client) *LogOutCompleteAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log out complete access params
func (o *LogOutCompleteAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LogOutCompleteAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
