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

// CatalogGPUDetails struct for CatalogGPUDetails
type CatalogGPUDetails struct {
	Count *int64 `json:"count,omitempty"`
	Brand *string `json:"brand,omitempty"`
	Memory *string `json:"memory,omitempty"`
}

// NewCatalogGPUDetails instantiates a new CatalogGPUDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogGPUDetails() *CatalogGPUDetails {
	this := CatalogGPUDetails{}
	return &this
}

// NewCatalogGPUDetailsWithDefaults instantiates a new CatalogGPUDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogGPUDetailsWithDefaults() *CatalogGPUDetails {
	this := CatalogGPUDetails{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CatalogGPUDetails) GetCount() int64 {
	if o == nil || isNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogGPUDetails) GetCountOk() (*int64, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CatalogGPUDetails) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *CatalogGPUDetails) SetCount(v int64) {
	o.Count = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *CatalogGPUDetails) GetBrand() string {
	if o == nil || isNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogGPUDetails) GetBrandOk() (*string, bool) {
	if o == nil || isNil(o.Brand) {
    return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *CatalogGPUDetails) HasBrand() bool {
	if o != nil && !isNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *CatalogGPUDetails) SetBrand(v string) {
	o.Brand = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CatalogGPUDetails) GetMemory() string {
	if o == nil || isNil(o.Memory) {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogGPUDetails) GetMemoryOk() (*string, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CatalogGPUDetails) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *CatalogGPUDetails) SetMemory(v string) {
	o.Memory = &v
}

func (o CatalogGPUDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !isNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogGPUDetails struct {
	value *CatalogGPUDetails
	isSet bool
}

func (v NullableCatalogGPUDetails) Get() *CatalogGPUDetails {
	return v.value
}

func (v *NullableCatalogGPUDetails) Set(val *CatalogGPUDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogGPUDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogGPUDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogGPUDetails(val *CatalogGPUDetails) *NullableCatalogGPUDetails {
	return &NullableCatalogGPUDetails{value: val, isSet: true}
}

func (v NullableCatalogGPUDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogGPUDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

