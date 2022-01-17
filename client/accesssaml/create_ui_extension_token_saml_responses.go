// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateUIExtensionTokenSamlReader is a Reader for the CreateUIExtensionTokenSaml structure.
type CreateUIExtensionTokenSamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUIExtensionTokenSamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUIExtensionTokenSamlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateUIExtensionTokenSamlForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUIExtensionTokenSamlConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUIExtensionTokenSamlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUIExtensionTokenSamlOK creates a CreateUIExtensionTokenSamlOK with default headers values
func NewCreateUIExtensionTokenSamlOK() *CreateUIExtensionTokenSamlOK {
	return &CreateUIExtensionTokenSamlOK{}
}

/*CreateUIExtensionTokenSamlOK handles this case with default header values.

successful operation
*/
type CreateUIExtensionTokenSamlOK struct {
	Payload string
}

func (o *CreateUIExtensionTokenSamlOK) Error() string {
	return fmt.Sprintf("[POST /access/saml/ui-extension-token][%d] createUiExtensionTokenSamlOK  %+v", 200, o.Payload)
}

func (o *CreateUIExtensionTokenSamlOK) GetPayload() string {
	return o.Payload
}

func (o *CreateUIExtensionTokenSamlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUIExtensionTokenSamlForbidden creates a CreateUIExtensionTokenSamlForbidden with default headers values
func NewCreateUIExtensionTokenSamlForbidden() *CreateUIExtensionTokenSamlForbidden {
	return &CreateUIExtensionTokenSamlForbidden{}
}

/*CreateUIExtensionTokenSamlForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateUIExtensionTokenSamlForbidden struct {
}

func (o *CreateUIExtensionTokenSamlForbidden) Error() string {
	return fmt.Sprintf("[POST /access/saml/ui-extension-token][%d] createUiExtensionTokenSamlForbidden ", 403)
}

func (o *CreateUIExtensionTokenSamlForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUIExtensionTokenSamlConflict creates a CreateUIExtensionTokenSamlConflict with default headers values
func NewCreateUIExtensionTokenSamlConflict() *CreateUIExtensionTokenSamlConflict {
	return &CreateUIExtensionTokenSamlConflict{}
}

/*CreateUIExtensionTokenSamlConflict handles this case with default header values.

Unable to create the download token because NiFi is not in the appropriate state. (i.e. may not have any tokens to grant or be configured to support username/password login)
*/
type CreateUIExtensionTokenSamlConflict struct {
}

func (o *CreateUIExtensionTokenSamlConflict) Error() string {
	return fmt.Sprintf("[POST /access/saml/ui-extension-token][%d] createUiExtensionTokenSamlConflict ", 409)
}

func (o *CreateUIExtensionTokenSamlConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUIExtensionTokenSamlInternalServerError creates a CreateUIExtensionTokenSamlInternalServerError with default headers values
func NewCreateUIExtensionTokenSamlInternalServerError() *CreateUIExtensionTokenSamlInternalServerError {
	return &CreateUIExtensionTokenSamlInternalServerError{}
}

/*CreateUIExtensionTokenSamlInternalServerError handles this case with default header values.

Unable to create download token because an unexpected error occurred.
*/
type CreateUIExtensionTokenSamlInternalServerError struct {
}

func (o *CreateUIExtensionTokenSamlInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/saml/ui-extension-token][%d] createUiExtensionTokenSamlInternalServerError ", 500)
}

func (o *CreateUIExtensionTokenSamlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}