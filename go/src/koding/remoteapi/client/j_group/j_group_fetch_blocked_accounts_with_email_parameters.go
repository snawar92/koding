package j_group

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

	"koding/remoteapi/models"
)

// NewJGroupFetchBlockedAccountsWithEmailParams creates a new JGroupFetchBlockedAccountsWithEmailParams object
// with the default values initialized.
func NewJGroupFetchBlockedAccountsWithEmailParams() *JGroupFetchBlockedAccountsWithEmailParams {
	var ()
	return &JGroupFetchBlockedAccountsWithEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchBlockedAccountsWithEmailParamsWithTimeout creates a new JGroupFetchBlockedAccountsWithEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchBlockedAccountsWithEmailParamsWithTimeout(timeout time.Duration) *JGroupFetchBlockedAccountsWithEmailParams {
	var ()
	return &JGroupFetchBlockedAccountsWithEmailParams{

		timeout: timeout,
	}
}

// NewJGroupFetchBlockedAccountsWithEmailParamsWithContext creates a new JGroupFetchBlockedAccountsWithEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchBlockedAccountsWithEmailParamsWithContext(ctx context.Context) *JGroupFetchBlockedAccountsWithEmailParams {
	var ()
	return &JGroupFetchBlockedAccountsWithEmailParams{

		Context: ctx,
	}
}

/*JGroupFetchBlockedAccountsWithEmailParams contains all the parameters to send to the API endpoint
for the j group fetch blocked accounts with email operation typically these are written to a http.Request
*/
type JGroupFetchBlockedAccountsWithEmailParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) WithTimeout(timeout time.Duration) *JGroupFetchBlockedAccountsWithEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) WithContext(ctx context.Context) *JGroupFetchBlockedAccountsWithEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) WithBody(body models.DefaultSelector) *JGroupFetchBlockedAccountsWithEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) WithID(id string) *JGroupFetchBlockedAccountsWithEmailParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch blocked accounts with email params
func (o *JGroupFetchBlockedAccountsWithEmailParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchBlockedAccountsWithEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
