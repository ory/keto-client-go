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
	"github.com/go-openapi/swag"
)

// NewListOryAccessControlPoliciesParams creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized.
func NewListOryAccessControlPoliciesParams() *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOryAccessControlPoliciesParamsWithTimeout creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOryAccessControlPoliciesParamsWithTimeout(timeout time.Duration) *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{

		timeout: timeout,
	}
}

// NewListOryAccessControlPoliciesParamsWithContext creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOryAccessControlPoliciesParamsWithContext(ctx context.Context) *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{

		Context: ctx,
	}
}

// NewListOryAccessControlPoliciesParamsWithHTTPClient creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOryAccessControlPoliciesParamsWithHTTPClient(client *http.Client) *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{
		HTTPClient: client,
	}
}

/*ListOryAccessControlPoliciesParams contains all the parameters to send to the API endpoint
for the list ory access control policies operation typically these are written to a http.Request
*/
type ListOryAccessControlPoliciesParams struct {

	/*Action
	  The action for which policies are to be listed.

	*/
	Action *string
	/*Flavor
	  The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact"

	*/
	Flavor string
	/*Limit
	  The maximum amount of policies returned.

	*/
	Limit *int64
	/*Offset
	  The offset from where to start looking.

	*/
	Offset *int64
	/*Resource
	  The resource for which the policies are to be listed.

	*/
	Resource *string
	/*Subject
	  The subject for whom the policies are to be listed.

	*/
	Subject *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithTimeout(timeout time.Duration) *ListOryAccessControlPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithContext(ctx context.Context) *ListOryAccessControlPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithHTTPClient(client *http.Client) *ListOryAccessControlPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithAction(action *string) *ListOryAccessControlPoliciesParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetAction(action *string) {
	o.Action = action
}

// WithFlavor adds the flavor to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithFlavor(flavor string) *ListOryAccessControlPoliciesParams {
	o.SetFlavor(flavor)
	return o
}

// SetFlavor adds the flavor to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetFlavor(flavor string) {
	o.Flavor = flavor
}

// WithLimit adds the limit to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithLimit(limit *int64) *ListOryAccessControlPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithOffset(offset *int64) *ListOryAccessControlPoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithResource adds the resource to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithResource(resource *string) *ListOryAccessControlPoliciesParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetResource(resource *string) {
	o.Resource = resource
}

// WithSubject adds the subject to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithSubject(subject *string) *ListOryAccessControlPoliciesParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetSubject(subject *string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *ListOryAccessControlPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string
		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {
			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}

	}

	// path param flavor
	if err := r.SetPathParam("flavor", o.Flavor); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Resource != nil {

		// query param resource
		var qrResource string
		if o.Resource != nil {
			qrResource = *o.Resource
		}
		qResource := qrResource
		if qResource != "" {
			if err := r.SetQueryParam("resource", qResource); err != nil {
				return err
			}
		}

	}

	if o.Subject != nil {

		// query param subject
		var qrSubject string
		if o.Subject != nil {
			qrSubject = *o.Subject
		}
		qSubject := qrSubject
		if qSubject != "" {
			if err := r.SetQueryParam("subject", qSubject); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
