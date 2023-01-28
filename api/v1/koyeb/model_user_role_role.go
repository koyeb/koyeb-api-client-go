/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// UserRoleRole the model 'UserRoleRole'
type UserRoleRole string

// List of UserRole.Role
const (
	USERROLEROLE_INVALID UserRoleRole = "INVALID"
	USERROLEROLE_OWNER UserRoleRole = "OWNER"
)

var allowedUserRoleRoleEnumValues = []UserRoleRole{
	"INVALID",
	"OWNER",
}

func (v *UserRoleRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserRoleRole(value)
	for _, existing := range allowedUserRoleRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserRoleRole", value)
}

// NewUserRoleRoleFromValue returns a pointer to a valid UserRoleRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserRoleRoleFromValue(v string) (*UserRoleRole, error) {
	ev := UserRoleRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserRoleRole: valid values are %v", v, allowedUserRoleRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserRoleRole) IsValid() bool {
	for _, existing := range allowedUserRoleRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserRole.Role value
func (v UserRoleRole) Ptr() *UserRoleRole {
	return &v
}

type NullableUserRoleRole struct {
	value *UserRoleRole
	isSet bool
}

func (v NullableUserRoleRole) Get() *UserRoleRole {
	return v.value
}

func (v *NullableUserRoleRole) Set(val *UserRoleRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRoleRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRoleRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRoleRole(val *UserRoleRole) *NullableUserRoleRole {
	return &NullableUserRoleRole{value: val, isSet: true}
}

func (v NullableUserRoleRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRoleRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

