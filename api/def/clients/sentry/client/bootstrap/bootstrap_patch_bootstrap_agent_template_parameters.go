// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

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

	"github.com/paralus/paralus/api/def/clients/sentry/models"
)

// NewBootstrapPatchBootstrapAgentTemplateParams creates a new BootstrapPatchBootstrapAgentTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBootstrapPatchBootstrapAgentTemplateParams() *BootstrapPatchBootstrapAgentTemplateParams {
	return &BootstrapPatchBootstrapAgentTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBootstrapPatchBootstrapAgentTemplateParamsWithTimeout creates a new BootstrapPatchBootstrapAgentTemplateParams object
// with the ability to set a timeout on a request.
func NewBootstrapPatchBootstrapAgentTemplateParamsWithTimeout(timeout time.Duration) *BootstrapPatchBootstrapAgentTemplateParams {
	return &BootstrapPatchBootstrapAgentTemplateParams{
		timeout: timeout,
	}
}

// NewBootstrapPatchBootstrapAgentTemplateParamsWithContext creates a new BootstrapPatchBootstrapAgentTemplateParams object
// with the ability to set a context for a request.
func NewBootstrapPatchBootstrapAgentTemplateParamsWithContext(ctx context.Context) *BootstrapPatchBootstrapAgentTemplateParams {
	return &BootstrapPatchBootstrapAgentTemplateParams{
		Context: ctx,
	}
}

// NewBootstrapPatchBootstrapAgentTemplateParamsWithHTTPClient creates a new BootstrapPatchBootstrapAgentTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBootstrapPatchBootstrapAgentTemplateParamsWithHTTPClient(client *http.Client) *BootstrapPatchBootstrapAgentTemplateParams {
	return &BootstrapPatchBootstrapAgentTemplateParams{
		HTTPClient: client,
	}
}

/*
BootstrapPatchBootstrapAgentTemplateParams contains all the parameters to send to the API endpoint

	for the bootstrap patch bootstrap agent template operation.

	Typically these are written to a http.Request.
*/
type BootstrapPatchBootstrapAgentTemplateParams struct {

	// Body.
	Body *models.SentryBootstrapAgentTemplate

	/* MetadataName.

	   name of the resource
	*/
	MetadataName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bootstrap patch bootstrap agent template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BootstrapPatchBootstrapAgentTemplateParams) WithDefaults() *BootstrapPatchBootstrapAgentTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bootstrap patch bootstrap agent template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BootstrapPatchBootstrapAgentTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) WithTimeout(timeout time.Duration) *BootstrapPatchBootstrapAgentTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) WithContext(ctx context.Context) *BootstrapPatchBootstrapAgentTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) WithHTTPClient(client *http.Client) *BootstrapPatchBootstrapAgentTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) WithBody(body *models.SentryBootstrapAgentTemplate) *BootstrapPatchBootstrapAgentTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) SetBody(body *models.SentryBootstrapAgentTemplate) {
	o.Body = body
}

// WithMetadataName adds the metadataName to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) WithMetadataName(metadataName string) *BootstrapPatchBootstrapAgentTemplateParams {
	o.SetMetadataName(metadataName)
	return o
}

// SetMetadataName adds the metadataName to the bootstrap patch bootstrap agent template params
func (o *BootstrapPatchBootstrapAgentTemplateParams) SetMetadataName(metadataName string) {
	o.MetadataName = metadataName
}

// WriteToRequest writes these params to a swagger request
func (o *BootstrapPatchBootstrapAgentTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param metadata.name
	if err := r.SetPathParam("metadata.name", o.MetadataName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
