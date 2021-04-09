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

// GetProcessorStatusHistoryReader is a Reader for the GetProcessorStatusHistory structure.
type GetProcessorStatusHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorStatusHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProcessorStatusHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProcessorStatusHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetProcessorStatusHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetProcessorStatusHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProcessorStatusHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetProcessorStatusHistoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProcessorStatusHistoryOK creates a GetProcessorStatusHistoryOK with default headers values
func NewGetProcessorStatusHistoryOK() *GetProcessorStatusHistoryOK {
	return &GetProcessorStatusHistoryOK{}
}

/*GetProcessorStatusHistoryOK handles this case with default header values.

successful operation
*/
type GetProcessorStatusHistoryOK struct {
	Payload *models.StatusHistoryEntity
}

func (o *GetProcessorStatusHistoryOK) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status/history][%d] getProcessorStatusHistoryOK  %+v", 200, o.Payload)
}

func (o *GetProcessorStatusHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusHistoryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorStatusHistoryBadRequest creates a GetProcessorStatusHistoryBadRequest with default headers values
func NewGetProcessorStatusHistoryBadRequest() *GetProcessorStatusHistoryBadRequest {
	return &GetProcessorStatusHistoryBadRequest{}
}

/*GetProcessorStatusHistoryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessorStatusHistoryBadRequest struct {
}

func (o *GetProcessorStatusHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status/history][%d] getProcessorStatusHistoryBadRequest ", 400)
}

func (o *GetProcessorStatusHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusHistoryUnauthorized creates a GetProcessorStatusHistoryUnauthorized with default headers values
func NewGetProcessorStatusHistoryUnauthorized() *GetProcessorStatusHistoryUnauthorized {
	return &GetProcessorStatusHistoryUnauthorized{}
}

/*GetProcessorStatusHistoryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetProcessorStatusHistoryUnauthorized struct {
}

func (o *GetProcessorStatusHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status/history][%d] getProcessorStatusHistoryUnauthorized ", 401)
}

func (o *GetProcessorStatusHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusHistoryForbidden creates a GetProcessorStatusHistoryForbidden with default headers values
func NewGetProcessorStatusHistoryForbidden() *GetProcessorStatusHistoryForbidden {
	return &GetProcessorStatusHistoryForbidden{}
}

/*GetProcessorStatusHistoryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetProcessorStatusHistoryForbidden struct {
}

func (o *GetProcessorStatusHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status/history][%d] getProcessorStatusHistoryForbidden ", 403)
}

func (o *GetProcessorStatusHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusHistoryNotFound creates a GetProcessorStatusHistoryNotFound with default headers values
func NewGetProcessorStatusHistoryNotFound() *GetProcessorStatusHistoryNotFound {
	return &GetProcessorStatusHistoryNotFound{}
}

/*GetProcessorStatusHistoryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetProcessorStatusHistoryNotFound struct {
}

func (o *GetProcessorStatusHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status/history][%d] getProcessorStatusHistoryNotFound ", 404)
}

func (o *GetProcessorStatusHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorStatusHistoryConflict creates a GetProcessorStatusHistoryConflict with default headers values
func NewGetProcessorStatusHistoryConflict() *GetProcessorStatusHistoryConflict {
	return &GetProcessorStatusHistoryConflict{}
}

/*GetProcessorStatusHistoryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessorStatusHistoryConflict struct {
}

func (o *GetProcessorStatusHistoryConflict) Error() string {
	return fmt.Sprintf("[GET /flow/processors/{id}/status/history][%d] getProcessorStatusHistoryConflict ", 409)
}

func (o *GetProcessorStatusHistoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}