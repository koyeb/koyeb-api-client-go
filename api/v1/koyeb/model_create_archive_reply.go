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

// CreateArchiveReply struct for CreateArchiveReply
type CreateArchiveReply struct {
	Archive *Archive `json:"archive,omitempty"`
}

// NewCreateArchiveReply instantiates a new CreateArchiveReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArchiveReply() *CreateArchiveReply {
	this := CreateArchiveReply{}
	return &this
}

// NewCreateArchiveReplyWithDefaults instantiates a new CreateArchiveReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArchiveReplyWithDefaults() *CreateArchiveReply {
	this := CreateArchiveReply{}
	return &this
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *CreateArchiveReply) GetArchive() Archive {
	if o == nil || isNil(o.Archive) {
		var ret Archive
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArchiveReply) GetArchiveOk() (*Archive, bool) {
	if o == nil || isNil(o.Archive) {
    return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *CreateArchiveReply) HasArchive() bool {
	if o != nil && !isNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given Archive and assigns it to the Archive field.
func (o *CreateArchiveReply) SetArchive(v Archive) {
	o.Archive = &v
}

func (o CreateArchiveReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	return json.Marshal(toSerialize)
}

type NullableCreateArchiveReply struct {
	value *CreateArchiveReply
	isSet bool
}

func (v NullableCreateArchiveReply) Get() *CreateArchiveReply {
	return v.value
}

func (v *NullableCreateArchiveReply) Set(val *CreateArchiveReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArchiveReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArchiveReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArchiveReply(val *CreateArchiveReply) *NullableCreateArchiveReply {
	return &NullableCreateArchiveReply{value: val, isSet: true}
}

func (v NullableCreateArchiveReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArchiveReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


