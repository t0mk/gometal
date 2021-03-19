// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewFindEventsParams creates a new FindEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindEventsParams() *FindEventsParams {
	return &FindEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindEventsParamsWithTimeout creates a new FindEventsParams object
// with the ability to set a timeout on a request.
func NewFindEventsParamsWithTimeout(timeout time.Duration) *FindEventsParams {
	return &FindEventsParams{
		timeout: timeout,
	}
}

// NewFindEventsParamsWithContext creates a new FindEventsParams object
// with the ability to set a context for a request.
func NewFindEventsParamsWithContext(ctx context.Context) *FindEventsParams {
	return &FindEventsParams{
		Context: ctx,
	}
}

// NewFindEventsParamsWithHTTPClient creates a new FindEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindEventsParamsWithHTTPClient(client *http.Client) *FindEventsParams {
	return &FindEventsParams{
		HTTPClient: client,
	}
}

/* FindEventsParams contains all the parameters to send to the API endpoint
   for the find events operation.

   Typically these are written to a http.Request.
*/
type FindEventsParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* Include.

	   Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects.
	*/
	Include []string

	/* Page.

	   Page to return

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* PerPage.

	   Items returned per page

	   Format: int32
	   Default: 10
	*/
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindEventsParams) WithDefaults() *FindEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindEventsParams) SetDefaults() {
	var (
		pageDefault = int32(1)

		perPageDefault = int32(10)
	)

	val := FindEventsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find events params
func (o *FindEventsParams) WithTimeout(timeout time.Duration) *FindEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find events params
func (o *FindEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find events params
func (o *FindEventsParams) WithContext(ctx context.Context) *FindEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find events params
func (o *FindEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find events params
func (o *FindEventsParams) WithHTTPClient(client *http.Client) *FindEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find events params
func (o *FindEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find events params
func (o *FindEventsParams) WithExclude(exclude []string) *FindEventsParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find events params
func (o *FindEventsParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the find events params
func (o *FindEventsParams) WithInclude(include []string) *FindEventsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find events params
func (o *FindEventsParams) SetInclude(include []string) {
	o.Include = include
}

// WithPage adds the page to the find events params
func (o *FindEventsParams) WithPage(page *int32) *FindEventsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the find events params
func (o *FindEventsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the find events params
func (o *FindEventsParams) WithPerPage(perPage *int32) *FindEventsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the find events params
func (o *FindEventsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *FindEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFindEvents binds the parameter exclude
func (o *FindEventsParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamFindEvents binds the parameter include
func (o *FindEventsParams) bindParamInclude(formats strfmt.Registry) []string {
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