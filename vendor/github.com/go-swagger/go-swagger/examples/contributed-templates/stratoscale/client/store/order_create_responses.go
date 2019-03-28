// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/go-swagger/go-swagger/examples/contributed-templates/stratoscale/models"
)

// OrderCreateReader is a Reader for the OrderCreate structure.
type OrderCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderCreateOK creates a OrderCreateOK with default headers values
func NewOrderCreateOK() *OrderCreateOK {
	return &OrderCreateOK{}
}

/*OrderCreateOK handles this case with default header values.

successful operation
*/
type OrderCreateOK struct {
	Payload *models.Order
}

func (o *OrderCreateOK) Error() string {
	return fmt.Sprintf("[POST /store/order][%d] orderCreateOK  %+v", 200, o.Payload)
}

func (o *OrderCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCreateBadRequest creates a OrderCreateBadRequest with default headers values
func NewOrderCreateBadRequest() *OrderCreateBadRequest {
	return &OrderCreateBadRequest{}
}

/*OrderCreateBadRequest handles this case with default header values.

Invalid Order
*/
type OrderCreateBadRequest struct {
}

func (o *OrderCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /store/order][%d] orderCreateBadRequest ", 400)
}

func (o *OrderCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
