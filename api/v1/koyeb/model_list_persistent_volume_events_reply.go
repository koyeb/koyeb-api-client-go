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

// ListPersistentVolumeEventsReply struct for ListPersistentVolumeEventsReply
type ListPersistentVolumeEventsReply struct {
	Events []PersistentVolumeEvent `json:"events,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Order *string `json:"order,omitempty"`
	HasNext *bool `json:"has_next,omitempty"`
}

// NewListPersistentVolumeEventsReply instantiates a new ListPersistentVolumeEventsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPersistentVolumeEventsReply() *ListPersistentVolumeEventsReply {
	this := ListPersistentVolumeEventsReply{}
	return &this
}

// NewListPersistentVolumeEventsReplyWithDefaults instantiates a new ListPersistentVolumeEventsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPersistentVolumeEventsReplyWithDefaults() *ListPersistentVolumeEventsReply {
	this := ListPersistentVolumeEventsReply{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ListPersistentVolumeEventsReply) GetEvents() []PersistentVolumeEvent {
	if o == nil || isNil(o.Events) {
		var ret []PersistentVolumeEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersistentVolumeEventsReply) GetEventsOk() ([]PersistentVolumeEvent, bool) {
	if o == nil || isNil(o.Events) {
    return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ListPersistentVolumeEventsReply) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []PersistentVolumeEvent and assigns it to the Events field.
func (o *ListPersistentVolumeEventsReply) SetEvents(v []PersistentVolumeEvent) {
	o.Events = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListPersistentVolumeEventsReply) GetLimit() int64 {
	if o == nil || isNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersistentVolumeEventsReply) GetLimitOk() (*int64, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListPersistentVolumeEventsReply) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListPersistentVolumeEventsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListPersistentVolumeEventsReply) GetOffset() int64 {
	if o == nil || isNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersistentVolumeEventsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || isNil(o.Offset) {
    return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListPersistentVolumeEventsReply) HasOffset() bool {
	if o != nil && !isNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListPersistentVolumeEventsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ListPersistentVolumeEventsReply) GetOrder() string {
	if o == nil || isNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersistentVolumeEventsReply) GetOrderOk() (*string, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ListPersistentVolumeEventsReply) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ListPersistentVolumeEventsReply) SetOrder(v string) {
	o.Order = &v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *ListPersistentVolumeEventsReply) GetHasNext() bool {
	if o == nil || isNil(o.HasNext) {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersistentVolumeEventsReply) GetHasNextOk() (*bool, bool) {
	if o == nil || isNil(o.HasNext) {
    return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *ListPersistentVolumeEventsReply) HasHasNext() bool {
	if o != nil && !isNil(o.HasNext) {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *ListPersistentVolumeEventsReply) SetHasNext(v bool) {
	o.HasNext = &v
}

func (o ListPersistentVolumeEventsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.HasNext) {
		toSerialize["has_next"] = o.HasNext
	}
	return json.Marshal(toSerialize)
}

type NullableListPersistentVolumeEventsReply struct {
	value *ListPersistentVolumeEventsReply
	isSet bool
}

func (v NullableListPersistentVolumeEventsReply) Get() *ListPersistentVolumeEventsReply {
	return v.value
}

func (v *NullableListPersistentVolumeEventsReply) Set(val *ListPersistentVolumeEventsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListPersistentVolumeEventsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListPersistentVolumeEventsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPersistentVolumeEventsReply(val *ListPersistentVolumeEventsReply) *NullableListPersistentVolumeEventsReply {
	return &NullableListPersistentVolumeEventsReply{value: val, isSet: true}
}

func (v NullableListPersistentVolumeEventsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPersistentVolumeEventsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

