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

// RegionalDeploymentRawSource struct for RegionalDeploymentRawSource
type RegionalDeploymentRawSource struct {
	Content *string `json:"content,omitempty"`
}

// NewRegionalDeploymentRawSource instantiates a new RegionalDeploymentRawSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalDeploymentRawSource() *RegionalDeploymentRawSource {
	this := RegionalDeploymentRawSource{}
	return &this
}

// NewRegionalDeploymentRawSourceWithDefaults instantiates a new RegionalDeploymentRawSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalDeploymentRawSourceWithDefaults() *RegionalDeploymentRawSource {
	this := RegionalDeploymentRawSource{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *RegionalDeploymentRawSource) GetContent() string {
	if o == nil || isNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentRawSource) GetContentOk() (*string, bool) {
	if o == nil || isNil(o.Content) {
    return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *RegionalDeploymentRawSource) HasContent() bool {
	if o != nil && !isNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *RegionalDeploymentRawSource) SetContent(v string) {
	o.Content = &v
}

func (o RegionalDeploymentRawSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableRegionalDeploymentRawSource struct {
	value *RegionalDeploymentRawSource
	isSet bool
}

func (v NullableRegionalDeploymentRawSource) Get() *RegionalDeploymentRawSource {
	return v.value
}

func (v *NullableRegionalDeploymentRawSource) Set(val *RegionalDeploymentRawSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentRawSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentRawSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentRawSource(val *RegionalDeploymentRawSource) *NullableRegionalDeploymentRawSource {
	return &NullableRegionalDeploymentRawSource{value: val, isSet: true}
}

func (v NullableRegionalDeploymentRawSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentRawSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


