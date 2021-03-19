// Code generated by go-swagger; DO NOT EDIT.

package facilities

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

// NewFindFacilitiesParams creates a new FindFacilitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindFacilitiesParams() *FindFacilitiesParams {
	return &FindFacilitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindFacilitiesParamsWithTimeout creates a new FindFacilitiesParams object
// with the ability to set a timeout on a request.
func NewFindFacilitiesParamsWithTimeout(timeout time.Duration) *FindFacilitiesParams {
	return &FindFacilitiesParams{
		timeout: timeout,
	}
}

// NewFindFacilitiesParamsWithContext creates a new FindFacilitiesParams object
// with the ability to set a context for a request.
func NewFindFacilitiesParamsWithContext(ctx context.Context) *FindFacilitiesParams {
	return &FindFacilitiesParams{
		Context: ctx,
	}
}

// NewFindFacilitiesParamsWithHTTPClient creates a new FindFacilitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindFacilitiesParamsWithHTTPClient(client *http.Client) *FindFacilitiesParams {
	return &FindFacilitiesParams{
		HTTPClient: client,
	}
}

/* FindFacilitiesParams contains all the parameters to send to the API endpoint
   for the find facilities operation.

   Typically these are written to a http.Request.
*/
type FindFacilitiesParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.

	   Default: ["address"]
	*/
	Exclude []string

	/* Include.

	   Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects.
	*/
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find facilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindFacilitiesParams) WithDefaults() *FindFacilitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find facilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindFacilitiesParams) SetDefaults() {
	var (
		excludeDefault = []string{"address"}
	)

	val := FindFacilitiesParams{
		Exclude: excludeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find facilities params
func (o *FindFacilitiesParams) WithTimeout(timeout time.Duration) *FindFacilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find facilities params
func (o *FindFacilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find facilities params
func (o *FindFacilitiesParams) WithContext(ctx context.Context) *FindFacilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find facilities params
func (o *FindFacilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find facilities params
func (o *FindFacilitiesParams) WithHTTPClient(client *http.Client) *FindFacilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find facilities params
func (o *FindFacilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find facilities params
func (o *FindFacilitiesParams) WithExclude(exclude []string) *FindFacilitiesParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find facilities params
func (o *FindFacilitiesParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the find facilities params
func (o *FindFacilitiesParams) WithInclude(include []string) *FindFacilitiesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find facilities params
func (o *FindFacilitiesParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindFacilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamFindFacilities binds the parameter exclude
func (o *FindFacilitiesParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamFindFacilities binds the parameter include
func (o *FindFacilitiesParams) bindParamInclude(formats strfmt.Registry) []string {
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