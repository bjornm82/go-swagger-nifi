// Code generated by go-swagger; DO NOT EDIT.

package labels

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

// NewGetLabelParams creates a new GetLabelParams object
// with the default values initialized.
func NewGetLabelParams() *GetLabelParams {
	var ()
	return &GetLabelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelParamsWithTimeout creates a new GetLabelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLabelParamsWithTimeout(timeout time.Duration) *GetLabelParams {
	var ()
	return &GetLabelParams{

		timeout: timeout,
	}
}

// NewGetLabelParamsWithContext creates a new GetLabelParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLabelParamsWithContext(ctx context.Context) *GetLabelParams {
	var ()
	return &GetLabelParams{

		Context: ctx,
	}
}

// NewGetLabelParamsWithHTTPClient creates a new GetLabelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLabelParamsWithHTTPClient(client *http.Client) *GetLabelParams {
	var ()
	return &GetLabelParams{
		HTTPClient: client,
	}
}

/*GetLabelParams contains all the parameters to send to the API endpoint
for the get label operation typically these are written to a http.Request
*/
type GetLabelParams struct {

	/*ID
	  The label id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get label params
func (o *GetLabelParams) WithTimeout(timeout time.Duration) *GetLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get label params
func (o *GetLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get label params
func (o *GetLabelParams) WithContext(ctx context.Context) *GetLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get label params
func (o *GetLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get label params
func (o *GetLabelParams) WithHTTPClient(client *http.Client) *GetLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get label params
func (o *GetLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get label params
func (o *GetLabelParams) WithID(id string) *GetLabelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get label params
func (o *GetLabelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
