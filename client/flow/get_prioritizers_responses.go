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

// GetPrioritizersReader is a Reader for the GetPrioritizers structure.
type GetPrioritizersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrioritizersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrioritizersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPrioritizersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetPrioritizersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPrioritizersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetPrioritizersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrioritizersOK creates a GetPrioritizersOK with default headers values
func NewGetPrioritizersOK() *GetPrioritizersOK {
	return &GetPrioritizersOK{}
}

/*GetPrioritizersOK handles this case with default header values.

successful operation
*/
type GetPrioritizersOK struct {
	Payload *models.PrioritizerTypesEntity
}

func (o *GetPrioritizersOK) Error() string {
	return fmt.Sprintf("[GET /flow/prioritizers][%d] getPrioritizersOK  %+v", 200, o.Payload)
}

func (o *GetPrioritizersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrioritizerTypesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrioritizersBadRequest creates a GetPrioritizersBadRequest with default headers values
func NewGetPrioritizersBadRequest() *GetPrioritizersBadRequest {
	return &GetPrioritizersBadRequest{}
}

/*GetPrioritizersBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetPrioritizersBadRequest struct {
}

func (o *GetPrioritizersBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/prioritizers][%d] getPrioritizersBadRequest ", 400)
}

func (o *GetPrioritizersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrioritizersUnauthorized creates a GetPrioritizersUnauthorized with default headers values
func NewGetPrioritizersUnauthorized() *GetPrioritizersUnauthorized {
	return &GetPrioritizersUnauthorized{}
}

/*GetPrioritizersUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetPrioritizersUnauthorized struct {
}

func (o *GetPrioritizersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/prioritizers][%d] getPrioritizersUnauthorized ", 401)
}

func (o *GetPrioritizersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrioritizersForbidden creates a GetPrioritizersForbidden with default headers values
func NewGetPrioritizersForbidden() *GetPrioritizersForbidden {
	return &GetPrioritizersForbidden{}
}

/*GetPrioritizersForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetPrioritizersForbidden struct {
}

func (o *GetPrioritizersForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/prioritizers][%d] getPrioritizersForbidden ", 403)
}

func (o *GetPrioritizersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrioritizersConflict creates a GetPrioritizersConflict with default headers values
func NewGetPrioritizersConflict() *GetPrioritizersConflict {
	return &GetPrioritizersConflict{}
}

/*GetPrioritizersConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetPrioritizersConflict struct {
}

func (o *GetPrioritizersConflict) Error() string {
	return fmt.Sprintf("[GET /flow/prioritizers][%d] getPrioritizersConflict ", 409)
}

func (o *GetPrioritizersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}