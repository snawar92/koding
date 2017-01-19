package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JAPIToken j Api token
// swagger:model JApiToken
type JAPIToken struct {

	// id
	ID string `json:"_id,omitempty"`

	// code
	// Required: true
	Code *string `json:"code"`

	// created at
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// group
	// Required: true
	Group *string `json:"group"`

	// origin Id
	// Required: true
	OriginID *string `json:"originId"`
}

// Validate validates this j Api token
func (m *JAPIToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOriginID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JAPIToken) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *JAPIToken) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *JAPIToken) validateOriginID(formats strfmt.Registry) error {

	if err := validate.Required("originId", "body", m.OriginID); err != nil {
		return err
	}

	return nil
}
