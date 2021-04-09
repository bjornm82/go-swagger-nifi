// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// DeleteUpdateRequestReader is a Reader for the DeleteUpdateRequest structure.
type DeleteUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUpdateRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteUpdateRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteUpdateRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUpdateRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteUpdateRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUpdateRequestOK creates a DeleteUpdateRequestOK with default headers values
func NewDeleteUpdateRequestOK() *DeleteUpdateRequestOK {
	return &DeleteUpdateRequestOK{}
}

/*DeleteUpdateRequestOK handles this case with default header values.

successful operation
*/
type DeleteUpdateRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

func (o *DeleteUpdateRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUpdateRequestBadRequest creates a DeleteUpdateRequestBadRequest with default headers values
func NewDeleteUpdateRequestBadRequest() *DeleteUpdateRequestBadRequest {
	return &DeleteUpdateRequestBadRequest{}
}

/*DeleteUpdateRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteUpdateRequestBadRequest struct {
}

func (o *DeleteUpdateRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestBadRequest ", 400)
}

func (o *DeleteUpdateRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestUnauthorized creates a DeleteUpdateRequestUnauthorized with default headers values
func NewDeleteUpdateRequestUnauthorized() *DeleteUpdateRequestUnauthorized {
	return &DeleteUpdateRequestUnauthorized{}
}

/*DeleteUpdateRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type DeleteUpdateRequestUnauthorized struct {
}

func (o *DeleteUpdateRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestUnauthorized ", 401)
}

func (o *DeleteUpdateRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestForbidden creates a DeleteUpdateRequestForbidden with default headers values
func NewDeleteUpdateRequestForbidden() *DeleteUpdateRequestForbidden {
	return &DeleteUpdateRequestForbidden{}
}

/*DeleteUpdateRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type DeleteUpdateRequestForbidden struct {
}

func (o *DeleteUpdateRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestForbidden ", 403)
}

func (o *DeleteUpdateRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestNotFound creates a DeleteUpdateRequestNotFound with default headers values
func NewDeleteUpdateRequestNotFound() *DeleteUpdateRequestNotFound {
	return &DeleteUpdateRequestNotFound{}
}

/*DeleteUpdateRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type DeleteUpdateRequestNotFound struct {
}

func (o *DeleteUpdateRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestNotFound ", 404)
}

func (o *DeleteUpdateRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestConflict creates a DeleteUpdateRequestConflict with default headers values
func NewDeleteUpdateRequestConflict() *DeleteUpdateRequestConflict {
	return &DeleteUpdateRequestConflict{}
}

/*DeleteUpdateRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteUpdateRequestConflict struct {
}

func (o *DeleteUpdateRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestConflict ", 409)
}

func (o *DeleteUpdateRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
