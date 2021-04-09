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

// GetInputPortStatusReader is a Reader for the GetInputPortStatus structure.
type GetInputPortStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInputPortStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInputPortStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInputPortStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInputPortStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInputPortStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInputPortStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetInputPortStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInputPortStatusOK creates a GetInputPortStatusOK with default headers values
func NewGetInputPortStatusOK() *GetInputPortStatusOK {
	return &GetInputPortStatusOK{}
}

/*GetInputPortStatusOK handles this case with default header values.

successful operation
*/
type GetInputPortStatusOK struct {
	Payload *models.PortStatusEntity
}

func (o *GetInputPortStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusOK  %+v", 200, o.Payload)
}

func (o *GetInputPortStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInputPortStatusBadRequest creates a GetInputPortStatusBadRequest with default headers values
func NewGetInputPortStatusBadRequest() *GetInputPortStatusBadRequest {
	return &GetInputPortStatusBadRequest{}
}

/*GetInputPortStatusBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetInputPortStatusBadRequest struct {
}

func (o *GetInputPortStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusBadRequest ", 400)
}

func (o *GetInputPortStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusUnauthorized creates a GetInputPortStatusUnauthorized with default headers values
func NewGetInputPortStatusUnauthorized() *GetInputPortStatusUnauthorized {
	return &GetInputPortStatusUnauthorized{}
}

/*GetInputPortStatusUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetInputPortStatusUnauthorized struct {
}

func (o *GetInputPortStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusUnauthorized ", 401)
}

func (o *GetInputPortStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusForbidden creates a GetInputPortStatusForbidden with default headers values
func NewGetInputPortStatusForbidden() *GetInputPortStatusForbidden {
	return &GetInputPortStatusForbidden{}
}

/*GetInputPortStatusForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetInputPortStatusForbidden struct {
}

func (o *GetInputPortStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusForbidden ", 403)
}

func (o *GetInputPortStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusNotFound creates a GetInputPortStatusNotFound with default headers values
func NewGetInputPortStatusNotFound() *GetInputPortStatusNotFound {
	return &GetInputPortStatusNotFound{}
}

/*GetInputPortStatusNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetInputPortStatusNotFound struct {
}

func (o *GetInputPortStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusNotFound ", 404)
}

func (o *GetInputPortStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusConflict creates a GetInputPortStatusConflict with default headers values
func NewGetInputPortStatusConflict() *GetInputPortStatusConflict {
	return &GetInputPortStatusConflict{}
}

/*GetInputPortStatusConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetInputPortStatusConflict struct {
}

func (o *GetInputPortStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusConflict ", 409)
}

func (o *GetInputPortStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
