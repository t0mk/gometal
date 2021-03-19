// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// GetInterconnectionReader is a Reader for the GetInterconnection structure.
type GetInterconnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInterconnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInterconnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetInterconnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInterconnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInterconnectionOK creates a GetInterconnectionOK with default headers values
func NewGetInterconnectionOK() *GetInterconnectionOK {
	return &GetInterconnectionOK{}
}

/* GetInterconnectionOK describes a response with status code 200, with default header values.

ok
*/
type GetInterconnectionOK struct {
	Payload *types.Interconnection
}

func (o *GetInterconnectionOK) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}][%d] getInterconnectionOK  %+v", 200, o.Payload)
}
func (o *GetInterconnectionOK) GetPayload() *types.Interconnection {
	return o.Payload
}

func (o *GetInterconnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Interconnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInterconnectionForbidden creates a GetInterconnectionForbidden with default headers values
func NewGetInterconnectionForbidden() *GetInterconnectionForbidden {
	return &GetInterconnectionForbidden{}
}

/* GetInterconnectionForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetInterconnectionForbidden struct {
}

func (o *GetInterconnectionForbidden) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}][%d] getInterconnectionForbidden ", 403)
}

func (o *GetInterconnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInterconnectionNotFound creates a GetInterconnectionNotFound with default headers values
func NewGetInterconnectionNotFound() *GetInterconnectionNotFound {
	return &GetInterconnectionNotFound{}
}

/* GetInterconnectionNotFound describes a response with status code 404, with default header values.

not found
*/
type GetInterconnectionNotFound struct {
}

func (o *GetInterconnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}][%d] getInterconnectionNotFound ", 404)
}

func (o *GetInterconnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}