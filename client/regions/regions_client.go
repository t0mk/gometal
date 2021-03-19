// Code generated by go-swagger; DO NOT EDIT.

package regions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new regions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for regions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FindRegions(params *FindRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRegionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FindRegions retrieves all regions

  Returns all regions.
*/
func (a *Client) FindRegions(params *FindRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindRegionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findRegions",
		Method:             "GET",
		PathPattern:        "/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindRegionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findRegions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}