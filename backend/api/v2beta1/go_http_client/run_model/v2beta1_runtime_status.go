// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2beta1RuntimeStatus Timestamped representation of a runtime state with an optional error.
// swagger:model v2beta1RuntimeStatus
type V2beta1RuntimeStatus struct {

	// The error that occurred during the state. May be set when the state is
	// any of the non-final states (PENDING/RUNNING/CANCELING) or FAILED state.
	// If the state is FAILED, the error here is final and not going to be
	// retried. If the state is a non-final state, the error indicates that a
	// system-error being retried.
	Error *RPCStatus `json:"error,omitempty"`

	// The state of a runtime instance.
	State V2beta1RuntimeState `json:"state,omitempty"`

	// Update time of this state.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this v2beta1 runtime status
func (m *V2beta1RuntimeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2beta1RuntimeStatus) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *V2beta1RuntimeStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *V2beta1RuntimeStatus) validateUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2beta1RuntimeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2beta1RuntimeStatus) UnmarshalBinary(b []byte) error {
	var res V2beta1RuntimeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}