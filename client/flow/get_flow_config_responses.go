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

// GetFlowConfigReader is a Reader for the GetFlowConfig structure.
type GetFlowConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlowConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFlowConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFlowConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetFlowConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetFlowConfigConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowConfigOK creates a GetFlowConfigOK with default headers values
func NewGetFlowConfigOK() *GetFlowConfigOK {
	return &GetFlowConfigOK{}
}

/*GetFlowConfigOK handles this case with default header values.

successful operation
*/
type GetFlowConfigOK struct {
	Payload *models.FlowConfigurationEntity
}

func (o *GetFlowConfigOK) Error() string {
	return fmt.Sprintf("[GET /flow/config][%d] getFlowConfigOK  %+v", 200, o.Payload)
}

func (o *GetFlowConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowConfigurationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowConfigBadRequest creates a GetFlowConfigBadRequest with default headers values
func NewGetFlowConfigBadRequest() *GetFlowConfigBadRequest {
	return &GetFlowConfigBadRequest{}
}

/*GetFlowConfigBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFlowConfigBadRequest struct {
}

func (o *GetFlowConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/config][%d] getFlowConfigBadRequest ", 400)
}

func (o *GetFlowConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowConfigUnauthorized creates a GetFlowConfigUnauthorized with default headers values
func NewGetFlowConfigUnauthorized() *GetFlowConfigUnauthorized {
	return &GetFlowConfigUnauthorized{}
}

/*GetFlowConfigUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetFlowConfigUnauthorized struct {
}

func (o *GetFlowConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/config][%d] getFlowConfigUnauthorized ", 401)
}

func (o *GetFlowConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowConfigForbidden creates a GetFlowConfigForbidden with default headers values
func NewGetFlowConfigForbidden() *GetFlowConfigForbidden {
	return &GetFlowConfigForbidden{}
}

/*GetFlowConfigForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetFlowConfigForbidden struct {
}

func (o *GetFlowConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/config][%d] getFlowConfigForbidden ", 403)
}

func (o *GetFlowConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowConfigConflict creates a GetFlowConfigConflict with default headers values
func NewGetFlowConfigConflict() *GetFlowConfigConflict {
	return &GetFlowConfigConflict{}
}

/*GetFlowConfigConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetFlowConfigConflict struct {
}

func (o *GetFlowConfigConflict) Error() string {
	return fmt.Sprintf("[GET /flow/config][%d] getFlowConfigConflict ", 409)
}

func (o *GetFlowConfigConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
