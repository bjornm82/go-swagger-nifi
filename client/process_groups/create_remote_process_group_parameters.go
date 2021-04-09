// Code generated by go-swagger; DO NOT EDIT.

package process_groups

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

	models "github.com/bjornm82/swagger-nifi/models"
)

// NewCreateRemoteProcessGroupParams creates a new CreateRemoteProcessGroupParams object
// with the default values initialized.
func NewCreateRemoteProcessGroupParams() *CreateRemoteProcessGroupParams {
	var ()
	return &CreateRemoteProcessGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRemoteProcessGroupParamsWithTimeout creates a new CreateRemoteProcessGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRemoteProcessGroupParamsWithTimeout(timeout time.Duration) *CreateRemoteProcessGroupParams {
	var ()
	return &CreateRemoteProcessGroupParams{

		timeout: timeout,
	}
}

// NewCreateRemoteProcessGroupParamsWithContext creates a new CreateRemoteProcessGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRemoteProcessGroupParamsWithContext(ctx context.Context) *CreateRemoteProcessGroupParams {
	var ()
	return &CreateRemoteProcessGroupParams{

		Context: ctx,
	}
}

// NewCreateRemoteProcessGroupParamsWithHTTPClient creates a new CreateRemoteProcessGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRemoteProcessGroupParamsWithHTTPClient(client *http.Client) *CreateRemoteProcessGroupParams {
	var ()
	return &CreateRemoteProcessGroupParams{
		HTTPClient: client,
	}
}

/*CreateRemoteProcessGroupParams contains all the parameters to send to the API endpoint
for the create remote process group operation typically these are written to a http.Request
*/
type CreateRemoteProcessGroupParams struct {

	/*Body
	  The remote process group configuration details.

	*/
	Body *models.RemoteProcessGroupEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create remote process group params
func (o *CreateRemoteProcessGroupParams) WithTimeout(timeout time.Duration) *CreateRemoteProcessGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create remote process group params
func (o *CreateRemoteProcessGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create remote process group params
func (o *CreateRemoteProcessGroupParams) WithContext(ctx context.Context) *CreateRemoteProcessGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create remote process group params
func (o *CreateRemoteProcessGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create remote process group params
func (o *CreateRemoteProcessGroupParams) WithHTTPClient(client *http.Client) *CreateRemoteProcessGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create remote process group params
func (o *CreateRemoteProcessGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create remote process group params
func (o *CreateRemoteProcessGroupParams) WithBody(body *models.RemoteProcessGroupEntity) *CreateRemoteProcessGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create remote process group params
func (o *CreateRemoteProcessGroupParams) SetBody(body *models.RemoteProcessGroupEntity) {
	o.Body = body
}

// WithID adds the id to the create remote process group params
func (o *CreateRemoteProcessGroupParams) WithID(id string) *CreateRemoteProcessGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create remote process group params
func (o *CreateRemoteProcessGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRemoteProcessGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
