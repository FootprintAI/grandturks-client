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

// NewKafeidoServicePutProjectPipelineParams creates a new KafeidoServicePutProjectPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServicePutProjectPipelineParams() *KafeidoServicePutProjectPipelineParams {
	return &KafeidoServicePutProjectPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServicePutProjectPipelineParamsWithTimeout creates a new KafeidoServicePutProjectPipelineParams object
// with the ability to set a timeout on a request.
func NewKafeidoServicePutProjectPipelineParamsWithTimeout(timeout time.Duration) *KafeidoServicePutProjectPipelineParams {
	return &KafeidoServicePutProjectPipelineParams{
		timeout: timeout,
	}
}

// NewKafeidoServicePutProjectPipelineParamsWithContext creates a new KafeidoServicePutProjectPipelineParams object
// with the ability to set a context for a request.
func NewKafeidoServicePutProjectPipelineParamsWithContext(ctx context.Context) *KafeidoServicePutProjectPipelineParams {
	return &KafeidoServicePutProjectPipelineParams{
		Context: ctx,
	}
}

// NewKafeidoServicePutProjectPipelineParamsWithHTTPClient creates a new KafeidoServicePutProjectPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServicePutProjectPipelineParamsWithHTTPClient(client *http.Client) *KafeidoServicePutProjectPipelineParams {
	return &KafeidoServicePutProjectPipelineParams{
		HTTPClient: client,
	}
}

/*
KafeidoServicePutProjectPipelineParams contains all the parameters to send to the API endpoint

	for the kafeido service put project pipeline operation.

	Typically these are written to a http.Request.
*/
type KafeidoServicePutProjectPipelineParams struct {

	// Body.
	Body KafeidoServicePutProjectPipelineBody

	// NamedPipelineID.
	NamedPipelineID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service put project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServicePutProjectPipelineParams) WithDefaults() *KafeidoServicePutProjectPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service put project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServicePutProjectPipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) WithTimeout(timeout time.Duration) *KafeidoServicePutProjectPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) WithContext(ctx context.Context) *KafeidoServicePutProjectPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) WithHTTPClient(client *http.Client) *KafeidoServicePutProjectPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) WithBody(body KafeidoServicePutProjectPipelineBody) *KafeidoServicePutProjectPipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) SetBody(body KafeidoServicePutProjectPipelineBody) {
	o.Body = body
}

// WithNamedPipelineID adds the namedPipelineID to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) WithNamedPipelineID(namedPipelineID string) *KafeidoServicePutProjectPipelineParams {
	o.SetNamedPipelineID(namedPipelineID)
	return o
}

// SetNamedPipelineID adds the namedPipelineId to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) SetNamedPipelineID(namedPipelineID string) {
	o.NamedPipelineID = namedPipelineID
}

// WithProjectID adds the projectID to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) WithProjectID(projectID string) *KafeidoServicePutProjectPipelineParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service put project pipeline params
func (o *KafeidoServicePutProjectPipelineParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServicePutProjectPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
