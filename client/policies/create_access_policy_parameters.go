// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewCreateAccessPolicyParams creates a new CreateAccessPolicyParams object
// with the default values initialized.
func NewCreateAccessPolicyParams() *CreateAccessPolicyParams {
	var ()
	return &CreateAccessPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccessPolicyParamsWithTimeout creates a new CreateAccessPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccessPolicyParamsWithTimeout(timeout time.Duration) *CreateAccessPolicyParams {
	var ()
	return &CreateAccessPolicyParams{

		timeout: timeout,
	}
}

// NewCreateAccessPolicyParamsWithContext creates a new CreateAccessPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccessPolicyParamsWithContext(ctx context.Context) *CreateAccessPolicyParams {
	var ()
	return &CreateAccessPolicyParams{

		Context: ctx,
	}
}

// NewCreateAccessPolicyParamsWithHTTPClient creates a new CreateAccessPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccessPolicyParamsWithHTTPClient(client *http.Client) *CreateAccessPolicyParams {
	var ()
	return &CreateAccessPolicyParams{
		HTTPClient: client,
	}
}

/*CreateAccessPolicyParams contains all the parameters to send to the API endpoint
for the create access policy operation typically these are written to a http.Request
*/
type CreateAccessPolicyParams struct {

	/*Body
	  The access policy configuration details.

	*/
	Body *models.AccessPolicyEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create access policy params
func (o *CreateAccessPolicyParams) WithTimeout(timeout time.Duration) *CreateAccessPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create access policy params
func (o *CreateAccessPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create access policy params
func (o *CreateAccessPolicyParams) WithContext(ctx context.Context) *CreateAccessPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create access policy params
func (o *CreateAccessPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create access policy params
func (o *CreateAccessPolicyParams) WithHTTPClient(client *http.Client) *CreateAccessPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create access policy params
func (o *CreateAccessPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create access policy params
func (o *CreateAccessPolicyParams) WithBody(body *models.AccessPolicyEntity) *CreateAccessPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create access policy params
func (o *CreateAccessPolicyParams) SetBody(body *models.AccessPolicyEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccessPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
