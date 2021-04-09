// Code generated by go-swagger; DO NOT EDIT.

package system_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSystemDiagnosticsParams creates a new GetSystemDiagnosticsParams object
// with the default values initialized.
func NewGetSystemDiagnosticsParams() *GetSystemDiagnosticsParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetSystemDiagnosticsParams{
		Nodewise: &nodewiseDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemDiagnosticsParamsWithTimeout creates a new GetSystemDiagnosticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemDiagnosticsParamsWithTimeout(timeout time.Duration) *GetSystemDiagnosticsParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetSystemDiagnosticsParams{
		Nodewise: &nodewiseDefault,

		timeout: timeout,
	}
}

// NewGetSystemDiagnosticsParamsWithContext creates a new GetSystemDiagnosticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemDiagnosticsParamsWithContext(ctx context.Context) *GetSystemDiagnosticsParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetSystemDiagnosticsParams{
		Nodewise: &nodewiseDefault,

		Context: ctx,
	}
}

// NewGetSystemDiagnosticsParamsWithHTTPClient creates a new GetSystemDiagnosticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemDiagnosticsParamsWithHTTPClient(client *http.Client) *GetSystemDiagnosticsParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetSystemDiagnosticsParams{
		Nodewise:   &nodewiseDefault,
		HTTPClient: client,
	}
}

/*GetSystemDiagnosticsParams contains all the parameters to send to the API endpoint
for the get system diagnostics operation typically these are written to a http.Request
*/
type GetSystemDiagnosticsParams struct {

	/*ClusterNodeID
	  The id of the node where to get the status.

	*/
	ClusterNodeID *string
	/*Nodewise
	  Whether or not to include the breakdown per node. Optional, defaults to false

	*/
	Nodewise *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) WithTimeout(timeout time.Duration) *GetSystemDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) WithContext(ctx context.Context) *GetSystemDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) WithHTTPClient(client *http.Client) *GetSystemDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) WithClusterNodeID(clusterNodeID *string) *GetSystemDiagnosticsParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithNodewise adds the nodewise to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) WithNodewise(nodewise *bool) *GetSystemDiagnosticsParams {
	o.SetNodewise(nodewise)
	return o
}

// SetNodewise adds the nodewise to the get system diagnostics params
func (o *GetSystemDiagnosticsParams) SetNodewise(nodewise *bool) {
	o.Nodewise = nodewise
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterNodeID != nil {

		// query param clusterNodeId
		var qrClusterNodeID string
		if o.ClusterNodeID != nil {
			qrClusterNodeID = *o.ClusterNodeID
		}
		qClusterNodeID := qrClusterNodeID
		if qClusterNodeID != "" {
			if err := r.SetQueryParam("clusterNodeId", qClusterNodeID); err != nil {
				return err
			}
		}

	}

	if o.Nodewise != nil {

		// query param nodewise
		var qrNodewise bool
		if o.Nodewise != nil {
			qrNodewise = *o.Nodewise
		}
		qNodewise := swag.FormatBool(qrNodewise)
		if qNodewise != "" {
			if err := r.SetQueryParam("nodewise", qNodewise); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
