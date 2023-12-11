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

// MembersSummary struct for MembersSummary
type MembersSummary struct {
	Total *string `json:"total,omitempty"`
	InvitationsByStatus *map[string]string `json:"invitations_by_status,omitempty"`
}

// NewMembersSummary instantiates a new MembersSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembersSummary() *MembersSummary {
	this := MembersSummary{}
	return &this
}

// NewMembersSummaryWithDefaults instantiates a new MembersSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembersSummaryWithDefaults() *MembersSummary {
	this := MembersSummary{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MembersSummary) GetTotal() string {
	if o == nil || isNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembersSummary) GetTotalOk() (*string, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MembersSummary) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *MembersSummary) SetTotal(v string) {
	o.Total = &v
}

// GetInvitationsByStatus returns the InvitationsByStatus field value if set, zero value otherwise.
func (o *MembersSummary) GetInvitationsByStatus() map[string]string {
	if o == nil || isNil(o.InvitationsByStatus) {
		var ret map[string]string
		return ret
	}
	return *o.InvitationsByStatus
}

// GetInvitationsByStatusOk returns a tuple with the InvitationsByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembersSummary) GetInvitationsByStatusOk() (*map[string]string, bool) {
	if o == nil || isNil(o.InvitationsByStatus) {
    return nil, false
	}
	return o.InvitationsByStatus, true
}

// HasInvitationsByStatus returns a boolean if a field has been set.
func (o *MembersSummary) HasInvitationsByStatus() bool {
	if o != nil && !isNil(o.InvitationsByStatus) {
		return true
	}

	return false
}

// SetInvitationsByStatus gets a reference to the given map[string]string and assigns it to the InvitationsByStatus field.
func (o *MembersSummary) SetInvitationsByStatus(v map[string]string) {
	o.InvitationsByStatus = &v
}

func (o MembersSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.InvitationsByStatus) {
		toSerialize["invitations_by_status"] = o.InvitationsByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableMembersSummary struct {
	value *MembersSummary
	isSet bool
}

func (v NullableMembersSummary) Get() *MembersSummary {
	return v.value
}

func (v *NullableMembersSummary) Set(val *MembersSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMembersSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMembersSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembersSummary(val *MembersSummary) *NullableMembersSummary {
	return &NullableMembersSummary{value: val, isSet: true}
}

func (v NullableMembersSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembersSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


