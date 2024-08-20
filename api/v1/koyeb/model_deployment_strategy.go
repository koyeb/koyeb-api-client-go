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

// DeploymentStrategy struct for DeploymentStrategy
type DeploymentStrategy struct {
	Type *DeploymentStrategyType `json:"type,omitempty"`
}

// NewDeploymentStrategy instantiates a new DeploymentStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStrategy() *DeploymentStrategy {
	this := DeploymentStrategy{}
	var type_ DeploymentStrategyType = DEPLOYMENTSTRATEGYTYPE_INVALID
	this.Type = &type_
	return &this
}

// NewDeploymentStrategyWithDefaults instantiates a new DeploymentStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStrategyWithDefaults() *DeploymentStrategy {
	this := DeploymentStrategy{}
	var type_ DeploymentStrategyType = DEPLOYMENTSTRATEGYTYPE_INVALID
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentStrategy) GetType() DeploymentStrategyType {
	if o == nil || isNil(o.Type) {
		var ret DeploymentStrategyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStrategy) GetTypeOk() (*DeploymentStrategyType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentStrategy) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DeploymentStrategyType and assigns it to the Type field.
func (o *DeploymentStrategy) SetType(v DeploymentStrategyType) {
	o.Type = &v
}

func (o DeploymentStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStrategy struct {
	value *DeploymentStrategy
	isSet bool
}

func (v NullableDeploymentStrategy) Get() *DeploymentStrategy {
	return v.value
}

func (v *NullableDeploymentStrategy) Set(val *DeploymentStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStrategy(val *DeploymentStrategy) *NullableDeploymentStrategy {
	return &NullableDeploymentStrategy{value: val, isSet: true}
}

func (v NullableDeploymentStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


