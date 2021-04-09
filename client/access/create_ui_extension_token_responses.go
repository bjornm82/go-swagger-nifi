// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateUIExtensionTokenReader is a Reader for the CreateUIExtensionToken structure.
type CreateUIExtensionTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUIExtensionTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUIExtensionTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewCreateUIExtensionTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateUIExtensionTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateUIExtensionTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUIExtensionTokenOK creates a CreateUIExtensionTokenOK with default headers values
func NewCreateUIExtensionTokenOK() *CreateUIExtensionTokenOK {
	return &CreateUIExtensionTokenOK{}
}

/*CreateUIExtensionTokenOK handles this case with default header values.

successful operation
*/
type CreateUIExtensionTokenOK struct {
	Payload string
}

func (o *CreateUIExtensionTokenOK) Error() string {
	return fmt.Sprintf("[POST /access/ui-extension-token][%d] createUiExtensionTokenOK  %+v", 200, o.Payload)
}

func (o *CreateUIExtensionTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUIExtensionTokenForbidden creates a CreateUIExtensionTokenForbidden with default headers values
func NewCreateUIExtensionTokenForbidden() *CreateUIExtensionTokenForbidden {
	return &CreateUIExtensionTokenForbidden{}
}

/*CreateUIExtensionTokenForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateUIExtensionTokenForbidden struct {
}

func (o *CreateUIExtensionTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /access/ui-extension-token][%d] createUiExtensionTokenForbidden ", 403)
}

func (o *CreateUIExtensionTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUIExtensionTokenConflict creates a CreateUIExtensionTokenConflict with default headers values
func NewCreateUIExtensionTokenConflict() *CreateUIExtensionTokenConflict {
	return &CreateUIExtensionTokenConflict{}
}

/*CreateUIExtensionTokenConflict handles this case with default header values.

Unable to create the download token because NiFi is not in the appropriate state. (i.e. may not have any tokens to grant or be configured to support username/password login)
*/
type CreateUIExtensionTokenConflict struct {
}

func (o *CreateUIExtensionTokenConflict) Error() string {
	return fmt.Sprintf("[POST /access/ui-extension-token][%d] createUiExtensionTokenConflict ", 409)
}

func (o *CreateUIExtensionTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUIExtensionTokenInternalServerError creates a CreateUIExtensionTokenInternalServerError with default headers values
func NewCreateUIExtensionTokenInternalServerError() *CreateUIExtensionTokenInternalServerError {
	return &CreateUIExtensionTokenInternalServerError{}
}

/*CreateUIExtensionTokenInternalServerError handles this case with default header values.

Unable to create download token because an unexpected error occurred.
*/
type CreateUIExtensionTokenInternalServerError struct {
}

func (o *CreateUIExtensionTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/ui-extension-token][%d] createUiExtensionTokenInternalServerError ", 500)
}

func (o *CreateUIExtensionTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
