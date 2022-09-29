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

// RegionUsage struct for RegionUsage
type RegionUsage struct {
	Instances *map[string]InstanceUsage `json:"instances,omitempty"`
}

// NewRegionUsage instantiates a new RegionUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionUsage() *RegionUsage {
	this := RegionUsage{}
	return &this
}

// NewRegionUsageWithDefaults instantiates a new RegionUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionUsageWithDefaults() *RegionUsage {
	this := RegionUsage{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *RegionUsage) GetInstances() map[string]InstanceUsage {
	if o == nil || o.Instances == nil {
		var ret map[string]InstanceUsage
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionUsage) GetInstancesOk() (*map[string]InstanceUsage, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *RegionUsage) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given map[string]InstanceUsage and assigns it to the Instances field.
func (o *RegionUsage) SetInstances(v map[string]InstanceUsage) {
	o.Instances = &v
}

func (o RegionUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	return json.Marshal(toSerialize)
}

type NullableRegionUsage struct {
	value *RegionUsage
	isSet bool
}

func (v NullableRegionUsage) Get() *RegionUsage {
	return v.value
}

func (v *NullableRegionUsage) Set(val *RegionUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionUsage(val *RegionUsage) *NullableRegionUsage {
	return &NullableRegionUsage{value: val, isSet: true}
}

func (v NullableRegionUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


