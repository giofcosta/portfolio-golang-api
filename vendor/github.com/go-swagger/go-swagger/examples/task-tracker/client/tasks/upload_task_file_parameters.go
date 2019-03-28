// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadTaskFileParams creates a new UploadTaskFileParams object
// with the default values initialized.
func NewUploadTaskFileParams() *UploadTaskFileParams {
	var ()
	return &UploadTaskFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadTaskFileParamsWithTimeout creates a new UploadTaskFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadTaskFileParamsWithTimeout(timeout time.Duration) *UploadTaskFileParams {
	var ()
	return &UploadTaskFileParams{

		timeout: timeout,
	}
}

// NewUploadTaskFileParamsWithContext creates a new UploadTaskFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadTaskFileParamsWithContext(ctx context.Context) *UploadTaskFileParams {
	var ()
	return &UploadTaskFileParams{

		Context: ctx,
	}
}

// NewUploadTaskFileParamsWithHTTPClient creates a new UploadTaskFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadTaskFileParamsWithHTTPClient(client *http.Client) *UploadTaskFileParams {
	var ()
	return &UploadTaskFileParams{
		HTTPClient: client,
	}
}

/*UploadTaskFileParams contains all the parameters to send to the API endpoint
for the upload task file operation typically these are written to a http.Request
*/
type UploadTaskFileParams struct {

	/*Description
	  Extra information describing the file

	*/
	Description *string
	/*File
	  The file to upload

	*/
	File runtime.NamedReadCloser
	/*ID
	  The id of the item

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload task file params
func (o *UploadTaskFileParams) WithTimeout(timeout time.Duration) *UploadTaskFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload task file params
func (o *UploadTaskFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload task file params
func (o *UploadTaskFileParams) WithContext(ctx context.Context) *UploadTaskFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload task file params
func (o *UploadTaskFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload task file params
func (o *UploadTaskFileParams) WithHTTPClient(client *http.Client) *UploadTaskFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload task file params
func (o *UploadTaskFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the upload task file params
func (o *UploadTaskFileParams) WithDescription(description *string) *UploadTaskFileParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the upload task file params
func (o *UploadTaskFileParams) SetDescription(description *string) {
	o.Description = description
}

// WithFile adds the file to the upload task file params
func (o *UploadTaskFileParams) WithFile(file runtime.NamedReadCloser) *UploadTaskFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload task file params
func (o *UploadTaskFileParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithID adds the id to the upload task file params
func (o *UploadTaskFileParams) WithID(id int64) *UploadTaskFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the upload task file params
func (o *UploadTaskFileParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UploadTaskFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	if o.File != nil {

		if o.File != nil {

			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}

		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
