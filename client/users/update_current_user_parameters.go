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

	"github.com/t0mk/gometal/types"
)

// NewUpdateCurrentUserParams creates a new UpdateCurrentUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCurrentUserParams() *UpdateCurrentUserParams {
	return &UpdateCurrentUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCurrentUserParamsWithTimeout creates a new UpdateCurrentUserParams object
// with the ability to set a timeout on a request.
func NewUpdateCurrentUserParamsWithTimeout(timeout time.Duration) *UpdateCurrentUserParams {
	return &UpdateCurrentUserParams{
		timeout: timeout,
	}
}

// NewUpdateCurrentUserParamsWithContext creates a new UpdateCurrentUserParams object
// with the ability to set a context for a request.
func NewUpdateCurrentUserParamsWithContext(ctx context.Context) *UpdateCurrentUserParams {
	return &UpdateCurrentUserParams{
		Context: ctx,
	}
}

// NewUpdateCurrentUserParamsWithHTTPClient creates a new UpdateCurrentUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCurrentUserParamsWithHTTPClient(client *http.Client) *UpdateCurrentUserParams {
	return &UpdateCurrentUserParams{
		HTTPClient: client,
	}
}

/* UpdateCurrentUserParams contains all the parameters to send to the API endpoint
   for the update current user operation.

   Typically these are written to a http.Request.
*/
type UpdateCurrentUserParams struct {

	/* User.

	   User to update
	*/
	User *types.UserUpdateInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update current user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentUserParams) WithDefaults() *UpdateCurrentUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update current user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update current user params
func (o *UpdateCurrentUserParams) WithTimeout(timeout time.Duration) *UpdateCurrentUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update current user params
func (o *UpdateCurrentUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update current user params
func (o *UpdateCurrentUserParams) WithContext(ctx context.Context) *UpdateCurrentUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update current user params
func (o *UpdateCurrentUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update current user params
func (o *UpdateCurrentUserParams) WithHTTPClient(client *http.Client) *UpdateCurrentUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update current user params
func (o *UpdateCurrentUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the update current user params
func (o *UpdateCurrentUserParams) WithUser(user *types.UserUpdateInput) *UpdateCurrentUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the update current user params
func (o *UpdateCurrentUserParams) SetUser(user *types.UserUpdateInput) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCurrentUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.User != nil {
		if err := r.SetBodyParam(o.User); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}