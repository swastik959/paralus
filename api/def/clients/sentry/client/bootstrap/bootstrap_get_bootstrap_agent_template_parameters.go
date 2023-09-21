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
	"github.com/go-openapi/swag"
)

// NewBootstrapGetBootstrapAgentTemplateParams creates a new BootstrapGetBootstrapAgentTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBootstrapGetBootstrapAgentTemplateParams() *BootstrapGetBootstrapAgentTemplateParams {
	return &BootstrapGetBootstrapAgentTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBootstrapGetBootstrapAgentTemplateParamsWithTimeout creates a new BootstrapGetBootstrapAgentTemplateParams object
// with the ability to set a timeout on a request.
func NewBootstrapGetBootstrapAgentTemplateParamsWithTimeout(timeout time.Duration) *BootstrapGetBootstrapAgentTemplateParams {
	return &BootstrapGetBootstrapAgentTemplateParams{
		timeout: timeout,
	}
}

// NewBootstrapGetBootstrapAgentTemplateParamsWithContext creates a new BootstrapGetBootstrapAgentTemplateParams object
// with the ability to set a context for a request.
func NewBootstrapGetBootstrapAgentTemplateParamsWithContext(ctx context.Context) *BootstrapGetBootstrapAgentTemplateParams {
	return &BootstrapGetBootstrapAgentTemplateParams{
		Context: ctx,
	}
}

// NewBootstrapGetBootstrapAgentTemplateParamsWithHTTPClient creates a new BootstrapGetBootstrapAgentTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBootstrapGetBootstrapAgentTemplateParamsWithHTTPClient(client *http.Client) *BootstrapGetBootstrapAgentTemplateParams {
	return &BootstrapGetBootstrapAgentTemplateParams{
		HTTPClient: client,
	}
}

/*
BootstrapGetBootstrapAgentTemplateParams contains all the parameters to send to the API endpoint

	for the bootstrap get bootstrap agent template operation.

	Typically these are written to a http.Request.
*/
type BootstrapGetBootstrapAgentTemplateParams struct {

	/* APIVersion.

	   API Version. API Version of the resource

	   Default: "infra.k8smgmt.io/v3"
	*/
	APIVersion *string

	/* Kind.

	   Kind. Kind of the resource

	   Default: "BootstrapAgentTemplate"
	*/
	Kind *string

	/* MetadataDescription.

	   Description. description of the resource
	*/
	MetadataDescription *string

	/* MetadataDisplayName.

	   Display Name. display name of the resource
	*/
	MetadataDisplayName *string

	// MetadataID.
	MetadataID *string

	// MetadataModifiedAt.
	//
	// Format: date-time
	MetadataModifiedAt *strfmt.DateTime

	/* MetadataName.

	   name of the resource
	*/
	MetadataName string

	/* MetadataOrganization.

	   Organization. Organization to which the resource belongs
	*/
	MetadataOrganization *string

	/* MetadataPartner.

	   Partner. Partner to which the resource belongs
	*/
	MetadataPartner *string

	/* MetadataProject.

	   Project. Project of the resource
	*/
	MetadataProject *string

	// SpecAutoApprove.
	SpecAutoApprove *bool

	// SpecAutoRegister.
	SpecAutoRegister *bool

	// SpecIgnoreMultipleRegister.
	SpecIgnoreMultipleRegister *bool

	// SpecInClusterTemplate.
	SpecInClusterTemplate *string

	// SpecInfraRef.
	SpecInfraRef *string

	// SpecOutOfClusterTemplate.
	SpecOutOfClusterTemplate *string

	// SpecTemplateType.
	//
	// Default: "TemplateTypeNotSet"
	SpecTemplateType *string

	// SpecToken.
	SpecToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bootstrap get bootstrap agent template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BootstrapGetBootstrapAgentTemplateParams) WithDefaults() *BootstrapGetBootstrapAgentTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bootstrap get bootstrap agent template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BootstrapGetBootstrapAgentTemplateParams) SetDefaults() {
	var (
		aPIVersionDefault = string("infra.k8smgmt.io/v3")

		kindDefault = string("BootstrapAgentTemplate")

		specTemplateTypeDefault = string("TemplateTypeNotSet")
	)

	val := BootstrapGetBootstrapAgentTemplateParams{
		APIVersion:       &aPIVersionDefault,
		Kind:             &kindDefault,
		SpecTemplateType: &specTemplateTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithTimeout(timeout time.Duration) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithContext(ctx context.Context) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithHTTPClient(client *http.Client) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithAPIVersion(aPIVersion *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithKind adds the kind to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithKind(kind *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithMetadataDescription adds the metadataDescription to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataDescription(metadataDescription *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataDescription(metadataDescription)
	return o
}

// SetMetadataDescription adds the metadataDescription to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataDescription(metadataDescription *string) {
	o.MetadataDescription = metadataDescription
}

// WithMetadataDisplayName adds the metadataDisplayName to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataDisplayName(metadataDisplayName *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataDisplayName(metadataDisplayName)
	return o
}

// SetMetadataDisplayName adds the metadataDisplayName to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataDisplayName(metadataDisplayName *string) {
	o.MetadataDisplayName = metadataDisplayName
}

// WithMetadataID adds the metadataID to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataID(metadataID *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataID(metadataID)
	return o
}

// SetMetadataID adds the metadataId to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataID(metadataID *string) {
	o.MetadataID = metadataID
}

// WithMetadataModifiedAt adds the metadataModifiedAt to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataModifiedAt(metadataModifiedAt *strfmt.DateTime) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataModifiedAt(metadataModifiedAt)
	return o
}

// SetMetadataModifiedAt adds the metadataModifiedAt to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataModifiedAt(metadataModifiedAt *strfmt.DateTime) {
	o.MetadataModifiedAt = metadataModifiedAt
}

// WithMetadataName adds the metadataName to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataName(metadataName string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataName(metadataName)
	return o
}

// SetMetadataName adds the metadataName to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataName(metadataName string) {
	o.MetadataName = metadataName
}

// WithMetadataOrganization adds the metadataOrganization to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataOrganization(metadataOrganization *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataOrganization(metadataOrganization)
	return o
}

// SetMetadataOrganization adds the metadataOrganization to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataOrganization(metadataOrganization *string) {
	o.MetadataOrganization = metadataOrganization
}

// WithMetadataPartner adds the metadataPartner to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataPartner(metadataPartner *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataPartner(metadataPartner)
	return o
}

// SetMetadataPartner adds the metadataPartner to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataPartner(metadataPartner *string) {
	o.MetadataPartner = metadataPartner
}

// WithMetadataProject adds the metadataProject to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithMetadataProject(metadataProject *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetMetadataProject(metadataProject)
	return o
}

// SetMetadataProject adds the metadataProject to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetMetadataProject(metadataProject *string) {
	o.MetadataProject = metadataProject
}

// WithSpecAutoApprove adds the specAutoApprove to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecAutoApprove(specAutoApprove *bool) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecAutoApprove(specAutoApprove)
	return o
}

// SetSpecAutoApprove adds the specAutoApprove to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecAutoApprove(specAutoApprove *bool) {
	o.SpecAutoApprove = specAutoApprove
}

// WithSpecAutoRegister adds the specAutoRegister to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecAutoRegister(specAutoRegister *bool) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecAutoRegister(specAutoRegister)
	return o
}

// SetSpecAutoRegister adds the specAutoRegister to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecAutoRegister(specAutoRegister *bool) {
	o.SpecAutoRegister = specAutoRegister
}

// WithSpecIgnoreMultipleRegister adds the specIgnoreMultipleRegister to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecIgnoreMultipleRegister(specIgnoreMultipleRegister *bool) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecIgnoreMultipleRegister(specIgnoreMultipleRegister)
	return o
}

// SetSpecIgnoreMultipleRegister adds the specIgnoreMultipleRegister to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecIgnoreMultipleRegister(specIgnoreMultipleRegister *bool) {
	o.SpecIgnoreMultipleRegister = specIgnoreMultipleRegister
}

// WithSpecInClusterTemplate adds the specInClusterTemplate to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecInClusterTemplate(specInClusterTemplate *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecInClusterTemplate(specInClusterTemplate)
	return o
}

// SetSpecInClusterTemplate adds the specInClusterTemplate to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecInClusterTemplate(specInClusterTemplate *string) {
	o.SpecInClusterTemplate = specInClusterTemplate
}

// WithSpecInfraRef adds the specInfraRef to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecInfraRef(specInfraRef *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecInfraRef(specInfraRef)
	return o
}

// SetSpecInfraRef adds the specInfraRef to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecInfraRef(specInfraRef *string) {
	o.SpecInfraRef = specInfraRef
}

// WithSpecOutOfClusterTemplate adds the specOutOfClusterTemplate to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecOutOfClusterTemplate(specOutOfClusterTemplate *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecOutOfClusterTemplate(specOutOfClusterTemplate)
	return o
}

// SetSpecOutOfClusterTemplate adds the specOutOfClusterTemplate to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecOutOfClusterTemplate(specOutOfClusterTemplate *string) {
	o.SpecOutOfClusterTemplate = specOutOfClusterTemplate
}

// WithSpecTemplateType adds the specTemplateType to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecTemplateType(specTemplateType *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecTemplateType(specTemplateType)
	return o
}

// SetSpecTemplateType adds the specTemplateType to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecTemplateType(specTemplateType *string) {
	o.SpecTemplateType = specTemplateType
}

// WithSpecToken adds the specToken to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) WithSpecToken(specToken *string) *BootstrapGetBootstrapAgentTemplateParams {
	o.SetSpecToken(specToken)
	return o
}

// SetSpecToken adds the specToken to the bootstrap get bootstrap agent template params
func (o *BootstrapGetBootstrapAgentTemplateParams) SetSpecToken(specToken *string) {
	o.SpecToken = specToken
}

// WriteToRequest writes these params to a swagger request
func (o *BootstrapGetBootstrapAgentTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	if o.MetadataDescription != nil {

		// query param metadata.description
		var qrMetadataDescription string

		if o.MetadataDescription != nil {
			qrMetadataDescription = *o.MetadataDescription
		}
		qMetadataDescription := qrMetadataDescription
		if qMetadataDescription != "" {

			if err := r.SetQueryParam("metadata.description", qMetadataDescription); err != nil {
				return err
			}
		}
	}

	if o.MetadataDisplayName != nil {

		// query param metadata.displayName
		var qrMetadataDisplayName string

		if o.MetadataDisplayName != nil {
			qrMetadataDisplayName = *o.MetadataDisplayName
		}
		qMetadataDisplayName := qrMetadataDisplayName
		if qMetadataDisplayName != "" {

			if err := r.SetQueryParam("metadata.displayName", qMetadataDisplayName); err != nil {
				return err
			}
		}
	}

	if o.MetadataID != nil {

		// query param metadata.id
		var qrMetadataID string

		if o.MetadataID != nil {
			qrMetadataID = *o.MetadataID
		}
		qMetadataID := qrMetadataID
		if qMetadataID != "" {

			if err := r.SetQueryParam("metadata.id", qMetadataID); err != nil {
				return err
			}
		}
	}

	if o.MetadataModifiedAt != nil {

		// query param metadata.modifiedAt
		var qrMetadataModifiedAt strfmt.DateTime

		if o.MetadataModifiedAt != nil {
			qrMetadataModifiedAt = *o.MetadataModifiedAt
		}
		qMetadataModifiedAt := qrMetadataModifiedAt.String()
		if qMetadataModifiedAt != "" {

			if err := r.SetQueryParam("metadata.modifiedAt", qMetadataModifiedAt); err != nil {
				return err
			}
		}
	}

	// path param metadata.name
	if err := r.SetPathParam("metadata.name", o.MetadataName); err != nil {
		return err
	}

	if o.MetadataOrganization != nil {

		// query param metadata.organization
		var qrMetadataOrganization string

		if o.MetadataOrganization != nil {
			qrMetadataOrganization = *o.MetadataOrganization
		}
		qMetadataOrganization := qrMetadataOrganization
		if qMetadataOrganization != "" {

			if err := r.SetQueryParam("metadata.organization", qMetadataOrganization); err != nil {
				return err
			}
		}
	}

	if o.MetadataPartner != nil {

		// query param metadata.partner
		var qrMetadataPartner string

		if o.MetadataPartner != nil {
			qrMetadataPartner = *o.MetadataPartner
		}
		qMetadataPartner := qrMetadataPartner
		if qMetadataPartner != "" {

			if err := r.SetQueryParam("metadata.partner", qMetadataPartner); err != nil {
				return err
			}
		}
	}

	if o.MetadataProject != nil {

		// query param metadata.project
		var qrMetadataProject string

		if o.MetadataProject != nil {
			qrMetadataProject = *o.MetadataProject
		}
		qMetadataProject := qrMetadataProject
		if qMetadataProject != "" {

			if err := r.SetQueryParam("metadata.project", qMetadataProject); err != nil {
				return err
			}
		}
	}

	if o.SpecAutoApprove != nil {

		// query param spec.autoApprove
		var qrSpecAutoApprove bool

		if o.SpecAutoApprove != nil {
			qrSpecAutoApprove = *o.SpecAutoApprove
		}
		qSpecAutoApprove := swag.FormatBool(qrSpecAutoApprove)
		if qSpecAutoApprove != "" {

			if err := r.SetQueryParam("spec.autoApprove", qSpecAutoApprove); err != nil {
				return err
			}
		}
	}

	if o.SpecAutoRegister != nil {

		// query param spec.autoRegister
		var qrSpecAutoRegister bool

		if o.SpecAutoRegister != nil {
			qrSpecAutoRegister = *o.SpecAutoRegister
		}
		qSpecAutoRegister := swag.FormatBool(qrSpecAutoRegister)
		if qSpecAutoRegister != "" {

			if err := r.SetQueryParam("spec.autoRegister", qSpecAutoRegister); err != nil {
				return err
			}
		}
	}

	if o.SpecIgnoreMultipleRegister != nil {

		// query param spec.ignoreMultipleRegister
		var qrSpecIgnoreMultipleRegister bool

		if o.SpecIgnoreMultipleRegister != nil {
			qrSpecIgnoreMultipleRegister = *o.SpecIgnoreMultipleRegister
		}
		qSpecIgnoreMultipleRegister := swag.FormatBool(qrSpecIgnoreMultipleRegister)
		if qSpecIgnoreMultipleRegister != "" {

			if err := r.SetQueryParam("spec.ignoreMultipleRegister", qSpecIgnoreMultipleRegister); err != nil {
				return err
			}
		}
	}

	if o.SpecInClusterTemplate != nil {

		// query param spec.inClusterTemplate
		var qrSpecInClusterTemplate string

		if o.SpecInClusterTemplate != nil {
			qrSpecInClusterTemplate = *o.SpecInClusterTemplate
		}
		qSpecInClusterTemplate := qrSpecInClusterTemplate
		if qSpecInClusterTemplate != "" {

			if err := r.SetQueryParam("spec.inClusterTemplate", qSpecInClusterTemplate); err != nil {
				return err
			}
		}
	}

	if o.SpecInfraRef != nil {

		// query param spec.infraRef
		var qrSpecInfraRef string

		if o.SpecInfraRef != nil {
			qrSpecInfraRef = *o.SpecInfraRef
		}
		qSpecInfraRef := qrSpecInfraRef
		if qSpecInfraRef != "" {

			if err := r.SetQueryParam("spec.infraRef", qSpecInfraRef); err != nil {
				return err
			}
		}
	}

	if o.SpecOutOfClusterTemplate != nil {

		// query param spec.outOfClusterTemplate
		var qrSpecOutOfClusterTemplate string

		if o.SpecOutOfClusterTemplate != nil {
			qrSpecOutOfClusterTemplate = *o.SpecOutOfClusterTemplate
		}
		qSpecOutOfClusterTemplate := qrSpecOutOfClusterTemplate
		if qSpecOutOfClusterTemplate != "" {

			if err := r.SetQueryParam("spec.outOfClusterTemplate", qSpecOutOfClusterTemplate); err != nil {
				return err
			}
		}
	}

	if o.SpecTemplateType != nil {

		// query param spec.templateType
		var qrSpecTemplateType string

		if o.SpecTemplateType != nil {
			qrSpecTemplateType = *o.SpecTemplateType
		}
		qSpecTemplateType := qrSpecTemplateType
		if qSpecTemplateType != "" {

			if err := r.SetQueryParam("spec.templateType", qSpecTemplateType); err != nil {
				return err
			}
		}
	}

	if o.SpecToken != nil {

		// query param spec.token
		var qrSpecToken string

		if o.SpecToken != nil {
			qrSpecToken = *o.SpecToken
		}
		qSpecToken := qrSpecToken
		if qSpecToken != "" {

			if err := r.SetQueryParam("spec.token", qSpecToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
