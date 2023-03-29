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

// checks if the CreateSecretReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSecretReply{}

// CreateSecretReply struct for CreateSecretReply
type CreateSecretReply struct {
	Secret *Secret `json:"secret,omitempty"`
}

// NewCreateSecretReply instantiates a new CreateSecretReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecretReply() *CreateSecretReply {
	this := CreateSecretReply{}
	return &this
}

// NewCreateSecretReplyWithDefaults instantiates a new CreateSecretReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecretReplyWithDefaults() *CreateSecretReply {
	this := CreateSecretReply{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreateSecretReply) GetSecret() Secret {
	if o == nil || IsNil(o.Secret) {
		var ret Secret
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecretReply) GetSecretOk() (*Secret, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreateSecretReply) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given Secret and assigns it to the Secret field.
func (o *CreateSecretReply) SetSecret(v Secret) {
	o.Secret = &v
}

func (o CreateSecretReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSecretReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableCreateSecretReply struct {
	value *CreateSecretReply
	isSet bool
}

func (v NullableCreateSecretReply) Get() *CreateSecretReply {
	return v.value
}

func (v *NullableCreateSecretReply) Set(val *CreateSecretReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecretReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecretReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecretReply(val *CreateSecretReply) *NullableCreateSecretReply {
	return &NullableCreateSecretReply{value: val, isSet: true}
}

func (v NullableCreateSecretReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecretReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


