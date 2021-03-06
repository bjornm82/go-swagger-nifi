// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// RemoveLabelReader is a Reader for the RemoveLabel structure.
type RemoveLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveLabelOK creates a RemoveLabelOK with default headers values
func NewRemoveLabelOK() *RemoveLabelOK {
	return &RemoveLabelOK{}
}

/*RemoveLabelOK handles this case with default header values.

successful operation
*/
type RemoveLabelOK struct {
	Payload *models.LabelEntity
}

func (o *RemoveLabelOK) Error() string {
	return fmt.Sprintf("[DELETE /labels/{id}][%d] removeLabelOK  %+v", 200, o.Payload)
}

func (o *RemoveLabelOK) GetPayload() *models.LabelEntity {
	return o.Payload
}

func (o *RemoveLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLabelBadRequest creates a RemoveLabelBadRequest with default headers values
func NewRemoveLabelBadRequest() *RemoveLabelBadRequest {
	return &RemoveLabelBadRequest{}
}

/*RemoveLabelBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveLabelBadRequest struct {
}

func (o *RemoveLabelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /labels/{id}][%d] removeLabelBadRequest ", 400)
}

func (o *RemoveLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveLabelUnauthorized creates a RemoveLabelUnauthorized with default headers values
func NewRemoveLabelUnauthorized() *RemoveLabelUnauthorized {
	return &RemoveLabelUnauthorized{}
}

/*RemoveLabelUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveLabelUnauthorized struct {
}

func (o *RemoveLabelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /labels/{id}][%d] removeLabelUnauthorized ", 401)
}

func (o *RemoveLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveLabelForbidden creates a RemoveLabelForbidden with default headers values
func NewRemoveLabelForbidden() *RemoveLabelForbidden {
	return &RemoveLabelForbidden{}
}

/*RemoveLabelForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveLabelForbidden struct {
}

func (o *RemoveLabelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /labels/{id}][%d] removeLabelForbidden ", 403)
}

func (o *RemoveLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveLabelNotFound creates a RemoveLabelNotFound with default headers values
func NewRemoveLabelNotFound() *RemoveLabelNotFound {
	return &RemoveLabelNotFound{}
}

/*RemoveLabelNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveLabelNotFound struct {
}

func (o *RemoveLabelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /labels/{id}][%d] removeLabelNotFound ", 404)
}

func (o *RemoveLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveLabelConflict creates a RemoveLabelConflict with default headers values
func NewRemoveLabelConflict() *RemoveLabelConflict {
	return &RemoveLabelConflict{}
}

/*RemoveLabelConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveLabelConflict struct {
}

func (o *RemoveLabelConflict) Error() string {
	return fmt.Sprintf("[DELETE /labels/{id}][%d] removeLabelConflict ", 409)
}

func (o *RemoveLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
