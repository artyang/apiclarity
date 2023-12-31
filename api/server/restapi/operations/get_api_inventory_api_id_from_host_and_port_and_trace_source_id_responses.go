// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/apiclarity/api/server/models"
)

// GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOKCode is the HTTP code returned for type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK
const GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOKCode int = 200

/*GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK Success

swagger:response getApiInventoryApiIdFromHostAndPortAndTraceSourceIdOK
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK struct {

	/*api id
	  In: Body
	*/
	Payload uint32 `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK creates GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK with default headers values
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK() *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK {

	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK{}
}

// WithPayload adds the payload to the get Api inventory Api Id from host and port and trace source Id o k response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK) WithPayload(payload uint32) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api inventory Api Id from host and port and trace source Id o k response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK) SetPayload(payload uint32) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFoundCode is the HTTP code returned for type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound
const GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFoundCode int = 404

/*GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound API ID Not Found

swagger:response getApiInventoryApiIdFromHostAndPortAndTraceSourceIdNotFound
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound creates GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound with default headers values
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound() *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound {

	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound{}
}

// WithPayload adds the payload to the get Api inventory Api Id from host and port and trace source Id not found response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound) WithPayload(payload *models.APIResponse) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api inventory Api Id from host and port and trace source Id not found response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault unknown error

swagger:response getApiInventoryApiIdFromHostAndPortAndTraceSourceIdDefault
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault creates GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault with default headers values
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault(code int) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get API inventory API ID from host and port and trace source ID default response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) WithStatusCode(code int) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get API inventory API ID from host and port and trace source ID default response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get API inventory API ID from host and port and trace source ID default response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) WithPayload(payload *models.APIResponse) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get API inventory API ID from host and port and trace source ID default response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
