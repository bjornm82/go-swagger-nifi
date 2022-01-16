// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

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
)

// NewExtendInputPortTransactionTTLParams creates a new ExtendInputPortTransactionTTLParams object
// with the default values initialized.
func NewExtendInputPortTransactionTTLParams() *ExtendInputPortTransactionTTLParams {
	var ()
	return &ExtendInputPortTransactionTTLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtendInputPortTransactionTTLParamsWithTimeout creates a new ExtendInputPortTransactionTTLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtendInputPortTransactionTTLParamsWithTimeout(timeout time.Duration) *ExtendInputPortTransactionTTLParams {
	var ()
	return &ExtendInputPortTransactionTTLParams{

		timeout: timeout,
	}
}

// NewExtendInputPortTransactionTTLParamsWithContext creates a new ExtendInputPortTransactionTTLParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtendInputPortTransactionTTLParamsWithContext(ctx context.Context) *ExtendInputPortTransactionTTLParams {
	var ()
	return &ExtendInputPortTransactionTTLParams{

		Context: ctx,
	}
}

// NewExtendInputPortTransactionTTLParamsWithHTTPClient creates a new ExtendInputPortTransactionTTLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtendInputPortTransactionTTLParamsWithHTTPClient(client *http.Client) *ExtendInputPortTransactionTTLParams {
	var ()
	return &ExtendInputPortTransactionTTLParams{
		HTTPClient: client,
	}
}

/*ExtendInputPortTransactionTTLParams contains all the parameters to send to the API endpoint
for the extend input port transaction TTL operation typically these are written to a http.Request
*/
type ExtendInputPortTransactionTTLParams struct {

	/*PortID*/
	PortID string
	/*TransactionID*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) WithTimeout(timeout time.Duration) *ExtendInputPortTransactionTTLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) WithContext(ctx context.Context) *ExtendInputPortTransactionTTLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) WithHTTPClient(client *http.Client) *ExtendInputPortTransactionTTLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPortID adds the portID to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) WithPortID(portID string) *ExtendInputPortTransactionTTLParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) SetPortID(portID string) {
	o.PortID = portID
}

// WithTransactionID adds the transactionID to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) WithTransactionID(transactionID string) *ExtendInputPortTransactionTTLParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the extend input port transaction TTL params
func (o *ExtendInputPortTransactionTTLParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ExtendInputPortTransactionTTLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param portId
	if err := r.SetPathParam("portId", o.PortID); err != nil {
		return err
	}

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
