// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/apiclarity/api/client/models"
)

// GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDReader is a Reader for the GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceID structure.
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK creates a GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK with default headers values
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK() *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK{}
}

/* GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK describes a response with status code 200, with default header values.

Success
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK struct {
	Payload uint32
}

func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK) Error() string {
	return fmt.Sprintf("[GET /apiInventory/apiId/fromHostAndPortAndTraceSourceID][%d] getApiInventoryApiIdFromHostAndPortAndTraceSourceIdOK  %+v", 200, o.Payload)
}
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK) GetPayload() uint32 {
	return o.Payload
}

func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound creates a GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound with default headers values
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound() *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound{}
}

/* GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound describes a response with status code 404, with default header values.

API ID Not Found
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound struct {
	Payload *models.APIResponse
}

func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /apiInventory/apiId/fromHostAndPortAndTraceSourceID][%d] getApiInventoryApiIdFromHostAndPortAndTraceSourceIdNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault creates a GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault with default headers values
func NewGetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault(code int) *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault {
	return &GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault{
		_statusCode: code,
	}
}

/* GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get API inventory API ID from host and port and trace source ID default response
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) Error() string {
	return fmt.Sprintf("[GET /apiInventory/apiId/fromHostAndPortAndTraceSourceID][%d] GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceID default  %+v", o._statusCode, o.Payload)
}
func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetAPIInventoryAPIIDFromHostAndPortAndTraceSourceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
