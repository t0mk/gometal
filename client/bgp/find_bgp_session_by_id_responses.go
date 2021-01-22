// Code generated by go-swagger; DO NOT EDIT.

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindBGPSessionByIDReader is a Reader for the FindBGPSessionByID structure.
type FindBGPSessionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindBGPSessionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindBGPSessionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindBGPSessionByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindBGPSessionByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindBGPSessionByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindBGPSessionByIDOK creates a FindBGPSessionByIDOK with default headers values
func NewFindBGPSessionByIDOK() *FindBGPSessionByIDOK {
	return &FindBGPSessionByIDOK{}
}

/* FindBGPSessionByIDOK describes a response with status code 200, with default header values.

ok
*/
type FindBGPSessionByIDOK struct {
	Payload *types.BGPSession
}

func (o *FindBGPSessionByIDOK) Error() string {
	return fmt.Sprintf("[GET /bgp/sessions/{id}][%d] findBgpSessionByIdOK  %+v", 200, o.Payload)
}
func (o *FindBGPSessionByIDOK) GetPayload() *types.BGPSession {
	return o.Payload
}

func (o *FindBGPSessionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.BGPSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindBGPSessionByIDUnauthorized creates a FindBGPSessionByIDUnauthorized with default headers values
func NewFindBGPSessionByIDUnauthorized() *FindBGPSessionByIDUnauthorized {
	return &FindBGPSessionByIDUnauthorized{}
}

/* FindBGPSessionByIDUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindBGPSessionByIDUnauthorized struct {
}

func (o *FindBGPSessionByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bgp/sessions/{id}][%d] findBgpSessionByIdUnauthorized ", 401)
}

func (o *FindBGPSessionByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindBGPSessionByIDForbidden creates a FindBGPSessionByIDForbidden with default headers values
func NewFindBGPSessionByIDForbidden() *FindBGPSessionByIDForbidden {
	return &FindBGPSessionByIDForbidden{}
}

/* FindBGPSessionByIDForbidden describes a response with status code 403, with default header values.

forbidden
*/
type FindBGPSessionByIDForbidden struct {
}

func (o *FindBGPSessionByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /bgp/sessions/{id}][%d] findBgpSessionByIdForbidden ", 403)
}

func (o *FindBGPSessionByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindBGPSessionByIDNotFound creates a FindBGPSessionByIDNotFound with default headers values
func NewFindBGPSessionByIDNotFound() *FindBGPSessionByIDNotFound {
	return &FindBGPSessionByIDNotFound{}
}

/* FindBGPSessionByIDNotFound describes a response with status code 404, with default header values.

not found
*/
type FindBGPSessionByIDNotFound struct {
}

func (o *FindBGPSessionByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /bgp/sessions/{id}][%d] findBgpSessionByIdNotFound ", 404)
}

func (o *FindBGPSessionByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}