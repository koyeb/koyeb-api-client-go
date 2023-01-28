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

// CreateOrganizationInvitationRequest struct for CreateOrganizationInvitationRequest
type CreateOrganizationInvitationRequest struct {
	Email *string `json:"email,omitempty"`
}

// NewCreateOrganizationInvitationRequest instantiates a new CreateOrganizationInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInvitationRequest() *CreateOrganizationInvitationRequest {
	this := CreateOrganizationInvitationRequest{}
	return &this
}

// NewCreateOrganizationInvitationRequestWithDefaults instantiates a new CreateOrganizationInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInvitationRequestWithDefaults() *CreateOrganizationInvitationRequest {
	this := CreateOrganizationInvitationRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateOrganizationInvitationRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInvitationRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateOrganizationInvitationRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateOrganizationInvitationRequest) SetEmail(v string) {
	o.Email = &v
}

func (o CreateOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInvitationRequest struct {
	value *CreateOrganizationInvitationRequest
	isSet bool
}

func (v NullableCreateOrganizationInvitationRequest) Get() *CreateOrganizationInvitationRequest {
	return v.value
}

func (v *NullableCreateOrganizationInvitationRequest) Set(val *CreateOrganizationInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInvitationRequest(val *CreateOrganizationInvitationRequest) *NullableCreateOrganizationInvitationRequest {
	return &NullableCreateOrganizationInvitationRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


