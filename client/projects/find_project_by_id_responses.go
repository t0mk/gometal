// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindProjectByIDReader is a Reader for the FindProjectByID structure.
type FindProjectByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindProjectByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindProjectByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindProjectByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindProjectByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindProjectByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindProjectByIDOK creates a FindProjectByIDOK with default headers values
func NewFindProjectByIDOK() *FindProjectByIDOK {
	return &FindProjectByIDOK{}
}

/* FindProjectByIDOK describes a response with status code 200, with default header values.

ok
*/
type FindProjectByIDOK struct {
	Payload *types.Project
}

func (o *FindProjectByIDOK) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] findProjectByIdOK  %+v", 200, o.Payload)
}
func (o *FindProjectByIDOK) GetPayload() *types.Project {
	return o.Payload
}

func (o *FindProjectByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindProjectByIDUnauthorized creates a FindProjectByIDUnauthorized with default headers values
func NewFindProjectByIDUnauthorized() *FindProjectByIDUnauthorized {
	return &FindProjectByIDUnauthorized{}
}

/* FindProjectByIDUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindProjectByIDUnauthorized struct {
}

func (o *FindProjectByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] findProjectByIdUnauthorized ", 401)
}

func (o *FindProjectByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindProjectByIDForbidden creates a FindProjectByIDForbidden with default headers values
func NewFindProjectByIDForbidden() *FindProjectByIDForbidden {
	return &FindProjectByIDForbidden{}
}

/* FindProjectByIDForbidden describes a response with status code 403, with default header values.

forbidden
*/
type FindProjectByIDForbidden struct {
}

func (o *FindProjectByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] findProjectByIdForbidden ", 403)
}

func (o *FindProjectByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindProjectByIDNotFound creates a FindProjectByIDNotFound with default headers values
func NewFindProjectByIDNotFound() *FindProjectByIDNotFound {
	return &FindProjectByIDNotFound{}
}

/* FindProjectByIDNotFound describes a response with status code 404, with default header values.

not found
*/
type FindProjectByIDNotFound struct {
}

func (o *FindProjectByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] findProjectByIdNotFound ", 404)
}

func (o *FindProjectByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}