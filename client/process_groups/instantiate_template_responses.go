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

// InstantiateTemplateReader is a Reader for the InstantiateTemplate structure.
type InstantiateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstantiateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInstantiateTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewInstantiateTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewInstantiateTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewInstantiateTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewInstantiateTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewInstantiateTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstantiateTemplateOK creates a InstantiateTemplateOK with default headers values
func NewInstantiateTemplateOK() *InstantiateTemplateOK {
	return &InstantiateTemplateOK{}
}

/*InstantiateTemplateOK handles this case with default header values.

successful operation
*/
type InstantiateTemplateOK struct {
	Payload *models.FlowEntity
}

func (o *InstantiateTemplateOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateOK  %+v", 200, o.Payload)
}

func (o *InstantiateTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantiateTemplateBadRequest creates a InstantiateTemplateBadRequest with default headers values
func NewInstantiateTemplateBadRequest() *InstantiateTemplateBadRequest {
	return &InstantiateTemplateBadRequest{}
}

/*InstantiateTemplateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type InstantiateTemplateBadRequest struct {
}

func (o *InstantiateTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateBadRequest ", 400)
}

func (o *InstantiateTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateUnauthorized creates a InstantiateTemplateUnauthorized with default headers values
func NewInstantiateTemplateUnauthorized() *InstantiateTemplateUnauthorized {
	return &InstantiateTemplateUnauthorized{}
}

/*InstantiateTemplateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type InstantiateTemplateUnauthorized struct {
}

func (o *InstantiateTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateUnauthorized ", 401)
}

func (o *InstantiateTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateForbidden creates a InstantiateTemplateForbidden with default headers values
func NewInstantiateTemplateForbidden() *InstantiateTemplateForbidden {
	return &InstantiateTemplateForbidden{}
}

/*InstantiateTemplateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type InstantiateTemplateForbidden struct {
}

func (o *InstantiateTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateForbidden ", 403)
}

func (o *InstantiateTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateNotFound creates a InstantiateTemplateNotFound with default headers values
func NewInstantiateTemplateNotFound() *InstantiateTemplateNotFound {
	return &InstantiateTemplateNotFound{}
}

/*InstantiateTemplateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type InstantiateTemplateNotFound struct {
}

func (o *InstantiateTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateNotFound ", 404)
}

func (o *InstantiateTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateConflict creates a InstantiateTemplateConflict with default headers values
func NewInstantiateTemplateConflict() *InstantiateTemplateConflict {
	return &InstantiateTemplateConflict{}
}

/*InstantiateTemplateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type InstantiateTemplateConflict struct {
}

func (o *InstantiateTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateConflict ", 409)
}

func (o *InstantiateTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
