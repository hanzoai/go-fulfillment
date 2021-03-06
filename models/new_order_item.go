// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewOrderItem new order item
// swagger:model newOrderItem
type NewOrderItem struct {

	// Using USD, a per unit value of a single SKU. If your declaring a kit this is the sum total for a single kit. In both scenarios we will multiply the `declaredValue` against the `quantity` for customs declaration.
	// Required: true
	DeclaredValue *float64 `json:"declaredValue"`

	// quantity
	// Required: true
	// Minimum: 1
	Quantity *int64 `json:"quantity"`

	// sku
	// Required: true
	Sku *string `json:"sku"`
}

// Validate validates this new order item
func (m *NewOrderItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeclaredValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewOrderItem) validateDeclaredValue(formats strfmt.Registry) error {

	if err := validate.Required("declaredValue", "body", m.DeclaredValue); err != nil {
		return err
	}

	return nil
}

func (m *NewOrderItem) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.MinimumInt("quantity", "body", int64(*m.Quantity), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *NewOrderItem) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewOrderItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewOrderItem) UnmarshalBinary(b []byte) error {
	var res NewOrderItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
