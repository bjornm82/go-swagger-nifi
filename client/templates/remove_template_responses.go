// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// RemoveTemplateReader is a Reader for the RemoveTemplate structure.
type RemoveTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveTemplateOK creates a RemoveTemplateOK with default headers values
func NewRemoveTemplateOK() *RemoveTemplateOK {
	return &RemoveTemplateOK{}
}

/*RemoveTemplateOK handles this case with default header values.

successful operation
*/
type RemoveTemplateOK struct {
	Payload *models.TemplateEntity
}

func (o *RemoveTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /templates/{id}][%d] removeTemplateOK  %+v", 200, o.Payload)
}

func (o *RemoveTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTemplateBadRequest creates a RemoveTemplateBadRequest with default headers values
func NewRemoveTemplateBadRequest() *RemoveTemplateBadRequest {
	return &RemoveTemplateBadRequest{}
}

/*RemoveTemplateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveTemplateBadRequest struct {
}

func (o *RemoveTemplateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /templates/{id}][%d] removeTemplateBadRequest ", 400)
}

func (o *RemoveTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveTemplateUnauthorized creates a RemoveTemplateUnauthorized with default headers values
func NewRemoveTemplateUnauthorized() *RemoveTemplateUnauthorized {
	return &RemoveTemplateUnauthorized{}
}

/*RemoveTemplateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveTemplateUnauthorized struct {
}

func (o *RemoveTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /templates/{id}][%d] removeTemplateUnauthorized ", 401)
}

func (o *RemoveTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveTemplateForbidden creates a RemoveTemplateForbidden with default headers values
func NewRemoveTemplateForbidden() *RemoveTemplateForbidden {
	return &RemoveTemplateForbidden{}
}

/*RemoveTemplateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveTemplateForbidden struct {
}

func (o *RemoveTemplateForbidden) Error() string {
	return fmt.Sprintf("[DELETE /templates/{id}][%d] removeTemplateForbidden ", 403)
}

func (o *RemoveTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveTemplateNotFound creates a RemoveTemplateNotFound with default headers values
func NewRemoveTemplateNotFound() *RemoveTemplateNotFound {
	return &RemoveTemplateNotFound{}
}

/*RemoveTemplateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveTemplateNotFound struct {
}

func (o *RemoveTemplateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /templates/{id}][%d] removeTemplateNotFound ", 404)
}

func (o *RemoveTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveTemplateConflict creates a RemoveTemplateConflict with default headers values
func NewRemoveTemplateConflict() *RemoveTemplateConflict {
	return &RemoveTemplateConflict{}
}

/*RemoveTemplateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveTemplateConflict struct {
}

func (o *RemoveTemplateConflict) Error() string {
	return fmt.Sprintf("[DELETE /templates/{id}][%d] removeTemplateConflict ", 409)
}

func (o *RemoveTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}