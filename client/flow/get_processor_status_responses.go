// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// GetProcessorStatusReader is a Reader for the GetProcessorStatus structure.
type GetProcessorStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProcessorStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProcessorStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetProcessorStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetProcessorStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProcessorStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetProcessorStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProcessorStatusOK creates a GetProcessorStatusOK with default headers values
func NewGetProcessorStatusOK() *GetProcessorStatusOK {
	return &GetProcessorStatusOK{}
}

/*GetProcessorStatusOK handles this case with default header values.

successful operation
*/
type GetProcessorStatusOK struct {
	Payload *models.ProcessorStatusEntity
}

func (o *GetProcessorStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status][%d] getProcessorStatusOK  %+v", 200, o.Payload)
}

func (o *GetProcessorStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorStatusBadRequest creates a GetProcessorStatusBadRequest with default headers values
func NewGetProcessorStatusBadRequest() *GetProcessorStatusBadRequest {
	return &GetProcessorStatusBadRequest{}
}

/*GetProcessorStatusBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessorStatusBadRequest struct {
}

func (o *GetProcessorStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status][%d] getProcessorStatusBadRequest ", 400)
}

func (o *GetProcessorStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusUnauthorized creates a GetProcessorStatusUnauthorized with default headers values
func NewGetProcessorStatusUnauthorized() *GetProcessorStatusUnauthorized {
	return &GetProcessorStatusUnauthorized{}
}

/*GetProcessorStatusUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetProcessorStatusUnauthorized struct {
}

func (o *GetProcessorStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status][%d] getProcessorStatusUnauthorized ", 401)
}

func (o *GetProcessorStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusForbidden creates a GetProcessorStatusForbidden with default headers values
func NewGetProcessorStatusForbidden() *GetProcessorStatusForbidden {
	return &GetProcessorStatusForbidden{}
}

/*GetProcessorStatusForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetProcessorStatusForbidden struct {
}

func (o *GetProcessorStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status][%d] getProcessorStatusForbidden ", 403)
}

func (o *GetProcessorStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusNotFound creates a GetProcessorStatusNotFound with default headers values
func NewGetProcessorStatusNotFound() *GetProcessorStatusNotFound {
	return &GetProcessorStatusNotFound{}
}

/*GetProcessorStatusNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetProcessorStatusNotFound struct {
}

func (o *GetProcessorStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status][%d] getProcessorStatusNotFound ", 404)
}

func (o *GetProcessorStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusConflict creates a GetProcessorStatusConflict with default headers values
func NewGetProcessorStatusConflict() *GetProcessorStatusConflict {
	return &GetProcessorStatusConflict{}
}

/*GetProcessorStatusConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessorStatusConflict struct {
}

func (o *GetProcessorStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status][%d] getProcessorStatusConflict ", 409)
}

func (o *GetProcessorStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
