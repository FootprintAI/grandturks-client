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

// NewKafeidoGetProjectPipelineParams creates a new KafeidoGetProjectPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoGetProjectPipelineParams() *KafeidoGetProjectPipelineParams {
	return &KafeidoGetProjectPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoGetProjectPipelineParamsWithTimeout creates a new KafeidoGetProjectPipelineParams object
// with the ability to set a timeout on a request.
func NewKafeidoGetProjectPipelineParamsWithTimeout(timeout time.Duration) *KafeidoGetProjectPipelineParams {
	return &KafeidoGetProjectPipelineParams{
		timeout: timeout,
	}
}

// NewKafeidoGetProjectPipelineParamsWithContext creates a new KafeidoGetProjectPipelineParams object
// with the ability to set a context for a request.
func NewKafeidoGetProjectPipelineParamsWithContext(ctx context.Context) *KafeidoGetProjectPipelineParams {
	return &KafeidoGetProjectPipelineParams{
		Context: ctx,
	}
}

// NewKafeidoGetProjectPipelineParamsWithHTTPClient creates a new KafeidoGetProjectPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoGetProjectPipelineParamsWithHTTPClient(client *http.Client) *KafeidoGetProjectPipelineParams {
	return &KafeidoGetProjectPipelineParams{
		HTTPClient: client,
	}
}

/*
KafeidoGetProjectPipelineParams contains all the parameters to send to the API endpoint

	for the kafeido get project pipeline operation.

	Typically these are written to a http.Request.
*/
type KafeidoGetProjectPipelineParams struct {

	// NamedPipelineID.
	NamedPipelineID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido get project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoGetProjectPipelineParams) WithDefaults() *KafeidoGetProjectPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido get project pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoGetProjectPipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) WithTimeout(timeout time.Duration) *KafeidoGetProjectPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) WithContext(ctx context.Context) *KafeidoGetProjectPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) WithHTTPClient(client *http.Client) *KafeidoGetProjectPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamedPipelineID adds the namedPipelineID to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) WithNamedPipelineID(namedPipelineID string) *KafeidoGetProjectPipelineParams {
	o.SetNamedPipelineID(namedPipelineID)
	return o
}

// SetNamedPipelineID adds the namedPipelineId to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) SetNamedPipelineID(namedPipelineID string) {
	o.NamedPipelineID = namedPipelineID
}

// WithProjectID adds the projectID to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) WithProjectID(projectID string) *KafeidoGetProjectPipelineParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido get project pipeline params
func (o *KafeidoGetProjectPipelineParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoGetProjectPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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