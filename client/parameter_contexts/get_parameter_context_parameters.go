// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

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

// NewGetParameterContextParams creates a new GetParameterContextParams object
// with the default values initialized.
func NewGetParameterContextParams() *GetParameterContextParams {
	var ()
	return &GetParameterContextParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParameterContextParamsWithTimeout creates a new GetParameterContextParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParameterContextParamsWithTimeout(timeout time.Duration) *GetParameterContextParams {
	var ()
	return &GetParameterContextParams{

		timeout: timeout,
	}
}

// NewGetParameterContextParamsWithContext creates a new GetParameterContextParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParameterContextParamsWithContext(ctx context.Context) *GetParameterContextParams {
	var ()
	return &GetParameterContextParams{

		Context: ctx,
	}
}

// NewGetParameterContextParamsWithHTTPClient creates a new GetParameterContextParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetParameterContextParamsWithHTTPClient(client *http.Client) *GetParameterContextParams {
	var ()
	return &GetParameterContextParams{
		HTTPClient: client,
	}
}

/*GetParameterContextParams contains all the parameters to send to the API endpoint
for the get parameter context operation typically these are written to a http.Request
*/
type GetParameterContextParams struct {

	/*ID
	  The ID of the Parameter Context

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get parameter context params
func (o *GetParameterContextParams) WithTimeout(timeout time.Duration) *GetParameterContextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parameter context params
func (o *GetParameterContextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parameter context params
func (o *GetParameterContextParams) WithContext(ctx context.Context) *GetParameterContextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parameter context params
func (o *GetParameterContextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parameter context params
func (o *GetParameterContextParams) WithHTTPClient(client *http.Client) *GetParameterContextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parameter context params
func (o *GetParameterContextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get parameter context params
func (o *GetParameterContextParams) WithID(id string) *GetParameterContextParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get parameter context params
func (o *GetParameterContextParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetParameterContextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
