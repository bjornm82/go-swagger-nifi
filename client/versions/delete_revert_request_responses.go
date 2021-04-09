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

// DeleteRevertRequestReader is a Reader for the DeleteRevertRequest structure.
type DeleteRevertRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRevertRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteRevertRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteRevertRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteRevertRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteRevertRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteRevertRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteRevertRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRevertRequestOK creates a DeleteRevertRequestOK with default headers values
func NewDeleteRevertRequestOK() *DeleteRevertRequestOK {
	return &DeleteRevertRequestOK{}
}

/*DeleteRevertRequestOK handles this case with default header values.

successful operation
*/
type DeleteRevertRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

func (o *DeleteRevertRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteRevertRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRevertRequestBadRequest creates a DeleteRevertRequestBadRequest with default headers values
func NewDeleteRevertRequestBadRequest() *DeleteRevertRequestBadRequest {
	return &DeleteRevertRequestBadRequest{}
}

/*DeleteRevertRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteRevertRequestBadRequest struct {
}

func (o *DeleteRevertRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestBadRequest ", 400)
}

func (o *DeleteRevertRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestUnauthorized creates a DeleteRevertRequestUnauthorized with default headers values
func NewDeleteRevertRequestUnauthorized() *DeleteRevertRequestUnauthorized {
	return &DeleteRevertRequestUnauthorized{}
}

/*DeleteRevertRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type DeleteRevertRequestUnauthorized struct {
}

func (o *DeleteRevertRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestUnauthorized ", 401)
}

func (o *DeleteRevertRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestForbidden creates a DeleteRevertRequestForbidden with default headers values
func NewDeleteRevertRequestForbidden() *DeleteRevertRequestForbidden {
	return &DeleteRevertRequestForbidden{}
}

/*DeleteRevertRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type DeleteRevertRequestForbidden struct {
}

func (o *DeleteRevertRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestForbidden ", 403)
}

func (o *DeleteRevertRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestNotFound creates a DeleteRevertRequestNotFound with default headers values
func NewDeleteRevertRequestNotFound() *DeleteRevertRequestNotFound {
	return &DeleteRevertRequestNotFound{}
}

/*DeleteRevertRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type DeleteRevertRequestNotFound struct {
}

func (o *DeleteRevertRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestNotFound ", 404)
}

func (o *DeleteRevertRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestConflict creates a DeleteRevertRequestConflict with default headers values
func NewDeleteRevertRequestConflict() *DeleteRevertRequestConflict {
	return &DeleteRevertRequestConflict{}
}

/*DeleteRevertRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteRevertRequestConflict struct {
}

func (o *DeleteRevertRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestConflict ", 409)
}

func (o *DeleteRevertRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
