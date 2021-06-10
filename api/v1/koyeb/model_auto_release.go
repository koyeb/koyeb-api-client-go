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
)

// AutoRelease struct for AutoRelease
type AutoRelease struct {
	Groups *[]AutoReleaseGroup `json:"groups,omitempty"`
}

// NewAutoRelease instantiates a new AutoRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoRelease() *AutoRelease {
	this := AutoRelease{}
	return &this
}

// NewAutoReleaseWithDefaults instantiates a new AutoRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoReleaseWithDefaults() *AutoRelease {
	this := AutoRelease{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *AutoRelease) GetGroups() []AutoReleaseGroup {
	if o == nil || o.Groups == nil {
		var ret []AutoReleaseGroup
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRelease) GetGroupsOk() (*[]AutoReleaseGroup, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *AutoRelease) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []AutoReleaseGroup and assigns it to the Groups field.
func (o *AutoRelease) SetGroups(v []AutoReleaseGroup) {
	o.Groups = &v
}

func (o AutoRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableAutoRelease struct {
	value *AutoRelease
	isSet bool
}

func (v NullableAutoRelease) Get() *AutoRelease {
	return v.value
}

func (v *NullableAutoRelease) Set(val *AutoRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoRelease(val *AutoRelease) *NullableAutoRelease {
	return &NullableAutoRelease{value: val, isSet: true}
}

func (v NullableAutoRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


