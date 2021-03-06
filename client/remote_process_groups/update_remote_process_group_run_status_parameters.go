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

	"github.com/bjornm82/go-swagger-nifi/models"
)

// NewUpdateRemoteProcessGroupRunStatusParams creates a new UpdateRemoteProcessGroupRunStatusParams object
// with the default values initialized.
func NewUpdateRemoteProcessGroupRunStatusParams() *UpdateRemoteProcessGroupRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupRunStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRemoteProcessGroupRunStatusParamsWithTimeout creates a new UpdateRemoteProcessGroupRunStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRemoteProcessGroupRunStatusParamsWithTimeout(timeout time.Duration) *UpdateRemoteProcessGroupRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupRunStatusParams{

		timeout: timeout,
	}
}

// NewUpdateRemoteProcessGroupRunStatusParamsWithContext creates a new UpdateRemoteProcessGroupRunStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRemoteProcessGroupRunStatusParamsWithContext(ctx context.Context) *UpdateRemoteProcessGroupRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupRunStatusParams{

		Context: ctx,
	}
}

// NewUpdateRemoteProcessGroupRunStatusParamsWithHTTPClient creates a new UpdateRemoteProcessGroupRunStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRemoteProcessGroupRunStatusParamsWithHTTPClient(client *http.Client) *UpdateRemoteProcessGroupRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupRunStatusParams{
		HTTPClient: client,
	}
}

/*UpdateRemoteProcessGroupRunStatusParams contains all the parameters to send to the API endpoint
for the update remote process group run status operation typically these are written to a http.Request
*/
type UpdateRemoteProcessGroupRunStatusParams struct {

	/*Body
	  The remote process group run status.

	*/
	Body *models.RemotePortRunStatusEntity
	/*ID
	  The remote process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) WithTimeout(timeout time.Duration) *UpdateRemoteProcessGroupRunStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) WithContext(ctx context.Context) *UpdateRemoteProcessGroupRunStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) WithHTTPClient(client *http.Client) *UpdateRemoteProcessGroupRunStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) WithBody(body *models.RemotePortRunStatusEntity) *UpdateRemoteProcessGroupRunStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) SetBody(body *models.RemotePortRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) WithID(id string) *UpdateRemoteProcessGroupRunStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update remote process group run status params
func (o *UpdateRemoteProcessGroupRunStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRemoteProcessGroupRunStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
