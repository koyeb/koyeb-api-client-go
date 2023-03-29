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

// checks if the GetDeploymentReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeploymentReply{}

// GetDeploymentReply struct for GetDeploymentReply
type GetDeploymentReply struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

// NewGetDeploymentReply instantiates a new GetDeploymentReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeploymentReply() *GetDeploymentReply {
	this := GetDeploymentReply{}
	return &this
}

// NewGetDeploymentReplyWithDefaults instantiates a new GetDeploymentReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeploymentReplyWithDefaults() *GetDeploymentReply {
	this := GetDeploymentReply{}
	return &this
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *GetDeploymentReply) GetDeployment() Deployment {
	if o == nil || IsNil(o.Deployment) {
		var ret Deployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeploymentReply) GetDeploymentOk() (*Deployment, bool) {
	if o == nil || IsNil(o.Deployment) {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *GetDeploymentReply) HasDeployment() bool {
	if o != nil && !IsNil(o.Deployment) {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given Deployment and assigns it to the Deployment field.
func (o *GetDeploymentReply) SetDeployment(v Deployment) {
	o.Deployment = &v
}

func (o GetDeploymentReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeploymentReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deployment) {
		toSerialize["deployment"] = o.Deployment
	}
	return toSerialize, nil
}

type NullableGetDeploymentReply struct {
	value *GetDeploymentReply
	isSet bool
}

func (v NullableGetDeploymentReply) Get() *GetDeploymentReply {
	return v.value
}

func (v *NullableGetDeploymentReply) Set(val *GetDeploymentReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeploymentReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeploymentReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeploymentReply(val *GetDeploymentReply) *NullableGetDeploymentReply {
	return &NullableGetDeploymentReply{value: val, isSet: true}
}

func (v NullableGetDeploymentReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeploymentReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


