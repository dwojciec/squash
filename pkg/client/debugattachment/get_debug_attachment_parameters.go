// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDebugAttachmentParams creates a new GetDebugAttachmentParams object
// with the default values initialized.
func NewGetDebugAttachmentParams() *GetDebugAttachmentParams {
	var ()
	return &GetDebugAttachmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebugAttachmentParamsWithTimeout creates a new GetDebugAttachmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDebugAttachmentParamsWithTimeout(timeout time.Duration) *GetDebugAttachmentParams {
	var ()
	return &GetDebugAttachmentParams{

		timeout: timeout,
	}
}

// NewGetDebugAttachmentParamsWithContext creates a new GetDebugAttachmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDebugAttachmentParamsWithContext(ctx context.Context) *GetDebugAttachmentParams {
	var ()
	return &GetDebugAttachmentParams{

		Context: ctx,
	}
}

// NewGetDebugAttachmentParamsWithHTTPClient creates a new GetDebugAttachmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDebugAttachmentParamsWithHTTPClient(client *http.Client) *GetDebugAttachmentParams {
	var ()
	return &GetDebugAttachmentParams{
		HTTPClient: client,
	}
}

/*GetDebugAttachmentParams contains all the parameters to send to the API endpoint
for the get debug attachment operation typically these are written to a http.Request
*/
type GetDebugAttachmentParams struct {

	/*DebugAttachmentID
	  ID of config to return

	*/
	DebugAttachmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get debug attachment params
func (o *GetDebugAttachmentParams) WithTimeout(timeout time.Duration) *GetDebugAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debug attachment params
func (o *GetDebugAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debug attachment params
func (o *GetDebugAttachmentParams) WithContext(ctx context.Context) *GetDebugAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debug attachment params
func (o *GetDebugAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debug attachment params
func (o *GetDebugAttachmentParams) WithHTTPClient(client *http.Client) *GetDebugAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debug attachment params
func (o *GetDebugAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDebugAttachmentID adds the debugAttachmentID to the get debug attachment params
func (o *GetDebugAttachmentParams) WithDebugAttachmentID(debugAttachmentID string) *GetDebugAttachmentParams {
	o.SetDebugAttachmentID(debugAttachmentID)
	return o
}

// SetDebugAttachmentID adds the debugAttachmentId to the get debug attachment params
func (o *GetDebugAttachmentParams) SetDebugAttachmentID(debugAttachmentID string) {
	o.DebugAttachmentID = debugAttachmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebugAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param debugAttachmentId
	if err := r.SetPathParam("debugAttachmentId", o.DebugAttachmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
