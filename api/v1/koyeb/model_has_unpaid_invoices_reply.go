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

// HasUnpaidInvoicesReply struct for HasUnpaidInvoicesReply
type HasUnpaidInvoicesReply struct {
	HasUnpaidInvoices *bool `json:"has_unpaid_invoices,omitempty"`
}

// NewHasUnpaidInvoicesReply instantiates a new HasUnpaidInvoicesReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasUnpaidInvoicesReply() *HasUnpaidInvoicesReply {
	this := HasUnpaidInvoicesReply{}
	return &this
}

// NewHasUnpaidInvoicesReplyWithDefaults instantiates a new HasUnpaidInvoicesReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasUnpaidInvoicesReplyWithDefaults() *HasUnpaidInvoicesReply {
	this := HasUnpaidInvoicesReply{}
	return &this
}

// GetHasUnpaidInvoices returns the HasUnpaidInvoices field value if set, zero value otherwise.
func (o *HasUnpaidInvoicesReply) GetHasUnpaidInvoices() bool {
	if o == nil || isNil(o.HasUnpaidInvoices) {
		var ret bool
		return ret
	}
	return *o.HasUnpaidInvoices
}

// GetHasUnpaidInvoicesOk returns a tuple with the HasUnpaidInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasUnpaidInvoicesReply) GetHasUnpaidInvoicesOk() (*bool, bool) {
	if o == nil || isNil(o.HasUnpaidInvoices) {
    return nil, false
	}
	return o.HasUnpaidInvoices, true
}

// HasHasUnpaidInvoices returns a boolean if a field has been set.
func (o *HasUnpaidInvoicesReply) HasHasUnpaidInvoices() bool {
	if o != nil && !isNil(o.HasUnpaidInvoices) {
		return true
	}

	return false
}

// SetHasUnpaidInvoices gets a reference to the given bool and assigns it to the HasUnpaidInvoices field.
func (o *HasUnpaidInvoicesReply) SetHasUnpaidInvoices(v bool) {
	o.HasUnpaidInvoices = &v
}

func (o HasUnpaidInvoicesReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasUnpaidInvoices) {
		toSerialize["has_unpaid_invoices"] = o.HasUnpaidInvoices
	}
	return json.Marshal(toSerialize)
}

type NullableHasUnpaidInvoicesReply struct {
	value *HasUnpaidInvoicesReply
	isSet bool
}

func (v NullableHasUnpaidInvoicesReply) Get() *HasUnpaidInvoicesReply {
	return v.value
}

func (v *NullableHasUnpaidInvoicesReply) Set(val *HasUnpaidInvoicesReply) {
	v.value = val
	v.isSet = true
}

func (v NullableHasUnpaidInvoicesReply) IsSet() bool {
	return v.isSet
}

func (v *NullableHasUnpaidInvoicesReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasUnpaidInvoicesReply(val *HasUnpaidInvoicesReply) *NullableHasUnpaidInvoicesReply {
	return &NullableHasUnpaidInvoicesReply{value: val, isSet: true}
}

func (v NullableHasUnpaidInvoicesReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasUnpaidInvoicesReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


