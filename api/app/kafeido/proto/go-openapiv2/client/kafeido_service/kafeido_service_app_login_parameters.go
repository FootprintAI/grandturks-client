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

// NewKafeidoServiceAppLoginParams creates a new KafeidoServiceAppLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoServiceAppLoginParams() *KafeidoServiceAppLoginParams {
	return &KafeidoServiceAppLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoServiceAppLoginParamsWithTimeout creates a new KafeidoServiceAppLoginParams object
// with the ability to set a timeout on a request.
func NewKafeidoServiceAppLoginParamsWithTimeout(timeout time.Duration) *KafeidoServiceAppLoginParams {
	return &KafeidoServiceAppLoginParams{
		timeout: timeout,
	}
}

// NewKafeidoServiceAppLoginParamsWithContext creates a new KafeidoServiceAppLoginParams object
// with the ability to set a context for a request.
func NewKafeidoServiceAppLoginParamsWithContext(ctx context.Context) *KafeidoServiceAppLoginParams {
	return &KafeidoServiceAppLoginParams{
		Context: ctx,
	}
}

// NewKafeidoServiceAppLoginParamsWithHTTPClient creates a new KafeidoServiceAppLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoServiceAppLoginParamsWithHTTPClient(client *http.Client) *KafeidoServiceAppLoginParams {
	return &KafeidoServiceAppLoginParams{
		HTTPClient: client,
	}
}

/*
KafeidoServiceAppLoginParams contains all the parameters to send to the API endpoint

	for the kafeido service app login operation.

	Typically these are written to a http.Request.
*/
type KafeidoServiceAppLoginParams struct {

	// Body.
	Body *models.KafeidoAppLoginRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido service app login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceAppLoginParams) WithDefaults() *KafeidoServiceAppLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido service app login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoServiceAppLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) WithTimeout(timeout time.Duration) *KafeidoServiceAppLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) WithContext(ctx context.Context) *KafeidoServiceAppLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) WithHTTPClient(client *http.Client) *KafeidoServiceAppLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) WithBody(body *models.KafeidoAppLoginRequest) *KafeidoServiceAppLoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido service app login params
func (o *KafeidoServiceAppLoginParams) SetBody(body *models.KafeidoAppLoginRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoServiceAppLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
