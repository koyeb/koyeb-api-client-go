/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// checks if the GetUserOrganizationInvitationReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserOrganizationInvitationReply{}

// GetUserOrganizationInvitationReply struct for GetUserOrganizationInvitationReply
type GetUserOrganizationInvitationReply struct {
	Invitation *OrganizationInvitation `json:"invitation,omitempty"`
}

// NewGetUserOrganizationInvitationReply instantiates a new GetUserOrganizationInvitationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserOrganizationInvitationReply() *GetUserOrganizationInvitationReply {
	this := GetUserOrganizationInvitationReply{}
	return &this
}

// NewGetUserOrganizationInvitationReplyWithDefaults instantiates a new GetUserOrganizationInvitationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserOrganizationInvitationReplyWithDefaults() *GetUserOrganizationInvitationReply {
	this := GetUserOrganizationInvitationReply{}
	return &this
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *GetUserOrganizationInvitationReply) GetInvitation() OrganizationInvitation {
	if o == nil || IsNil(o.Invitation) {
		var ret OrganizationInvitation
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserOrganizationInvitationReply) GetInvitationOk() (*OrganizationInvitation, bool) {
	if o == nil || IsNil(o.Invitation) {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *GetUserOrganizationInvitationReply) HasInvitation() bool {
	if o != nil && !IsNil(o.Invitation) {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given OrganizationInvitation and assigns it to the Invitation field.
func (o *GetUserOrganizationInvitationReply) SetInvitation(v OrganizationInvitation) {
	o.Invitation = &v
}

func (o GetUserOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserOrganizationInvitationReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invitation) {
		toSerialize["invitation"] = o.Invitation
	}
	return toSerialize, nil
}

type NullableGetUserOrganizationInvitationReply struct {
	value *GetUserOrganizationInvitationReply
	isSet bool
}

func (v NullableGetUserOrganizationInvitationReply) Get() *GetUserOrganizationInvitationReply {
	return v.value
}

func (v *NullableGetUserOrganizationInvitationReply) Set(val *GetUserOrganizationInvitationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserOrganizationInvitationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserOrganizationInvitationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserOrganizationInvitationReply(val *GetUserOrganizationInvitationReply) *NullableGetUserOrganizationInvitationReply {
	return &NullableGetUserOrganizationInvitationReply{value: val, isSet: true}
}

func (v NullableGetUserOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserOrganizationInvitationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


