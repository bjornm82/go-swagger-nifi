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

// GetControllerServicesFromGroupReader is a Reader for the GetControllerServicesFromGroup structure.
type GetControllerServicesFromGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControllerServicesFromGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetControllerServicesFromGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetControllerServicesFromGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetControllerServicesFromGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetControllerServicesFromGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetControllerServicesFromGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetControllerServicesFromGroupOK creates a GetControllerServicesFromGroupOK with default headers values
func NewGetControllerServicesFromGroupOK() *GetControllerServicesFromGroupOK {
	return &GetControllerServicesFromGroupOK{}
}

/*GetControllerServicesFromGroupOK handles this case with default header values.

successful operation
*/
type GetControllerServicesFromGroupOK struct {
	Payload *models.ControllerServicesEntity
}

func (o *GetControllerServicesFromGroupOK) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/controller-services][%d] getControllerServicesFromGroupOK  %+v", 200, o.Payload)
}

func (o *GetControllerServicesFromGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServicesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControllerServicesFromGroupBadRequest creates a GetControllerServicesFromGroupBadRequest with default headers values
func NewGetControllerServicesFromGroupBadRequest() *GetControllerServicesFromGroupBadRequest {
	return &GetControllerServicesFromGroupBadRequest{}
}

/*GetControllerServicesFromGroupBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetControllerServicesFromGroupBadRequest struct {
}

func (o *GetControllerServicesFromGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/controller-services][%d] getControllerServicesFromGroupBadRequest ", 400)
}

func (o *GetControllerServicesFromGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServicesFromGroupUnauthorized creates a GetControllerServicesFromGroupUnauthorized with default headers values
func NewGetControllerServicesFromGroupUnauthorized() *GetControllerServicesFromGroupUnauthorized {
	return &GetControllerServicesFromGroupUnauthorized{}
}

/*GetControllerServicesFromGroupUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetControllerServicesFromGroupUnauthorized struct {
}

func (o *GetControllerServicesFromGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/controller-services][%d] getControllerServicesFromGroupUnauthorized ", 401)
}

func (o *GetControllerServicesFromGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServicesFromGroupForbidden creates a GetControllerServicesFromGroupForbidden with default headers values
func NewGetControllerServicesFromGroupForbidden() *GetControllerServicesFromGroupForbidden {
	return &GetControllerServicesFromGroupForbidden{}
}

/*GetControllerServicesFromGroupForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetControllerServicesFromGroupForbidden struct {
}

func (o *GetControllerServicesFromGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/controller-services][%d] getControllerServicesFromGroupForbidden ", 403)
}

func (o *GetControllerServicesFromGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServicesFromGroupConflict creates a GetControllerServicesFromGroupConflict with default headers values
func NewGetControllerServicesFromGroupConflict() *GetControllerServicesFromGroupConflict {
	return &GetControllerServicesFromGroupConflict{}
}

/*GetControllerServicesFromGroupConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetControllerServicesFromGroupConflict struct {
}

func (o *GetControllerServicesFromGroupConflict) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/controller-services][%d] getControllerServicesFromGroupConflict ", 409)
}

func (o *GetControllerServicesFromGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
