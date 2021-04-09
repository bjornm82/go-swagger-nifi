// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

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

// NewReceiveFlowFilesParams creates a new ReceiveFlowFilesParams object
// with the default values initialized.
func NewReceiveFlowFilesParams() *ReceiveFlowFilesParams {
	var ()
	return &ReceiveFlowFilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReceiveFlowFilesParamsWithTimeout creates a new ReceiveFlowFilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReceiveFlowFilesParamsWithTimeout(timeout time.Duration) *ReceiveFlowFilesParams {
	var ()
	return &ReceiveFlowFilesParams{

		timeout: timeout,
	}
}

// NewReceiveFlowFilesParamsWithContext creates a new ReceiveFlowFilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewReceiveFlowFilesParamsWithContext(ctx context.Context) *ReceiveFlowFilesParams {
	var ()
	return &ReceiveFlowFilesParams{

		Context: ctx,
	}
}

// NewReceiveFlowFilesParamsWithHTTPClient creates a new ReceiveFlowFilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReceiveFlowFilesParamsWithHTTPClient(client *http.Client) *ReceiveFlowFilesParams {
	var ()
	return &ReceiveFlowFilesParams{
		HTTPClient: client,
	}
}

/*ReceiveFlowFilesParams contains all the parameters to send to the API endpoint
for the receive flow files operation typically these are written to a http.Request
*/
type ReceiveFlowFilesParams struct {

	/*PortID
	  The input port id.

	*/
	PortID string
	/*TransactionID*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the receive flow files params
func (o *ReceiveFlowFilesParams) WithTimeout(timeout time.Duration) *ReceiveFlowFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the receive flow files params
func (o *ReceiveFlowFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the receive flow files params
func (o *ReceiveFlowFilesParams) WithContext(ctx context.Context) *ReceiveFlowFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the receive flow files params
func (o *ReceiveFlowFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the receive flow files params
func (o *ReceiveFlowFilesParams) WithHTTPClient(client *http.Client) *ReceiveFlowFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the receive flow files params
func (o *ReceiveFlowFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPortID adds the portID to the receive flow files params
func (o *ReceiveFlowFilesParams) WithPortID(portID string) *ReceiveFlowFilesParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the receive flow files params
func (o *ReceiveFlowFilesParams) SetPortID(portID string) {
	o.PortID = portID
}

// WithTransactionID adds the transactionID to the receive flow files params
func (o *ReceiveFlowFilesParams) WithTransactionID(transactionID string) *ReceiveFlowFilesParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the receive flow files params
func (o *ReceiveFlowFilesParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ReceiveFlowFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
