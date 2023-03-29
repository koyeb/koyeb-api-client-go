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

// checks if the ListDeploymentsReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeploymentsReply{}

// ListDeploymentsReply struct for ListDeploymentsReply
type ListDeploymentsReply struct {
	Deployments []DeploymentListItem `json:"deployments,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListDeploymentsReply instantiates a new ListDeploymentsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeploymentsReply() *ListDeploymentsReply {
	this := ListDeploymentsReply{}
	return &this
}

// NewListDeploymentsReplyWithDefaults instantiates a new ListDeploymentsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeploymentsReplyWithDefaults() *ListDeploymentsReply {
	this := ListDeploymentsReply{}
	return &this
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *ListDeploymentsReply) GetDeployments() []DeploymentListItem {
	if o == nil || IsNil(o.Deployments) {
		var ret []DeploymentListItem
		return ret
	}
	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeploymentsReply) GetDeploymentsOk() ([]DeploymentListItem, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *ListDeploymentsReply) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []DeploymentListItem and assigns it to the Deployments field.
func (o *ListDeploymentsReply) SetDeployments(v []DeploymentListItem) {
	o.Deployments = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListDeploymentsReply) GetLimit() int64 {
	if o == nil || IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeploymentsReply) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListDeploymentsReply) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListDeploymentsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListDeploymentsReply) GetOffset() int64 {
	if o == nil || IsNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeploymentsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListDeploymentsReply) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListDeploymentsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListDeploymentsReply) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeploymentsReply) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListDeploymentsReply) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListDeploymentsReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListDeploymentsReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDeploymentsReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableListDeploymentsReply struct {
	value *ListDeploymentsReply
	isSet bool
}

func (v NullableListDeploymentsReply) Get() *ListDeploymentsReply {
	return v.value
}

func (v *NullableListDeploymentsReply) Set(val *ListDeploymentsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeploymentsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeploymentsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeploymentsReply(val *ListDeploymentsReply) *NullableListDeploymentsReply {
	return &NullableListDeploymentsReply{value: val, isSet: true}
}

func (v NullableListDeploymentsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeploymentsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


