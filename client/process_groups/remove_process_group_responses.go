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

// RemoveProcessGroupReader is a Reader for the RemoveProcessGroup structure.
type RemoveProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveProcessGroupOK creates a RemoveProcessGroupOK with default headers values
func NewRemoveProcessGroupOK() *RemoveProcessGroupOK {
	return &RemoveProcessGroupOK{}
}

/*RemoveProcessGroupOK handles this case with default header values.

successful operation
*/
type RemoveProcessGroupOK struct {
	Payload *models.ProcessGroupEntity
}

func (o *RemoveProcessGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{id}][%d] removeProcessGroupOK  %+v", 200, o.Payload)
}

func (o *RemoveProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProcessGroupBadRequest creates a RemoveProcessGroupBadRequest with default headers values
func NewRemoveProcessGroupBadRequest() *RemoveProcessGroupBadRequest {
	return &RemoveProcessGroupBadRequest{}
}

/*RemoveProcessGroupBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveProcessGroupBadRequest struct {
}

func (o *RemoveProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{id}][%d] removeProcessGroupBadRequest ", 400)
}

func (o *RemoveProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveProcessGroupUnauthorized creates a RemoveProcessGroupUnauthorized with default headers values
func NewRemoveProcessGroupUnauthorized() *RemoveProcessGroupUnauthorized {
	return &RemoveProcessGroupUnauthorized{}
}

/*RemoveProcessGroupUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveProcessGroupUnauthorized struct {
}

func (o *RemoveProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{id}][%d] removeProcessGroupUnauthorized ", 401)
}

func (o *RemoveProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveProcessGroupForbidden creates a RemoveProcessGroupForbidden with default headers values
func NewRemoveProcessGroupForbidden() *RemoveProcessGroupForbidden {
	return &RemoveProcessGroupForbidden{}
}

/*RemoveProcessGroupForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveProcessGroupForbidden struct {
}

func (o *RemoveProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{id}][%d] removeProcessGroupForbidden ", 403)
}

func (o *RemoveProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveProcessGroupNotFound creates a RemoveProcessGroupNotFound with default headers values
func NewRemoveProcessGroupNotFound() *RemoveProcessGroupNotFound {
	return &RemoveProcessGroupNotFound{}
}

/*RemoveProcessGroupNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveProcessGroupNotFound struct {
}

func (o *RemoveProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{id}][%d] removeProcessGroupNotFound ", 404)
}

func (o *RemoveProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveProcessGroupConflict creates a RemoveProcessGroupConflict with default headers values
func NewRemoveProcessGroupConflict() *RemoveProcessGroupConflict {
	return &RemoveProcessGroupConflict{}
}

/*RemoveProcessGroupConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveProcessGroupConflict struct {
}

func (o *RemoveProcessGroupConflict) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{id}][%d] removeProcessGroupConflict ", 409)
}

func (o *RemoveProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}