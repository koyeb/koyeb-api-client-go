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

// AcceptOrganizationInvitationRequest struct for AcceptOrganizationInvitationRequest
type AcceptOrganizationInvitationRequest struct {
	Id *string `json:"id,omitempty"`
}

// NewAcceptOrganizationInvitationRequest instantiates a new AcceptOrganizationInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptOrganizationInvitationRequest() *AcceptOrganizationInvitationRequest {
	this := AcceptOrganizationInvitationRequest{}
	return &this
}

// NewAcceptOrganizationInvitationRequestWithDefaults instantiates a new AcceptOrganizationInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptOrganizationInvitationRequestWithDefaults() *AcceptOrganizationInvitationRequest {
	this := AcceptOrganizationInvitationRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AcceptOrganizationInvitationRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOrganizationInvitationRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AcceptOrganizationInvitationRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AcceptOrganizationInvitationRequest) SetId(v string) {
	o.Id = &v
}

func (o AcceptOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptOrganizationInvitationRequest struct {
	value *AcceptOrganizationInvitationRequest
	isSet bool
}

func (v NullableAcceptOrganizationInvitationRequest) Get() *AcceptOrganizationInvitationRequest {
	return v.value
}

func (v *NullableAcceptOrganizationInvitationRequest) Set(val *AcceptOrganizationInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptOrganizationInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptOrganizationInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptOrganizationInvitationRequest(val *AcceptOrganizationInvitationRequest) *NullableAcceptOrganizationInvitationRequest {
	return &NullableAcceptOrganizationInvitationRequest{value: val, isSet: true}
}

func (v NullableAcceptOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptOrganizationInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


