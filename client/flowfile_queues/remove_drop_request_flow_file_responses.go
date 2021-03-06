// Code generated by go-swagger; DO NOT EDIT.

package flowfile_queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// RemoveDropRequestFlowFileReader is a Reader for the RemoveDropRequestFlowFile structure.
type RemoveDropRequestFlowFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDropRequestFlowFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveDropRequestFlowFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveDropRequestFlowFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveDropRequestFlowFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveDropRequestFlowFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveDropRequestFlowFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveDropRequestFlowFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveDropRequestFlowFileOK creates a RemoveDropRequestFlowFileOK with default headers values
func NewRemoveDropRequestFlowFileOK() *RemoveDropRequestFlowFileOK {
	return &RemoveDropRequestFlowFileOK{}
}

/*RemoveDropRequestFlowFileOK handles this case with default header values.

successful operation
*/
type RemoveDropRequestFlowFileOK struct {
	Payload *models.DropRequestEntity
}

func (o *RemoveDropRequestFlowFileOK) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] removeDropRequestFlowFileOK  %+v", 200, o.Payload)
}

func (o *RemoveDropRequestFlowFileOK) GetPayload() *models.DropRequestEntity {
	return o.Payload
}

func (o *RemoveDropRequestFlowFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DropRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveDropRequestFlowFileBadRequest creates a RemoveDropRequestFlowFileBadRequest with default headers values
func NewRemoveDropRequestFlowFileBadRequest() *RemoveDropRequestFlowFileBadRequest {
	return &RemoveDropRequestFlowFileBadRequest{}
}

/*RemoveDropRequestFlowFileBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveDropRequestFlowFileBadRequest struct {
}

func (o *RemoveDropRequestFlowFileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] removeDropRequestFlowFileBadRequest ", 400)
}

func (o *RemoveDropRequestFlowFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDropRequestFlowFileUnauthorized creates a RemoveDropRequestFlowFileUnauthorized with default headers values
func NewRemoveDropRequestFlowFileUnauthorized() *RemoveDropRequestFlowFileUnauthorized {
	return &RemoveDropRequestFlowFileUnauthorized{}
}

/*RemoveDropRequestFlowFileUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveDropRequestFlowFileUnauthorized struct {
}

func (o *RemoveDropRequestFlowFileUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] removeDropRequestFlowFileUnauthorized ", 401)
}

func (o *RemoveDropRequestFlowFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDropRequestFlowFileForbidden creates a RemoveDropRequestFlowFileForbidden with default headers values
func NewRemoveDropRequestFlowFileForbidden() *RemoveDropRequestFlowFileForbidden {
	return &RemoveDropRequestFlowFileForbidden{}
}

/*RemoveDropRequestFlowFileForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveDropRequestFlowFileForbidden struct {
}

func (o *RemoveDropRequestFlowFileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] removeDropRequestFlowFileForbidden ", 403)
}

func (o *RemoveDropRequestFlowFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDropRequestFlowFileNotFound creates a RemoveDropRequestFlowFileNotFound with default headers values
func NewRemoveDropRequestFlowFileNotFound() *RemoveDropRequestFlowFileNotFound {
	return &RemoveDropRequestFlowFileNotFound{}
}

/*RemoveDropRequestFlowFileNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveDropRequestFlowFileNotFound struct {
}

func (o *RemoveDropRequestFlowFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] removeDropRequestFlowFileNotFound ", 404)
}

func (o *RemoveDropRequestFlowFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDropRequestFlowFileConflict creates a RemoveDropRequestFlowFileConflict with default headers values
func NewRemoveDropRequestFlowFileConflict() *RemoveDropRequestFlowFileConflict {
	return &RemoveDropRequestFlowFileConflict{}
}

/*RemoveDropRequestFlowFileConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveDropRequestFlowFileConflict struct {
}

func (o *RemoveDropRequestFlowFileConflict) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] removeDropRequestFlowFileConflict ", 409)
}

func (o *RemoveDropRequestFlowFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
