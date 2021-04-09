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

// CreateLabelReader is a Reader for the CreateLabel structure.
type CreateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLabelOK creates a CreateLabelOK with default headers values
func NewCreateLabelOK() *CreateLabelOK {
	return &CreateLabelOK{}
}

/*CreateLabelOK handles this case with default header values.

successful operation
*/
type CreateLabelOK struct {
	Payload *models.LabelEntity
}

func (o *CreateLabelOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/labels][%d] createLabelOK  %+v", 200, o.Payload)
}

func (o *CreateLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLabelBadRequest creates a CreateLabelBadRequest with default headers values
func NewCreateLabelBadRequest() *CreateLabelBadRequest {
	return &CreateLabelBadRequest{}
}

/*CreateLabelBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateLabelBadRequest struct {
}

func (o *CreateLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/labels][%d] createLabelBadRequest ", 400)
}

func (o *CreateLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLabelUnauthorized creates a CreateLabelUnauthorized with default headers values
func NewCreateLabelUnauthorized() *CreateLabelUnauthorized {
	return &CreateLabelUnauthorized{}
}

/*CreateLabelUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateLabelUnauthorized struct {
}

func (o *CreateLabelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/labels][%d] createLabelUnauthorized ", 401)
}

func (o *CreateLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLabelForbidden creates a CreateLabelForbidden with default headers values
func NewCreateLabelForbidden() *CreateLabelForbidden {
	return &CreateLabelForbidden{}
}

/*CreateLabelForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateLabelForbidden struct {
}

func (o *CreateLabelForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/labels][%d] createLabelForbidden ", 403)
}

func (o *CreateLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLabelNotFound creates a CreateLabelNotFound with default headers values
func NewCreateLabelNotFound() *CreateLabelNotFound {
	return &CreateLabelNotFound{}
}

/*CreateLabelNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreateLabelNotFound struct {
}

func (o *CreateLabelNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/labels][%d] createLabelNotFound ", 404)
}

func (o *CreateLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLabelConflict creates a CreateLabelConflict with default headers values
func NewCreateLabelConflict() *CreateLabelConflict {
	return &CreateLabelConflict{}
}

/*CreateLabelConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateLabelConflict struct {
}

func (o *CreateLabelConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/labels][%d] createLabelConflict ", 409)
}

func (o *CreateLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}