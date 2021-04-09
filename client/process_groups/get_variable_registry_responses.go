// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// GetVariableRegistryReader is a Reader for the GetVariableRegistry structure.
type GetVariableRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVariableRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVariableRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetVariableRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetVariableRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetVariableRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetVariableRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetVariableRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVariableRegistryOK creates a GetVariableRegistryOK with default headers values
func NewGetVariableRegistryOK() *GetVariableRegistryOK {
	return &GetVariableRegistryOK{}
}

/*GetVariableRegistryOK handles this case with default header values.

successful operation
*/
type GetVariableRegistryOK struct {
	Payload *models.VariableRegistryEntity
}

func (o *GetVariableRegistryOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/variable-registry][%d] getVariableRegistryOK  %+v", 200, o.Payload)
}

func (o *GetVariableRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableRegistryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVariableRegistryBadRequest creates a GetVariableRegistryBadRequest with default headers values
func NewGetVariableRegistryBadRequest() *GetVariableRegistryBadRequest {
	return &GetVariableRegistryBadRequest{}
}

/*GetVariableRegistryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetVariableRegistryBadRequest struct {
}

func (o *GetVariableRegistryBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/variable-registry][%d] getVariableRegistryBadRequest ", 400)
}

func (o *GetVariableRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryUnauthorized creates a GetVariableRegistryUnauthorized with default headers values
func NewGetVariableRegistryUnauthorized() *GetVariableRegistryUnauthorized {
	return &GetVariableRegistryUnauthorized{}
}

/*GetVariableRegistryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetVariableRegistryUnauthorized struct {
}

func (o *GetVariableRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/variable-registry][%d] getVariableRegistryUnauthorized ", 401)
}

func (o *GetVariableRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryForbidden creates a GetVariableRegistryForbidden with default headers values
func NewGetVariableRegistryForbidden() *GetVariableRegistryForbidden {
	return &GetVariableRegistryForbidden{}
}

/*GetVariableRegistryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetVariableRegistryForbidden struct {
}

func (o *GetVariableRegistryForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/variable-registry][%d] getVariableRegistryForbidden ", 403)
}

func (o *GetVariableRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryNotFound creates a GetVariableRegistryNotFound with default headers values
func NewGetVariableRegistryNotFound() *GetVariableRegistryNotFound {
	return &GetVariableRegistryNotFound{}
}

/*GetVariableRegistryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetVariableRegistryNotFound struct {
}

func (o *GetVariableRegistryNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/variable-registry][%d] getVariableRegistryNotFound ", 404)
}

func (o *GetVariableRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryConflict creates a GetVariableRegistryConflict with default headers values
func NewGetVariableRegistryConflict() *GetVariableRegistryConflict {
	return &GetVariableRegistryConflict{}
}

/*GetVariableRegistryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetVariableRegistryConflict struct {
}

func (o *GetVariableRegistryConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/variable-registry][%d] getVariableRegistryConflict ", 409)
}

func (o *GetVariableRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
