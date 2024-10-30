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

// NewKafeidoServiceDeleteDataSourceParams creates a new KafeidoServiceDeleteDataSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceDeleteDataSourceParams() *KafeidoServiceDeleteDataSourceParams {
	return &KafeidoServiceDeleteDataSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceDeleteDataSourceParamsWithTimeout creates a new KafeidoServiceDeleteDataSourceParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceDeleteDataSourceParamsWithTimeout(timeout time.Duration) *KafeidoServiceDeleteDataSourceParams {
	return &KafeidoServiceDeleteDataSourceParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceDeleteDataSourceParamsWithContext creates a new KafeidoServiceDeleteDataSourceParams object
// with the ability to set a context for a request.
func NewKafeidoServiceDeleteDataSourceParamsWithContext(ctx context.Context) *KafeidoServiceDeleteDataSourceParams {
	return &KafeidoServiceDeleteDataSourceParams{
		Context: ctx,
	}
}

// NewKafeidoServiceDeleteDataSourceParamsWithHTTPClient creates a new KafeidoServiceDeleteDataSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceDeleteDataSourceParamsWithHTTPClient(client *http.Client) *KafeidoServiceDeleteDataSourceParams {
	return &KafeidoServiceDeleteDataSourceParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceDeleteDataSourceParams contains all the parameters to send to the API endpoint

	for the kafeido service delete data source operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceDeleteDataSourceParams struct {

	// DataSourceID.
	DataSourceID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service delete data source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceDeleteDataSourceParams) WithDefaults() *KafeidoServiceDeleteDataSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service delete data source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceDeleteDataSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) WithTimeout(timeout time.Duration) *KafeidoServiceDeleteDataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) WithContext(ctx context.Context) *KafeidoServiceDeleteDataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) WithHTTPClient(client *http.Client) *KafeidoServiceDeleteDataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataSourceID adds the dataSourceID to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) WithDataSourceID(dataSourceID string) *KafeidoServiceDeleteDataSourceParams {
	o.SetDataSourceID(dataSourceID)
	return o
}

// SetDataSourceID adds the dataSourceId to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) SetDataSourceID(dataSourceID string) {
	o.DataSourceID = dataSourceID
}

// WithProjectID adds the projectID to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) WithProjectID(projectID string) *KafeidoServiceDeleteDataSourceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kafeido service delete data source params
func (o *KafeidoServiceDeleteDataSourceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceDeleteDataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dataSourceId
	if err := r.SetPathParam("dataSourceId", o.DataSourceID); err != nil {
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