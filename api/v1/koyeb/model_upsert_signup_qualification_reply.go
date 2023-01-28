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

// UpsertSignupQualificationReply struct for UpsertSignupQualificationReply
type UpsertSignupQualificationReply struct {
	Organization *Organization `json:"organization,omitempty"`
}

// NewUpsertSignupQualificationReply instantiates a new UpsertSignupQualificationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertSignupQualificationReply() *UpsertSignupQualificationReply {
	this := UpsertSignupQualificationReply{}
	return &this
}

// NewUpsertSignupQualificationReplyWithDefaults instantiates a new UpsertSignupQualificationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertSignupQualificationReplyWithDefaults() *UpsertSignupQualificationReply {
	this := UpsertSignupQualificationReply{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *UpsertSignupQualificationReply) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertSignupQualificationReply) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UpsertSignupQualificationReply) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *UpsertSignupQualificationReply) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o UpsertSignupQualificationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableUpsertSignupQualificationReply struct {
	value *UpsertSignupQualificationReply
	isSet bool
}

func (v NullableUpsertSignupQualificationReply) Get() *UpsertSignupQualificationReply {
	return v.value
}

func (v *NullableUpsertSignupQualificationReply) Set(val *UpsertSignupQualificationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertSignupQualificationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertSignupQualificationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertSignupQualificationReply(val *UpsertSignupQualificationReply) *NullableUpsertSignupQualificationReply {
	return &NullableUpsertSignupQualificationReply{value: val, isSet: true}
}

func (v NullableUpsertSignupQualificationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertSignupQualificationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

