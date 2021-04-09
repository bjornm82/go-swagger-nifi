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

// GetRemoteProcessGroupStatusHistoryReader is a Reader for the GetRemoteProcessGroupStatusHistory structure.
type GetRemoteProcessGroupStatusHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteProcessGroupStatusHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRemoteProcessGroupStatusHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRemoteProcessGroupStatusHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRemoteProcessGroupStatusHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRemoteProcessGroupStatusHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRemoteProcessGroupStatusHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetRemoteProcessGroupStatusHistoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRemoteProcessGroupStatusHistoryOK creates a GetRemoteProcessGroupStatusHistoryOK with default headers values
func NewGetRemoteProcessGroupStatusHistoryOK() *GetRemoteProcessGroupStatusHistoryOK {
	return &GetRemoteProcessGroupStatusHistoryOK{}
}

/*GetRemoteProcessGroupStatusHistoryOK handles this case with default header values.

successful operation
*/
type GetRemoteProcessGroupStatusHistoryOK struct {
	Payload *models.StatusHistoryEntity
}

func (o *GetRemoteProcessGroupStatusHistoryOK) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryOK  %+v", 200, o.Payload)
}

func (o *GetRemoteProcessGroupStatusHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusHistoryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryBadRequest creates a GetRemoteProcessGroupStatusHistoryBadRequest with default headers values
func NewGetRemoteProcessGroupStatusHistoryBadRequest() *GetRemoteProcessGroupStatusHistoryBadRequest {
	return &GetRemoteProcessGroupStatusHistoryBadRequest{}
}

/*GetRemoteProcessGroupStatusHistoryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRemoteProcessGroupStatusHistoryBadRequest struct {
}

func (o *GetRemoteProcessGroupStatusHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryBadRequest ", 400)
}

func (o *GetRemoteProcessGroupStatusHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryUnauthorized creates a GetRemoteProcessGroupStatusHistoryUnauthorized with default headers values
func NewGetRemoteProcessGroupStatusHistoryUnauthorized() *GetRemoteProcessGroupStatusHistoryUnauthorized {
	return &GetRemoteProcessGroupStatusHistoryUnauthorized{}
}

/*GetRemoteProcessGroupStatusHistoryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetRemoteProcessGroupStatusHistoryUnauthorized struct {
}

func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryUnauthorized ", 401)
}

func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryForbidden creates a GetRemoteProcessGroupStatusHistoryForbidden with default headers values
func NewGetRemoteProcessGroupStatusHistoryForbidden() *GetRemoteProcessGroupStatusHistoryForbidden {
	return &GetRemoteProcessGroupStatusHistoryForbidden{}
}

/*GetRemoteProcessGroupStatusHistoryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetRemoteProcessGroupStatusHistoryForbidden struct {
}

func (o *GetRemoteProcessGroupStatusHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryForbidden ", 403)
}

func (o *GetRemoteProcessGroupStatusHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryNotFound creates a GetRemoteProcessGroupStatusHistoryNotFound with default headers values
func NewGetRemoteProcessGroupStatusHistoryNotFound() *GetRemoteProcessGroupStatusHistoryNotFound {
	return &GetRemoteProcessGroupStatusHistoryNotFound{}
}

/*GetRemoteProcessGroupStatusHistoryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetRemoteProcessGroupStatusHistoryNotFound struct {
}

func (o *GetRemoteProcessGroupStatusHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryNotFound ", 404)
}

func (o *GetRemoteProcessGroupStatusHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryConflict creates a GetRemoteProcessGroupStatusHistoryConflict with default headers values
func NewGetRemoteProcessGroupStatusHistoryConflict() *GetRemoteProcessGroupStatusHistoryConflict {
	return &GetRemoteProcessGroupStatusHistoryConflict{}
}

/*GetRemoteProcessGroupStatusHistoryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRemoteProcessGroupStatusHistoryConflict struct {
}

func (o *GetRemoteProcessGroupStatusHistoryConflict) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryConflict ", 409)
}

func (o *GetRemoteProcessGroupStatusHistoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}