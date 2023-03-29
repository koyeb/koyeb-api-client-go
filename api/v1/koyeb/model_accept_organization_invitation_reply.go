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

// checks if the AcceptOrganizationInvitationReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptOrganizationInvitationReply{}

// AcceptOrganizationInvitationReply struct for AcceptOrganizationInvitationReply
type AcceptOrganizationInvitationReply struct {
	Invitation *OrganizationInvitation `json:"invitation,omitempty"`
}

// NewAcceptOrganizationInvitationReply instantiates a new AcceptOrganizationInvitationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptOrganizationInvitationReply() *AcceptOrganizationInvitationReply {
	this := AcceptOrganizationInvitationReply{}
	return &this
}

// NewAcceptOrganizationInvitationReplyWithDefaults instantiates a new AcceptOrganizationInvitationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptOrganizationInvitationReplyWithDefaults() *AcceptOrganizationInvitationReply {
	this := AcceptOrganizationInvitationReply{}
	return &this
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *AcceptOrganizationInvitationReply) GetInvitation() OrganizationInvitation {
	if o == nil || IsNil(o.Invitation) {
		var ret OrganizationInvitation
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOrganizationInvitationReply) GetInvitationOk() (*OrganizationInvitation, bool) {
	if o == nil || IsNil(o.Invitation) {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *AcceptOrganizationInvitationReply) HasInvitation() bool {
	if o != nil && !IsNil(o.Invitation) {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given OrganizationInvitation and assigns it to the Invitation field.
func (o *AcceptOrganizationInvitationReply) SetInvitation(v OrganizationInvitation) {
	o.Invitation = &v
}

func (o AcceptOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptOrganizationInvitationReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invitation) {
		toSerialize["invitation"] = o.Invitation
	}
	return toSerialize, nil
}

type NullableAcceptOrganizationInvitationReply struct {
	value *AcceptOrganizationInvitationReply
	isSet bool
}

func (v NullableAcceptOrganizationInvitationReply) Get() *AcceptOrganizationInvitationReply {
	return v.value
}

func (v *NullableAcceptOrganizationInvitationReply) Set(val *AcceptOrganizationInvitationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptOrganizationInvitationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptOrganizationInvitationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptOrganizationInvitationReply(val *AcceptOrganizationInvitationReply) *NullableAcceptOrganizationInvitationReply {
	return &NullableAcceptOrganizationInvitationReply{value: val, isSet: true}
}

func (v NullableAcceptOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptOrganizationInvitationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


