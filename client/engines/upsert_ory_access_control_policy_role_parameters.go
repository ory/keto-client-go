// Code generated by go-swagger; DO NOT EDIT.

package engines

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

	"github.com/ory/keto-client-go/models"
)

// NewUpsertOryAccessControlPolicyRoleParams creates a new UpsertOryAccessControlPolicyRoleParams object
// with the default values initialized.
func NewUpsertOryAccessControlPolicyRoleParams() *UpsertOryAccessControlPolicyRoleParams {
	var ()
	return &UpsertOryAccessControlPolicyRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertOryAccessControlPolicyRoleParamsWithTimeout creates a new UpsertOryAccessControlPolicyRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpsertOryAccessControlPolicyRoleParamsWithTimeout(timeout time.Duration) *UpsertOryAccessControlPolicyRoleParams {
	var ()
	return &UpsertOryAccessControlPolicyRoleParams{

		timeout: timeout,
	}
}

// NewUpsertOryAccessControlPolicyRoleParamsWithContext creates a new UpsertOryAccessControlPolicyRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpsertOryAccessControlPolicyRoleParamsWithContext(ctx context.Context) *UpsertOryAccessControlPolicyRoleParams {
	var ()
	return &UpsertOryAccessControlPolicyRoleParams{

		Context: ctx,
	}
}

// NewUpsertOryAccessControlPolicyRoleParamsWithHTTPClient creates a new UpsertOryAccessControlPolicyRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpsertOryAccessControlPolicyRoleParamsWithHTTPClient(client *http.Client) *UpsertOryAccessControlPolicyRoleParams {
	var ()
	return &UpsertOryAccessControlPolicyRoleParams{
		HTTPClient: client,
	}
}

/*UpsertOryAccessControlPolicyRoleParams contains all the parameters to send to the API endpoint
for the upsert ory access control policy role operation typically these are written to a http.Request
*/
type UpsertOryAccessControlPolicyRoleParams struct {

	/*Body*/
	Body *models.OryAccessControlPolicyRole
	/*Flavor
	  The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".

	*/
	Flavor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) WithTimeout(timeout time.Duration) *UpsertOryAccessControlPolicyRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) WithContext(ctx context.Context) *UpsertOryAccessControlPolicyRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) WithHTTPClient(client *http.Client) *UpsertOryAccessControlPolicyRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) WithBody(body *models.OryAccessControlPolicyRole) *UpsertOryAccessControlPolicyRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) SetBody(body *models.OryAccessControlPolicyRole) {
	o.Body = body
}

// WithFlavor adds the flavor to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) WithFlavor(flavor string) *UpsertOryAccessControlPolicyRoleParams {
	o.SetFlavor(flavor)
	return o
}

// SetFlavor adds the flavor to the upsert ory access control policy role params
func (o *UpsertOryAccessControlPolicyRoleParams) SetFlavor(flavor string) {
	o.Flavor = flavor
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertOryAccessControlPolicyRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param flavor
	if err := r.SetPathParam("flavor", o.Flavor); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
