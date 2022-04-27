// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/apiclarity/apiclarity/api/client/models"
)

// DeleteAPIInventoryAPIIDSpecsProvidedSpecReader is a Reader for the DeleteAPIInventoryAPIIDSpecsProvidedSpec structure.
type DeleteAPIInventoryAPIIDSpecsProvidedSpecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIInventoryAPIIDSpecsProvidedSpecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAPIInventoryAPIIDSpecsProvidedSpecDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAPIInventoryAPIIDSpecsProvidedSpecOK creates a DeleteAPIInventoryAPIIDSpecsProvidedSpecOK with default headers values
func NewDeleteAPIInventoryAPIIDSpecsProvidedSpecOK() *DeleteAPIInventoryAPIIDSpecsProvidedSpecOK {
	return &DeleteAPIInventoryAPIIDSpecsProvidedSpecOK{}
}

/* DeleteAPIInventoryAPIIDSpecsProvidedSpecOK describes a response with status code 200, with default header values.

Success
*/
type DeleteAPIInventoryAPIIDSpecsProvidedSpecOK struct {
	Payload interface{}
}

func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecOK) Error() string {
	return fmt.Sprintf("[DELETE /apiInventory/{apiId}/specs/providedSpec][%d] deleteApiInventoryApiIdSpecsProvidedSpecOK  %+v", 200, o.Payload)
}
func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIInventoryAPIIDSpecsProvidedSpecDefault creates a DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault with default headers values
func NewDeleteAPIInventoryAPIIDSpecsProvidedSpecDefault(code int) *DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault {
	return &DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault{
		_statusCode: code,
	}
}

/* DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault describes a response with status code -1, with default header values.

unknown error
*/
type DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the delete API inventory API ID specs provided spec default response
func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault) Error() string {
	return fmt.Sprintf("[DELETE /apiInventory/{apiId}/specs/providedSpec][%d] DeleteAPIInventoryAPIIDSpecsProvidedSpec default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *DeleteAPIInventoryAPIIDSpecsProvidedSpecDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
