// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// CommitInputPortTransactionReader is a Reader for the CommitInputPortTransaction structure.
type CommitInputPortTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitInputPortTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCommitInputPortTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCommitInputPortTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCommitInputPortTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCommitInputPortTransactionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCommitInputPortTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCommitInputPortTransactionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCommitInputPortTransactionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommitInputPortTransactionOK creates a CommitInputPortTransactionOK with default headers values
func NewCommitInputPortTransactionOK() *CommitInputPortTransactionOK {
	return &CommitInputPortTransactionOK{}
}

/*CommitInputPortTransactionOK handles this case with default header values.

successful operation
*/
type CommitInputPortTransactionOK struct {
	Payload *models.TransactionResultEntity
}

func (o *CommitInputPortTransactionOK) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionOK  %+v", 200, o.Payload)
}

func (o *CommitInputPortTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionResultEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitInputPortTransactionBadRequest creates a CommitInputPortTransactionBadRequest with default headers values
func NewCommitInputPortTransactionBadRequest() *CommitInputPortTransactionBadRequest {
	return &CommitInputPortTransactionBadRequest{}
}

/*CommitInputPortTransactionBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CommitInputPortTransactionBadRequest struct {
}

func (o *CommitInputPortTransactionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionBadRequest ", 400)
}

func (o *CommitInputPortTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitInputPortTransactionUnauthorized creates a CommitInputPortTransactionUnauthorized with default headers values
func NewCommitInputPortTransactionUnauthorized() *CommitInputPortTransactionUnauthorized {
	return &CommitInputPortTransactionUnauthorized{}
}

/*CommitInputPortTransactionUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CommitInputPortTransactionUnauthorized struct {
}

func (o *CommitInputPortTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionUnauthorized ", 401)
}

func (o *CommitInputPortTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitInputPortTransactionForbidden creates a CommitInputPortTransactionForbidden with default headers values
func NewCommitInputPortTransactionForbidden() *CommitInputPortTransactionForbidden {
	return &CommitInputPortTransactionForbidden{}
}

/*CommitInputPortTransactionForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CommitInputPortTransactionForbidden struct {
}

func (o *CommitInputPortTransactionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionForbidden ", 403)
}

func (o *CommitInputPortTransactionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitInputPortTransactionNotFound creates a CommitInputPortTransactionNotFound with default headers values
func NewCommitInputPortTransactionNotFound() *CommitInputPortTransactionNotFound {
	return &CommitInputPortTransactionNotFound{}
}

/*CommitInputPortTransactionNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CommitInputPortTransactionNotFound struct {
}

func (o *CommitInputPortTransactionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionNotFound ", 404)
}

func (o *CommitInputPortTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitInputPortTransactionConflict creates a CommitInputPortTransactionConflict with default headers values
func NewCommitInputPortTransactionConflict() *CommitInputPortTransactionConflict {
	return &CommitInputPortTransactionConflict{}
}

/*CommitInputPortTransactionConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CommitInputPortTransactionConflict struct {
}

func (o *CommitInputPortTransactionConflict) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionConflict ", 409)
}

func (o *CommitInputPortTransactionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitInputPortTransactionServiceUnavailable creates a CommitInputPortTransactionServiceUnavailable with default headers values
func NewCommitInputPortTransactionServiceUnavailable() *CommitInputPortTransactionServiceUnavailable {
	return &CommitInputPortTransactionServiceUnavailable{}
}

/*CommitInputPortTransactionServiceUnavailable handles this case with default header values.

NiFi instance is not ready for serving request, or temporarily overloaded. Retrying the same request later may be successful
*/
type CommitInputPortTransactionServiceUnavailable struct {
}

func (o *CommitInputPortTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/input-ports/{portId}/transactions/{transactionId}][%d] commitInputPortTransactionServiceUnavailable ", 503)
}

func (o *CommitInputPortTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
