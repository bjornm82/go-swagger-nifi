// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// GetDropAllFlowfilesRequestReader is a Reader for the GetDropAllFlowfilesRequest structure.
type GetDropAllFlowfilesRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDropAllFlowfilesRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDropAllFlowfilesRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDropAllFlowfilesRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetDropAllFlowfilesRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDropAllFlowfilesRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDropAllFlowfilesRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetDropAllFlowfilesRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDropAllFlowfilesRequestOK creates a GetDropAllFlowfilesRequestOK with default headers values
func NewGetDropAllFlowfilesRequestOK() *GetDropAllFlowfilesRequestOK {
	return &GetDropAllFlowfilesRequestOK{}
}

/*GetDropAllFlowfilesRequestOK handles this case with default header values.

successful operation
*/
type GetDropAllFlowfilesRequestOK struct {
	Payload *models.DropRequestEntity
}

func (o *GetDropAllFlowfilesRequestOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/empty-all-connections-requests/{drop-request-id}][%d] getDropAllFlowfilesRequestOK  %+v", 200, o.Payload)
}

func (o *GetDropAllFlowfilesRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DropRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDropAllFlowfilesRequestBadRequest creates a GetDropAllFlowfilesRequestBadRequest with default headers values
func NewGetDropAllFlowfilesRequestBadRequest() *GetDropAllFlowfilesRequestBadRequest {
	return &GetDropAllFlowfilesRequestBadRequest{}
}

/*GetDropAllFlowfilesRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetDropAllFlowfilesRequestBadRequest struct {
}

func (o *GetDropAllFlowfilesRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/empty-all-connections-requests/{drop-request-id}][%d] getDropAllFlowfilesRequestBadRequest ", 400)
}

func (o *GetDropAllFlowfilesRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropAllFlowfilesRequestUnauthorized creates a GetDropAllFlowfilesRequestUnauthorized with default headers values
func NewGetDropAllFlowfilesRequestUnauthorized() *GetDropAllFlowfilesRequestUnauthorized {
	return &GetDropAllFlowfilesRequestUnauthorized{}
}

/*GetDropAllFlowfilesRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetDropAllFlowfilesRequestUnauthorized struct {
}

func (o *GetDropAllFlowfilesRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/empty-all-connections-requests/{drop-request-id}][%d] getDropAllFlowfilesRequestUnauthorized ", 401)
}

func (o *GetDropAllFlowfilesRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropAllFlowfilesRequestForbidden creates a GetDropAllFlowfilesRequestForbidden with default headers values
func NewGetDropAllFlowfilesRequestForbidden() *GetDropAllFlowfilesRequestForbidden {
	return &GetDropAllFlowfilesRequestForbidden{}
}

/*GetDropAllFlowfilesRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetDropAllFlowfilesRequestForbidden struct {
}

func (o *GetDropAllFlowfilesRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/empty-all-connections-requests/{drop-request-id}][%d] getDropAllFlowfilesRequestForbidden ", 403)
}

func (o *GetDropAllFlowfilesRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropAllFlowfilesRequestNotFound creates a GetDropAllFlowfilesRequestNotFound with default headers values
func NewGetDropAllFlowfilesRequestNotFound() *GetDropAllFlowfilesRequestNotFound {
	return &GetDropAllFlowfilesRequestNotFound{}
}

/*GetDropAllFlowfilesRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetDropAllFlowfilesRequestNotFound struct {
}

func (o *GetDropAllFlowfilesRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/empty-all-connections-requests/{drop-request-id}][%d] getDropAllFlowfilesRequestNotFound ", 404)
}

func (o *GetDropAllFlowfilesRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropAllFlowfilesRequestConflict creates a GetDropAllFlowfilesRequestConflict with default headers values
func NewGetDropAllFlowfilesRequestConflict() *GetDropAllFlowfilesRequestConflict {
	return &GetDropAllFlowfilesRequestConflict{}
}

/*GetDropAllFlowfilesRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetDropAllFlowfilesRequestConflict struct {
}

func (o *GetDropAllFlowfilesRequestConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/empty-all-connections-requests/{drop-request-id}][%d] getDropAllFlowfilesRequestConflict ", 409)
}

func (o *GetDropAllFlowfilesRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
