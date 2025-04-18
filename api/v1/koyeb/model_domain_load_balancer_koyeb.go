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

// DomainLoadBalancerKoyeb struct for DomainLoadBalancerKoyeb
type DomainLoadBalancerKoyeb struct {
	RequestTimeoutSeconds *int64 `json:"request_timeout_seconds,omitempty"`
}

// NewDomainLoadBalancerKoyeb instantiates a new DomainLoadBalancerKoyeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainLoadBalancerKoyeb() *DomainLoadBalancerKoyeb {
	this := DomainLoadBalancerKoyeb{}
	return &this
}

// NewDomainLoadBalancerKoyebWithDefaults instantiates a new DomainLoadBalancerKoyeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainLoadBalancerKoyebWithDefaults() *DomainLoadBalancerKoyeb {
	this := DomainLoadBalancerKoyeb{}
	return &this
}

// GetRequestTimeoutSeconds returns the RequestTimeoutSeconds field value if set, zero value otherwise.
func (o *DomainLoadBalancerKoyeb) GetRequestTimeoutSeconds() int64 {
	if o == nil || isNil(o.RequestTimeoutSeconds) {
		var ret int64
		return ret
	}
	return *o.RequestTimeoutSeconds
}

// GetRequestTimeoutSecondsOk returns a tuple with the RequestTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLoadBalancerKoyeb) GetRequestTimeoutSecondsOk() (*int64, bool) {
	if o == nil || isNil(o.RequestTimeoutSeconds) {
    return nil, false
	}
	return o.RequestTimeoutSeconds, true
}

// HasRequestTimeoutSeconds returns a boolean if a field has been set.
func (o *DomainLoadBalancerKoyeb) HasRequestTimeoutSeconds() bool {
	if o != nil && !isNil(o.RequestTimeoutSeconds) {
		return true
	}

	return false
}

// SetRequestTimeoutSeconds gets a reference to the given int64 and assigns it to the RequestTimeoutSeconds field.
func (o *DomainLoadBalancerKoyeb) SetRequestTimeoutSeconds(v int64) {
	o.RequestTimeoutSeconds = &v
}

func (o DomainLoadBalancerKoyeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequestTimeoutSeconds) {
		toSerialize["request_timeout_seconds"] = o.RequestTimeoutSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableDomainLoadBalancerKoyeb struct {
	value *DomainLoadBalancerKoyeb
	isSet bool
}

func (v NullableDomainLoadBalancerKoyeb) Get() *DomainLoadBalancerKoyeb {
	return v.value
}

func (v *NullableDomainLoadBalancerKoyeb) Set(val *DomainLoadBalancerKoyeb) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainLoadBalancerKoyeb) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainLoadBalancerKoyeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainLoadBalancerKoyeb(val *DomainLoadBalancerKoyeb) *NullableDomainLoadBalancerKoyeb {
	return &NullableDomainLoadBalancerKoyeb{value: val, isSet: true}
}

func (v NullableDomainLoadBalancerKoyeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainLoadBalancerKoyeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


