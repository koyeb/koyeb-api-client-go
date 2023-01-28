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

// DeclineOrganizationInvitationReply struct for DeclineOrganizationInvitationReply
type DeclineOrganizationInvitationReply struct {
	Invitation *OrganizationInvitation `json:"invitation,omitempty"`
}

// NewDeclineOrganizationInvitationReply instantiates a new DeclineOrganizationInvitationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeclineOrganizationInvitationReply() *DeclineOrganizationInvitationReply {
	this := DeclineOrganizationInvitationReply{}
	return &this
}

// NewDeclineOrganizationInvitationReplyWithDefaults instantiates a new DeclineOrganizationInvitationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeclineOrganizationInvitationReplyWithDefaults() *DeclineOrganizationInvitationReply {
	this := DeclineOrganizationInvitationReply{}
	return &this
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *DeclineOrganizationInvitationReply) GetInvitation() OrganizationInvitation {
	if o == nil || o.Invitation == nil {
		var ret OrganizationInvitation
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeclineOrganizationInvitationReply) GetInvitationOk() (*OrganizationInvitation, bool) {
	if o == nil || o.Invitation == nil {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *DeclineOrganizationInvitationReply) HasInvitation() bool {
	if o != nil && o.Invitation != nil {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given OrganizationInvitation and assigns it to the Invitation field.
func (o *DeclineOrganizationInvitationReply) SetInvitation(v OrganizationInvitation) {
	o.Invitation = &v
}

func (o DeclineOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invitation != nil {
		toSerialize["invitation"] = o.Invitation
	}
	return json.Marshal(toSerialize)
}

type NullableDeclineOrganizationInvitationReply struct {
	value *DeclineOrganizationInvitationReply
	isSet bool
}

func (v NullableDeclineOrganizationInvitationReply) Get() *DeclineOrganizationInvitationReply {
	return v.value
}

func (v *NullableDeclineOrganizationInvitationReply) Set(val *DeclineOrganizationInvitationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableDeclineOrganizationInvitationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableDeclineOrganizationInvitationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeclineOrganizationInvitationReply(val *DeclineOrganizationInvitationReply) *NullableDeclineOrganizationInvitationReply {
	return &NullableDeclineOrganizationInvitationReply{value: val, isSet: true}
}

func (v NullableDeclineOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeclineOrganizationInvitationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

