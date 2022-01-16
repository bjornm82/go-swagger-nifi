// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewUpdateUserGroupParams creates a new UpdateUserGroupParams object
// with the default values initialized.
func NewUpdateUserGroupParams() *UpdateUserGroupParams {
	var ()
	return &UpdateUserGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserGroupParamsWithTimeout creates a new UpdateUserGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserGroupParamsWithTimeout(timeout time.Duration) *UpdateUserGroupParams {
	var ()
	return &UpdateUserGroupParams{

		timeout: timeout,
	}
}

// NewUpdateUserGroupParamsWithContext creates a new UpdateUserGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserGroupParamsWithContext(ctx context.Context) *UpdateUserGroupParams {
	var ()
	return &UpdateUserGroupParams{

		Context: ctx,
	}
}

// NewUpdateUserGroupParamsWithHTTPClient creates a new UpdateUserGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserGroupParamsWithHTTPClient(client *http.Client) *UpdateUserGroupParams {
	var ()
	return &UpdateUserGroupParams{
		HTTPClient: client,
	}
}

/*UpdateUserGroupParams contains all the parameters to send to the API endpoint
for the update user group operation typically these are written to a http.Request
*/
type UpdateUserGroupParams struct {

	/*Body
	  The user group configuration details.

	*/
	Body *models.UserGroupEntity
	/*ID
	  The user group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user group params
func (o *UpdateUserGroupParams) WithTimeout(timeout time.Duration) *UpdateUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user group params
func (o *UpdateUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user group params
func (o *UpdateUserGroupParams) WithContext(ctx context.Context) *UpdateUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user group params
func (o *UpdateUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user group params
func (o *UpdateUserGroupParams) WithHTTPClient(client *http.Client) *UpdateUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user group params
func (o *UpdateUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user group params
func (o *UpdateUserGroupParams) WithBody(body *models.UserGroupEntity) *UpdateUserGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user group params
func (o *UpdateUserGroupParams) SetBody(body *models.UserGroupEntity) {
	o.Body = body
}

// WithID adds the id to the update user group params
func (o *UpdateUserGroupParams) WithID(id string) *UpdateUserGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update user group params
func (o *UpdateUserGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
