// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostControlNewDiscoveredAPIsHandlerFunc turns a function with the right signature into a post control new discovered a p is handler
type PostControlNewDiscoveredAPIsHandlerFunc func(PostControlNewDiscoveredAPIsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostControlNewDiscoveredAPIsHandlerFunc) Handle(params PostControlNewDiscoveredAPIsParams) middleware.Responder {
	return fn(params)
}

// PostControlNewDiscoveredAPIsHandler interface for that can handle valid post control new discovered a p is params
type PostControlNewDiscoveredAPIsHandler interface {
	Handle(PostControlNewDiscoveredAPIsParams) middleware.Responder
}

// NewPostControlNewDiscoveredAPIs creates a new http.Handler for the post control new discovered a p is operation
func NewPostControlNewDiscoveredAPIs(ctx *middleware.Context, handler PostControlNewDiscoveredAPIsHandler) *PostControlNewDiscoveredAPIs {
	return &PostControlNewDiscoveredAPIs{Context: ctx, Handler: handler}
}

/* PostControlNewDiscoveredAPIs swagger:route POST /control/newDiscoveredAPIs postControlNewDiscoveredAPIs

Allows a client to notify APIClarity about new APIs.

This allows a client (a gateway for example) to notify APIclarity about newly discovered APIs. If one of the APIs already exists, it is ignored.

*/
type PostControlNewDiscoveredAPIs struct {
	Context *middleware.Context
	Handler PostControlNewDiscoveredAPIsHandler
}

func (o *PostControlNewDiscoveredAPIs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostControlNewDiscoveredAPIsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostControlNewDiscoveredAPIsBody post control new discovered a p is body
//
// swagger:model PostControlNewDiscoveredAPIsBody
type PostControlNewDiscoveredAPIsBody struct {

	// List of discovered APIs, format of hostname:port
	// Required: true
	Hosts []string `json:"hosts"`
}

// Validate validates this post control new discovered a p is body
func (o *PostControlNewDiscoveredAPIsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostControlNewDiscoveredAPIsBody) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"hosts", "body", o.Hosts); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post control new discovered a p is body based on context it is used
func (o *PostControlNewDiscoveredAPIsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostControlNewDiscoveredAPIsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostControlNewDiscoveredAPIsBody) UnmarshalBinary(b []byte) error {
	var res PostControlNewDiscoveredAPIsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
