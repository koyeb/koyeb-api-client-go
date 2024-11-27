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

// DeploymentScalingTargetSleepIdleDelay struct for DeploymentScalingTargetSleepIdleDelay
type DeploymentScalingTargetSleepIdleDelay struct {
	// Delay in seconds after which a service which received 0 request is scaled to 0.
	Value *int64 `json:"value,omitempty"`
}

// NewDeploymentScalingTargetSleepIdleDelay instantiates a new DeploymentScalingTargetSleepIdleDelay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentScalingTargetSleepIdleDelay() *DeploymentScalingTargetSleepIdleDelay {
	this := DeploymentScalingTargetSleepIdleDelay{}
	return &this
}

// NewDeploymentScalingTargetSleepIdleDelayWithDefaults instantiates a new DeploymentScalingTargetSleepIdleDelay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentScalingTargetSleepIdleDelayWithDefaults() *DeploymentScalingTargetSleepIdleDelay {
	this := DeploymentScalingTargetSleepIdleDelay{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeploymentScalingTargetSleepIdleDelay) GetValue() int64 {
	if o == nil || isNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTargetSleepIdleDelay) GetValueOk() (*int64, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeploymentScalingTargetSleepIdleDelay) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *DeploymentScalingTargetSleepIdleDelay) SetValue(v int64) {
	o.Value = &v
}

func (o DeploymentScalingTargetSleepIdleDelay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentScalingTargetSleepIdleDelay struct {
	value *DeploymentScalingTargetSleepIdleDelay
	isSet bool
}

func (v NullableDeploymentScalingTargetSleepIdleDelay) Get() *DeploymentScalingTargetSleepIdleDelay {
	return v.value
}

func (v *NullableDeploymentScalingTargetSleepIdleDelay) Set(val *DeploymentScalingTargetSleepIdleDelay) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentScalingTargetSleepIdleDelay) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentScalingTargetSleepIdleDelay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentScalingTargetSleepIdleDelay(val *DeploymentScalingTargetSleepIdleDelay) *NullableDeploymentScalingTargetSleepIdleDelay {
	return &NullableDeploymentScalingTargetSleepIdleDelay{value: val, isSet: true}
}

func (v NullableDeploymentScalingTargetSleepIdleDelay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentScalingTargetSleepIdleDelay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

