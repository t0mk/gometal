// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FacilityList facility list
//
// swagger:model FacilityList
type FacilityList struct {

	// facilities
	Facilities []*Facility `json:"facilities"`
}

// Validate validates this facility list
func (m *FacilityList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityList) validateFacilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Facilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Facilities); i++ {
		if swag.IsZero(m.Facilities[i]) { // not required
			continue
		}

		if m.Facilities[i] != nil {
			if err := m.Facilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this facility list based on the context it is used
func (m *FacilityList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFacilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityList) contextValidateFacilities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Facilities); i++ {

		if m.Facilities[i] != nil {
			if err := m.Facilities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacilityList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityList) UnmarshalBinary(b []byte) error {
	var res FacilityList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}