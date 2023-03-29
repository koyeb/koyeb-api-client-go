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

// checks if the OAuthCallbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthCallbackRequest{}

// OAuthCallbackRequest struct for OAuthCallbackRequest
type OAuthCallbackRequest struct {
	State *string `json:"state,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewOAuthCallbackRequest instantiates a new OAuthCallbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthCallbackRequest() *OAuthCallbackRequest {
	this := OAuthCallbackRequest{}
	return &this
}

// NewOAuthCallbackRequestWithDefaults instantiates a new OAuthCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthCallbackRequestWithDefaults() *OAuthCallbackRequest {
	this := OAuthCallbackRequest{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OAuthCallbackRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCallbackRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OAuthCallbackRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OAuthCallbackRequest) SetState(v string) {
	o.State = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OAuthCallbackRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthCallbackRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OAuthCallbackRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OAuthCallbackRequest) SetCode(v string) {
	o.Code = &v
}

func (o OAuthCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthCallbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableOAuthCallbackRequest struct {
	value *OAuthCallbackRequest
	isSet bool
}

func (v NullableOAuthCallbackRequest) Get() *OAuthCallbackRequest {
	return v.value
}

func (v *NullableOAuthCallbackRequest) Set(val *OAuthCallbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthCallbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthCallbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthCallbackRequest(val *OAuthCallbackRequest) *NullableOAuthCallbackRequest {
	return &NullableOAuthCallbackRequest{value: val, isSet: true}
}

func (v NullableOAuthCallbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthCallbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


