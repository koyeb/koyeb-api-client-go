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

// ExecCommandIO struct for ExecCommandIO
type ExecCommandIO struct {
	// Data is base64 encoded
	Data *string `json:"data,omitempty"`
	// Indicate last data frame
	Close *bool `json:"close,omitempty"`
}

// NewExecCommandIO instantiates a new ExecCommandIO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecCommandIO() *ExecCommandIO {
	this := ExecCommandIO{}
	return &this
}

// NewExecCommandIOWithDefaults instantiates a new ExecCommandIO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecCommandIOWithDefaults() *ExecCommandIO {
	this := ExecCommandIO{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExecCommandIO) GetData() string {
	if o == nil || isNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandIO) GetDataOk() (*string, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExecCommandIO) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ExecCommandIO) SetData(v string) {
	o.Data = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *ExecCommandIO) GetClose() bool {
	if o == nil || isNil(o.Close) {
		var ret bool
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandIO) GetCloseOk() (*bool, bool) {
	if o == nil || isNil(o.Close) {
    return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *ExecCommandIO) HasClose() bool {
	if o != nil && !isNil(o.Close) {
		return true
	}

	return false
}

// SetClose gets a reference to the given bool and assigns it to the Close field.
func (o *ExecCommandIO) SetClose(v bool) {
	o.Close = &v
}

func (o ExecCommandIO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Close) {
		toSerialize["close"] = o.Close
	}
	return json.Marshal(toSerialize)
}

type NullableExecCommandIO struct {
	value *ExecCommandIO
	isSet bool
}

func (v NullableExecCommandIO) Get() *ExecCommandIO {
	return v.value
}

func (v *NullableExecCommandIO) Set(val *ExecCommandIO) {
	v.value = val
	v.isSet = true
}

func (v NullableExecCommandIO) IsSet() bool {
	return v.isSet
}

func (v *NullableExecCommandIO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecCommandIO(val *ExecCommandIO) *NullableExecCommandIO {
	return &NullableExecCommandIO{value: val, isSet: true}
}

func (v NullableExecCommandIO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecCommandIO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


