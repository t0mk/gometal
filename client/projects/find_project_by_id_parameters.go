// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewFindProjectByIDParams creates a new FindProjectByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindProjectByIDParams() *FindProjectByIDParams {
	return &FindProjectByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindProjectByIDParamsWithTimeout creates a new FindProjectByIDParams object
// with the ability to set a timeout on a request.
func NewFindProjectByIDParamsWithTimeout(timeout time.Duration) *FindProjectByIDParams {
	return &FindProjectByIDParams{
		timeout: timeout,
	}
}

// NewFindProjectByIDParamsWithContext creates a new FindProjectByIDParams object
// with the ability to set a context for a request.
func NewFindProjectByIDParamsWithContext(ctx context.Context) *FindProjectByIDParams {
	return &FindProjectByIDParams{
		Context: ctx,
	}
}

// NewFindProjectByIDParamsWithHTTPClient creates a new FindProjectByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindProjectByIDParamsWithHTTPClient(client *http.Client) *FindProjectByIDParams {
	return &FindProjectByIDParams{
		HTTPClient: client,
	}
}

/* FindProjectByIDParams contains all the parameters to send to the API endpoint
   for the find project by Id operation.

   Typically these are written to a http.Request.
*/
type FindProjectByIDParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* ID.

	   Project UUID

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

// WithDefaults hydrates default values in the find project by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectByIDParams) WithDefaults() *FindProjectByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find project by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find project by Id params
func (o *FindProjectByIDParams) WithTimeout(timeout time.Duration) *FindProjectByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find project by Id params
func (o *FindProjectByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find project by Id params
func (o *FindProjectByIDParams) WithContext(ctx context.Context) *FindProjectByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find project by Id params
func (o *FindProjectByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find project by Id params
func (o *FindProjectByIDParams) WithHTTPClient(client *http.Client) *FindProjectByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find project by Id params
func (o *FindProjectByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find project by Id params
func (o *FindProjectByIDParams) WithExclude(exclude []string) *FindProjectByIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find project by Id params
func (o *FindProjectByIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the find project by Id params
func (o *FindProjectByIDParams) WithID(id strfmt.UUID) *FindProjectByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find project by Id params
func (o *FindProjectByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find project by Id params
func (o *FindProjectByIDParams) WithInclude(include []string) *FindProjectByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find project by Id params
func (o *FindProjectByIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindProjectByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamFindProjectByID binds the parameter exclude
func (o *FindProjectByIDParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamFindProjectByID binds the parameter include
func (o *FindProjectByIDParams) bindParamInclude(formats strfmt.Registry) []string {
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