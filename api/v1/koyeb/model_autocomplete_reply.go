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

// AutocompleteReply struct for AutocompleteReply
type AutocompleteReply struct {
	Secrets []string `json:"secrets,omitempty"`
	UserEnv []string `json:"user_env,omitempty"`
	SystemEnv []string `json:"system_env,omitempty"`
}

// NewAutocompleteReply instantiates a new AutocompleteReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteReply() *AutocompleteReply {
	this := AutocompleteReply{}
	return &this
}

// NewAutocompleteReplyWithDefaults instantiates a new AutocompleteReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteReplyWithDefaults() *AutocompleteReply {
	this := AutocompleteReply{}
	return &this
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *AutocompleteReply) GetSecrets() []string {
	if o == nil || isNil(o.Secrets) {
		var ret []string
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteReply) GetSecretsOk() ([]string, bool) {
	if o == nil || isNil(o.Secrets) {
    return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *AutocompleteReply) HasSecrets() bool {
	if o != nil && !isNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []string and assigns it to the Secrets field.
func (o *AutocompleteReply) SetSecrets(v []string) {
	o.Secrets = v
}

// GetUserEnv returns the UserEnv field value if set, zero value otherwise.
func (o *AutocompleteReply) GetUserEnv() []string {
	if o == nil || isNil(o.UserEnv) {
		var ret []string
		return ret
	}
	return o.UserEnv
}

// GetUserEnvOk returns a tuple with the UserEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteReply) GetUserEnvOk() ([]string, bool) {
	if o == nil || isNil(o.UserEnv) {
    return nil, false
	}
	return o.UserEnv, true
}

// HasUserEnv returns a boolean if a field has been set.
func (o *AutocompleteReply) HasUserEnv() bool {
	if o != nil && !isNil(o.UserEnv) {
		return true
	}

	return false
}

// SetUserEnv gets a reference to the given []string and assigns it to the UserEnv field.
func (o *AutocompleteReply) SetUserEnv(v []string) {
	o.UserEnv = v
}

// GetSystemEnv returns the SystemEnv field value if set, zero value otherwise.
func (o *AutocompleteReply) GetSystemEnv() []string {
	if o == nil || isNil(o.SystemEnv) {
		var ret []string
		return ret
	}
	return o.SystemEnv
}

// GetSystemEnvOk returns a tuple with the SystemEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteReply) GetSystemEnvOk() ([]string, bool) {
	if o == nil || isNil(o.SystemEnv) {
    return nil, false
	}
	return o.SystemEnv, true
}

// HasSystemEnv returns a boolean if a field has been set.
func (o *AutocompleteReply) HasSystemEnv() bool {
	if o != nil && !isNil(o.SystemEnv) {
		return true
	}

	return false
}

// SetSystemEnv gets a reference to the given []string and assigns it to the SystemEnv field.
func (o *AutocompleteReply) SetSystemEnv(v []string) {
	o.SystemEnv = v
}

func (o AutocompleteReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	if !isNil(o.UserEnv) {
		toSerialize["user_env"] = o.UserEnv
	}
	if !isNil(o.SystemEnv) {
		toSerialize["system_env"] = o.SystemEnv
	}
	return json.Marshal(toSerialize)
}

type NullableAutocompleteReply struct {
	value *AutocompleteReply
	isSet bool
}

func (v NullableAutocompleteReply) Get() *AutocompleteReply {
	return v.value
}

func (v *NullableAutocompleteReply) Set(val *AutocompleteReply) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteReply) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteReply(val *AutocompleteReply) *NullableAutocompleteReply {
	return &NullableAutocompleteReply{value: val, isSet: true}
}

func (v NullableAutocompleteReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


