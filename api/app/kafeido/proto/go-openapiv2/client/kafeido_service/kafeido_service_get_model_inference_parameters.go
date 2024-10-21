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

// NewKafeidoServiceGetModelInferenceParams creates a new KafeidoServiceGetModelInferenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceGetModelInferenceParams() *KafeidoServiceGetModelInferenceParams {
	return &KafeidoServiceGetModelInferenceParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceGetModelInferenceParamsWithTimeout creates a new KafeidoServiceGetModelInferenceParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceGetModelInferenceParamsWithTimeout(timeout time.Duration) *KafeidoServiceGetModelInferenceParams {
	return &KafeidoServiceGetModelInferenceParams{
		requestTimeout: timeout,
	}
}

// NewKafeidoServiceGetModelInferenceParamsWithContext creates a new KafeidoServiceGetModelInferenceParams object
// with the ability to set a context for a request.
func NewKafeidoServiceGetModelInferenceParamsWithContext(ctx context.Context) *KafeidoServiceGetModelInferenceParams {
	return &KafeidoServiceGetModelInferenceParams{
		Context: ctx,
	}
}

// NewKafeidoServiceGetModelInferenceParamsWithHTTPClient creates a new KafeidoServiceGetModelInferenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceGetModelInferenceParamsWithHTTPClient(client *http.Client) *KafeidoServiceGetModelInferenceParams {
	return &KafeidoServiceGetModelInferenceParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceGetModelInferenceParams contains all the parameters to send to the API endpoint

	for the kafeido service get model inference operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceGetModelInferenceParams struct {

	// ModelInferenceID.
	ModelInferenceID string

	// ProjectID.
	ProjectID string

	// Timeout.
	Timeout *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the kafeido service get model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceGetModelInferenceParams) WithDefaults() *KafeidoServiceGetModelInferenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service get model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceGetModelInferenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithRequestTimeout adds the timeout to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) WithRequestTimeout(timeout time.Duration) *KafeidoServiceGetModelInferenceParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) WithContext(ctx context.Context) *KafeidoServiceGetModelInferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) WithHTTPClient(client *http.Client) *KafeidoServiceGetModelInferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelInferenceID adds the modelInferenceID to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) WithModelInferenceID(modelInferenceID string) *KafeidoServiceGetModelInferenceParams {
	o.SetModelInferenceID(modelInferenceID)
	return o
}

// SetModelInferenceID adds the modelInferenceId to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) SetModelInferenceID(modelInferenceID string) {
	o.ModelInferenceID = modelInferenceID
}

// WithProjectID adds the projectID to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) WithProjectID(projectID string) *KafeidoServiceGetModelInferenceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithTimeout adds the timeout to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) WithTimeout(timeout *string) *KafeidoServiceGetModelInferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service get model inference params
func (o *KafeidoServiceGetModelInferenceParams) SetTimeout(timeout *string) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceGetModelInferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// path param modelInferenceId
	if err := r.SetPathParam("modelInferenceId", o.ModelInferenceID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout string

		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := qrTimeout
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
