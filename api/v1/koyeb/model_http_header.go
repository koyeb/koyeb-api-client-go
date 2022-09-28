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

// HTTPHeader struct for HTTPHeader
type HTTPHeader struct {
	Key *string `json:"key,omitempty"`
	Values *[]string `json:"values,omitempty"`
}

// NewHTTPHeader instantiates a new HTTPHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTPHeader() *HTTPHeader {
	this := HTTPHeader{}
	return &this
}

// NewHTTPHeaderWithDefaults instantiates a new HTTPHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPHeaderWithDefaults() *HTTPHeader {
	this := HTTPHeader{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *HTTPHeader) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPHeader) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *HTTPHeader) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *HTTPHeader) SetKey(v string) {
	o.Key = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *HTTPHeader) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPHeader) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *HTTPHeader) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *HTTPHeader) SetValues(v []string) {
	o.Values = &v
}

func (o HTTPHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableHTTPHeader struct {
	value *HTTPHeader
	isSet bool
}

func (v NullableHTTPHeader) Get() *HTTPHeader {
	return v.value
}

func (v *NullableHTTPHeader) Set(val *HTTPHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTPHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableHTTPHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTPHeader(val *HTTPHeader) *NullableHTTPHeader {
	return &NullableHTTPHeader{value: val, isSet: true}
}

func (v NullableHTTPHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTPHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


