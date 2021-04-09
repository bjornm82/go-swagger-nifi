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

// CreatePortTransactionReader is a Reader for the CreatePortTransaction structure.
type CreatePortTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePortTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePortTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePortTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreatePortTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreatePortTransactionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreatePortTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreatePortTransactionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreatePortTransactionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePortTransactionOK creates a CreatePortTransactionOK with default headers values
func NewCreatePortTransactionOK() *CreatePortTransactionOK {
	return &CreatePortTransactionOK{}
}

/*CreatePortTransactionOK handles this case with default header values.

successful operation
*/
type CreatePortTransactionOK struct {
	Payload *models.TransactionResultEntity
}

func (o *CreatePortTransactionOK) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionOK  %+v", 200, o.Payload)
}

func (o *CreatePortTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionResultEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePortTransactionBadRequest creates a CreatePortTransactionBadRequest with default headers values
func NewCreatePortTransactionBadRequest() *CreatePortTransactionBadRequest {
	return &CreatePortTransactionBadRequest{}
}

/*CreatePortTransactionBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreatePortTransactionBadRequest struct {
}

func (o *CreatePortTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionBadRequest ", 400)
}

func (o *CreatePortTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePortTransactionUnauthorized creates a CreatePortTransactionUnauthorized with default headers values
func NewCreatePortTransactionUnauthorized() *CreatePortTransactionUnauthorized {
	return &CreatePortTransactionUnauthorized{}
}

/*CreatePortTransactionUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreatePortTransactionUnauthorized struct {
}

func (o *CreatePortTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionUnauthorized ", 401)
}

func (o *CreatePortTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePortTransactionForbidden creates a CreatePortTransactionForbidden with default headers values
func NewCreatePortTransactionForbidden() *CreatePortTransactionForbidden {
	return &CreatePortTransactionForbidden{}
}

/*CreatePortTransactionForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreatePortTransactionForbidden struct {
}

func (o *CreatePortTransactionForbidden) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionForbidden ", 403)
}

func (o *CreatePortTransactionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePortTransactionNotFound creates a CreatePortTransactionNotFound with default headers values
func NewCreatePortTransactionNotFound() *CreatePortTransactionNotFound {
	return &CreatePortTransactionNotFound{}
}

/*CreatePortTransactionNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreatePortTransactionNotFound struct {
}

func (o *CreatePortTransactionNotFound) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionNotFound ", 404)
}

func (o *CreatePortTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePortTransactionConflict creates a CreatePortTransactionConflict with default headers values
func NewCreatePortTransactionConflict() *CreatePortTransactionConflict {
	return &CreatePortTransactionConflict{}
}

/*CreatePortTransactionConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreatePortTransactionConflict struct {
}

func (o *CreatePortTransactionConflict) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionConflict ", 409)
}

func (o *CreatePortTransactionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePortTransactionServiceUnavailable creates a CreatePortTransactionServiceUnavailable with default headers values
func NewCreatePortTransactionServiceUnavailable() *CreatePortTransactionServiceUnavailable {
	return &CreatePortTransactionServiceUnavailable{}
}

/*CreatePortTransactionServiceUnavailable handles this case with default header values.

NiFi instance is not ready for serving request, or temporarily overloaded. Retrying the same request later may be successful
*/
type CreatePortTransactionServiceUnavailable struct {
}

func (o *CreatePortTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /data-transfer/{portType}/{portId}/transactions][%d] createPortTransactionServiceUnavailable ", 503)
}

func (o *CreatePortTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
