// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteInterconnectionParams creates a new DeleteInterconnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInterconnectionParams() *DeleteInterconnectionParams {
	return &DeleteInterconnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInterconnectionParamsWithTimeout creates a new DeleteInterconnectionParams object
// with the ability to set a timeout on a request.
func NewDeleteInterconnectionParamsWithTimeout(timeout time.Duration) *DeleteInterconnectionParams {
	return &DeleteInterconnectionParams{
		timeout: timeout,
	}
}

// NewDeleteInterconnectionParamsWithContext creates a new DeleteInterconnectionParams object
// with the ability to set a context for a request.
func NewDeleteInterconnectionParamsWithContext(ctx context.Context) *DeleteInterconnectionParams {
	return &DeleteInterconnectionParams{
		Context: ctx,
	}
}

// NewDeleteInterconnectionParamsWithHTTPClient creates a new DeleteInterconnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInterconnectionParamsWithHTTPClient(client *http.Client) *DeleteInterconnectionParams {
	return &DeleteInterconnectionParams{
		HTTPClient: client,
	}
}

/* DeleteInterconnectionParams contains all the parameters to send to the API endpoint
   for the delete interconnection operation.

   Typically these are written to a http.Request.
*/
type DeleteInterconnectionParams struct {

	/* ConnectionID.

	   Connection UUID

	   Format: uuid
	*/
	ConnectionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete interconnection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInterconnectionParams) WithDefaults() *DeleteInterconnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete interconnection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInterconnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete interconnection params
func (o *DeleteInterconnectionParams) WithTimeout(timeout time.Duration) *DeleteInterconnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete interconnection params
func (o *DeleteInterconnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete interconnection params
func (o *DeleteInterconnectionParams) WithContext(ctx context.Context) *DeleteInterconnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete interconnection params
func (o *DeleteInterconnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete interconnection params
func (o *DeleteInterconnectionParams) WithHTTPClient(client *http.Client) *DeleteInterconnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete interconnection params
func (o *DeleteInterconnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the delete interconnection params
func (o *DeleteInterconnectionParams) WithConnectionID(connectionID strfmt.UUID) *DeleteInterconnectionParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the delete interconnection params
func (o *DeleteInterconnectionParams) SetConnectionID(connectionID strfmt.UUID) {
	o.ConnectionID = connectionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInterconnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}