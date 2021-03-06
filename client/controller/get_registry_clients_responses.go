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

// GetRegistryClientsReader is a Reader for the GetRegistryClients structure.
type GetRegistryClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegistryClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegistryClientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistryClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistryClientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRegistryClientsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRegistryClientsOK creates a GetRegistryClientsOK with default headers values
func NewGetRegistryClientsOK() *GetRegistryClientsOK {
	return &GetRegistryClientsOK{}
}

/*GetRegistryClientsOK handles this case with default header values.

successful operation
*/
type GetRegistryClientsOK struct {
	Payload *models.RegistryClientsEntity
}

func (o *GetRegistryClientsOK) Error() string {
	return fmt.Sprintf("[GET /controller/registry-clients][%d] getRegistryClientsOK  %+v", 200, o.Payload)
}

func (o *GetRegistryClientsOK) GetPayload() *models.RegistryClientsEntity {
	return o.Payload
}

func (o *GetRegistryClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistryClientsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryClientsBadRequest creates a GetRegistryClientsBadRequest with default headers values
func NewGetRegistryClientsBadRequest() *GetRegistryClientsBadRequest {
	return &GetRegistryClientsBadRequest{}
}

/*GetRegistryClientsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRegistryClientsBadRequest struct {
}

func (o *GetRegistryClientsBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller/registry-clients][%d] getRegistryClientsBadRequest ", 400)
}

func (o *GetRegistryClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientsUnauthorized creates a GetRegistryClientsUnauthorized with default headers values
func NewGetRegistryClientsUnauthorized() *GetRegistryClientsUnauthorized {
	return &GetRegistryClientsUnauthorized{}
}

/*GetRegistryClientsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetRegistryClientsUnauthorized struct {
}

func (o *GetRegistryClientsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller/registry-clients][%d] getRegistryClientsUnauthorized ", 401)
}

func (o *GetRegistryClientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientsForbidden creates a GetRegistryClientsForbidden with default headers values
func NewGetRegistryClientsForbidden() *GetRegistryClientsForbidden {
	return &GetRegistryClientsForbidden{}
}

/*GetRegistryClientsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetRegistryClientsForbidden struct {
}

func (o *GetRegistryClientsForbidden) Error() string {
	return fmt.Sprintf("[GET /controller/registry-clients][%d] getRegistryClientsForbidden ", 403)
}

func (o *GetRegistryClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientsNotFound creates a GetRegistryClientsNotFound with default headers values
func NewGetRegistryClientsNotFound() *GetRegistryClientsNotFound {
	return &GetRegistryClientsNotFound{}
}

/*GetRegistryClientsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetRegistryClientsNotFound struct {
}

func (o *GetRegistryClientsNotFound) Error() string {
	return fmt.Sprintf("[GET /controller/registry-clients][%d] getRegistryClientsNotFound ", 404)
}

func (o *GetRegistryClientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientsConflict creates a GetRegistryClientsConflict with default headers values
func NewGetRegistryClientsConflict() *GetRegistryClientsConflict {
	return &GetRegistryClientsConflict{}
}

/*GetRegistryClientsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRegistryClientsConflict struct {
}

func (o *GetRegistryClientsConflict) Error() string {
	return fmt.Sprintf("[GET /controller/registry-clients][%d] getRegistryClientsConflict ", 409)
}

func (o *GetRegistryClientsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
