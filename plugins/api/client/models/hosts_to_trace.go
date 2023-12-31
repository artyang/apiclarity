// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostsToTrace List of hosts to trace
//
// swagger:model HostsToTrace
type HostsToTrace struct {

	// hosts
	// Required: true
	Hosts HostsList `json:"hosts"`
}

// Validate validates this hosts to trace
func (m *HostsToTrace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostsToTrace) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("hosts", "body", m.Hosts); err != nil {
		return err
	}

	if err := m.Hosts.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hosts")
		}
		return err
	}

	return nil
}

// ContextValidate validate this hosts to trace based on the context it is used
func (m *HostsToTrace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostsToTrace) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Hosts.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hosts")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostsToTrace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostsToTrace) UnmarshalBinary(b []byte) error {
	var res HostsToTrace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
