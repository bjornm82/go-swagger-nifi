// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// TransferFlowFilesReader is a Reader for the TransferFlowFiles structure.
type TransferFlowFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferFlowFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTransferFlowFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTransferFlowFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTransferFlowFilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTransferFlowFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTransferFlowFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewTransferFlowFilesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewTransferFlowFilesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTransferFlowFilesOK creates a TransferFlowFilesOK with default headers values
func NewTransferFlowFilesOK() *TransferFlowFilesOK {
	return &TransferFlowFilesOK{}
}

/*TransferFlowFilesOK handles this case with default header values.

There is no flow file to return.
*/
type TransferFlowFilesOK struct {
	Payload models.StreamingOutput
}

func (o *TransferFlowFilesOK) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesOK  %+v", 200, o.Payload)
}

func (o *TransferFlowFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferFlowFilesBadRequest creates a TransferFlowFilesBadRequest with default headers values
func NewTransferFlowFilesBadRequest() *TransferFlowFilesBadRequest {
	return &TransferFlowFilesBadRequest{}
}

/*TransferFlowFilesBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type TransferFlowFilesBadRequest struct {
}

func (o *TransferFlowFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesBadRequest ", 400)
}

func (o *TransferFlowFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferFlowFilesUnauthorized creates a TransferFlowFilesUnauthorized with default headers values
func NewTransferFlowFilesUnauthorized() *TransferFlowFilesUnauthorized {
	return &TransferFlowFilesUnauthorized{}
}

/*TransferFlowFilesUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type TransferFlowFilesUnauthorized struct {
}

func (o *TransferFlowFilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesUnauthorized ", 401)
}

func (o *TransferFlowFilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferFlowFilesForbidden creates a TransferFlowFilesForbidden with default headers values
func NewTransferFlowFilesForbidden() *TransferFlowFilesForbidden {
	return &TransferFlowFilesForbidden{}
}

/*TransferFlowFilesForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type TransferFlowFilesForbidden struct {
}

func (o *TransferFlowFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesForbidden ", 403)
}

func (o *TransferFlowFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferFlowFilesNotFound creates a TransferFlowFilesNotFound with default headers values
func NewTransferFlowFilesNotFound() *TransferFlowFilesNotFound {
	return &TransferFlowFilesNotFound{}
}

/*TransferFlowFilesNotFound handles this case with default header values.

The specified resource could not be found.
*/
type TransferFlowFilesNotFound struct {
}

func (o *TransferFlowFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesNotFound ", 404)
}

func (o *TransferFlowFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferFlowFilesConflict creates a TransferFlowFilesConflict with default headers values
func NewTransferFlowFilesConflict() *TransferFlowFilesConflict {
	return &TransferFlowFilesConflict{}
}

/*TransferFlowFilesConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type TransferFlowFilesConflict struct {
}

func (o *TransferFlowFilesConflict) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesConflict ", 409)
}

func (o *TransferFlowFilesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferFlowFilesServiceUnavailable creates a TransferFlowFilesServiceUnavailable with default headers values
func NewTransferFlowFilesServiceUnavailable() *TransferFlowFilesServiceUnavailable {
	return &TransferFlowFilesServiceUnavailable{}
}

/*TransferFlowFilesServiceUnavailable handles this case with default header values.

NiFi instance is not ready for serving request, or temporarily overloaded. Retrying the same request later may be successful
*/
type TransferFlowFilesServiceUnavailable struct {
}

func (o *TransferFlowFilesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files][%d] transferFlowFilesServiceUnavailable ", 503)
}

func (o *TransferFlowFilesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}