// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// HostsList List of hosts
//
// swagger:model HostsList
type HostsList []string

// Validate validates this hosts list
func (m HostsList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hosts list based on context it is used
func (m HostsList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
