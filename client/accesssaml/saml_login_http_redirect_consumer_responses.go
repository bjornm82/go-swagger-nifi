// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SamlLoginHTTPRedirectConsumerReader is a Reader for the SamlLoginHTTPRedirectConsumer structure.
type SamlLoginHTTPRedirectConsumerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlLoginHTTPRedirectConsumerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSamlLoginHTTPRedirectConsumerDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSamlLoginHTTPRedirectConsumerDefault creates a SamlLoginHTTPRedirectConsumerDefault with default headers values
func NewSamlLoginHTTPRedirectConsumerDefault(code int) *SamlLoginHTTPRedirectConsumerDefault {
	return &SamlLoginHTTPRedirectConsumerDefault{
		_statusCode: code,
	}
}

/*SamlLoginHTTPRedirectConsumerDefault handles this case with default header values.

successful operation
*/
type SamlLoginHTTPRedirectConsumerDefault struct {
	_statusCode int
}

// Code gets the status code for the saml login Http redirect consumer default response
func (o *SamlLoginHTTPRedirectConsumerDefault) Code() int {
	return o._statusCode
}

func (o *SamlLoginHTTPRedirectConsumerDefault) Error() string {
	return fmt.Sprintf("[GET /access/saml/login/consumer][%d] samlLoginHttpRedirectConsumer default ", o._statusCode)
}

func (o *SamlLoginHTTPRedirectConsumerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
