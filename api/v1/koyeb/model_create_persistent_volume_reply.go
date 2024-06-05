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

// CreatePersistentVolumeReply struct for CreatePersistentVolumeReply
type CreatePersistentVolumeReply struct {
	Volume *PersistentVolume `json:"volume,omitempty"`
}

// NewCreatePersistentVolumeReply instantiates a new CreatePersistentVolumeReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePersistentVolumeReply() *CreatePersistentVolumeReply {
	this := CreatePersistentVolumeReply{}
	return &this
}

// NewCreatePersistentVolumeReplyWithDefaults instantiates a new CreatePersistentVolumeReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePersistentVolumeReplyWithDefaults() *CreatePersistentVolumeReply {
	this := CreatePersistentVolumeReply{}
	return &this
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *CreatePersistentVolumeReply) GetVolume() PersistentVolume {
	if o == nil || isNil(o.Volume) {
		var ret PersistentVolume
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersistentVolumeReply) GetVolumeOk() (*PersistentVolume, bool) {
	if o == nil || isNil(o.Volume) {
    return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *CreatePersistentVolumeReply) HasVolume() bool {
	if o != nil && !isNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given PersistentVolume and assigns it to the Volume field.
func (o *CreatePersistentVolumeReply) SetVolume(v PersistentVolume) {
	o.Volume = &v
}

func (o CreatePersistentVolumeReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePersistentVolumeReply struct {
	value *CreatePersistentVolumeReply
	isSet bool
}

func (v NullableCreatePersistentVolumeReply) Get() *CreatePersistentVolumeReply {
	return v.value
}

func (v *NullableCreatePersistentVolumeReply) Set(val *CreatePersistentVolumeReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePersistentVolumeReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePersistentVolumeReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePersistentVolumeReply(val *CreatePersistentVolumeReply) *NullableCreatePersistentVolumeReply {
	return &NullableCreatePersistentVolumeReply{value: val, isSet: true}
}

func (v NullableCreatePersistentVolumeReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePersistentVolumeReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


