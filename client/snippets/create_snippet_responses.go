// Code generated by go-swagger; DO NOT EDIT.

package snippets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// CreateSnippetReader is a Reader for the CreateSnippet structure.
type CreateSnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSnippetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSnippetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateSnippetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateSnippetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateSnippetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSnippetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSnippetOK creates a CreateSnippetOK with default headers values
func NewCreateSnippetOK() *CreateSnippetOK {
	return &CreateSnippetOK{}
}

/*CreateSnippetOK handles this case with default header values.

successful operation
*/
type CreateSnippetOK struct {
	Payload *models.SnippetEntity
}

func (o *CreateSnippetOK) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] createSnippetOK  %+v", 200, o.Payload)
}

func (o *CreateSnippetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnippetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnippetBadRequest creates a CreateSnippetBadRequest with default headers values
func NewCreateSnippetBadRequest() *CreateSnippetBadRequest {
	return &CreateSnippetBadRequest{}
}

/*CreateSnippetBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateSnippetBadRequest struct {
}

func (o *CreateSnippetBadRequest) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] createSnippetBadRequest ", 400)
}

func (o *CreateSnippetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSnippetUnauthorized creates a CreateSnippetUnauthorized with default headers values
func NewCreateSnippetUnauthorized() *CreateSnippetUnauthorized {
	return &CreateSnippetUnauthorized{}
}

/*CreateSnippetUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateSnippetUnauthorized struct {
}

func (o *CreateSnippetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] createSnippetUnauthorized ", 401)
}

func (o *CreateSnippetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSnippetForbidden creates a CreateSnippetForbidden with default headers values
func NewCreateSnippetForbidden() *CreateSnippetForbidden {
	return &CreateSnippetForbidden{}
}

/*CreateSnippetForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateSnippetForbidden struct {
}

func (o *CreateSnippetForbidden) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] createSnippetForbidden ", 403)
}

func (o *CreateSnippetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSnippetNotFound creates a CreateSnippetNotFound with default headers values
func NewCreateSnippetNotFound() *CreateSnippetNotFound {
	return &CreateSnippetNotFound{}
}

/*CreateSnippetNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreateSnippetNotFound struct {
}

func (o *CreateSnippetNotFound) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] createSnippetNotFound ", 404)
}

func (o *CreateSnippetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSnippetConflict creates a CreateSnippetConflict with default headers values
func NewCreateSnippetConflict() *CreateSnippetConflict {
	return &CreateSnippetConflict{}
}

/*CreateSnippetConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateSnippetConflict struct {
}

func (o *CreateSnippetConflict) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] createSnippetConflict ", 409)
}

func (o *CreateSnippetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
