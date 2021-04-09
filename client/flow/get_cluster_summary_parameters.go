// Code generated by go-swagger; DO NOT EDIT.

package flow

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

// NewGetClusterSummaryParams creates a new GetClusterSummaryParams object
// with the default values initialized.
func NewGetClusterSummaryParams() *GetClusterSummaryParams {

	return &GetClusterSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterSummaryParamsWithTimeout creates a new GetClusterSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterSummaryParamsWithTimeout(timeout time.Duration) *GetClusterSummaryParams {

	return &GetClusterSummaryParams{

		timeout: timeout,
	}
}

// NewGetClusterSummaryParamsWithContext creates a new GetClusterSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterSummaryParamsWithContext(ctx context.Context) *GetClusterSummaryParams {

	return &GetClusterSummaryParams{

		Context: ctx,
	}
}

// NewGetClusterSummaryParamsWithHTTPClient creates a new GetClusterSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterSummaryParamsWithHTTPClient(client *http.Client) *GetClusterSummaryParams {

	return &GetClusterSummaryParams{
		HTTPClient: client,
	}
}

/*GetClusterSummaryParams contains all the parameters to send to the API endpoint
for the get cluster summary operation typically these are written to a http.Request
*/
type GetClusterSummaryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster summary params
func (o *GetClusterSummaryParams) WithTimeout(timeout time.Duration) *GetClusterSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster summary params
func (o *GetClusterSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster summary params
func (o *GetClusterSummaryParams) WithContext(ctx context.Context) *GetClusterSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster summary params
func (o *GetClusterSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster summary params
func (o *GetClusterSummaryParams) WithHTTPClient(client *http.Client) *GetClusterSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster summary params
func (o *GetClusterSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}