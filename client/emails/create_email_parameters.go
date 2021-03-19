// Code generated by go-swagger; DO NOT EDIT.

package emails

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

// NewCreateEmailParams creates a new CreateEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEmailParams() *CreateEmailParams {
	return &CreateEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEmailParamsWithTimeout creates a new CreateEmailParams object
// with the ability to set a timeout on a request.
func NewCreateEmailParamsWithTimeout(timeout time.Duration) *CreateEmailParams {
	return &CreateEmailParams{
		timeout: timeout,
	}
}

// NewCreateEmailParamsWithContext creates a new CreateEmailParams object
// with the ability to set a context for a request.
func NewCreateEmailParamsWithContext(ctx context.Context) *CreateEmailParams {
	return &CreateEmailParams{
		Context: ctx,
	}
}

// NewCreateEmailParamsWithHTTPClient creates a new CreateEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEmailParamsWithHTTPClient(client *http.Client) *CreateEmailParams {
	return &CreateEmailParams{
		HTTPClient: client,
	}
}

/* CreateEmailParams contains all the parameters to send to the API endpoint
   for the create email operation.

   Typically these are written to a http.Request.
*/
type CreateEmailParams struct {

	/* Email.

	   Email to create
	*/
	Email *types.CreateEmailInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEmailParams) WithDefaults() *CreateEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create email params
func (o *CreateEmailParams) WithTimeout(timeout time.Duration) *CreateEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create email params
func (o *CreateEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create email params
func (o *CreateEmailParams) WithContext(ctx context.Context) *CreateEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create email params
func (o *CreateEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create email params
func (o *CreateEmailParams) WithHTTPClient(client *http.Client) *CreateEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create email params
func (o *CreateEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the create email params
func (o *CreateEmailParams) WithEmail(email *types.CreateEmailInput) *CreateEmailParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the create email params
func (o *CreateEmailParams) SetEmail(email *types.CreateEmailInput) {
	o.Email = email
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Email != nil {
		if err := r.SetBodyParam(o.Email); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}