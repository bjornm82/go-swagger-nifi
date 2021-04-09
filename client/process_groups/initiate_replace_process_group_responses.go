// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// InitiateReplaceProcessGroupReader is a Reader for the InitiateReplaceProcessGroup structure.
type InitiateReplaceProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitiateReplaceProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewInitiateReplaceProcessGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewInitiateReplaceProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewInitiateReplaceProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewInitiateReplaceProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewInitiateReplaceProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewInitiateReplaceProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInitiateReplaceProcessGroupCreated creates a InitiateReplaceProcessGroupCreated with default headers values
func NewInitiateReplaceProcessGroupCreated() *InitiateReplaceProcessGroupCreated {
	return &InitiateReplaceProcessGroupCreated{}
}

/*InitiateReplaceProcessGroupCreated handles this case with default header values.

successful operation
*/
type InitiateReplaceProcessGroupCreated struct {
	Payload *models.ProcessGroupReplaceRequestEntity
}

func (o *InitiateReplaceProcessGroupCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/replace-requests][%d] initiateReplaceProcessGroupCreated  %+v", 201, o.Payload)
}

func (o *InitiateReplaceProcessGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupReplaceRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitiateReplaceProcessGroupBadRequest creates a InitiateReplaceProcessGroupBadRequest with default headers values
func NewInitiateReplaceProcessGroupBadRequest() *InitiateReplaceProcessGroupBadRequest {
	return &InitiateReplaceProcessGroupBadRequest{}
}

/*InitiateReplaceProcessGroupBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type InitiateReplaceProcessGroupBadRequest struct {
}

func (o *InitiateReplaceProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/replace-requests][%d] initiateReplaceProcessGroupBadRequest ", 400)
}

func (o *InitiateReplaceProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateReplaceProcessGroupUnauthorized creates a InitiateReplaceProcessGroupUnauthorized with default headers values
func NewInitiateReplaceProcessGroupUnauthorized() *InitiateReplaceProcessGroupUnauthorized {
	return &InitiateReplaceProcessGroupUnauthorized{}
}

/*InitiateReplaceProcessGroupUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type InitiateReplaceProcessGroupUnauthorized struct {
}

func (o *InitiateReplaceProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/replace-requests][%d] initiateReplaceProcessGroupUnauthorized ", 401)
}

func (o *InitiateReplaceProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateReplaceProcessGroupForbidden creates a InitiateReplaceProcessGroupForbidden with default headers values
func NewInitiateReplaceProcessGroupForbidden() *InitiateReplaceProcessGroupForbidden {
	return &InitiateReplaceProcessGroupForbidden{}
}

/*InitiateReplaceProcessGroupForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type InitiateReplaceProcessGroupForbidden struct {
}

func (o *InitiateReplaceProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/replace-requests][%d] initiateReplaceProcessGroupForbidden ", 403)
}

func (o *InitiateReplaceProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateReplaceProcessGroupNotFound creates a InitiateReplaceProcessGroupNotFound with default headers values
func NewInitiateReplaceProcessGroupNotFound() *InitiateReplaceProcessGroupNotFound {
	return &InitiateReplaceProcessGroupNotFound{}
}

/*InitiateReplaceProcessGroupNotFound handles this case with default header values.

The specified resource could not be found.
*/
type InitiateReplaceProcessGroupNotFound struct {
}

func (o *InitiateReplaceProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/replace-requests][%d] initiateReplaceProcessGroupNotFound ", 404)
}

func (o *InitiateReplaceProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateReplaceProcessGroupConflict creates a InitiateReplaceProcessGroupConflict with default headers values
func NewInitiateReplaceProcessGroupConflict() *InitiateReplaceProcessGroupConflict {
	return &InitiateReplaceProcessGroupConflict{}
}

/*InitiateReplaceProcessGroupConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type InitiateReplaceProcessGroupConflict struct {
}

func (o *InitiateReplaceProcessGroupConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/replace-requests][%d] initiateReplaceProcessGroupConflict ", 409)
}

func (o *InitiateReplaceProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
