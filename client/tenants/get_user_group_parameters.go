// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewGetUserGroupParams creates a new GetUserGroupParams object
// with the default values initialized.
func NewGetUserGroupParams() *GetUserGroupParams {
	var ()
	return &GetUserGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserGroupParamsWithTimeout creates a new GetUserGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserGroupParamsWithTimeout(timeout time.Duration) *GetUserGroupParams {
	var ()
	return &GetUserGroupParams{

		timeout: timeout,
	}
}

// NewGetUserGroupParamsWithContext creates a new GetUserGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserGroupParamsWithContext(ctx context.Context) *GetUserGroupParams {
	var ()
	return &GetUserGroupParams{

		Context: ctx,
	}
}

// NewGetUserGroupParamsWithHTTPClient creates a new GetUserGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserGroupParamsWithHTTPClient(client *http.Client) *GetUserGroupParams {
	var ()
	return &GetUserGroupParams{
		HTTPClient: client,
	}
}

/*GetUserGroupParams contains all the parameters to send to the API endpoint
for the get user group operation typically these are written to a http.Request
*/
type GetUserGroupParams struct {

	/*ID
	  The user group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user group params
func (o *GetUserGroupParams) WithTimeout(timeout time.Duration) *GetUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user group params
func (o *GetUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user group params
func (o *GetUserGroupParams) WithContext(ctx context.Context) *GetUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user group params
func (o *GetUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user group params
func (o *GetUserGroupParams) WithHTTPClient(client *http.Client) *GetUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user group params
func (o *GetUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get user group params
func (o *GetUserGroupParams) WithID(id string) *GetUserGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get user group params
func (o *GetUserGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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