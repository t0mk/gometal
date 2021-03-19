// Code generated by go-swagger; DO NOT EDIT.

package ports

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
	"github.com/go-openapi/swag"
)

// NewDisbondPortParams creates a new DisbondPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisbondPortParams() *DisbondPortParams {
	return &DisbondPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisbondPortParamsWithTimeout creates a new DisbondPortParams object
// with the ability to set a timeout on a request.
func NewDisbondPortParamsWithTimeout(timeout time.Duration) *DisbondPortParams {
	return &DisbondPortParams{
		timeout: timeout,
	}
}

// NewDisbondPortParamsWithContext creates a new DisbondPortParams object
// with the ability to set a context for a request.
func NewDisbondPortParamsWithContext(ctx context.Context) *DisbondPortParams {
	return &DisbondPortParams{
		Context: ctx,
	}
}

// NewDisbondPortParamsWithHTTPClient creates a new DisbondPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisbondPortParamsWithHTTPClient(client *http.Client) *DisbondPortParams {
	return &DisbondPortParams{
		HTTPClient: client,
	}
}

/* DisbondPortParams contains all the parameters to send to the API endpoint
   for the disbond port operation.

   Typically these are written to a http.Request.
*/
type DisbondPortParams struct {

	/* BulkDisable.

	   disable both ports
	*/
	BulkDisable *bool

	/* ID.

	   Port UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disbond port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisbondPortParams) WithDefaults() *DisbondPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disbond port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisbondPortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disbond port params
func (o *DisbondPortParams) WithTimeout(timeout time.Duration) *DisbondPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disbond port params
func (o *DisbondPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disbond port params
func (o *DisbondPortParams) WithContext(ctx context.Context) *DisbondPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disbond port params
func (o *DisbondPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disbond port params
func (o *DisbondPortParams) WithHTTPClient(client *http.Client) *DisbondPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disbond port params
func (o *DisbondPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBulkDisable adds the bulkDisable to the disbond port params
func (o *DisbondPortParams) WithBulkDisable(bulkDisable *bool) *DisbondPortParams {
	o.SetBulkDisable(bulkDisable)
	return o
}

// SetBulkDisable adds the bulkDisable to the disbond port params
func (o *DisbondPortParams) SetBulkDisable(bulkDisable *bool) {
	o.BulkDisable = bulkDisable
}

// WithID adds the id to the disbond port params
func (o *DisbondPortParams) WithID(id strfmt.UUID) *DisbondPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the disbond port params
func (o *DisbondPortParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DisbondPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BulkDisable != nil {

		// query param bulk_disable
		var qrBulkDisable bool

		if o.BulkDisable != nil {
			qrBulkDisable = *o.BulkDisable
		}
		qBulkDisable := swag.FormatBool(qrBulkDisable)
		if qBulkDisable != "" {

			if err := r.SetQueryParam("bulk_disable", qBulkDisable); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}