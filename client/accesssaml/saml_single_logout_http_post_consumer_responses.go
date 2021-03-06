// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SamlSingleLogoutHTTPPostConsumerReader is a Reader for the SamlSingleLogoutHTTPPostConsumer structure.
type SamlSingleLogoutHTTPPostConsumerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlSingleLogoutHTTPPostConsumerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSamlSingleLogoutHTTPPostConsumerDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSamlSingleLogoutHTTPPostConsumerDefault creates a SamlSingleLogoutHTTPPostConsumerDefault with default headers values
func NewSamlSingleLogoutHTTPPostConsumerDefault(code int) *SamlSingleLogoutHTTPPostConsumerDefault {
	return &SamlSingleLogoutHTTPPostConsumerDefault{
		_statusCode: code,
	}
}

/*SamlSingleLogoutHTTPPostConsumerDefault handles this case with default header values.

successful operation
*/
type SamlSingleLogoutHTTPPostConsumerDefault struct {
	_statusCode int
}

// Code gets the status code for the saml single logout Http post consumer default response
func (o *SamlSingleLogoutHTTPPostConsumerDefault) Code() int {
	return o._statusCode
}

func (o *SamlSingleLogoutHTTPPostConsumerDefault) Error() string {
	return fmt.Sprintf("[POST /access/saml/single-logout/consumer][%d] samlSingleLogoutHttpPostConsumer default ", o._statusCode)
}

func (o *SamlSingleLogoutHTTPPostConsumerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
