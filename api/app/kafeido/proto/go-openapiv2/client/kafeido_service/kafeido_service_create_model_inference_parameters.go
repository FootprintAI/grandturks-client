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

// NewKafeidoServiceCreateModelInferenceParams creates a new KafeidoServiceCreateModelInferenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceCreateModelInferenceParams() *KafeidoServiceCreateModelInferenceParams {
	return &KafeidoServiceCreateModelInferenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceCreateModelInferenceParamsWithTimeout creates a new KafeidoServiceCreateModelInferenceParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceCreateModelInferenceParamsWithTimeout(timeout time.Duration) *KafeidoServiceCreateModelInferenceParams {
	return &KafeidoServiceCreateModelInferenceParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceCreateModelInferenceParamsWithContext creates a new KafeidoServiceCreateModelInferenceParams object
// with the ability to set a context for a request.
func NewKafeidoServiceCreateModelInferenceParamsWithContext(ctx context.Context) *KafeidoServiceCreateModelInferenceParams {
	return &KafeidoServiceCreateModelInferenceParams{
		Context: ctx,
	}
}

// NewKafeidoServiceCreateModelInferenceParamsWithHTTPClient creates a new KafeidoServiceCreateModelInferenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceCreateModelInferenceParamsWithHTTPClient(client *http.Client) *KafeidoServiceCreateModelInferenceParams {
	return &KafeidoServiceCreateModelInferenceParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceCreateModelInferenceParams contains all the parameters to send to the API endpoint

	for the kafeido service create model inference operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceCreateModelInferenceParams struct {

	// Body.
	Body KafeidoServiceCreateModelInferenceBody

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service create model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceCreateModelInferenceParams) WithDefaults() *KafeidoServiceCreateModelInferenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service create model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceCreateModelInferenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) WithTimeout(timeout time.Duration) *KafeidoServiceCreateModelInferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) WithContext(ctx context.Context) *KafeidoServiceCreateModelInferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) WithHTTPClient(client *http.Client) *KafeidoServiceCreateModelInferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) WithBody(body KafeidoServiceCreateModelInferenceBody) *KafeidoServiceCreateModelInferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) SetBody(body KafeidoServiceCreateModelInferenceBody) {
	o.Body = body
}

// WithProjectID adds the projectID to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) WithProjectID(projectID string) *KafeidoServiceCreateModelInferenceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service create model inference params
func (o *KafeidoServiceCreateModelInferenceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceCreateModelInferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
