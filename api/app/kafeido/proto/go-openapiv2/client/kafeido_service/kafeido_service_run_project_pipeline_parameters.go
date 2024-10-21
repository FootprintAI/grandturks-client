// Code generated by go-swagger; DO NOT EDIT.

package kafeido_service

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

// NewKafeidoServiceRunProjectPipelineParams creates a new KafeidoServiceRunProjectPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceRunProjectPipelineParams() *KafeidoServiceRunProjectPipelineParams {
	return &KafeidoServiceRunProjectPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceRunProjectPipelineParamsWithTimeout creates a new KafeidoServiceRunProjectPipelineParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceRunProjectPipelineParamsWithTimeout(timeout time.Duration) *KafeidoServiceRunProjectPipelineParams {
	return &KafeidoServiceRunProjectPipelineParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceRunProjectPipelineParamsWithContext creates a new KafeidoServiceRunProjectPipelineParams object
// with the ability to set a context for a request.
func NewKafeidoServiceRunProjectPipelineParamsWithContext(ctx context.Context) *KafeidoServiceRunProjectPipelineParams {
	return &KafeidoServiceRunProjectPipelineParams{
		Context: ctx,
	}
}

// NewKafeidoServiceRunProjectPipelineParamsWithHTTPClient creates a new KafeidoServiceRunProjectPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceRunProjectPipelineParamsWithHTTPClient(client *http.Client) *KafeidoServiceRunProjectPipelineParams {
	return &KafeidoServiceRunProjectPipelineParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceRunProjectPipelineParams contains all the parameters to send to the API endpoint

	for the kafeido service run project pipeline operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceRunProjectPipelineParams struct {

	// Body.
	Body KafeidoServiceRunProjectPipelineBody

	// NamedPipelineID.
	NamedPipelineID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service run project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceRunProjectPipelineParams) WithDefaults() *KafeidoServiceRunProjectPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service run project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceRunProjectPipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) WithTimeout(timeout time.Duration) *KafeidoServiceRunProjectPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) WithContext(ctx context.Context) *KafeidoServiceRunProjectPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) WithHTTPClient(client *http.Client) *KafeidoServiceRunProjectPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) WithBody(body KafeidoServiceRunProjectPipelineBody) *KafeidoServiceRunProjectPipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) SetBody(body KafeidoServiceRunProjectPipelineBody) {
	o.Body = body
}

// WithNamedPipelineID adds the namedPipelineID to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) WithNamedPipelineID(namedPipelineID string) *KafeidoServiceRunProjectPipelineParams {
	o.SetNamedPipelineID(namedPipelineID)
	return o
}

// SetNamedPipelineID adds the namedPipelineId to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) SetNamedPipelineID(namedPipelineID string) {
	o.NamedPipelineID = namedPipelineID
}

// WithProjectID adds the projectID to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) WithProjectID(projectID string) *KafeidoServiceRunProjectPipelineParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service run project pipeline params
func (o *KafeidoServiceRunProjectPipelineParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceRunProjectPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
