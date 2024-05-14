// Code generated by go-swagger; DO NOT EDIT.

package kafeido

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

// NewKafeidoRunProjectPipelineParams creates a new KafeidoRunProjectPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoRunProjectPipelineParams() *KafeidoRunProjectPipelineParams {
	return &KafeidoRunProjectPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoRunProjectPipelineParamsWithTimeout creates a new KafeidoRunProjectPipelineParams object
// with the ability to set a timeout on a request.
func NewKafeidoRunProjectPipelineParamsWithTimeout(timeout time.Duration) *KafeidoRunProjectPipelineParams {
	return &KafeidoRunProjectPipelineParams{
		timeout: timeout,
	}
}

// NewKafeidoRunProjectPipelineParamsWithContext creates a new KafeidoRunProjectPipelineParams object
// with the ability to set a context for a request.
func NewKafeidoRunProjectPipelineParamsWithContext(ctx context.Context) *KafeidoRunProjectPipelineParams {
	return &KafeidoRunProjectPipelineParams{
		Context: ctx,
	}
}

// NewKafeidoRunProjectPipelineParamsWithHTTPClient creates a new KafeidoRunProjectPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoRunProjectPipelineParamsWithHTTPClient(client *http.Client) *KafeidoRunProjectPipelineParams {
	return &KafeidoRunProjectPipelineParams{
		HTTPClient: client,
	}
}

/*
KafeidoRunProjectPipelineParams contains all the parameters to send to the API endpoint

	for the kafeido run project pipeline operation.

	Typically these are written to a http.Request.
*/
type KafeidoRunProjectPipelineParams struct {

	// Body.
	Body KafeidoRunProjectPipelineBody

	// NamedPipelineID.
	NamedPipelineID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido run project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoRunProjectPipelineParams) WithDefaults() *KafeidoRunProjectPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido run project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoRunProjectPipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) WithTimeout(timeout time.Duration) *KafeidoRunProjectPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) WithContext(ctx context.Context) *KafeidoRunProjectPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) WithHTTPClient(client *http.Client) *KafeidoRunProjectPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) WithBody(body KafeidoRunProjectPipelineBody) *KafeidoRunProjectPipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) SetBody(body KafeidoRunProjectPipelineBody) {
	o.Body = body
}

// WithNamedPipelineID adds the namedPipelineID to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) WithNamedPipelineID(namedPipelineID string) *KafeidoRunProjectPipelineParams {
	o.SetNamedPipelineID(namedPipelineID)
	return o
}

// SetNamedPipelineID adds the namedPipelineId to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) SetNamedPipelineID(namedPipelineID string) {
	o.NamedPipelineID = namedPipelineID
}

// WithProjectID adds the projectID to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) WithProjectID(projectID string) *KafeidoRunProjectPipelineParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido run project pipeline params
func (o *KafeidoRunProjectPipelineParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoRunProjectPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param namedPipelineId
	if err := r.SetPathParam("namedPipelineId", o.NamedPipelineID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}