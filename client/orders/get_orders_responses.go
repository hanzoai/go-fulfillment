// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "go-fulfillment/models"
)

// GetOrdersReader is a Reader for the GetOrders structure.
type GetOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrdersOK creates a GetOrdersOK with default headers values
func NewGetOrdersOK() *GetOrdersOK {
	return &GetOrdersOK{}
}

/*GetOrdersOK handles this case with default header values.

Found Orders
*/
type GetOrdersOK struct {
	Payload *models.OrdersArray
}

func (o *GetOrdersOK) Error() string {
	return fmt.Sprintf("[GET /orders][%d] getOrdersOK  %+v", 200, o.Payload)
}

func (o *GetOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrdersArray)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersNotFound creates a GetOrdersNotFound with default headers values
func NewGetOrdersNotFound() *GetOrdersNotFound {
	return &GetOrdersNotFound{}
}

/*GetOrdersNotFound handles this case with default header values.

No Orders Found
*/
type GetOrdersNotFound struct {
	Payload *models.StandardError
}

func (o *GetOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /orders][%d] getOrdersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
