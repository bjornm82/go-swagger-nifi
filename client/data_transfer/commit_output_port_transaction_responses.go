// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// CommitOutputPortTransactionReader is a Reader for the CommitOutputPortTransaction structure.
type CommitOutputPortTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitOutputPortTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitOutputPortTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCommitOutputPortTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCommitOutputPortTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCommitOutputPortTransactionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCommitOutputPortTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCommitOutputPortTransactionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCommitOutputPortTransactionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommitOutputPortTransactionOK creates a CommitOutputPortTransactionOK with default headers values
func NewCommitOutputPortTransactionOK() *CommitOutputPortTransactionOK {
	return &CommitOutputPortTransactionOK{}
}

/*CommitOutputPortTransactionOK handles this case with default header values.

successful operation
*/
type CommitOutputPortTransactionOK struct {
	Payload *models.TransactionResultEntity
}

func (o *CommitOutputPortTransactionOK) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionOK  %+v", 200, o.Payload)
}

func (o *CommitOutputPortTransactionOK) GetPayload() *models.TransactionResultEntity {
	return o.Payload
}

func (o *CommitOutputPortTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionResultEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitOutputPortTransactionBadRequest creates a CommitOutputPortTransactionBadRequest with default headers values
func NewCommitOutputPortTransactionBadRequest() *CommitOutputPortTransactionBadRequest {
	return &CommitOutputPortTransactionBadRequest{}
}

/*CommitOutputPortTransactionBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CommitOutputPortTransactionBadRequest struct {
}

func (o *CommitOutputPortTransactionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionBadRequest ", 400)
}

func (o *CommitOutputPortTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionUnauthorized creates a CommitOutputPortTransactionUnauthorized with default headers values
func NewCommitOutputPortTransactionUnauthorized() *CommitOutputPortTransactionUnauthorized {
	return &CommitOutputPortTransactionUnauthorized{}
}

/*CommitOutputPortTransactionUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CommitOutputPortTransactionUnauthorized struct {
}

func (o *CommitOutputPortTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionUnauthorized ", 401)
}

func (o *CommitOutputPortTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionForbidden creates a CommitOutputPortTransactionForbidden with default headers values
func NewCommitOutputPortTransactionForbidden() *CommitOutputPortTransactionForbidden {
	return &CommitOutputPortTransactionForbidden{}
}

/*CommitOutputPortTransactionForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CommitOutputPortTransactionForbidden struct {
}

func (o *CommitOutputPortTransactionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionForbidden ", 403)
}

func (o *CommitOutputPortTransactionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionNotFound creates a CommitOutputPortTransactionNotFound with default headers values
func NewCommitOutputPortTransactionNotFound() *CommitOutputPortTransactionNotFound {
	return &CommitOutputPortTransactionNotFound{}
}

/*CommitOutputPortTransactionNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CommitOutputPortTransactionNotFound struct {
}

func (o *CommitOutputPortTransactionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionNotFound ", 404)
}

func (o *CommitOutputPortTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionConflict creates a CommitOutputPortTransactionConflict with default headers values
func NewCommitOutputPortTransactionConflict() *CommitOutputPortTransactionConflict {
	return &CommitOutputPortTransactionConflict{}
}

/*CommitOutputPortTransactionConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CommitOutputPortTransactionConflict struct {
}

func (o *CommitOutputPortTransactionConflict) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionConflict ", 409)
}

func (o *CommitOutputPortTransactionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionServiceUnavailable creates a CommitOutputPortTransactionServiceUnavailable with default headers values
func NewCommitOutputPortTransactionServiceUnavailable() *CommitOutputPortTransactionServiceUnavailable {
	return &CommitOutputPortTransactionServiceUnavailable{}
}

/*CommitOutputPortTransactionServiceUnavailable handles this case with default header values.

NiFi instance is not ready for serving request, or temporarily overloaded. Retrying the same request later may be successful
*/
type CommitOutputPortTransactionServiceUnavailable struct {
}

func (o *CommitOutputPortTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionServiceUnavailable ", 503)
}

func (o *CommitOutputPortTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
