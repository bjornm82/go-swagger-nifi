// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// GetLoginConfigReader is a Reader for the GetLoginConfig structure.
type GetLoginConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoginConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoginConfigOK creates a GetLoginConfigOK with default headers values
func NewGetLoginConfigOK() *GetLoginConfigOK {
	return &GetLoginConfigOK{}
}

/*GetLoginConfigOK handles this case with default header values.

successful operation
*/
type GetLoginConfigOK struct {
	Payload *models.AccessConfigurationEntity
}

func (o *GetLoginConfigOK) Error() string {
	return fmt.Sprintf("[GET /access/config][%d] getLoginConfigOK  %+v", 200, o.Payload)
}

func (o *GetLoginConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessConfigurationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
