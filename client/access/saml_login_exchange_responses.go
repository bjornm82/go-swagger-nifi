// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SamlLoginExchangeReader is a Reader for the SamlLoginExchange structure.
type SamlLoginExchangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlLoginExchangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSamlLoginExchangeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSamlLoginExchangeCreated creates a SamlLoginExchangeCreated with default headers values
func NewSamlLoginExchangeCreated() *SamlLoginExchangeCreated {
	return &SamlLoginExchangeCreated{}
}

/*SamlLoginExchangeCreated handles this case with default header values.

successful operation
*/
type SamlLoginExchangeCreated struct {
	Payload string
}

func (o *SamlLoginExchangeCreated) Error() string {
	return fmt.Sprintf("[POST /access/saml/login/exchange][%d] samlLoginExchangeCreated  %+v", 201, o.Payload)
}

func (o *SamlLoginExchangeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
