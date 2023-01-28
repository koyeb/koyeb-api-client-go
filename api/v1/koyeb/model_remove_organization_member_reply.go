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

// RemoveOrganizationMemberReply struct for RemoveOrganizationMemberReply
type RemoveOrganizationMemberReply struct {
	Member *OrganizationMember `json:"member,omitempty"`
}

// NewRemoveOrganizationMemberReply instantiates a new RemoveOrganizationMemberReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveOrganizationMemberReply() *RemoveOrganizationMemberReply {
	this := RemoveOrganizationMemberReply{}
	return &this
}

// NewRemoveOrganizationMemberReplyWithDefaults instantiates a new RemoveOrganizationMemberReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveOrganizationMemberReplyWithDefaults() *RemoveOrganizationMemberReply {
	this := RemoveOrganizationMemberReply{}
	return &this
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *RemoveOrganizationMemberReply) GetMember() OrganizationMember {
	if o == nil || o.Member == nil {
		var ret OrganizationMember
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveOrganizationMemberReply) GetMemberOk() (*OrganizationMember, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *RemoveOrganizationMemberReply) HasMember() bool {
	if o != nil && o.Member != nil {
		return true
	}

	return false
}

// SetMember gets a reference to the given OrganizationMember and assigns it to the Member field.
func (o *RemoveOrganizationMemberReply) SetMember(v OrganizationMember) {
	o.Member = &v
}

func (o RemoveOrganizationMemberReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveOrganizationMemberReply struct {
	value *RemoveOrganizationMemberReply
	isSet bool
}

func (v NullableRemoveOrganizationMemberReply) Get() *RemoveOrganizationMemberReply {
	return v.value
}

func (v *NullableRemoveOrganizationMemberReply) Set(val *RemoveOrganizationMemberReply) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveOrganizationMemberReply) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveOrganizationMemberReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveOrganizationMemberReply(val *RemoveOrganizationMemberReply) *NullableRemoveOrganizationMemberReply {
	return &NullableRemoveOrganizationMemberReply{value: val, isSet: true}
}

func (v NullableRemoveOrganizationMemberReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveOrganizationMemberReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

