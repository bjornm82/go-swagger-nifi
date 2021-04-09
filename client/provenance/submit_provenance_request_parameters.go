// Code generated by go-swagger; DO NOT EDIT.

package provenance

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

// NewSubmitProvenanceRequestParams creates a new SubmitProvenanceRequestParams object
// with the default values initialized.
func NewSubmitProvenanceRequestParams() *SubmitProvenanceRequestParams {
	var ()
	return &SubmitProvenanceRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitProvenanceRequestParamsWithTimeout creates a new SubmitProvenanceRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubmitProvenanceRequestParamsWithTimeout(timeout time.Duration) *SubmitProvenanceRequestParams {
	var ()
	return &SubmitProvenanceRequestParams{

		timeout: timeout,
	}
}

// NewSubmitProvenanceRequestParamsWithContext creates a new SubmitProvenanceRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubmitProvenanceRequestParamsWithContext(ctx context.Context) *SubmitProvenanceRequestParams {
	var ()
	return &SubmitProvenanceRequestParams{

		Context: ctx,
	}
}

// NewSubmitProvenanceRequestParamsWithHTTPClient creates a new SubmitProvenanceRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubmitProvenanceRequestParamsWithHTTPClient(client *http.Client) *SubmitProvenanceRequestParams {
	var ()
	return &SubmitProvenanceRequestParams{
		HTTPClient: client,
	}
}

/*SubmitProvenanceRequestParams contains all the parameters to send to the API endpoint
for the submit provenance request operation typically these are written to a http.Request
*/
type SubmitProvenanceRequestParams struct {

	/*Body
	  The provenance query details.

	*/
	Body *models.ProvenanceEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the submit provenance request params
func (o *SubmitProvenanceRequestParams) WithTimeout(timeout time.Duration) *SubmitProvenanceRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit provenance request params
func (o *SubmitProvenanceRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit provenance request params
func (o *SubmitProvenanceRequestParams) WithContext(ctx context.Context) *SubmitProvenanceRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit provenance request params
func (o *SubmitProvenanceRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit provenance request params
func (o *SubmitProvenanceRequestParams) WithHTTPClient(client *http.Client) *SubmitProvenanceRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit provenance request params
func (o *SubmitProvenanceRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit provenance request params
func (o *SubmitProvenanceRequestParams) WithBody(body *models.ProvenanceEntity) *SubmitProvenanceRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit provenance request params
func (o *SubmitProvenanceRequestParams) SetBody(body *models.ProvenanceEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitProvenanceRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
