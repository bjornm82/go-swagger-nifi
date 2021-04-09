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

// GetFlowsReader is a Reader for the GetFlows structure.
type GetFlowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFlowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFlowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFlowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFlowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetFlowsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowsOK creates a GetFlowsOK with default headers values
func NewGetFlowsOK() *GetFlowsOK {
	return &GetFlowsOK{}
}

/*GetFlowsOK handles this case with default header values.

successful operation
*/
type GetFlowsOK struct {
	Payload *models.VersionedFlowsEntity
}

func (o *GetFlowsOK) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows][%d] getFlowsOK  %+v", 200, o.Payload)
}

func (o *GetFlowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsBadRequest creates a GetFlowsBadRequest with default headers values
func NewGetFlowsBadRequest() *GetFlowsBadRequest {
	return &GetFlowsBadRequest{}
}

/*GetFlowsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFlowsBadRequest struct {
}

func (o *GetFlowsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows][%d] getFlowsBadRequest ", 400)
}

func (o *GetFlowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowsUnauthorized creates a GetFlowsUnauthorized with default headers values
func NewGetFlowsUnauthorized() *GetFlowsUnauthorized {
	return &GetFlowsUnauthorized{}
}

/*GetFlowsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetFlowsUnauthorized struct {
}

func (o *GetFlowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows][%d] getFlowsUnauthorized ", 401)
}

func (o *GetFlowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowsForbidden creates a GetFlowsForbidden with default headers values
func NewGetFlowsForbidden() *GetFlowsForbidden {
	return &GetFlowsForbidden{}
}

/*GetFlowsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetFlowsForbidden struct {
}

func (o *GetFlowsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows][%d] getFlowsForbidden ", 403)
}

func (o *GetFlowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowsNotFound creates a GetFlowsNotFound with default headers values
func NewGetFlowsNotFound() *GetFlowsNotFound {
	return &GetFlowsNotFound{}
}

/*GetFlowsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetFlowsNotFound struct {
}

func (o *GetFlowsNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows][%d] getFlowsNotFound ", 404)
}

func (o *GetFlowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowsConflict creates a GetFlowsConflict with default headers values
func NewGetFlowsConflict() *GetFlowsConflict {
	return &GetFlowsConflict{}
}

/*GetFlowsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetFlowsConflict struct {
}

func (o *GetFlowsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows][%d] getFlowsConflict ", 409)
}

func (o *GetFlowsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
