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

// NewKafeidoServiceDeleteModelInferenceParams creates a new KafeidoServiceDeleteModelInferenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceDeleteModelInferenceParams() *KafeidoServiceDeleteModelInferenceParams {
	return &KafeidoServiceDeleteModelInferenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceDeleteModelInferenceParamsWithTimeout creates a new KafeidoServiceDeleteModelInferenceParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceDeleteModelInferenceParamsWithTimeout(timeout time.Duration) *KafeidoServiceDeleteModelInferenceParams {
	return &KafeidoServiceDeleteModelInferenceParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceDeleteModelInferenceParamsWithContext creates a new KafeidoServiceDeleteModelInferenceParams object
// with the ability to set a context for a request.
func NewKafeidoServiceDeleteModelInferenceParamsWithContext(ctx context.Context) *KafeidoServiceDeleteModelInferenceParams {
	return &KafeidoServiceDeleteModelInferenceParams{
		Context: ctx,
	}
}

// NewKafeidoServiceDeleteModelInferenceParamsWithHTTPClient creates a new KafeidoServiceDeleteModelInferenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceDeleteModelInferenceParamsWithHTTPClient(client *http.Client) *KafeidoServiceDeleteModelInferenceParams {
	return &KafeidoServiceDeleteModelInferenceParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceDeleteModelInferenceParams contains all the parameters to send to the API endpoint

	for the kafeido service delete model inference operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceDeleteModelInferenceParams struct {

	// ModelInferenceID.
	ModelInferenceID string

	// NamedPipelineNameForDelete.
	NamedPipelineNameForDelete *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service delete model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceDeleteModelInferenceParams) WithDefaults() *KafeidoServiceDeleteModelInferenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service delete model inference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceDeleteModelInferenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) WithTimeout(timeout time.Duration) *KafeidoServiceDeleteModelInferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) WithContext(ctx context.Context) *KafeidoServiceDeleteModelInferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) WithHTTPClient(client *http.Client) *KafeidoServiceDeleteModelInferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModelInferenceID adds the modelInferenceID to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) WithModelInferenceID(modelInferenceID string) *KafeidoServiceDeleteModelInferenceParams {
	o.SetModelInferenceID(modelInferenceID)
	return o
}

// SetModelInferenceID adds the modelInferenceId to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) SetModelInferenceID(modelInferenceID string) {
	o.ModelInferenceID = modelInferenceID
}

// WithNamedPipelineNameForDelete adds the namedPipelineNameForDelete to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) WithNamedPipelineNameForDelete(namedPipelineNameForDelete *string) *KafeidoServiceDeleteModelInferenceParams {
	o.SetNamedPipelineNameForDelete(namedPipelineNameForDelete)
	return o
}

// SetNamedPipelineNameForDelete adds the namedPipelineNameForDelete to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) SetNamedPipelineNameForDelete(namedPipelineNameForDelete *string) {
	o.NamedPipelineNameForDelete = namedPipelineNameForDelete
}

// WithProjectID adds the projectID to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) WithProjectID(projectID string) *KafeidoServiceDeleteModelInferenceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service delete model inference params
func (o *KafeidoServiceDeleteModelInferenceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceDeleteModelInferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modelInferenceId
	if err := r.SetPathParam("modelInferenceId", o.ModelInferenceID); err != nil {
		return err
	}

	if o.NamedPipelineNameForDelete != nil {

		// query param namedPipelineNameForDelete
		var qrNamedPipelineNameForDelete string

		if o.NamedPipelineNameForDelete != nil {
			qrNamedPipelineNameForDelete = *o.NamedPipelineNameForDelete
		}
		qNamedPipelineNameForDelete := qrNamedPipelineNameForDelete
		if qNamedPipelineNameForDelete != "" {

			if err := r.SetQueryParam("namedPipelineNameForDelete", qNamedPipelineNameForDelete); err != nil {
				return err
			}
		}
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
