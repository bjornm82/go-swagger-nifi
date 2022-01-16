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

// NewKnoxRequestParams creates a new KnoxRequestParams object
// with the default values initialized.
func NewKnoxRequestParams() *KnoxRequestParams {

	return &KnoxRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKnoxRequestParamsWithTimeout creates a new KnoxRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKnoxRequestParamsWithTimeout(timeout time.Duration) *KnoxRequestParams {

	return &KnoxRequestParams{

		timeout: timeout,
	}
}

// NewKnoxRequestParamsWithContext creates a new KnoxRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewKnoxRequestParamsWithContext(ctx context.Context) *KnoxRequestParams {

	return &KnoxRequestParams{

		Context: ctx,
	}
}

// NewKnoxRequestParamsWithHTTPClient creates a new KnoxRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKnoxRequestParamsWithHTTPClient(client *http.Client) *KnoxRequestParams {

	return &KnoxRequestParams{
		HTTPClient: client,
	}
}

/*KnoxRequestParams contains all the parameters to send to the API endpoint
for the knox request operation typically these are written to a http.Request
*/
type KnoxRequestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the knox request params
func (o *KnoxRequestParams) WithTimeout(timeout time.Duration) *KnoxRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the knox request params
func (o *KnoxRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the knox request params
func (o *KnoxRequestParams) WithContext(ctx context.Context) *KnoxRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the knox request params
func (o *KnoxRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the knox request params
func (o *KnoxRequestParams) WithHTTPClient(client *http.Client) *KnoxRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the knox request params
func (o *KnoxRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *KnoxRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
