// Code generated by go-swagger; DO NOT EDIT.

package accessoidc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetAccessStatusOidcReader is a Reader for the GetAccessStatusOidc structure.
type GetAccessStatusOidcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessStatusOidcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessStatusOidcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccessStatusOidcBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccessStatusOidcUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessStatusOidcForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAccessStatusOidcConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccessStatusOidcInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccessStatusOidcOK creates a GetAccessStatusOidcOK with default headers values
func NewGetAccessStatusOidcOK() *GetAccessStatusOidcOK {
	return &GetAccessStatusOidcOK{}
}

/*GetAccessStatusOidcOK handles this case with default header values.

successful operation
*/
type GetAccessStatusOidcOK struct {
	Payload *models.AccessStatusEntity
}

func (o *GetAccessStatusOidcOK) Error() string {
	return fmt.Sprintf("[GET /access/oidc][%d] getAccessStatusOidcOK  %+v", 200, o.Payload)
}

func (o *GetAccessStatusOidcOK) GetPayload() *models.AccessStatusEntity {
	return o.Payload
}

func (o *GetAccessStatusOidcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessStatusOidcBadRequest creates a GetAccessStatusOidcBadRequest with default headers values
func NewGetAccessStatusOidcBadRequest() *GetAccessStatusOidcBadRequest {
	return &GetAccessStatusOidcBadRequest{}
}

/*GetAccessStatusOidcBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetAccessStatusOidcBadRequest struct {
}

func (o *GetAccessStatusOidcBadRequest) Error() string {
	return fmt.Sprintf("[GET /access/oidc][%d] getAccessStatusOidcBadRequest ", 400)
}

func (o *GetAccessStatusOidcBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusOidcUnauthorized creates a GetAccessStatusOidcUnauthorized with default headers values
func NewGetAccessStatusOidcUnauthorized() *GetAccessStatusOidcUnauthorized {
	return &GetAccessStatusOidcUnauthorized{}
}

/*GetAccessStatusOidcUnauthorized handles this case with default header values.

Unable to determine access status because the client could not be authenticated.
*/
type GetAccessStatusOidcUnauthorized struct {
}

func (o *GetAccessStatusOidcUnauthorized) Error() string {
	return fmt.Sprintf("[GET /access/oidc][%d] getAccessStatusOidcUnauthorized ", 401)
}

func (o *GetAccessStatusOidcUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusOidcForbidden creates a GetAccessStatusOidcForbidden with default headers values
func NewGetAccessStatusOidcForbidden() *GetAccessStatusOidcForbidden {
	return &GetAccessStatusOidcForbidden{}
}

/*GetAccessStatusOidcForbidden handles this case with default header values.

Unable to determine access status because the client is not authorized to make this request.
*/
type GetAccessStatusOidcForbidden struct {
}

func (o *GetAccessStatusOidcForbidden) Error() string {
	return fmt.Sprintf("[GET /access/oidc][%d] getAccessStatusOidcForbidden ", 403)
}

func (o *GetAccessStatusOidcForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusOidcConflict creates a GetAccessStatusOidcConflict with default headers values
func NewGetAccessStatusOidcConflict() *GetAccessStatusOidcConflict {
	return &GetAccessStatusOidcConflict{}
}

/*GetAccessStatusOidcConflict handles this case with default header values.

Unable to determine access status because NiFi is not in the appropriate state.
*/
type GetAccessStatusOidcConflict struct {
}

func (o *GetAccessStatusOidcConflict) Error() string {
	return fmt.Sprintf("[GET /access/oidc][%d] getAccessStatusOidcConflict ", 409)
}

func (o *GetAccessStatusOidcConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusOidcInternalServerError creates a GetAccessStatusOidcInternalServerError with default headers values
func NewGetAccessStatusOidcInternalServerError() *GetAccessStatusOidcInternalServerError {
	return &GetAccessStatusOidcInternalServerError{}
}

/*GetAccessStatusOidcInternalServerError handles this case with default header values.

Unable to determine access status because an unexpected error occurred.
*/
type GetAccessStatusOidcInternalServerError struct {
}

func (o *GetAccessStatusOidcInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access/oidc][%d] getAccessStatusOidcInternalServerError ", 500)
}

func (o *GetAccessStatusOidcInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
