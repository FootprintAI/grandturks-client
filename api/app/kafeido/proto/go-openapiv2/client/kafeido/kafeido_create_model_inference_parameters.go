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

// NewKafeidoCreateModelInferenceParams creates a new KafeidoCreateModelInferenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoCreateModelInferenceParams() *KafeidoCreateModelInferenceParams {
	return &KafeidoCreateModelInferenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoCreateModelInferenceParamsWithTimeout creates a new KafeidoCreateModelInferenceParams object
// with the ability to set a timeout on a request.
func NewKafeidoCreateModelInferenceParamsWithTimeout(timeout time.Duration) *KafeidoCreateModelInferenceParams {
	return &KafeidoCreateModelInferenceParams{
		timeout: timeout,
	}
}

// NewKafeidoCreateModelInferenceParamsWithContext creates a new KafeidoCreateModelInferenceParams object
// with the ability to set a context for a request.
func NewKafeidoCreateModelInferenceParamsWithContext(ctx context.Context) *KafeidoCreateModelInferenceParams {
	return &KafeidoCreateModelInferenceParams{
		Context: ctx,
	}
}

// NewKafeidoCreateModelInferenceParamsWithHTTPClient creates a new KafeidoCreateModelInferenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoCreateModelInferenceParamsWithHTTPClient(client *http.Client) *KafeidoCreateModelInferenceParams {
	return &KafeidoCreateModelInferenceParams{
		HTTPClient: client,
	}
}

/*
KafeidoCreateModelInferenceParams contains all the parameters to send to the API endpoint

	for the kafeido create model inference operation.

	Typically these are written to a http.Request.
*/
type KafeidoCreateModelInferenceParams struct {

	// Body.
	Body KafeidoCreateModelInferenceBody

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido create model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoCreateModelInferenceParams) WithDefaults() *KafeidoCreateModelInferenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido create model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoCreateModelInferenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) WithTimeout(timeout time.Duration) *KafeidoCreateModelInferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) WithContext(ctx context.Context) *KafeidoCreateModelInferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) WithHTTPClient(client *http.Client) *KafeidoCreateModelInferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) WithBody(body KafeidoCreateModelInferenceBody) *KafeidoCreateModelInferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) SetBody(body KafeidoCreateModelInferenceBody) {
	o.Body = body
}

// WithProjectID adds the projectID to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) WithProjectID(projectID string) *KafeidoCreateModelInferenceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido create model inference params
func (o *KafeidoCreateModelInferenceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoCreateModelInferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
