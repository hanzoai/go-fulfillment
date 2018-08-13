// Code generated by go-swagger; DO NOT EDIT.

package orders

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

	models "go-fulfillment/models"
)

// NewCreateOrderParams creates a new CreateOrderParams object
// with the default values initialized.
func NewCreateOrderParams() *CreateOrderParams {
	var ()
	return &CreateOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrderParamsWithTimeout creates a new CreateOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrderParamsWithTimeout(timeout time.Duration) *CreateOrderParams {
	var ()
	return &CreateOrderParams{

		timeout: timeout,
	}
}

// NewCreateOrderParamsWithContext creates a new CreateOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrderParamsWithContext(ctx context.Context) *CreateOrderParams {
	var ()
	return &CreateOrderParams{

		Context: ctx,
	}
}

// NewCreateOrderParamsWithHTTPClient creates a new CreateOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrderParamsWithHTTPClient(client *http.Client) *CreateOrderParams {
	var ()
	return &CreateOrderParams{
		HTTPClient: client,
	}
}

/*CreateOrderParams contains all the parameters to send to the API endpoint
for the create order operation typically these are written to a http.Request
*/
type CreateOrderParams struct {

	/*Authorization
	  OAuth 2.0 Bearer Access Token

	*/
	Authorization string
	/*Body
	  The order to create

	*/
	Body *models.NewOrder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create order params
func (o *CreateOrderParams) WithTimeout(timeout time.Duration) *CreateOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create order params
func (o *CreateOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create order params
func (o *CreateOrderParams) WithContext(ctx context.Context) *CreateOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create order params
func (o *CreateOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create order params
func (o *CreateOrderParams) WithHTTPClient(client *http.Client) *CreateOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create order params
func (o *CreateOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create order params
func (o *CreateOrderParams) WithAuthorization(authorization string) *CreateOrderParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create order params
func (o *CreateOrderParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the create order params
func (o *CreateOrderParams) WithBody(body *models.NewOrder) *CreateOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create order params
func (o *CreateOrderParams) SetBody(body *models.NewOrder) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

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
