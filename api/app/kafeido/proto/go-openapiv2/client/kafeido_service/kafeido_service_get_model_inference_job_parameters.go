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

// NewKafeidoServiceGetModelInferenceJobParams creates a new KafeidoServiceGetModelInferenceJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceGetModelInferenceJobParams() *KafeidoServiceGetModelInferenceJobParams {
	return &KafeidoServiceGetModelInferenceJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceGetModelInferenceJobParamsWithTimeout creates a new KafeidoServiceGetModelInferenceJobParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceGetModelInferenceJobParamsWithTimeout(timeout time.Duration) *KafeidoServiceGetModelInferenceJobParams {
	return &KafeidoServiceGetModelInferenceJobParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceGetModelInferenceJobParamsWithContext creates a new KafeidoServiceGetModelInferenceJobParams object
// with the ability to set a context for a request.
func NewKafeidoServiceGetModelInferenceJobParamsWithContext(ctx context.Context) *KafeidoServiceGetModelInferenceJobParams {
	return &KafeidoServiceGetModelInferenceJobParams{
		Context: ctx,
	}
}

// NewKafeidoServiceGetModelInferenceJobParamsWithHTTPClient creates a new KafeidoServiceGetModelInferenceJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceGetModelInferenceJobParamsWithHTTPClient(client *http.Client) *KafeidoServiceGetModelInferenceJobParams {
	return &KafeidoServiceGetModelInferenceJobParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceGetModelInferenceJobParams contains all the parameters to send to the API endpoint

	for the kafeido service get model inference job operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceGetModelInferenceJobParams struct {

	// ModelInferenceJobID.
	ModelInferenceJobID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service get model inference job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceGetModelInferenceJobParams) WithDefaults() *KafeidoServiceGetModelInferenceJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service get model inference job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceGetModelInferenceJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) WithTimeout(timeout time.Duration) *KafeidoServiceGetModelInferenceJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) WithContext(ctx context.Context) *KafeidoServiceGetModelInferenceJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) WithHTTPClient(client *http.Client) *KafeidoServiceGetModelInferenceJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelInferenceJobID adds the modelInferenceJobID to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) WithModelInferenceJobID(modelInferenceJobID string) *KafeidoServiceGetModelInferenceJobParams {
	o.SetModelInferenceJobID(modelInferenceJobID)
	return o
}

// SetModelInferenceJobID adds the modelInferenceJobId to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) SetModelInferenceJobID(modelInferenceJobID string) {
	o.ModelInferenceJobID = modelInferenceJobID
}

// WithProjectID adds the projectID to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) WithProjectID(projectID string) *KafeidoServiceGetModelInferenceJobParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service get model inference job params
func (o *KafeidoServiceGetModelInferenceJobParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceGetModelInferenceJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modelInferenceJobId
	if err := r.SetPathParam("modelInferenceJobId", o.ModelInferenceJobID); err != nil {
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
