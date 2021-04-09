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

// NewCopySnippetParams creates a new CopySnippetParams object
// with the default values initialized.
func NewCopySnippetParams() *CopySnippetParams {
	var ()
	return &CopySnippetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCopySnippetParamsWithTimeout creates a new CopySnippetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCopySnippetParamsWithTimeout(timeout time.Duration) *CopySnippetParams {
	var ()
	return &CopySnippetParams{

		timeout: timeout,
	}
}

// NewCopySnippetParamsWithContext creates a new CopySnippetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCopySnippetParamsWithContext(ctx context.Context) *CopySnippetParams {
	var ()
	return &CopySnippetParams{

		Context: ctx,
	}
}

// NewCopySnippetParamsWithHTTPClient creates a new CopySnippetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCopySnippetParamsWithHTTPClient(client *http.Client) *CopySnippetParams {
	var ()
	return &CopySnippetParams{
		HTTPClient: client,
	}
}

/*CopySnippetParams contains all the parameters to send to the API endpoint
for the copy snippet operation typically these are written to a http.Request
*/
type CopySnippetParams struct {

	/*Body
	  The copy snippet request.

	*/
	Body *models.CopySnippetRequestEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the copy snippet params
func (o *CopySnippetParams) WithTimeout(timeout time.Duration) *CopySnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the copy snippet params
func (o *CopySnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the copy snippet params
func (o *CopySnippetParams) WithContext(ctx context.Context) *CopySnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the copy snippet params
func (o *CopySnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the copy snippet params
func (o *CopySnippetParams) WithHTTPClient(client *http.Client) *CopySnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the copy snippet params
func (o *CopySnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the copy snippet params
func (o *CopySnippetParams) WithBody(body *models.CopySnippetRequestEntity) *CopySnippetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the copy snippet params
func (o *CopySnippetParams) SetBody(body *models.CopySnippetRequestEntity) {
	o.Body = body
}

// WithID adds the id to the copy snippet params
func (o *CopySnippetParams) WithID(id string) *CopySnippetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the copy snippet params
func (o *CopySnippetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CopySnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
