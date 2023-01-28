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

// ListUserOrganizationInvitationsReply struct for ListUserOrganizationInvitationsReply
type ListUserOrganizationInvitationsReply struct {
	Invitations *[]OrganizationInvitation `json:"invitations,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListUserOrganizationInvitationsReply instantiates a new ListUserOrganizationInvitationsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserOrganizationInvitationsReply() *ListUserOrganizationInvitationsReply {
	this := ListUserOrganizationInvitationsReply{}
	return &this
}

// NewListUserOrganizationInvitationsReplyWithDefaults instantiates a new ListUserOrganizationInvitationsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserOrganizationInvitationsReplyWithDefaults() *ListUserOrganizationInvitationsReply {
	this := ListUserOrganizationInvitationsReply{}
	return &this
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *ListUserOrganizationInvitationsReply) GetInvitations() []OrganizationInvitation {
	if o == nil || o.Invitations == nil {
		var ret []OrganizationInvitation
		return ret
	}
	return *o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserOrganizationInvitationsReply) GetInvitationsOk() (*[]OrganizationInvitation, bool) {
	if o == nil || o.Invitations == nil {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *ListUserOrganizationInvitationsReply) HasInvitations() bool {
	if o != nil && o.Invitations != nil {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []OrganizationInvitation and assigns it to the Invitations field.
func (o *ListUserOrganizationInvitationsReply) SetInvitations(v []OrganizationInvitation) {
	o.Invitations = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListUserOrganizationInvitationsReply) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserOrganizationInvitationsReply) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListUserOrganizationInvitationsReply) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListUserOrganizationInvitationsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListUserOrganizationInvitationsReply) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserOrganizationInvitationsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListUserOrganizationInvitationsReply) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListUserOrganizationInvitationsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListUserOrganizationInvitationsReply) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserOrganizationInvitationsReply) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListUserOrganizationInvitationsReply) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListUserOrganizationInvitationsReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListUserOrganizationInvitationsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invitations != nil {
		toSerialize["invitations"] = o.Invitations
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableListUserOrganizationInvitationsReply struct {
	value *ListUserOrganizationInvitationsReply
	isSet bool
}

func (v NullableListUserOrganizationInvitationsReply) Get() *ListUserOrganizationInvitationsReply {
	return v.value
}

func (v *NullableListUserOrganizationInvitationsReply) Set(val *ListUserOrganizationInvitationsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserOrganizationInvitationsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserOrganizationInvitationsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserOrganizationInvitationsReply(val *ListUserOrganizationInvitationsReply) *NullableListUserOrganizationInvitationsReply {
	return &NullableListUserOrganizationInvitationsReply{value: val, isSet: true}
}

func (v NullableListUserOrganizationInvitationsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserOrganizationInvitationsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


