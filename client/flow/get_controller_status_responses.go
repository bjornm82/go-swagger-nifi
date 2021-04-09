// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// GetControllerStatusReader is a Reader for the GetControllerStatus structure.
type GetControllerStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControllerStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetControllerStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetControllerStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetControllerStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetControllerStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetControllerStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetControllerStatusOK creates a GetControllerStatusOK with default headers values
func NewGetControllerStatusOK() *GetControllerStatusOK {
	return &GetControllerStatusOK{}
}

/*GetControllerStatusOK handles this case with default header values.

successful operation
*/
type GetControllerStatusOK struct {
	Payload *models.ControllerStatusEntity
}

func (o *GetControllerStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/status][%d] getControllerStatusOK  %+v", 200, o.Payload)
}

func (o *GetControllerStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControllerStatusBadRequest creates a GetControllerStatusBadRequest with default headers values
func NewGetControllerStatusBadRequest() *GetControllerStatusBadRequest {
	return &GetControllerStatusBadRequest{}
}

/*GetControllerStatusBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetControllerStatusBadRequest struct {
}

func (o *GetControllerStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/status][%d] getControllerStatusBadRequest ", 400)
}

func (o *GetControllerStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerStatusUnauthorized creates a GetControllerStatusUnauthorized with default headers values
func NewGetControllerStatusUnauthorized() *GetControllerStatusUnauthorized {
	return &GetControllerStatusUnauthorized{}
}

/*GetControllerStatusUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetControllerStatusUnauthorized struct {
}

func (o *GetControllerStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/status][%d] getControllerStatusUnauthorized ", 401)
}

func (o *GetControllerStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerStatusForbidden creates a GetControllerStatusForbidden with default headers values
func NewGetControllerStatusForbidden() *GetControllerStatusForbidden {
	return &GetControllerStatusForbidden{}
}

/*GetControllerStatusForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetControllerStatusForbidden struct {
}

func (o *GetControllerStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/status][%d] getControllerStatusForbidden ", 403)
}

func (o *GetControllerStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerStatusConflict creates a GetControllerStatusConflict with default headers values
func NewGetControllerStatusConflict() *GetControllerStatusConflict {
	return &GetControllerStatusConflict{}
}

/*GetControllerStatusConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetControllerStatusConflict struct {
}

func (o *GetControllerStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/status][%d] getControllerStatusConflict ", 409)
}

func (o *GetControllerStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
