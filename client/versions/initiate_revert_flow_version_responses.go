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

// InitiateRevertFlowVersionReader is a Reader for the InitiateRevertFlowVersion structure.
type InitiateRevertFlowVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitiateRevertFlowVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInitiateRevertFlowVersionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInitiateRevertFlowVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInitiateRevertFlowVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInitiateRevertFlowVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInitiateRevertFlowVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewInitiateRevertFlowVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInitiateRevertFlowVersionCreated creates a InitiateRevertFlowVersionCreated with default headers values
func NewInitiateRevertFlowVersionCreated() *InitiateRevertFlowVersionCreated {
	return &InitiateRevertFlowVersionCreated{}
}

/*InitiateRevertFlowVersionCreated handles this case with default header values.

successful operation
*/
type InitiateRevertFlowVersionCreated struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

func (o *InitiateRevertFlowVersionCreated) Error() string {
	return fmt.Sprintf("[POST /versions/revert-requests/process-groups/{id}][%d] initiateRevertFlowVersionCreated  %+v", 201, o.Payload)
}

func (o *InitiateRevertFlowVersionCreated) GetPayload() *models.VersionedFlowUpdateRequestEntity {
	return o.Payload
}

func (o *InitiateRevertFlowVersionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitiateRevertFlowVersionBadRequest creates a InitiateRevertFlowVersionBadRequest with default headers values
func NewInitiateRevertFlowVersionBadRequest() *InitiateRevertFlowVersionBadRequest {
	return &InitiateRevertFlowVersionBadRequest{}
}

/*InitiateRevertFlowVersionBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type InitiateRevertFlowVersionBadRequest struct {
}

func (o *InitiateRevertFlowVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /versions/revert-requests/process-groups/{id}][%d] initiateRevertFlowVersionBadRequest ", 400)
}

func (o *InitiateRevertFlowVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateRevertFlowVersionUnauthorized creates a InitiateRevertFlowVersionUnauthorized with default headers values
func NewInitiateRevertFlowVersionUnauthorized() *InitiateRevertFlowVersionUnauthorized {
	return &InitiateRevertFlowVersionUnauthorized{}
}

/*InitiateRevertFlowVersionUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type InitiateRevertFlowVersionUnauthorized struct {
}

func (o *InitiateRevertFlowVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /versions/revert-requests/process-groups/{id}][%d] initiateRevertFlowVersionUnauthorized ", 401)
}

func (o *InitiateRevertFlowVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateRevertFlowVersionForbidden creates a InitiateRevertFlowVersionForbidden with default headers values
func NewInitiateRevertFlowVersionForbidden() *InitiateRevertFlowVersionForbidden {
	return &InitiateRevertFlowVersionForbidden{}
}

/*InitiateRevertFlowVersionForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type InitiateRevertFlowVersionForbidden struct {
}

func (o *InitiateRevertFlowVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /versions/revert-requests/process-groups/{id}][%d] initiateRevertFlowVersionForbidden ", 403)
}

func (o *InitiateRevertFlowVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateRevertFlowVersionNotFound creates a InitiateRevertFlowVersionNotFound with default headers values
func NewInitiateRevertFlowVersionNotFound() *InitiateRevertFlowVersionNotFound {
	return &InitiateRevertFlowVersionNotFound{}
}

/*InitiateRevertFlowVersionNotFound handles this case with default header values.

The specified resource could not be found.
*/
type InitiateRevertFlowVersionNotFound struct {
}

func (o *InitiateRevertFlowVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /versions/revert-requests/process-groups/{id}][%d] initiateRevertFlowVersionNotFound ", 404)
}

func (o *InitiateRevertFlowVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateRevertFlowVersionConflict creates a InitiateRevertFlowVersionConflict with default headers values
func NewInitiateRevertFlowVersionConflict() *InitiateRevertFlowVersionConflict {
	return &InitiateRevertFlowVersionConflict{}
}

/*InitiateRevertFlowVersionConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type InitiateRevertFlowVersionConflict struct {
}

func (o *InitiateRevertFlowVersionConflict) Error() string {
	return fmt.Sprintf("[POST /versions/revert-requests/process-groups/{id}][%d] initiateRevertFlowVersionConflict ", 409)
}

func (o *InitiateRevertFlowVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
