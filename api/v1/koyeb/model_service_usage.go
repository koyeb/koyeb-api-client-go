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

// checks if the ServiceUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceUsage{}

// ServiceUsage struct for ServiceUsage
type ServiceUsage struct {
	ServiceId *string `json:"service_id,omitempty"`
	ServiceName *string `json:"service_name,omitempty"`
	Regions *map[string]RegionUsage `json:"regions,omitempty"`
}

// NewServiceUsage instantiates a new ServiceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUsage() *ServiceUsage {
	this := ServiceUsage{}
	return &this
}

// NewServiceUsageWithDefaults instantiates a new ServiceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUsageWithDefaults() *ServiceUsage {
	this := ServiceUsage{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServiceUsage) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsage) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServiceUsage) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ServiceUsage) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceUsage) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsage) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceUsage) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceUsage) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *ServiceUsage) GetRegions() map[string]RegionUsage {
	if o == nil || IsNil(o.Regions) {
		var ret map[string]RegionUsage
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUsage) GetRegionsOk() (*map[string]RegionUsage, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *ServiceUsage) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given map[string]RegionUsage and assigns it to the Regions field.
func (o *ServiceUsage) SetRegions(v map[string]RegionUsage) {
	o.Regions = &v
}

func (o ServiceUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return toSerialize, nil
}

type NullableServiceUsage struct {
	value *ServiceUsage
	isSet bool
}

func (v NullableServiceUsage) Get() *ServiceUsage {
	return v.value
}

func (v *NullableServiceUsage) Set(val *ServiceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUsage(val *ServiceUsage) *NullableServiceUsage {
	return &NullableServiceUsage{value: val, isSet: true}
}

func (v NullableServiceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


