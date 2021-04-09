// Code generated by go-swagger; DO NOT EDIT.

package flowfile_queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// DeleteListingRequestReader is a Reader for the DeleteListingRequest structure.
type DeleteListingRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteListingRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteListingRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteListingRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteListingRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteListingRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteListingRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteListingRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteListingRequestOK creates a DeleteListingRequestOK with default headers values
func NewDeleteListingRequestOK() *DeleteListingRequestOK {
	return &DeleteListingRequestOK{}
}

/*DeleteListingRequestOK handles this case with default header values.

successful operation
*/
type DeleteListingRequestOK struct {
	Payload *models.ListingRequestEntity
}

func (o *DeleteListingRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteListingRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListingRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteListingRequestBadRequest creates a DeleteListingRequestBadRequest with default headers values
func NewDeleteListingRequestBadRequest() *DeleteListingRequestBadRequest {
	return &DeleteListingRequestBadRequest{}
}

/*DeleteListingRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteListingRequestBadRequest struct {
}

func (o *DeleteListingRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestBadRequest ", 400)
}

func (o *DeleteListingRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestUnauthorized creates a DeleteListingRequestUnauthorized with default headers values
func NewDeleteListingRequestUnauthorized() *DeleteListingRequestUnauthorized {
	return &DeleteListingRequestUnauthorized{}
}

/*DeleteListingRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type DeleteListingRequestUnauthorized struct {
}

func (o *DeleteListingRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestUnauthorized ", 401)
}

func (o *DeleteListingRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestForbidden creates a DeleteListingRequestForbidden with default headers values
func NewDeleteListingRequestForbidden() *DeleteListingRequestForbidden {
	return &DeleteListingRequestForbidden{}
}

/*DeleteListingRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type DeleteListingRequestForbidden struct {
}

func (o *DeleteListingRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestForbidden ", 403)
}

func (o *DeleteListingRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestNotFound creates a DeleteListingRequestNotFound with default headers values
func NewDeleteListingRequestNotFound() *DeleteListingRequestNotFound {
	return &DeleteListingRequestNotFound{}
}

/*DeleteListingRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type DeleteListingRequestNotFound struct {
}

func (o *DeleteListingRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestNotFound ", 404)
}

func (o *DeleteListingRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestConflict creates a DeleteListingRequestConflict with default headers values
func NewDeleteListingRequestConflict() *DeleteListingRequestConflict {
	return &DeleteListingRequestConflict{}
}

/*DeleteListingRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteListingRequestConflict struct {
}

func (o *DeleteListingRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestConflict ", 409)
}

func (o *DeleteListingRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
