// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// RemoveRemoteProcessGroupReader is a Reader for the RemoveRemoteProcessGroup structure.
type RemoveRemoteProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveRemoteProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveRemoteProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveRemoteProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveRemoteProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveRemoteProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveRemoteProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveRemoteProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveRemoteProcessGroupOK creates a RemoveRemoteProcessGroupOK with default headers values
func NewRemoveRemoteProcessGroupOK() *RemoveRemoteProcessGroupOK {
	return &RemoveRemoteProcessGroupOK{}
}

/*RemoveRemoteProcessGroupOK handles this case with default header values.

successful operation
*/
type RemoveRemoteProcessGroupOK struct {
	Payload *models.RemoteProcessGroupEntity
}

func (o *RemoveRemoteProcessGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /remote-process-groups/{id}][%d] removeRemoteProcessGroupOK  %+v", 200, o.Payload)
}

func (o *RemoveRemoteProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveRemoteProcessGroupBadRequest creates a RemoveRemoteProcessGroupBadRequest with default headers values
func NewRemoveRemoteProcessGroupBadRequest() *RemoveRemoteProcessGroupBadRequest {
	return &RemoveRemoteProcessGroupBadRequest{}
}

/*RemoveRemoteProcessGroupBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveRemoteProcessGroupBadRequest struct {
}

func (o *RemoveRemoteProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /remote-process-groups/{id}][%d] removeRemoteProcessGroupBadRequest ", 400)
}

func (o *RemoveRemoteProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRemoteProcessGroupUnauthorized creates a RemoveRemoteProcessGroupUnauthorized with default headers values
func NewRemoveRemoteProcessGroupUnauthorized() *RemoveRemoteProcessGroupUnauthorized {
	return &RemoveRemoteProcessGroupUnauthorized{}
}

/*RemoveRemoteProcessGroupUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveRemoteProcessGroupUnauthorized struct {
}

func (o *RemoveRemoteProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /remote-process-groups/{id}][%d] removeRemoteProcessGroupUnauthorized ", 401)
}

func (o *RemoveRemoteProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRemoteProcessGroupForbidden creates a RemoveRemoteProcessGroupForbidden with default headers values
func NewRemoveRemoteProcessGroupForbidden() *RemoveRemoteProcessGroupForbidden {
	return &RemoveRemoteProcessGroupForbidden{}
}

/*RemoveRemoteProcessGroupForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveRemoteProcessGroupForbidden struct {
}

func (o *RemoveRemoteProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /remote-process-groups/{id}][%d] removeRemoteProcessGroupForbidden ", 403)
}

func (o *RemoveRemoteProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRemoteProcessGroupNotFound creates a RemoveRemoteProcessGroupNotFound with default headers values
func NewRemoveRemoteProcessGroupNotFound() *RemoveRemoteProcessGroupNotFound {
	return &RemoveRemoteProcessGroupNotFound{}
}

/*RemoveRemoteProcessGroupNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveRemoteProcessGroupNotFound struct {
}

func (o *RemoveRemoteProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /remote-process-groups/{id}][%d] removeRemoteProcessGroupNotFound ", 404)
}

func (o *RemoveRemoteProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRemoteProcessGroupConflict creates a RemoveRemoteProcessGroupConflict with default headers values
func NewRemoveRemoteProcessGroupConflict() *RemoveRemoteProcessGroupConflict {
	return &RemoveRemoteProcessGroupConflict{}
}

/*RemoveRemoteProcessGroupConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveRemoteProcessGroupConflict struct {
}

func (o *RemoveRemoteProcessGroupConflict) Error() string {
	return fmt.Sprintf("[DELETE /remote-process-groups/{id}][%d] removeRemoteProcessGroupConflict ", 409)
}

func (o *RemoveRemoteProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
