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

// GetFlowFileReader is a Reader for the GetFlowFile structure.
type GetFlowFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetFlowFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowFileOK creates a GetFlowFileOK with default headers values
func NewGetFlowFileOK() *GetFlowFileOK {
	return &GetFlowFileOK{}
}

/*GetFlowFileOK handles this case with default header values.

successful operation
*/
type GetFlowFileOK struct {
	Payload *models.FlowFileEntity
}

func (o *GetFlowFileOK) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/flowfiles/{flowfile-uuid}][%d] getFlowFileOK  %+v", 200, o.Payload)
}

func (o *GetFlowFileOK) GetPayload() *models.FlowFileEntity {
	return o.Payload
}

func (o *GetFlowFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowFileEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowFileBadRequest creates a GetFlowFileBadRequest with default headers values
func NewGetFlowFileBadRequest() *GetFlowFileBadRequest {
	return &GetFlowFileBadRequest{}
}

/*GetFlowFileBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFlowFileBadRequest struct {
}

func (o *GetFlowFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/flowfiles/{flowfile-uuid}][%d] getFlowFileBadRequest ", 400)
}

func (o *GetFlowFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowFileUnauthorized creates a GetFlowFileUnauthorized with default headers values
func NewGetFlowFileUnauthorized() *GetFlowFileUnauthorized {
	return &GetFlowFileUnauthorized{}
}

/*GetFlowFileUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetFlowFileUnauthorized struct {
}

func (o *GetFlowFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/flowfiles/{flowfile-uuid}][%d] getFlowFileUnauthorized ", 401)
}

func (o *GetFlowFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowFileForbidden creates a GetFlowFileForbidden with default headers values
func NewGetFlowFileForbidden() *GetFlowFileForbidden {
	return &GetFlowFileForbidden{}
}

/*GetFlowFileForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetFlowFileForbidden struct {
}

func (o *GetFlowFileForbidden) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/flowfiles/{flowfile-uuid}][%d] getFlowFileForbidden ", 403)
}

func (o *GetFlowFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowFileNotFound creates a GetFlowFileNotFound with default headers values
func NewGetFlowFileNotFound() *GetFlowFileNotFound {
	return &GetFlowFileNotFound{}
}

/*GetFlowFileNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetFlowFileNotFound struct {
}

func (o *GetFlowFileNotFound) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/flowfiles/{flowfile-uuid}][%d] getFlowFileNotFound ", 404)
}

func (o *GetFlowFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowFileConflict creates a GetFlowFileConflict with default headers values
func NewGetFlowFileConflict() *GetFlowFileConflict {
	return &GetFlowFileConflict{}
}

/*GetFlowFileConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetFlowFileConflict struct {
}

func (o *GetFlowFileConflict) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/flowfiles/{flowfile-uuid}][%d] getFlowFileConflict ", 409)
}

func (o *GetFlowFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
