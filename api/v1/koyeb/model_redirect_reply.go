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

// RedirectReply struct for RedirectReply
type RedirectReply struct {
	Url *string `json:"url,omitempty"`
}

// NewRedirectReply instantiates a new RedirectReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectReply() *RedirectReply {
	this := RedirectReply{}
	return &this
}

// NewRedirectReplyWithDefaults instantiates a new RedirectReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectReplyWithDefaults() *RedirectReply {
	this := RedirectReply{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RedirectReply) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectReply) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RedirectReply) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RedirectReply) SetUrl(v string) {
	o.Url = &v
}

func (o RedirectReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectReply struct {
	value *RedirectReply
	isSet bool
}

func (v NullableRedirectReply) Get() *RedirectReply {
	return v.value
}

func (v *NullableRedirectReply) Set(val *RedirectReply) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectReply) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectReply(val *RedirectReply) *NullableRedirectReply {
	return &NullableRedirectReply{value: val, isSet: true}
}

func (v NullableRedirectReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

