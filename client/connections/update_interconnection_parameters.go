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

	"github.com/t0mk/gometal/types"
)

// NewUpdateInterconnectionParams creates a new UpdateInterconnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateInterconnectionParams() *UpdateInterconnectionParams {
	return &UpdateInterconnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInterconnectionParamsWithTimeout creates a new UpdateInterconnectionParams object
// with the ability to set a timeout on a request.
func NewUpdateInterconnectionParamsWithTimeout(timeout time.Duration) *UpdateInterconnectionParams {
	return &UpdateInterconnectionParams{
		timeout: timeout,
	}
}

// NewUpdateInterconnectionParamsWithContext creates a new UpdateInterconnectionParams object
// with the ability to set a context for a request.
func NewUpdateInterconnectionParamsWithContext(ctx context.Context) *UpdateInterconnectionParams {
	return &UpdateInterconnectionParams{
		Context: ctx,
	}
}

// NewUpdateInterconnectionParamsWithHTTPClient creates a new UpdateInterconnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateInterconnectionParamsWithHTTPClient(client *http.Client) *UpdateInterconnectionParams {
	return &UpdateInterconnectionParams{
		HTTPClient: client,
	}
}

/* UpdateInterconnectionParams contains all the parameters to send to the API endpoint
   for the update interconnection operation.

   Typically these are written to a http.Request.
*/
type UpdateInterconnectionParams struct {

	/* Connection.

	   Updated connection details
	*/
	Connection *types.InterconnectionUpdateInput

	/* ConnectionID.

	   Connection UUID

	   Format: uuid
	*/
	ConnectionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update interconnection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInterconnectionParams) WithDefaults() *UpdateInterconnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update interconnection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInterconnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update interconnection params
func (o *UpdateInterconnectionParams) WithTimeout(timeout time.Duration) *UpdateInterconnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update interconnection params
func (o *UpdateInterconnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update interconnection params
func (o *UpdateInterconnectionParams) WithContext(ctx context.Context) *UpdateInterconnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update interconnection params
func (o *UpdateInterconnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update interconnection params
func (o *UpdateInterconnectionParams) WithHTTPClient(client *http.Client) *UpdateInterconnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update interconnection params
func (o *UpdateInterconnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnection adds the connection to the update interconnection params
func (o *UpdateInterconnectionParams) WithConnection(connection *types.InterconnectionUpdateInput) *UpdateInterconnectionParams {
	o.SetConnection(connection)
	return o
}

// SetConnection adds the connection to the update interconnection params
func (o *UpdateInterconnectionParams) SetConnection(connection *types.InterconnectionUpdateInput) {
	o.Connection = connection
}

// WithConnectionID adds the connectionID to the update interconnection params
func (o *UpdateInterconnectionParams) WithConnectionID(connectionID strfmt.UUID) *UpdateInterconnectionParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the update interconnection params
func (o *UpdateInterconnectionParams) SetConnectionID(connectionID strfmt.UUID) {
	o.ConnectionID = connectionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInterconnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Connection != nil {
		if err := r.SetBodyParam(o.Connection); err != nil {
			return err
		}
	}

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}