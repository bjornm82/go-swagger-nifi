// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// GetValidationRequestReader is a Reader for the GetValidationRequest structure.
type GetValidationRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetValidationRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetValidationRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetValidationRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetValidationRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetValidationRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetValidationRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetValidationRequestOK creates a GetValidationRequestOK with default headers values
func NewGetValidationRequestOK() *GetValidationRequestOK {
	return &GetValidationRequestOK{}
}

/*GetValidationRequestOK handles this case with default header values.

successful operation
*/
type GetValidationRequestOK struct {
	Payload *models.ParameterContextValidationRequestEntity
}

func (o *GetValidationRequestOK) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/validation-requests/{id}][%d] getValidationRequestOK  %+v", 200, o.Payload)
}

func (o *GetValidationRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextValidationRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationRequestBadRequest creates a GetValidationRequestBadRequest with default headers values
func NewGetValidationRequestBadRequest() *GetValidationRequestBadRequest {
	return &GetValidationRequestBadRequest{}
}

/*GetValidationRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetValidationRequestBadRequest struct {
}

func (o *GetValidationRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/validation-requests/{id}][%d] getValidationRequestBadRequest ", 400)
}

func (o *GetValidationRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetValidationRequestUnauthorized creates a GetValidationRequestUnauthorized with default headers values
func NewGetValidationRequestUnauthorized() *GetValidationRequestUnauthorized {
	return &GetValidationRequestUnauthorized{}
}

/*GetValidationRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetValidationRequestUnauthorized struct {
}

func (o *GetValidationRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/validation-requests/{id}][%d] getValidationRequestUnauthorized ", 401)
}

func (o *GetValidationRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetValidationRequestForbidden creates a GetValidationRequestForbidden with default headers values
func NewGetValidationRequestForbidden() *GetValidationRequestForbidden {
	return &GetValidationRequestForbidden{}
}

/*GetValidationRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetValidationRequestForbidden struct {
}

func (o *GetValidationRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/validation-requests/{id}][%d] getValidationRequestForbidden ", 403)
}

func (o *GetValidationRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetValidationRequestNotFound creates a GetValidationRequestNotFound with default headers values
func NewGetValidationRequestNotFound() *GetValidationRequestNotFound {
	return &GetValidationRequestNotFound{}
}

/*GetValidationRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetValidationRequestNotFound struct {
}

func (o *GetValidationRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/validation-requests/{id}][%d] getValidationRequestNotFound ", 404)
}

func (o *GetValidationRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetValidationRequestConflict creates a GetValidationRequestConflict with default headers values
func NewGetValidationRequestConflict() *GetValidationRequestConflict {
	return &GetValidationRequestConflict{}
}

/*GetValidationRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetValidationRequestConflict struct {
}

func (o *GetValidationRequestConflict) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/validation-requests/{id}][%d] getValidationRequestConflict ", 409)
}

func (o *GetValidationRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}