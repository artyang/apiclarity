// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/openclarity/apiclarity/api/client/models"
)

// NewPostControlTraceSourcesParams creates a new PostControlTraceSourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostControlTraceSourcesParams() *PostControlTraceSourcesParams {
	return &PostControlTraceSourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostControlTraceSourcesParamsWithTimeout creates a new PostControlTraceSourcesParams object
// with the ability to set a timeout on a request.
func NewPostControlTraceSourcesParamsWithTimeout(timeout time.Duration) *PostControlTraceSourcesParams {
	return &PostControlTraceSourcesParams{
		timeout: timeout,
	}
}

// NewPostControlTraceSourcesParamsWithContext creates a new PostControlTraceSourcesParams object
// with the ability to set a context for a request.
func NewPostControlTraceSourcesParamsWithContext(ctx context.Context) *PostControlTraceSourcesParams {
	return &PostControlTraceSourcesParams{
		Context: ctx,
	}
}

// NewPostControlTraceSourcesParamsWithHTTPClient creates a new PostControlTraceSourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostControlTraceSourcesParamsWithHTTPClient(client *http.Client) *PostControlTraceSourcesParams {
	return &PostControlTraceSourcesParams{
		HTTPClient: client,
	}
}

/* PostControlTraceSourcesParams contains all the parameters to send to the API endpoint
   for the post control trace sources operation.

   Typically these are written to a http.Request.
*/
type PostControlTraceSourcesParams struct {

	// Body.
	Body *models.TraceSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post control trace sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostControlTraceSourcesParams) WithDefaults() *PostControlTraceSourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post control trace sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostControlTraceSourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post control trace sources params
func (o *PostControlTraceSourcesParams) WithTimeout(timeout time.Duration) *PostControlTraceSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post control trace sources params
func (o *PostControlTraceSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post control trace sources params
func (o *PostControlTraceSourcesParams) WithContext(ctx context.Context) *PostControlTraceSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post control trace sources params
func (o *PostControlTraceSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post control trace sources params
func (o *PostControlTraceSourcesParams) WithHTTPClient(client *http.Client) *PostControlTraceSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post control trace sources params
func (o *PostControlTraceSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post control trace sources params
func (o *PostControlTraceSourcesParams) WithBody(body *models.TraceSource) *PostControlTraceSourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post control trace sources params
func (o *PostControlTraceSourcesParams) SetBody(body *models.TraceSource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostControlTraceSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
