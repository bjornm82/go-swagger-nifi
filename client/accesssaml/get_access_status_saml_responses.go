// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetAccessStatusSamlReader is a Reader for the GetAccessStatusSaml structure.
type GetAccessStatusSamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessStatusSamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessStatusSamlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccessStatusSamlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccessStatusSamlUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessStatusSamlForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAccessStatusSamlConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccessStatusSamlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccessStatusSamlOK creates a GetAccessStatusSamlOK with default headers values
func NewGetAccessStatusSamlOK() *GetAccessStatusSamlOK {
	return &GetAccessStatusSamlOK{}
}

/*GetAccessStatusSamlOK handles this case with default header values.

successful operation
*/
type GetAccessStatusSamlOK struct {
	Payload *models.AccessStatusEntity
}

func (o *GetAccessStatusSamlOK) Error() string {
	return fmt.Sprintf("[GET /access/saml][%d] getAccessStatusSamlOK  %+v", 200, o.Payload)
}

func (o *GetAccessStatusSamlOK) GetPayload() *models.AccessStatusEntity {
	return o.Payload
}

func (o *GetAccessStatusSamlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessStatusSamlBadRequest creates a GetAccessStatusSamlBadRequest with default headers values
func NewGetAccessStatusSamlBadRequest() *GetAccessStatusSamlBadRequest {
	return &GetAccessStatusSamlBadRequest{}
}

/*GetAccessStatusSamlBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetAccessStatusSamlBadRequest struct {
}

func (o *GetAccessStatusSamlBadRequest) Error() string {
	return fmt.Sprintf("[GET /access/saml][%d] getAccessStatusSamlBadRequest ", 400)
}

func (o *GetAccessStatusSamlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusSamlUnauthorized creates a GetAccessStatusSamlUnauthorized with default headers values
func NewGetAccessStatusSamlUnauthorized() *GetAccessStatusSamlUnauthorized {
	return &GetAccessStatusSamlUnauthorized{}
}

/*GetAccessStatusSamlUnauthorized handles this case with default header values.

Unable to determine access status because the client could not be authenticated.
*/
type GetAccessStatusSamlUnauthorized struct {
}

func (o *GetAccessStatusSamlUnauthorized) Error() string {
	return fmt.Sprintf("[GET /access/saml][%d] getAccessStatusSamlUnauthorized ", 401)
}

func (o *GetAccessStatusSamlUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusSamlForbidden creates a GetAccessStatusSamlForbidden with default headers values
func NewGetAccessStatusSamlForbidden() *GetAccessStatusSamlForbidden {
	return &GetAccessStatusSamlForbidden{}
}

/*GetAccessStatusSamlForbidden handles this case with default header values.

Unable to determine access status because the client is not authorized to make this request.
*/
type GetAccessStatusSamlForbidden struct {
}

func (o *GetAccessStatusSamlForbidden) Error() string {
	return fmt.Sprintf("[GET /access/saml][%d] getAccessStatusSamlForbidden ", 403)
}

func (o *GetAccessStatusSamlForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusSamlConflict creates a GetAccessStatusSamlConflict with default headers values
func NewGetAccessStatusSamlConflict() *GetAccessStatusSamlConflict {
	return &GetAccessStatusSamlConflict{}
}

/*GetAccessStatusSamlConflict handles this case with default header values.

Unable to determine access status because NiFi is not in the appropriate state.
*/
type GetAccessStatusSamlConflict struct {
}

func (o *GetAccessStatusSamlConflict) Error() string {
	return fmt.Sprintf("[GET /access/saml][%d] getAccessStatusSamlConflict ", 409)
}

func (o *GetAccessStatusSamlConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessStatusSamlInternalServerError creates a GetAccessStatusSamlInternalServerError with default headers values
func NewGetAccessStatusSamlInternalServerError() *GetAccessStatusSamlInternalServerError {
	return &GetAccessStatusSamlInternalServerError{}
}

/*GetAccessStatusSamlInternalServerError handles this case with default header values.

Unable to determine access status because an unexpected error occurred.
*/
type GetAccessStatusSamlInternalServerError struct {
}

func (o *GetAccessStatusSamlInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access/saml][%d] getAccessStatusSamlInternalServerError ", 500)
}

func (o *GetAccessStatusSamlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}