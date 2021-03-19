// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewFindUserCustomdataParams creates a new FindUserCustomdataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindUserCustomdataParams() *FindUserCustomdataParams {
	return &FindUserCustomdataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindUserCustomdataParamsWithTimeout creates a new FindUserCustomdataParams object
// with the ability to set a timeout on a request.
func NewFindUserCustomdataParamsWithTimeout(timeout time.Duration) *FindUserCustomdataParams {
	return &FindUserCustomdataParams{
		timeout: timeout,
	}
}

// NewFindUserCustomdataParamsWithContext creates a new FindUserCustomdataParams object
// with the ability to set a context for a request.
func NewFindUserCustomdataParamsWithContext(ctx context.Context) *FindUserCustomdataParams {
	return &FindUserCustomdataParams{
		Context: ctx,
	}
}

// NewFindUserCustomdataParamsWithHTTPClient creates a new FindUserCustomdataParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindUserCustomdataParamsWithHTTPClient(client *http.Client) *FindUserCustomdataParams {
	return &FindUserCustomdataParams{
		HTTPClient: client,
	}
}

/* FindUserCustomdataParams contains all the parameters to send to the API endpoint
   for the find user customdata operation.

   Typically these are written to a http.Request.
*/
type FindUserCustomdataParams struct {

	/* ID.

	   User UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find user customdata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUserCustomdataParams) WithDefaults() *FindUserCustomdataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find user customdata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUserCustomdataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find user customdata params
func (o *FindUserCustomdataParams) WithTimeout(timeout time.Duration) *FindUserCustomdataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find user customdata params
func (o *FindUserCustomdataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find user customdata params
func (o *FindUserCustomdataParams) WithContext(ctx context.Context) *FindUserCustomdataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find user customdata params
func (o *FindUserCustomdataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find user customdata params
func (o *FindUserCustomdataParams) WithHTTPClient(client *http.Client) *FindUserCustomdataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find user customdata params
func (o *FindUserCustomdataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find user customdata params
func (o *FindUserCustomdataParams) WithID(id strfmt.UUID) *FindUserCustomdataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find user customdata params
func (o *FindUserCustomdataParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindUserCustomdataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}