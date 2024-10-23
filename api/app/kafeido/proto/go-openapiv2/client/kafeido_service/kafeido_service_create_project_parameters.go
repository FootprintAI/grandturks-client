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

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// NewKafeidoServiceCreateProjectParams creates a new KafeidoServiceCreateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceCreateProjectParams() *KafeidoServiceCreateProjectParams {
	return &KafeidoServiceCreateProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceCreateProjectParamsWithTimeout creates a new KafeidoServiceCreateProjectParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceCreateProjectParamsWithTimeout(timeout time.Duration) *KafeidoServiceCreateProjectParams {
	return &KafeidoServiceCreateProjectParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceCreateProjectParamsWithContext creates a new KafeidoServiceCreateProjectParams object
// with the ability to set a context for a request.
func NewKafeidoServiceCreateProjectParamsWithContext(ctx context.Context) *KafeidoServiceCreateProjectParams {
	return &KafeidoServiceCreateProjectParams{
		Context: ctx,
	}
}

// NewKafeidoServiceCreateProjectParamsWithHTTPClient creates a new KafeidoServiceCreateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceCreateProjectParamsWithHTTPClient(client *http.Client) *KafeidoServiceCreateProjectParams {
	return &KafeidoServiceCreateProjectParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceCreateProjectParams contains all the parameters to send to the API endpoint

	for the kafeido service create project operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceCreateProjectParams struct {

	// Body.
	Body *models.KafeidoCreateProjectRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceCreateProjectParams) WithDefaults() *KafeidoServiceCreateProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceCreateProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) WithTimeout(timeout time.Duration) *KafeidoServiceCreateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) WithContext(ctx context.Context) *KafeidoServiceCreateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) WithHTTPClient(client *http.Client) *KafeidoServiceCreateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) WithBody(body *models.KafeidoCreateProjectRequest) *KafeidoServiceCreateProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido service create project params
func (o *KafeidoServiceCreateProjectParams) SetBody(body *models.KafeidoCreateProjectRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceCreateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
