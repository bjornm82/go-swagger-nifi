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

// GetConnectionStatusHistoryReader is a Reader for the GetConnectionStatusHistory structure.
type GetConnectionStatusHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConnectionStatusHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConnectionStatusHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetConnectionStatusHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetConnectionStatusHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetConnectionStatusHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetConnectionStatusHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetConnectionStatusHistoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConnectionStatusHistoryOK creates a GetConnectionStatusHistoryOK with default headers values
func NewGetConnectionStatusHistoryOK() *GetConnectionStatusHistoryOK {
	return &GetConnectionStatusHistoryOK{}
}

/*GetConnectionStatusHistoryOK handles this case with default header values.

successful operation
*/
type GetConnectionStatusHistoryOK struct {
	Payload *models.StatusHistoryEntity
}

func (o *GetConnectionStatusHistoryOK) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/status/history][%d] getConnectionStatusHistoryOK  %+v", 200, o.Payload)
}

func (o *GetConnectionStatusHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusHistoryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionStatusHistoryBadRequest creates a GetConnectionStatusHistoryBadRequest with default headers values
func NewGetConnectionStatusHistoryBadRequest() *GetConnectionStatusHistoryBadRequest {
	return &GetConnectionStatusHistoryBadRequest{}
}

/*GetConnectionStatusHistoryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetConnectionStatusHistoryBadRequest struct {
}

func (o *GetConnectionStatusHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/status/history][%d] getConnectionStatusHistoryBadRequest ", 400)
}

func (o *GetConnectionStatusHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatusHistoryUnauthorized creates a GetConnectionStatusHistoryUnauthorized with default headers values
func NewGetConnectionStatusHistoryUnauthorized() *GetConnectionStatusHistoryUnauthorized {
	return &GetConnectionStatusHistoryUnauthorized{}
}

/*GetConnectionStatusHistoryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetConnectionStatusHistoryUnauthorized struct {
}

func (o *GetConnectionStatusHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/status/history][%d] getConnectionStatusHistoryUnauthorized ", 401)
}

func (o *GetConnectionStatusHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatusHistoryForbidden creates a GetConnectionStatusHistoryForbidden with default headers values
func NewGetConnectionStatusHistoryForbidden() *GetConnectionStatusHistoryForbidden {
	return &GetConnectionStatusHistoryForbidden{}
}

/*GetConnectionStatusHistoryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetConnectionStatusHistoryForbidden struct {
}

func (o *GetConnectionStatusHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/status/history][%d] getConnectionStatusHistoryForbidden ", 403)
}

func (o *GetConnectionStatusHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatusHistoryNotFound creates a GetConnectionStatusHistoryNotFound with default headers values
func NewGetConnectionStatusHistoryNotFound() *GetConnectionStatusHistoryNotFound {
	return &GetConnectionStatusHistoryNotFound{}
}

/*GetConnectionStatusHistoryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetConnectionStatusHistoryNotFound struct {
}

func (o *GetConnectionStatusHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/status/history][%d] getConnectionStatusHistoryNotFound ", 404)
}

func (o *GetConnectionStatusHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatusHistoryConflict creates a GetConnectionStatusHistoryConflict with default headers values
func NewGetConnectionStatusHistoryConflict() *GetConnectionStatusHistoryConflict {
	return &GetConnectionStatusHistoryConflict{}
}

/*GetConnectionStatusHistoryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetConnectionStatusHistoryConflict struct {
}

func (o *GetConnectionStatusHistoryConflict) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/status/history][%d] getConnectionStatusHistoryConflict ", 409)
}

func (o *GetConnectionStatusHistoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
