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

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// NewCreateInputPortParams creates a new CreateInputPortParams object
// with the default values initialized.
func NewCreateInputPortParams() *CreateInputPortParams {
	var ()
	return &CreateInputPortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInputPortParamsWithTimeout creates a new CreateInputPortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInputPortParamsWithTimeout(timeout time.Duration) *CreateInputPortParams {
	var ()
	return &CreateInputPortParams{

		timeout: timeout,
	}
}

// NewCreateInputPortParamsWithContext creates a new CreateInputPortParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInputPortParamsWithContext(ctx context.Context) *CreateInputPortParams {
	var ()
	return &CreateInputPortParams{

		Context: ctx,
	}
}

// NewCreateInputPortParamsWithHTTPClient creates a new CreateInputPortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInputPortParamsWithHTTPClient(client *http.Client) *CreateInputPortParams {
	var ()
	return &CreateInputPortParams{
		HTTPClient: client,
	}
}

/*CreateInputPortParams contains all the parameters to send to the API endpoint
for the create input port operation typically these are written to a http.Request
*/
type CreateInputPortParams struct {

	/*Body
	  The input port configuration details.

	*/
	Body *models.PortEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create input port params
func (o *CreateInputPortParams) WithTimeout(timeout time.Duration) *CreateInputPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create input port params
func (o *CreateInputPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create input port params
func (o *CreateInputPortParams) WithContext(ctx context.Context) *CreateInputPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create input port params
func (o *CreateInputPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create input port params
func (o *CreateInputPortParams) WithHTTPClient(client *http.Client) *CreateInputPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create input port params
func (o *CreateInputPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create input port params
func (o *CreateInputPortParams) WithBody(body *models.PortEntity) *CreateInputPortParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create input port params
func (o *CreateInputPortParams) SetBody(body *models.PortEntity) {
	o.Body = body
}

// WithID adds the id to the create input port params
func (o *CreateInputPortParams) WithID(id string) *CreateInputPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create input port params
func (o *CreateInputPortParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInputPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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