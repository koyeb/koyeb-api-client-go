/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// OrganizationDeactivationReason the model 'OrganizationDeactivationReason'
type OrganizationDeactivationReason string

// List of Organization.DeactivationReason
const (
	ORGANIZATIONDEACTIVATIONREASON_INVALID OrganizationDeactivationReason = "INVALID"
	ORGANIZATIONDEACTIVATIONREASON_REQUESTED_BY_OWNER OrganizationDeactivationReason = "REQUESTED_BY_OWNER"
	ORGANIZATIONDEACTIVATIONREASON_SUBSCRIPTION_TERMINATION OrganizationDeactivationReason = "SUBSCRIPTION_TERMINATION"
	ORGANIZATIONDEACTIVATIONREASON_LOCKED_BY_ADMIN OrganizationDeactivationReason = "LOCKED_BY_ADMIN"
	ORGANIZATIONDEACTIVATIONREASON_VERIFICATION_FAILED OrganizationDeactivationReason = "VERIFICATION_FAILED"
	ORGANIZATIONDEACTIVATIONREASON_TRIAL_DID_NOT_CONVERT OrganizationDeactivationReason = "TRIAL_DID_NOT_CONVERT"
)

// All allowed values of OrganizationDeactivationReason enum
var AllowedOrganizationDeactivationReasonEnumValues = []OrganizationDeactivationReason{
	"INVALID",
	"REQUESTED_BY_OWNER",
	"SUBSCRIPTION_TERMINATION",
	"LOCKED_BY_ADMIN",
	"VERIFICATION_FAILED",
	"TRIAL_DID_NOT_CONVERT",
}

func (v *OrganizationDeactivationReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationDeactivationReason(value)
	for _, existing := range AllowedOrganizationDeactivationReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationDeactivationReason", value)
}

// NewOrganizationDeactivationReasonFromValue returns a pointer to a valid OrganizationDeactivationReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationDeactivationReasonFromValue(v string) (*OrganizationDeactivationReason, error) {
	ev := OrganizationDeactivationReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationDeactivationReason: valid values are %v", v, AllowedOrganizationDeactivationReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationDeactivationReason) IsValid() bool {
	for _, existing := range AllowedOrganizationDeactivationReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Organization.DeactivationReason value
func (v OrganizationDeactivationReason) Ptr() *OrganizationDeactivationReason {
	return &v
}

type NullableOrganizationDeactivationReason struct {
	value *OrganizationDeactivationReason
	isSet bool
}

func (v NullableOrganizationDeactivationReason) Get() *OrganizationDeactivationReason {
	return v.value
}

func (v *NullableOrganizationDeactivationReason) Set(val *OrganizationDeactivationReason) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDeactivationReason) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDeactivationReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDeactivationReason(val *OrganizationDeactivationReason) *NullableOrganizationDeactivationReason {
	return &NullableOrganizationDeactivationReason{value: val, isSet: true}
}

func (v NullableOrganizationDeactivationReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDeactivationReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

