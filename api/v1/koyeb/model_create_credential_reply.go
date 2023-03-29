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

// checks if the CreateCredentialReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCredentialReply{}

// CreateCredentialReply struct for CreateCredentialReply
type CreateCredentialReply struct {
	Credential *Credential `json:"credential,omitempty"`
}

// NewCreateCredentialReply instantiates a new CreateCredentialReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCredentialReply() *CreateCredentialReply {
	this := CreateCredentialReply{}
	return &this
}

// NewCreateCredentialReplyWithDefaults instantiates a new CreateCredentialReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCredentialReplyWithDefaults() *CreateCredentialReply {
	this := CreateCredentialReply{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *CreateCredentialReply) GetCredential() Credential {
	if o == nil || IsNil(o.Credential) {
		var ret Credential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCredentialReply) GetCredentialOk() (*Credential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *CreateCredentialReply) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given Credential and assigns it to the Credential field.
func (o *CreateCredentialReply) SetCredential(v Credential) {
	o.Credential = &v
}

func (o CreateCredentialReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCredentialReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	return toSerialize, nil
}

type NullableCreateCredentialReply struct {
	value *CreateCredentialReply
	isSet bool
}

func (v NullableCreateCredentialReply) Get() *CreateCredentialReply {
	return v.value
}

func (v *NullableCreateCredentialReply) Set(val *CreateCredentialReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCredentialReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCredentialReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCredentialReply(val *CreateCredentialReply) *NullableCreateCredentialReply {
	return &NullableCreateCredentialReply{value: val, isSet: true}
}

func (v NullableCreateCredentialReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCredentialReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


