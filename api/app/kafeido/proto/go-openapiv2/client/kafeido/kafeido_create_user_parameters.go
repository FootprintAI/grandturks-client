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

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// NewKafeidoCreateUserParams creates a new KafeidoCreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafeidoCreateUserParams() *KafeidoCreateUserParams {
	return &KafeidoCreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafeidoCreateUserParamsWithTimeout creates a new KafeidoCreateUserParams object
// with the ability to set a timeout on a request.
func NewKafeidoCreateUserParamsWithTimeout(timeout time.Duration) *KafeidoCreateUserParams {
	return &KafeidoCreateUserParams{
		timeout: timeout,
	}
}

// NewKafeidoCreateUserParamsWithContext creates a new KafeidoCreateUserParams object
// with the ability to set a context for a request.
func NewKafeidoCreateUserParamsWithContext(ctx context.Context) *KafeidoCreateUserParams {
	return &KafeidoCreateUserParams{
		Context: ctx,
	}
}

// NewKafeidoCreateUserParamsWithHTTPClient creates a new KafeidoCreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafeidoCreateUserParamsWithHTTPClient(client *http.Client) *KafeidoCreateUserParams {
	return &KafeidoCreateUserParams{
		HTTPClient: client,
	}
}

/*
KafeidoCreateUserParams contains all the parameters to send to the API endpoint

	for the kafeido create user operation.

	Typically these are written to a http.Request.
*/
type KafeidoCreateUserParams struct {

	// Body.
	Body *models.KafeidoCreateUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafeido create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoCreateUserParams) WithDefaults() *KafeidoCreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafeido create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafeidoCreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafeido create user params
func (o *KafeidoCreateUserParams) WithTimeout(timeout time.Duration) *KafeidoCreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafeido create user params
func (o *KafeidoCreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafeido create user params
func (o *KafeidoCreateUserParams) WithContext(ctx context.Context) *KafeidoCreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafeido create user params
func (o *KafeidoCreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafeido create user params
func (o *KafeidoCreateUserParams) WithHTTPClient(client *http.Client) *KafeidoCreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafeido create user params
func (o *KafeidoCreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafeido create user params
func (o *KafeidoCreateUserParams) WithBody(body *models.KafeidoCreateUserRequest) *KafeidoCreateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafeido create user params
func (o *KafeidoCreateUserParams) SetBody(body *models.KafeidoCreateUserRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KafeidoCreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
