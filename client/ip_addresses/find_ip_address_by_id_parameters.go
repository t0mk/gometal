// Code generated by go-swagger; DO NOT EDIT.

package ip_addresses

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

// NewFindIPAddressByIDParams creates a new FindIPAddressByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindIPAddressByIDParams() *FindIPAddressByIDParams {
	return &FindIPAddressByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindIPAddressByIDParamsWithTimeout creates a new FindIPAddressByIDParams object
// with the ability to set a timeout on a request.
func NewFindIPAddressByIDParamsWithTimeout(timeout time.Duration) *FindIPAddressByIDParams {
	return &FindIPAddressByIDParams{
		timeout: timeout,
	}
}

// NewFindIPAddressByIDParamsWithContext creates a new FindIPAddressByIDParams object
// with the ability to set a context for a request.
func NewFindIPAddressByIDParamsWithContext(ctx context.Context) *FindIPAddressByIDParams {
	return &FindIPAddressByIDParams{
		Context: ctx,
	}
}

// NewFindIPAddressByIDParamsWithHTTPClient creates a new FindIPAddressByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindIPAddressByIDParamsWithHTTPClient(client *http.Client) *FindIPAddressByIDParams {
	return &FindIPAddressByIDParams{
		HTTPClient: client,
	}
}

/* FindIPAddressByIDParams contains all the parameters to send to the API endpoint
   for the find IP address by Id operation.

   Typically these are written to a http.Request.
*/
type FindIPAddressByIDParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* ID.

	   IP Address UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Include.

	   Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects.
	*/
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find IP address by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindIPAddressByIDParams) WithDefaults() *FindIPAddressByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find IP address by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindIPAddressByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find IP address by Id params
func (o *FindIPAddressByIDParams) WithTimeout(timeout time.Duration) *FindIPAddressByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find IP address by Id params
func (o *FindIPAddressByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find IP address by Id params
func (o *FindIPAddressByIDParams) WithContext(ctx context.Context) *FindIPAddressByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find IP address by Id params
func (o *FindIPAddressByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find IP address by Id params
func (o *FindIPAddressByIDParams) WithHTTPClient(client *http.Client) *FindIPAddressByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find IP address by Id params
func (o *FindIPAddressByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find IP address by Id params
func (o *FindIPAddressByIDParams) WithExclude(exclude []string) *FindIPAddressByIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find IP address by Id params
func (o *FindIPAddressByIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the find IP address by Id params
func (o *FindIPAddressByIDParams) WithID(id strfmt.UUID) *FindIPAddressByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find IP address by Id params
func (o *FindIPAddressByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find IP address by Id params
func (o *FindIPAddressByIDParams) WithInclude(include []string) *FindIPAddressByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find IP address by Id params
func (o *FindIPAddressByIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindIPAddressByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exclude != nil {

		// binding items for exclude
		joinedExclude := o.bindParamExclude(reg)

		// query array param exclude
		if err := r.SetQueryParam("exclude", joinedExclude...); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// binding items for include
		joinedInclude := o.bindParamInclude(reg)

		// query array param include
		if err := r.SetQueryParam("include", joinedInclude...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFindIPAddressByID binds the parameter exclude
func (o *FindIPAddressByIDParams) bindParamExclude(formats strfmt.Registry) []string {
	excludeIR := o.Exclude

	var excludeIC []string
	for _, excludeIIR := range excludeIR { // explode []string

		excludeIIV := excludeIIR // string as string
		excludeIC = append(excludeIC, excludeIIV)
	}

	// items.CollectionFormat: "csv"
	excludeIS := swag.JoinByFormat(excludeIC, "csv")

	return excludeIS
}

// bindParamFindIPAddressByID binds the parameter include
func (o *FindIPAddressByIDParams) bindParamInclude(formats strfmt.Registry) []string {
	includeIR := o.Include

	var includeIC []string
	for _, includeIIR := range includeIR { // explode []string

		includeIIV := includeIIR // string as string
		includeIC = append(includeIC, includeIIV)
	}

	// items.CollectionFormat: "csv"
	includeIS := swag.JoinByFormat(includeIC, "csv")

	return includeIS
}