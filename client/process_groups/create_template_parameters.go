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

// NewCreateTemplateParams creates a new CreateTemplateParams object
// with the default values initialized.
func NewCreateTemplateParams() *CreateTemplateParams {
	var ()
	return &CreateTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTemplateParamsWithTimeout creates a new CreateTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTemplateParamsWithTimeout(timeout time.Duration) *CreateTemplateParams {
	var ()
	return &CreateTemplateParams{

		timeout: timeout,
	}
}

// NewCreateTemplateParamsWithContext creates a new CreateTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTemplateParamsWithContext(ctx context.Context) *CreateTemplateParams {
	var ()
	return &CreateTemplateParams{

		Context: ctx,
	}
}

// NewCreateTemplateParamsWithHTTPClient creates a new CreateTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTemplateParamsWithHTTPClient(client *http.Client) *CreateTemplateParams {
	var ()
	return &CreateTemplateParams{
		HTTPClient: client,
	}
}

/*CreateTemplateParams contains all the parameters to send to the API endpoint
for the create template operation typically these are written to a http.Request
*/
type CreateTemplateParams struct {

	/*Body
	  The create template request.

	*/
	Body *models.CreateTemplateRequestEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create template params
func (o *CreateTemplateParams) WithTimeout(timeout time.Duration) *CreateTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create template params
func (o *CreateTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create template params
func (o *CreateTemplateParams) WithContext(ctx context.Context) *CreateTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create template params
func (o *CreateTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create template params
func (o *CreateTemplateParams) WithHTTPClient(client *http.Client) *CreateTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create template params
func (o *CreateTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create template params
func (o *CreateTemplateParams) WithBody(body *models.CreateTemplateRequestEntity) *CreateTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create template params
func (o *CreateTemplateParams) SetBody(body *models.CreateTemplateRequestEntity) {
	o.Body = body
}

// WithID adds the id to the create template params
func (o *CreateTemplateParams) WithID(id string) *CreateTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create template params
func (o *CreateTemplateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
