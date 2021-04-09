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

// NewInstantiateTemplateParams creates a new InstantiateTemplateParams object
// with the default values initialized.
func NewInstantiateTemplateParams() *InstantiateTemplateParams {
	var ()
	return &InstantiateTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstantiateTemplateParamsWithTimeout creates a new InstantiateTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstantiateTemplateParamsWithTimeout(timeout time.Duration) *InstantiateTemplateParams {
	var ()
	return &InstantiateTemplateParams{

		timeout: timeout,
	}
}

// NewInstantiateTemplateParamsWithContext creates a new InstantiateTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstantiateTemplateParamsWithContext(ctx context.Context) *InstantiateTemplateParams {
	var ()
	return &InstantiateTemplateParams{

		Context: ctx,
	}
}

// NewInstantiateTemplateParamsWithHTTPClient creates a new InstantiateTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstantiateTemplateParamsWithHTTPClient(client *http.Client) *InstantiateTemplateParams {
	var ()
	return &InstantiateTemplateParams{
		HTTPClient: client,
	}
}

/*InstantiateTemplateParams contains all the parameters to send to the API endpoint
for the instantiate template operation typically these are written to a http.Request
*/
type InstantiateTemplateParams struct {

	/*Body
	  The instantiate template request.

	*/
	Body *models.InstantiateTemplateRequestEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the instantiate template params
func (o *InstantiateTemplateParams) WithTimeout(timeout time.Duration) *InstantiateTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instantiate template params
func (o *InstantiateTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instantiate template params
func (o *InstantiateTemplateParams) WithContext(ctx context.Context) *InstantiateTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instantiate template params
func (o *InstantiateTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instantiate template params
func (o *InstantiateTemplateParams) WithHTTPClient(client *http.Client) *InstantiateTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instantiate template params
func (o *InstantiateTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the instantiate template params
func (o *InstantiateTemplateParams) WithBody(body *models.InstantiateTemplateRequestEntity) *InstantiateTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the instantiate template params
func (o *InstantiateTemplateParams) SetBody(body *models.InstantiateTemplateRequestEntity) {
	o.Body = body
}

// WithID adds the id to the instantiate template params
func (o *InstantiateTemplateParams) WithID(id string) *InstantiateTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instantiate template params
func (o *InstantiateTemplateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InstantiateTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
