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

// CreatePaymentAuthorizationReply struct for CreatePaymentAuthorizationReply
type CreatePaymentAuthorizationReply struct {
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

// NewCreatePaymentAuthorizationReply instantiates a new CreatePaymentAuthorizationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePaymentAuthorizationReply() *CreatePaymentAuthorizationReply {
	this := CreatePaymentAuthorizationReply{}
	return &this
}

// NewCreatePaymentAuthorizationReplyWithDefaults instantiates a new CreatePaymentAuthorizationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePaymentAuthorizationReplyWithDefaults() *CreatePaymentAuthorizationReply {
	this := CreatePaymentAuthorizationReply{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CreatePaymentAuthorizationReply) GetPaymentMethod() PaymentMethod {
	if o == nil || o.PaymentMethod == nil {
		var ret PaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePaymentAuthorizationReply) GetPaymentMethodOk() (*PaymentMethod, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CreatePaymentAuthorizationReply) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethod and assigns it to the PaymentMethod field.
func (o *CreatePaymentAuthorizationReply) SetPaymentMethod(v PaymentMethod) {
	o.PaymentMethod = &v
}

func (o CreatePaymentAuthorizationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePaymentAuthorizationReply struct {
	value *CreatePaymentAuthorizationReply
	isSet bool
}

func (v NullableCreatePaymentAuthorizationReply) Get() *CreatePaymentAuthorizationReply {
	return v.value
}

func (v *NullableCreatePaymentAuthorizationReply) Set(val *CreatePaymentAuthorizationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePaymentAuthorizationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePaymentAuthorizationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePaymentAuthorizationReply(val *CreatePaymentAuthorizationReply) *NullableCreatePaymentAuthorizationReply {
	return &NullableCreatePaymentAuthorizationReply{value: val, isSet: true}
}

func (v NullableCreatePaymentAuthorizationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePaymentAuthorizationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


