// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// InitiateVersionControlUpdateReader is a Reader for the InitiateVersionControlUpdate structure.
type InitiateVersionControlUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitiateVersionControlUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInitiateVersionControlUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInitiateVersionControlUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInitiateVersionControlUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInitiateVersionControlUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInitiateVersionControlUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewInitiateVersionControlUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInitiateVersionControlUpdateCreated creates a InitiateVersionControlUpdateCreated with default headers values
func NewInitiateVersionControlUpdateCreated() *InitiateVersionControlUpdateCreated {
	return &InitiateVersionControlUpdateCreated{}
}

/*InitiateVersionControlUpdateCreated handles this case with default header values.

successful operation
*/
type InitiateVersionControlUpdateCreated struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

func (o *InitiateVersionControlUpdateCreated) Error() string {
	return fmt.Sprintf("[POST /versions/update-requests/process-groups/{id}][%d] initiateVersionControlUpdateCreated  %+v", 201, o.Payload)
}

func (o *InitiateVersionControlUpdateCreated) GetPayload() *models.VersionedFlowUpdateRequestEntity {
	return o.Payload
}

func (o *InitiateVersionControlUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitiateVersionControlUpdateBadRequest creates a InitiateVersionControlUpdateBadRequest with default headers values
func NewInitiateVersionControlUpdateBadRequest() *InitiateVersionControlUpdateBadRequest {
	return &InitiateVersionControlUpdateBadRequest{}
}

/*InitiateVersionControlUpdateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type InitiateVersionControlUpdateBadRequest struct {
}

func (o *InitiateVersionControlUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /versions/update-requests/process-groups/{id}][%d] initiateVersionControlUpdateBadRequest ", 400)
}

func (o *InitiateVersionControlUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateVersionControlUpdateUnauthorized creates a InitiateVersionControlUpdateUnauthorized with default headers values
func NewInitiateVersionControlUpdateUnauthorized() *InitiateVersionControlUpdateUnauthorized {
	return &InitiateVersionControlUpdateUnauthorized{}
}

/*InitiateVersionControlUpdateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type InitiateVersionControlUpdateUnauthorized struct {
}

func (o *InitiateVersionControlUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /versions/update-requests/process-groups/{id}][%d] initiateVersionControlUpdateUnauthorized ", 401)
}

func (o *InitiateVersionControlUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateVersionControlUpdateForbidden creates a InitiateVersionControlUpdateForbidden with default headers values
func NewInitiateVersionControlUpdateForbidden() *InitiateVersionControlUpdateForbidden {
	return &InitiateVersionControlUpdateForbidden{}
}

/*InitiateVersionControlUpdateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type InitiateVersionControlUpdateForbidden struct {
}

func (o *InitiateVersionControlUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /versions/update-requests/process-groups/{id}][%d] initiateVersionControlUpdateForbidden ", 403)
}

func (o *InitiateVersionControlUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateVersionControlUpdateNotFound creates a InitiateVersionControlUpdateNotFound with default headers values
func NewInitiateVersionControlUpdateNotFound() *InitiateVersionControlUpdateNotFound {
	return &InitiateVersionControlUpdateNotFound{}
}

/*InitiateVersionControlUpdateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type InitiateVersionControlUpdateNotFound struct {
}

func (o *InitiateVersionControlUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /versions/update-requests/process-groups/{id}][%d] initiateVersionControlUpdateNotFound ", 404)
}

func (o *InitiateVersionControlUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateVersionControlUpdateConflict creates a InitiateVersionControlUpdateConflict with default headers values
func NewInitiateVersionControlUpdateConflict() *InitiateVersionControlUpdateConflict {
	return &InitiateVersionControlUpdateConflict{}
}

/*InitiateVersionControlUpdateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type InitiateVersionControlUpdateConflict struct {
}

func (o *InitiateVersionControlUpdateConflict) Error() string {
	return fmt.Sprintf("[POST /versions/update-requests/process-groups/{id}][%d] initiateVersionControlUpdateConflict ", 409)
}

func (o *InitiateVersionControlUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
