// Code generated by go-swagger; DO NOT EDIT.

package funnel

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
	"github.com/go-openapi/swag"
)

// NewRemoveFunnelParams creates a new RemoveFunnelParams object
// with the default values initialized.
func NewRemoveFunnelParams() *RemoveFunnelParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &RemoveFunnelParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveFunnelParamsWithTimeout creates a new RemoveFunnelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveFunnelParamsWithTimeout(timeout time.Duration) *RemoveFunnelParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &RemoveFunnelParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		timeout: timeout,
	}
}

// NewRemoveFunnelParamsWithContext creates a new RemoveFunnelParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveFunnelParamsWithContext(ctx context.Context) *RemoveFunnelParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &RemoveFunnelParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		Context: ctx,
	}
}

// NewRemoveFunnelParamsWithHTTPClient creates a new RemoveFunnelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveFunnelParamsWithHTTPClient(client *http.Client) *RemoveFunnelParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &RemoveFunnelParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
		HTTPClient:                   client,
	}
}

/*RemoveFunnelParams contains all the parameters to send to the API endpoint
for the remove funnel operation typically these are written to a http.Request
*/
type RemoveFunnelParams struct {

	/*ClientID
	  If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.

	*/
	ClientID *string
	/*DisconnectedNodeAcknowledged
	  Acknowledges that this node is disconnected to allow for mutable requests to proceed.

	*/
	DisconnectedNodeAcknowledged *bool
	/*ID
	  The funnel id.

	*/
	ID string
	/*Version
	  The revision is used to verify the client is working with the latest version of the flow.

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove funnel params
func (o *RemoveFunnelParams) WithTimeout(timeout time.Duration) *RemoveFunnelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove funnel params
func (o *RemoveFunnelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove funnel params
func (o *RemoveFunnelParams) WithContext(ctx context.Context) *RemoveFunnelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove funnel params
func (o *RemoveFunnelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove funnel params
func (o *RemoveFunnelParams) WithHTTPClient(client *http.Client) *RemoveFunnelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove funnel params
func (o *RemoveFunnelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the remove funnel params
func (o *RemoveFunnelParams) WithClientID(clientID *string) *RemoveFunnelParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the remove funnel params
func (o *RemoveFunnelParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the remove funnel params
func (o *RemoveFunnelParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *RemoveFunnelParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the remove funnel params
func (o *RemoveFunnelParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithID adds the id to the remove funnel params
func (o *RemoveFunnelParams) WithID(id string) *RemoveFunnelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove funnel params
func (o *RemoveFunnelParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the remove funnel params
func (o *RemoveFunnelParams) WithVersion(version *string) *RemoveFunnelParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the remove funnel params
func (o *RemoveFunnelParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveFunnelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// query param clientId
		var qrClientID string
		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {
			if err := r.SetQueryParam("clientId", qClientID); err != nil {
				return err
			}
		}

	}

	if o.DisconnectedNodeAcknowledged != nil {

		// query param disconnectedNodeAcknowledged
		var qrDisconnectedNodeAcknowledged bool
		if o.DisconnectedNodeAcknowledged != nil {
			qrDisconnectedNodeAcknowledged = *o.DisconnectedNodeAcknowledged
		}
		qDisconnectedNodeAcknowledged := swag.FormatBool(qrDisconnectedNodeAcknowledged)
		if qDisconnectedNodeAcknowledged != "" {
			if err := r.SetQueryParam("disconnectedNodeAcknowledged", qDisconnectedNodeAcknowledged); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
