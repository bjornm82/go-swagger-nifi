// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetLoginConfigSamlReader is a Reader for the GetLoginConfigSaml structure.
type GetLoginConfigSamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginConfigSamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoginConfigSamlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoginConfigSamlOK creates a GetLoginConfigSamlOK with default headers values
func NewGetLoginConfigSamlOK() *GetLoginConfigSamlOK {
	return &GetLoginConfigSamlOK{}
}

/*GetLoginConfigSamlOK handles this case with default header values.

successful operation
*/
type GetLoginConfigSamlOK struct {
	Payload *models.AccessConfigurationEntity
}

func (o *GetLoginConfigSamlOK) Error() string {
	return fmt.Sprintf("[GET /access/saml/config][%d] getLoginConfigSamlOK  %+v", 200, o.Payload)
}

func (o *GetLoginConfigSamlOK) GetPayload() *models.AccessConfigurationEntity {
	return o.Payload
}

func (o *GetLoginConfigSamlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessConfigurationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}