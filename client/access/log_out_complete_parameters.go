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

// NewLogOutCompleteParams creates a new LogOutCompleteParams object
// with the default values initialized.
func NewLogOutCompleteParams() *LogOutCompleteParams {

	return &LogOutCompleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogOutCompleteParamsWithTimeout creates a new LogOutCompleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogOutCompleteParamsWithTimeout(timeout time.Duration) *LogOutCompleteParams {

	return &LogOutCompleteParams{

		timeout: timeout,
	}
}

// NewLogOutCompleteParamsWithContext creates a new LogOutCompleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogOutCompleteParamsWithContext(ctx context.Context) *LogOutCompleteParams {

	return &LogOutCompleteParams{

		Context: ctx,
	}
}

// NewLogOutCompleteParamsWithHTTPClient creates a new LogOutCompleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogOutCompleteParamsWithHTTPClient(client *http.Client) *LogOutCompleteParams {

	return &LogOutCompleteParams{
		HTTPClient: client,
	}
}

/*LogOutCompleteParams contains all the parameters to send to the API endpoint
for the log out complete operation typically these are written to a http.Request
*/
type LogOutCompleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the log out complete params
func (o *LogOutCompleteParams) WithTimeout(timeout time.Duration) *LogOutCompleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log out complete params
func (o *LogOutCompleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log out complete params
func (o *LogOutCompleteParams) WithContext(ctx context.Context) *LogOutCompleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log out complete params
func (o *LogOutCompleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log out complete params
func (o *LogOutCompleteParams) WithHTTPClient(client *http.Client) *LogOutCompleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log out complete params
func (o *LogOutCompleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LogOutCompleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}