// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// CopySnippetReader is a Reader for the CopySnippet structure.
type CopySnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopySnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCopySnippetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCopySnippetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCopySnippetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCopySnippetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopySnippetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCopySnippetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCopySnippetOK creates a CopySnippetOK with default headers values
func NewCopySnippetOK() *CopySnippetOK {
	return &CopySnippetOK{}
}

/*CopySnippetOK handles this case with default header values.

successful operation
*/
type CopySnippetOK struct {
	Payload *models.FlowEntity
}

func (o *CopySnippetOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/snippet-instance][%d] copySnippetOK  %+v", 200, o.Payload)
}

func (o *CopySnippetOK) GetPayload() *models.FlowEntity {
	return o.Payload
}

func (o *CopySnippetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopySnippetBadRequest creates a CopySnippetBadRequest with default headers values
func NewCopySnippetBadRequest() *CopySnippetBadRequest {
	return &CopySnippetBadRequest{}
}

/*CopySnippetBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CopySnippetBadRequest struct {
}

func (o *CopySnippetBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/snippet-instance][%d] copySnippetBadRequest ", 400)
}

func (o *CopySnippetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopySnippetUnauthorized creates a CopySnippetUnauthorized with default headers values
func NewCopySnippetUnauthorized() *CopySnippetUnauthorized {
	return &CopySnippetUnauthorized{}
}

/*CopySnippetUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CopySnippetUnauthorized struct {
}

func (o *CopySnippetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/snippet-instance][%d] copySnippetUnauthorized ", 401)
}

func (o *CopySnippetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopySnippetForbidden creates a CopySnippetForbidden with default headers values
func NewCopySnippetForbidden() *CopySnippetForbidden {
	return &CopySnippetForbidden{}
}

/*CopySnippetForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CopySnippetForbidden struct {
}

func (o *CopySnippetForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/snippet-instance][%d] copySnippetForbidden ", 403)
}

func (o *CopySnippetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopySnippetNotFound creates a CopySnippetNotFound with default headers values
func NewCopySnippetNotFound() *CopySnippetNotFound {
	return &CopySnippetNotFound{}
}

/*CopySnippetNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CopySnippetNotFound struct {
}

func (o *CopySnippetNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/snippet-instance][%d] copySnippetNotFound ", 404)
}

func (o *CopySnippetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopySnippetConflict creates a CopySnippetConflict with default headers values
func NewCopySnippetConflict() *CopySnippetConflict {
	return &CopySnippetConflict{}
}

/*CopySnippetConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CopySnippetConflict struct {
}

func (o *CopySnippetConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/snippet-instance][%d] copySnippetConflict ", 409)
}

func (o *CopySnippetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
