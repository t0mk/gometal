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

// NewCreateProjectInterconnectionParams creates a new CreateProjectInterconnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProjectInterconnectionParams() *CreateProjectInterconnectionParams {
	return &CreateProjectInterconnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectInterconnectionParamsWithTimeout creates a new CreateProjectInterconnectionParams object
// with the ability to set a timeout on a request.
func NewCreateProjectInterconnectionParamsWithTimeout(timeout time.Duration) *CreateProjectInterconnectionParams {
	return &CreateProjectInterconnectionParams{
		timeout: timeout,
	}
}

// NewCreateProjectInterconnectionParamsWithContext creates a new CreateProjectInterconnectionParams object
// with the ability to set a context for a request.
func NewCreateProjectInterconnectionParamsWithContext(ctx context.Context) *CreateProjectInterconnectionParams {
	return &CreateProjectInterconnectionParams{
		Context: ctx,
	}
}

// NewCreateProjectInterconnectionParamsWithHTTPClient creates a new CreateProjectInterconnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProjectInterconnectionParamsWithHTTPClient(client *http.Client) *CreateProjectInterconnectionParams {
	return &CreateProjectInterconnectionParams{
		HTTPClient: client,
	}
}

/* CreateProjectInterconnectionParams contains all the parameters to send to the API endpoint
   for the create project interconnection operation.

   Typically these are written to a http.Request.
*/
type CreateProjectInterconnectionParams struct {

	/* Connection.

	   Connection details
	*/
	Connection *types.InterconnectionCreateInput

	/* ProjectID.

	   UUID of the project

	   Format: uuid
	*/
	ProjectID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create project interconnection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectInterconnectionParams) WithDefaults() *CreateProjectInterconnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project interconnection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectInterconnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create project interconnection params
func (o *CreateProjectInterconnectionParams) WithTimeout(timeout time.Duration) *CreateProjectInterconnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project interconnection params
func (o *CreateProjectInterconnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project interconnection params
func (o *CreateProjectInterconnectionParams) WithContext(ctx context.Context) *CreateProjectInterconnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project interconnection params
func (o *CreateProjectInterconnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project interconnection params
func (o *CreateProjectInterconnectionParams) WithHTTPClient(client *http.Client) *CreateProjectInterconnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project interconnection params
func (o *CreateProjectInterconnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnection adds the connection to the create project interconnection params
func (o *CreateProjectInterconnectionParams) WithConnection(connection *types.InterconnectionCreateInput) *CreateProjectInterconnectionParams {
	o.SetConnection(connection)
	return o
}

// SetConnection adds the connection to the create project interconnection params
func (o *CreateProjectInterconnectionParams) SetConnection(connection *types.InterconnectionCreateInput) {
	o.Connection = connection
}

// WithProjectID adds the projectID to the create project interconnection params
func (o *CreateProjectInterconnectionParams) WithProjectID(projectID strfmt.UUID) *CreateProjectInterconnectionParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create project interconnection params
func (o *CreateProjectInterconnectionParams) SetProjectID(projectID strfmt.UUID) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectInterconnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Connection != nil {
		if err := r.SetBodyParam(o.Connection); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}