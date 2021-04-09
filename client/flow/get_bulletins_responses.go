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

// GetBulletinsReader is a Reader for the GetBulletins structure.
type GetBulletinsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBulletinsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBulletinsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetBulletinsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetBulletinsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetBulletinsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetBulletinsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetBulletinsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBulletinsOK creates a GetBulletinsOK with default headers values
func NewGetBulletinsOK() *GetBulletinsOK {
	return &GetBulletinsOK{}
}

/*GetBulletinsOK handles this case with default header values.

successful operation
*/
type GetBulletinsOK struct {
	Payload *models.ControllerBulletinsEntity
}

func (o *GetBulletinsOK) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsOK  %+v", 200, o.Payload)
}

func (o *GetBulletinsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerBulletinsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBulletinsBadRequest creates a GetBulletinsBadRequest with default headers values
func NewGetBulletinsBadRequest() *GetBulletinsBadRequest {
	return &GetBulletinsBadRequest{}
}

/*GetBulletinsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetBulletinsBadRequest struct {
}

func (o *GetBulletinsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsBadRequest ", 400)
}

func (o *GetBulletinsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsUnauthorized creates a GetBulletinsUnauthorized with default headers values
func NewGetBulletinsUnauthorized() *GetBulletinsUnauthorized {
	return &GetBulletinsUnauthorized{}
}

/*GetBulletinsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetBulletinsUnauthorized struct {
}

func (o *GetBulletinsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsUnauthorized ", 401)
}

func (o *GetBulletinsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsForbidden creates a GetBulletinsForbidden with default headers values
func NewGetBulletinsForbidden() *GetBulletinsForbidden {
	return &GetBulletinsForbidden{}
}

/*GetBulletinsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetBulletinsForbidden struct {
}

func (o *GetBulletinsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsForbidden ", 403)
}

func (o *GetBulletinsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsNotFound creates a GetBulletinsNotFound with default headers values
func NewGetBulletinsNotFound() *GetBulletinsNotFound {
	return &GetBulletinsNotFound{}
}

/*GetBulletinsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetBulletinsNotFound struct {
}

func (o *GetBulletinsNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsNotFound ", 404)
}

func (o *GetBulletinsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsConflict creates a GetBulletinsConflict with default headers values
func NewGetBulletinsConflict() *GetBulletinsConflict {
	return &GetBulletinsConflict{}
}

/*GetBulletinsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetBulletinsConflict struct {
}

func (o *GetBulletinsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsConflict ", 409)
}

func (o *GetBulletinsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}