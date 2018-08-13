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

// NewContactInfo new contact info
// swagger:model newContactInfo
type NewContactInfo struct {

	// Address Line 1
	// Required: true
	Address1 *string `json:"address1"`

	// Address Line 2
	Address2 string `json:"address2,omitempty"`

	// City
	// Required: true
	AddressLocality *string `json:"addressLocality"`

	// Province / State
	// Required: true
	AddressRegion *string `json:"addressRegion"`

	// company name
	CompanyName string `json:"companyName,omitempty"`

	// Country, for best results please provide the two character ISO code
	// Required: true
	Country *string `json:"country"`

	// Email, required for international shipments
	// Required: true
	Email *string `json:"email"`

	// first name
	// Required: true
	FirstName *string `json:"firstName"`

	// last name
	// Required: true
	LastName *string `json:"lastName"`

	// Phone number, required for international shipments
	// Required: true
	Phone *string `json:"phone"`

	// Postal Code / Zip
	PostalCode string `json:"postalCode,omitempty"`
}

// Validate validates this new contact info
func (m *NewContactInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressLocality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewContactInfo) validateAddress1(formats strfmt.Registry) error {

	if err := validate.Required("address1", "body", m.Address1); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validateAddressLocality(formats strfmt.Registry) error {

	if err := validate.Required("addressLocality", "body", m.AddressLocality); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validateAddressRegion(formats strfmt.Registry) error {

	if err := validate.Required("addressRegion", "body", m.AddressRegion); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *NewContactInfo) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewContactInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewContactInfo) UnmarshalBinary(b []byte) error {
	var res NewContactInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}