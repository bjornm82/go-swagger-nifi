// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetNodeStatusHistoryReader is a Reader for the GetNodeStatusHistory structure.
type GetNodeStatusHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeStatusHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeStatusHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNodeStatusHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNodeStatusHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeStatusHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNodeStatusHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetNodeStatusHistoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNodeStatusHistoryOK creates a GetNodeStatusHistoryOK with default headers values
func NewGetNodeStatusHistoryOK() *GetNodeStatusHistoryOK {
	return &GetNodeStatusHistoryOK{}
}

/*GetNodeStatusHistoryOK handles this case with default header values.

successful operation
*/
type GetNodeStatusHistoryOK struct {
	Payload *models.ComponentHistoryEntity
}

func (o *GetNodeStatusHistoryOK) Error() string {
	return fmt.Sprintf("[GET /controller/status/history][%d] getNodeStatusHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNodeStatusHistoryOK) GetPayload() *models.ComponentHistoryEntity {
	return o.Payload
}

func (o *GetNodeStatusHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentHistoryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeStatusHistoryBadRequest creates a GetNodeStatusHistoryBadRequest with default headers values
func NewGetNodeStatusHistoryBadRequest() *GetNodeStatusHistoryBadRequest {
	return &GetNodeStatusHistoryBadRequest{}
}

/*GetNodeStatusHistoryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetNodeStatusHistoryBadRequest struct {
}

func (o *GetNodeStatusHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller/status/history][%d] getNodeStatusHistoryBadRequest ", 400)
}

func (o *GetNodeStatusHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeStatusHistoryUnauthorized creates a GetNodeStatusHistoryUnauthorized with default headers values
func NewGetNodeStatusHistoryUnauthorized() *GetNodeStatusHistoryUnauthorized {
	return &GetNodeStatusHistoryUnauthorized{}
}

/*GetNodeStatusHistoryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetNodeStatusHistoryUnauthorized struct {
}

func (o *GetNodeStatusHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller/status/history][%d] getNodeStatusHistoryUnauthorized ", 401)
}

func (o *GetNodeStatusHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeStatusHistoryForbidden creates a GetNodeStatusHistoryForbidden with default headers values
func NewGetNodeStatusHistoryForbidden() *GetNodeStatusHistoryForbidden {
	return &GetNodeStatusHistoryForbidden{}
}

/*GetNodeStatusHistoryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetNodeStatusHistoryForbidden struct {
}

func (o *GetNodeStatusHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /controller/status/history][%d] getNodeStatusHistoryForbidden ", 403)
}

func (o *GetNodeStatusHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeStatusHistoryNotFound creates a GetNodeStatusHistoryNotFound with default headers values
func NewGetNodeStatusHistoryNotFound() *GetNodeStatusHistoryNotFound {
	return &GetNodeStatusHistoryNotFound{}
}

/*GetNodeStatusHistoryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetNodeStatusHistoryNotFound struct {
}

func (o *GetNodeStatusHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /controller/status/history][%d] getNodeStatusHistoryNotFound ", 404)
}

func (o *GetNodeStatusHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeStatusHistoryConflict creates a GetNodeStatusHistoryConflict with default headers values
func NewGetNodeStatusHistoryConflict() *GetNodeStatusHistoryConflict {
	return &GetNodeStatusHistoryConflict{}
}

/*GetNodeStatusHistoryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetNodeStatusHistoryConflict struct {
}

func (o *GetNodeStatusHistoryConflict) Error() string {
	return fmt.Sprintf("[GET /controller/status/history][%d] getNodeStatusHistoryConflict ", 409)
}

func (o *GetNodeStatusHistoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
