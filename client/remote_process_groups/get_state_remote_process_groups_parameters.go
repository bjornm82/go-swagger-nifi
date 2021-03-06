// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

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

// NewGetStateRemoteProcessGroupsParams creates a new GetStateRemoteProcessGroupsParams object
// with the default values initialized.
func NewGetStateRemoteProcessGroupsParams() *GetStateRemoteProcessGroupsParams {
	var ()
	return &GetStateRemoteProcessGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStateRemoteProcessGroupsParamsWithTimeout creates a new GetStateRemoteProcessGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStateRemoteProcessGroupsParamsWithTimeout(timeout time.Duration) *GetStateRemoteProcessGroupsParams {
	var ()
	return &GetStateRemoteProcessGroupsParams{

		timeout: timeout,
	}
}

// NewGetStateRemoteProcessGroupsParamsWithContext creates a new GetStateRemoteProcessGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStateRemoteProcessGroupsParamsWithContext(ctx context.Context) *GetStateRemoteProcessGroupsParams {
	var ()
	return &GetStateRemoteProcessGroupsParams{

		Context: ctx,
	}
}

// NewGetStateRemoteProcessGroupsParamsWithHTTPClient creates a new GetStateRemoteProcessGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStateRemoteProcessGroupsParamsWithHTTPClient(client *http.Client) *GetStateRemoteProcessGroupsParams {
	var ()
	return &GetStateRemoteProcessGroupsParams{
		HTTPClient: client,
	}
}

/*GetStateRemoteProcessGroupsParams contains all the parameters to send to the API endpoint
for the get state remote process groups operation typically these are written to a http.Request
*/
type GetStateRemoteProcessGroupsParams struct {

	/*ID
	  The processor id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) WithTimeout(timeout time.Duration) *GetStateRemoteProcessGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) WithContext(ctx context.Context) *GetStateRemoteProcessGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) WithHTTPClient(client *http.Client) *GetStateRemoteProcessGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) WithID(id string) *GetStateRemoteProcessGroupsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get state remote process groups params
func (o *GetStateRemoteProcessGroupsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStateRemoteProcessGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
