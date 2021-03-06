// Code generated by go-swagger; DO NOT EDIT.

package snippets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateSnippetReader is a Reader for the UpdateSnippet structure.
type UpdateSnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnippetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSnippetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateSnippetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSnippetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSnippetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSnippetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSnippetOK creates a UpdateSnippetOK with default headers values
func NewUpdateSnippetOK() *UpdateSnippetOK {
	return &UpdateSnippetOK{}
}

/*UpdateSnippetOK handles this case with default header values.

successful operation
*/
type UpdateSnippetOK struct {
	Payload *models.SnippetEntity
}

func (o *UpdateSnippetOK) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetOK  %+v", 200, o.Payload)
}

func (o *UpdateSnippetOK) GetPayload() *models.SnippetEntity {
	return o.Payload
}

func (o *UpdateSnippetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnippetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnippetBadRequest creates a UpdateSnippetBadRequest with default headers values
func NewUpdateSnippetBadRequest() *UpdateSnippetBadRequest {
	return &UpdateSnippetBadRequest{}
}

/*UpdateSnippetBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateSnippetBadRequest struct {
}

func (o *UpdateSnippetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetBadRequest ", 400)
}

func (o *UpdateSnippetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetUnauthorized creates a UpdateSnippetUnauthorized with default headers values
func NewUpdateSnippetUnauthorized() *UpdateSnippetUnauthorized {
	return &UpdateSnippetUnauthorized{}
}

/*UpdateSnippetUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateSnippetUnauthorized struct {
}

func (o *UpdateSnippetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetUnauthorized ", 401)
}

func (o *UpdateSnippetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetForbidden creates a UpdateSnippetForbidden with default headers values
func NewUpdateSnippetForbidden() *UpdateSnippetForbidden {
	return &UpdateSnippetForbidden{}
}

/*UpdateSnippetForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateSnippetForbidden struct {
}

func (o *UpdateSnippetForbidden) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetForbidden ", 403)
}

func (o *UpdateSnippetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetNotFound creates a UpdateSnippetNotFound with default headers values
func NewUpdateSnippetNotFound() *UpdateSnippetNotFound {
	return &UpdateSnippetNotFound{}
}

/*UpdateSnippetNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateSnippetNotFound struct {
}

func (o *UpdateSnippetNotFound) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetNotFound ", 404)
}

func (o *UpdateSnippetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetConflict creates a UpdateSnippetConflict with default headers values
func NewUpdateSnippetConflict() *UpdateSnippetConflict {
	return &UpdateSnippetConflict{}
}

/*UpdateSnippetConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateSnippetConflict struct {
}

func (o *UpdateSnippetConflict) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetConflict ", 409)
}

func (o *UpdateSnippetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
