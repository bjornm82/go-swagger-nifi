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

// GetRemoteProcessGroupStatusReader is a Reader for the GetRemoteProcessGroupStatus structure.
type GetRemoteProcessGroupStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteProcessGroupStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRemoteProcessGroupStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRemoteProcessGroupStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRemoteProcessGroupStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRemoteProcessGroupStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRemoteProcessGroupStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetRemoteProcessGroupStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRemoteProcessGroupStatusOK creates a GetRemoteProcessGroupStatusOK with default headers values
func NewGetRemoteProcessGroupStatusOK() *GetRemoteProcessGroupStatusOK {
	return &GetRemoteProcessGroupStatusOK{}
}

/*GetRemoteProcessGroupStatusOK handles this case with default header values.

successful operation
*/
type GetRemoteProcessGroupStatusOK struct {
	Payload *models.RemoteProcessGroupStatusEntity
}

func (o *GetRemoteProcessGroupStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status][%d] getRemoteProcessGroupStatusOK  %+v", 200, o.Payload)
}

func (o *GetRemoteProcessGroupStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteProcessGroupStatusBadRequest creates a GetRemoteProcessGroupStatusBadRequest with default headers values
func NewGetRemoteProcessGroupStatusBadRequest() *GetRemoteProcessGroupStatusBadRequest {
	return &GetRemoteProcessGroupStatusBadRequest{}
}

/*GetRemoteProcessGroupStatusBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRemoteProcessGroupStatusBadRequest struct {
}

func (o *GetRemoteProcessGroupStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status][%d] getRemoteProcessGroupStatusBadRequest ", 400)
}

func (o *GetRemoteProcessGroupStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusUnauthorized creates a GetRemoteProcessGroupStatusUnauthorized with default headers values
func NewGetRemoteProcessGroupStatusUnauthorized() *GetRemoteProcessGroupStatusUnauthorized {
	return &GetRemoteProcessGroupStatusUnauthorized{}
}

/*GetRemoteProcessGroupStatusUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetRemoteProcessGroupStatusUnauthorized struct {
}

func (o *GetRemoteProcessGroupStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status][%d] getRemoteProcessGroupStatusUnauthorized ", 401)
}

func (o *GetRemoteProcessGroupStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusForbidden creates a GetRemoteProcessGroupStatusForbidden with default headers values
func NewGetRemoteProcessGroupStatusForbidden() *GetRemoteProcessGroupStatusForbidden {
	return &GetRemoteProcessGroupStatusForbidden{}
}

/*GetRemoteProcessGroupStatusForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetRemoteProcessGroupStatusForbidden struct {
}

func (o *GetRemoteProcessGroupStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status][%d] getRemoteProcessGroupStatusForbidden ", 403)
}

func (o *GetRemoteProcessGroupStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusNotFound creates a GetRemoteProcessGroupStatusNotFound with default headers values
func NewGetRemoteProcessGroupStatusNotFound() *GetRemoteProcessGroupStatusNotFound {
	return &GetRemoteProcessGroupStatusNotFound{}
}

/*GetRemoteProcessGroupStatusNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetRemoteProcessGroupStatusNotFound struct {
}

func (o *GetRemoteProcessGroupStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status][%d] getRemoteProcessGroupStatusNotFound ", 404)
}

func (o *GetRemoteProcessGroupStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusConflict creates a GetRemoteProcessGroupStatusConflict with default headers values
func NewGetRemoteProcessGroupStatusConflict() *GetRemoteProcessGroupStatusConflict {
	return &GetRemoteProcessGroupStatusConflict{}
}

/*GetRemoteProcessGroupStatusConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRemoteProcessGroupStatusConflict struct {
}

func (o *GetRemoteProcessGroupStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status][%d] getRemoteProcessGroupStatusConflict ", 409)
}

func (o *GetRemoteProcessGroupStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
