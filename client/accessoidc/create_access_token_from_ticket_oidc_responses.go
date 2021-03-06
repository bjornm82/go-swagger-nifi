// Code generated by go-swagger; DO NOT EDIT.

package accessoidc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateAccessTokenFromTicketOidcReader is a Reader for the CreateAccessTokenFromTicketOidc structure.
type CreateAccessTokenFromTicketOidcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccessTokenFromTicketOidcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAccessTokenFromTicketOidcCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccessTokenFromTicketOidcBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAccessTokenFromTicketOidcUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAccessTokenFromTicketOidcConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAccessTokenFromTicketOidcInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccessTokenFromTicketOidcCreated creates a CreateAccessTokenFromTicketOidcCreated with default headers values
func NewCreateAccessTokenFromTicketOidcCreated() *CreateAccessTokenFromTicketOidcCreated {
	return &CreateAccessTokenFromTicketOidcCreated{}
}

/*CreateAccessTokenFromTicketOidcCreated handles this case with default header values.

successful operation
*/
type CreateAccessTokenFromTicketOidcCreated struct {
	Payload string
}

func (o *CreateAccessTokenFromTicketOidcCreated) Error() string {
	return fmt.Sprintf("[POST /access/oidc/kerberos][%d] createAccessTokenFromTicketOidcCreated  %+v", 201, o.Payload)
}

func (o *CreateAccessTokenFromTicketOidcCreated) GetPayload() string {
	return o.Payload
}

func (o *CreateAccessTokenFromTicketOidcCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccessTokenFromTicketOidcBadRequest creates a CreateAccessTokenFromTicketOidcBadRequest with default headers values
func NewCreateAccessTokenFromTicketOidcBadRequest() *CreateAccessTokenFromTicketOidcBadRequest {
	return &CreateAccessTokenFromTicketOidcBadRequest{}
}

/*CreateAccessTokenFromTicketOidcBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateAccessTokenFromTicketOidcBadRequest struct {
}

func (o *CreateAccessTokenFromTicketOidcBadRequest) Error() string {
	return fmt.Sprintf("[POST /access/oidc/kerberos][%d] createAccessTokenFromTicketOidcBadRequest ", 400)
}

func (o *CreateAccessTokenFromTicketOidcBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenFromTicketOidcUnauthorized creates a CreateAccessTokenFromTicketOidcUnauthorized with default headers values
func NewCreateAccessTokenFromTicketOidcUnauthorized() *CreateAccessTokenFromTicketOidcUnauthorized {
	return &CreateAccessTokenFromTicketOidcUnauthorized{}
}

/*CreateAccessTokenFromTicketOidcUnauthorized handles this case with default header values.

NiFi was unable to complete the request because it did not contain a valid Kerberos ticket in the Authorization header. Retry this request after initializing a ticket with kinit and ensuring your browser is configured to support SPNEGO.
*/
type CreateAccessTokenFromTicketOidcUnauthorized struct {
}

func (o *CreateAccessTokenFromTicketOidcUnauthorized) Error() string {
	return fmt.Sprintf("[POST /access/oidc/kerberos][%d] createAccessTokenFromTicketOidcUnauthorized ", 401)
}

func (o *CreateAccessTokenFromTicketOidcUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenFromTicketOidcConflict creates a CreateAccessTokenFromTicketOidcConflict with default headers values
func NewCreateAccessTokenFromTicketOidcConflict() *CreateAccessTokenFromTicketOidcConflict {
	return &CreateAccessTokenFromTicketOidcConflict{}
}

/*CreateAccessTokenFromTicketOidcConflict handles this case with default header values.

Unable to create access token because NiFi is not in the appropriate state. (i.e. may not be configured to support Kerberos login.
*/
type CreateAccessTokenFromTicketOidcConflict struct {
}

func (o *CreateAccessTokenFromTicketOidcConflict) Error() string {
	return fmt.Sprintf("[POST /access/oidc/kerberos][%d] createAccessTokenFromTicketOidcConflict ", 409)
}

func (o *CreateAccessTokenFromTicketOidcConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenFromTicketOidcInternalServerError creates a CreateAccessTokenFromTicketOidcInternalServerError with default headers values
func NewCreateAccessTokenFromTicketOidcInternalServerError() *CreateAccessTokenFromTicketOidcInternalServerError {
	return &CreateAccessTokenFromTicketOidcInternalServerError{}
}

/*CreateAccessTokenFromTicketOidcInternalServerError handles this case with default header values.

Unable to create access token because an unexpected error occurred.
*/
type CreateAccessTokenFromTicketOidcInternalServerError struct {
}

func (o *CreateAccessTokenFromTicketOidcInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/oidc/kerberos][%d] createAccessTokenFromTicketOidcInternalServerError ", 500)
}

func (o *CreateAccessTokenFromTicketOidcInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
