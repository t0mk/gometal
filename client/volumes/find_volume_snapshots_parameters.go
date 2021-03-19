// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewFindVolumeSnapshotsParams creates a new FindVolumeSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindVolumeSnapshotsParams() *FindVolumeSnapshotsParams {
	return &FindVolumeSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindVolumeSnapshotsParamsWithTimeout creates a new FindVolumeSnapshotsParams object
// with the ability to set a timeout on a request.
func NewFindVolumeSnapshotsParamsWithTimeout(timeout time.Duration) *FindVolumeSnapshotsParams {
	return &FindVolumeSnapshotsParams{
		timeout: timeout,
	}
}

// NewFindVolumeSnapshotsParamsWithContext creates a new FindVolumeSnapshotsParams object
// with the ability to set a context for a request.
func NewFindVolumeSnapshotsParamsWithContext(ctx context.Context) *FindVolumeSnapshotsParams {
	return &FindVolumeSnapshotsParams{
		Context: ctx,
	}
}

// NewFindVolumeSnapshotsParamsWithHTTPClient creates a new FindVolumeSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindVolumeSnapshotsParamsWithHTTPClient(client *http.Client) *FindVolumeSnapshotsParams {
	return &FindVolumeSnapshotsParams{
		HTTPClient: client,
	}
}

/* FindVolumeSnapshotsParams contains all the parameters to send to the API endpoint
   for the find volume snapshots operation.

   Typically these are written to a http.Request.
*/
type FindVolumeSnapshotsParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* ID.

	   Volume UUID

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

// WithDefaults hydrates default values in the find volume snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindVolumeSnapshotsParams) WithDefaults() *FindVolumeSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find volume snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindVolumeSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) WithTimeout(timeout time.Duration) *FindVolumeSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) WithContext(ctx context.Context) *FindVolumeSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) WithHTTPClient(client *http.Client) *FindVolumeSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) WithExclude(exclude []string) *FindVolumeSnapshotsParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) WithID(id strfmt.UUID) *FindVolumeSnapshotsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) WithInclude(include []string) *FindVolumeSnapshotsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find volume snapshots params
func (o *FindVolumeSnapshotsParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindVolumeSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamFindVolumeSnapshots binds the parameter exclude
func (o *FindVolumeSnapshotsParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamFindVolumeSnapshots binds the parameter include
func (o *FindVolumeSnapshotsParams) bindParamInclude(formats strfmt.Registry) []string {
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