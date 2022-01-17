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

// CreateAccessTokenFromTicketSamlReader is a Reader for the CreateAccessTokenFromTicketSaml structure.
type CreateAccessTokenFromTicketSamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccessTokenFromTicketSamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAccessTokenFromTicketSamlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccessTokenFromTicketSamlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAccessTokenFromTicketSamlUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAccessTokenFromTicketSamlConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAccessTokenFromTicketSamlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccessTokenFromTicketSamlOK creates a CreateAccessTokenFromTicketSamlOK with default headers values
func NewCreateAccessTokenFromTicketSamlOK() *CreateAccessTokenFromTicketSamlOK {
	return &CreateAccessTokenFromTicketSamlOK{}
}

/*CreateAccessTokenFromTicketSamlOK handles this case with default header values.

successful operation
*/
type CreateAccessTokenFromTicketSamlOK struct {
	Payload string
}

func (o *CreateAccessTokenFromTicketSamlOK) Error() string {
	return fmt.Sprintf("[POST /access/saml/kerberos][%d] createAccessTokenFromTicketSamlOK  %+v", 200, o.Payload)
}

func (o *CreateAccessTokenFromTicketSamlOK) GetPayload() string {
	return o.Payload
}

func (o *CreateAccessTokenFromTicketSamlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccessTokenFromTicketSamlBadRequest creates a CreateAccessTokenFromTicketSamlBadRequest with default headers values
func NewCreateAccessTokenFromTicketSamlBadRequest() *CreateAccessTokenFromTicketSamlBadRequest {
	return &CreateAccessTokenFromTicketSamlBadRequest{}
}

/*CreateAccessTokenFromTicketSamlBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateAccessTokenFromTicketSamlBadRequest struct {
}

func (o *CreateAccessTokenFromTicketSamlBadRequest) Error() string {
	return fmt.Sprintf("[POST /access/saml/kerberos][%d] createAccessTokenFromTicketSamlBadRequest ", 400)
}

func (o *CreateAccessTokenFromTicketSamlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenFromTicketSamlUnauthorized creates a CreateAccessTokenFromTicketSamlUnauthorized with default headers values
func NewCreateAccessTokenFromTicketSamlUnauthorized() *CreateAccessTokenFromTicketSamlUnauthorized {
	return &CreateAccessTokenFromTicketSamlUnauthorized{}
}

/*CreateAccessTokenFromTicketSamlUnauthorized handles this case with default header values.

NiFi was unable to complete the request because it did not contain a valid Kerberos ticket in the Authorization header. Retry this request after initializing a ticket with kinit and ensuring your browser is configured to support SPNEGO.
*/
type CreateAccessTokenFromTicketSamlUnauthorized struct {
}

func (o *CreateAccessTokenFromTicketSamlUnauthorized) Error() string {
	return fmt.Sprintf("[POST /access/saml/kerberos][%d] createAccessTokenFromTicketSamlUnauthorized ", 401)
}

func (o *CreateAccessTokenFromTicketSamlUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenFromTicketSamlConflict creates a CreateAccessTokenFromTicketSamlConflict with default headers values
func NewCreateAccessTokenFromTicketSamlConflict() *CreateAccessTokenFromTicketSamlConflict {
	return &CreateAccessTokenFromTicketSamlConflict{}
}

/*CreateAccessTokenFromTicketSamlConflict handles this case with default header values.

Unable to create access token because NiFi is not in the appropriate state. (i.e. may not be configured to support Kerberos login.
*/
type CreateAccessTokenFromTicketSamlConflict struct {
}

func (o *CreateAccessTokenFromTicketSamlConflict) Error() string {
	return fmt.Sprintf("[POST /access/saml/kerberos][%d] createAccessTokenFromTicketSamlConflict ", 409)
}

func (o *CreateAccessTokenFromTicketSamlConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenFromTicketSamlInternalServerError creates a CreateAccessTokenFromTicketSamlInternalServerError with default headers values
func NewCreateAccessTokenFromTicketSamlInternalServerError() *CreateAccessTokenFromTicketSamlInternalServerError {
	return &CreateAccessTokenFromTicketSamlInternalServerError{}
}

/*CreateAccessTokenFromTicketSamlInternalServerError handles this case with default header values.

Unable to create access token because an unexpected error occurred.
*/
type CreateAccessTokenFromTicketSamlInternalServerError struct {
}

func (o *CreateAccessTokenFromTicketSamlInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/saml/kerberos][%d] createAccessTokenFromTicketSamlInternalServerError ", 500)
}

func (o *CreateAccessTokenFromTicketSamlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}