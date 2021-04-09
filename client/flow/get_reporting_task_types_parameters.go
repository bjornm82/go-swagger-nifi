// Code generated by go-swagger; DO NOT EDIT.

package flow

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
)

// NewGetReportingTaskTypesParams creates a new GetReportingTaskTypesParams object
// with the default values initialized.
func NewGetReportingTaskTypesParams() *GetReportingTaskTypesParams {
	var ()
	return &GetReportingTaskTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportingTaskTypesParamsWithTimeout creates a new GetReportingTaskTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReportingTaskTypesParamsWithTimeout(timeout time.Duration) *GetReportingTaskTypesParams {
	var ()
	return &GetReportingTaskTypesParams{

		timeout: timeout,
	}
}

// NewGetReportingTaskTypesParamsWithContext creates a new GetReportingTaskTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReportingTaskTypesParamsWithContext(ctx context.Context) *GetReportingTaskTypesParams {
	var ()
	return &GetReportingTaskTypesParams{

		Context: ctx,
	}
}

// NewGetReportingTaskTypesParamsWithHTTPClient creates a new GetReportingTaskTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReportingTaskTypesParamsWithHTTPClient(client *http.Client) *GetReportingTaskTypesParams {
	var ()
	return &GetReportingTaskTypesParams{
		HTTPClient: client,
	}
}

/*GetReportingTaskTypesParams contains all the parameters to send to the API endpoint
for the get reporting task types operation typically these are written to a http.Request
*/
type GetReportingTaskTypesParams struct {

	/*BundleArtifactFilter
	  If specified, will only return types that are a member of this bundle artifact.

	*/
	BundleArtifactFilter *string
	/*BundleGroupFilter
	  If specified, will only return types that are a member of this bundle group.

	*/
	BundleGroupFilter *string
	/*Type
	  If specified, will only return types whose fully qualified classname matches.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reporting task types params
func (o *GetReportingTaskTypesParams) WithTimeout(timeout time.Duration) *GetReportingTaskTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reporting task types params
func (o *GetReportingTaskTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reporting task types params
func (o *GetReportingTaskTypesParams) WithContext(ctx context.Context) *GetReportingTaskTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reporting task types params
func (o *GetReportingTaskTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reporting task types params
func (o *GetReportingTaskTypesParams) WithHTTPClient(client *http.Client) *GetReportingTaskTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reporting task types params
func (o *GetReportingTaskTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleArtifactFilter adds the bundleArtifactFilter to the get reporting task types params
func (o *GetReportingTaskTypesParams) WithBundleArtifactFilter(bundleArtifactFilter *string) *GetReportingTaskTypesParams {
	o.SetBundleArtifactFilter(bundleArtifactFilter)
	return o
}

// SetBundleArtifactFilter adds the bundleArtifactFilter to the get reporting task types params
func (o *GetReportingTaskTypesParams) SetBundleArtifactFilter(bundleArtifactFilter *string) {
	o.BundleArtifactFilter = bundleArtifactFilter
}

// WithBundleGroupFilter adds the bundleGroupFilter to the get reporting task types params
func (o *GetReportingTaskTypesParams) WithBundleGroupFilter(bundleGroupFilter *string) *GetReportingTaskTypesParams {
	o.SetBundleGroupFilter(bundleGroupFilter)
	return o
}

// SetBundleGroupFilter adds the bundleGroupFilter to the get reporting task types params
func (o *GetReportingTaskTypesParams) SetBundleGroupFilter(bundleGroupFilter *string) {
	o.BundleGroupFilter = bundleGroupFilter
}

// WithType adds the typeVar to the get reporting task types params
func (o *GetReportingTaskTypesParams) WithType(typeVar *string) *GetReportingTaskTypesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get reporting task types params
func (o *GetReportingTaskTypesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportingTaskTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BundleArtifactFilter != nil {

		// query param bundleArtifactFilter
		var qrBundleArtifactFilter string
		if o.BundleArtifactFilter != nil {
			qrBundleArtifactFilter = *o.BundleArtifactFilter
		}
		qBundleArtifactFilter := qrBundleArtifactFilter
		if qBundleArtifactFilter != "" {
			if err := r.SetQueryParam("bundleArtifactFilter", qBundleArtifactFilter); err != nil {
				return err
			}
		}

	}

	if o.BundleGroupFilter != nil {

		// query param bundleGroupFilter
		var qrBundleGroupFilter string
		if o.BundleGroupFilter != nil {
			qrBundleGroupFilter = *o.BundleGroupFilter
		}
		qBundleGroupFilter := qrBundleGroupFilter
		if qBundleGroupFilter != "" {
			if err := r.SetQueryParam("bundleGroupFilter", qBundleGroupFilter); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
