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
	"fmt"
)

// PaymentMethodStatus the model 'PaymentMethodStatus'
type PaymentMethodStatus string

// List of PaymentMethod.Status
const (
	PAYMENTMETHODSTATUS_INVALID PaymentMethodStatus = "INVALID"
	PAYMENTMETHODSTATUS_CREATED PaymentMethodStatus = "CREATED"
	PAYMENTMETHODSTATUS_AUTHORIZED PaymentMethodStatus = "AUTHORIZED"
	PAYMENTMETHODSTATUS_DECLINED PaymentMethodStatus = "DECLINED"
	PAYMENTMETHODSTATUS_CANCELED PaymentMethodStatus = "CANCELED"
	PAYMENTMETHODSTATUS_EXPIRED PaymentMethodStatus = "EXPIRED"
	PAYMENTMETHODSTATUS_UNCHECKED PaymentMethodStatus = "UNCHECKED"
)

var allowedPaymentMethodStatusEnumValues = []PaymentMethodStatus{
	"INVALID",
	"CREATED",
	"AUTHORIZED",
	"DECLINED",
	"CANCELED",
	"EXPIRED",
	"UNCHECKED",
}

func (v *PaymentMethodStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodStatus(value)
	for _, existing := range allowedPaymentMethodStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentMethodStatus", value)
}

// NewPaymentMethodStatusFromValue returns a pointer to a valid PaymentMethodStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodStatusFromValue(v string) (*PaymentMethodStatus, error) {
	ev := PaymentMethodStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodStatus: valid values are %v", v, allowedPaymentMethodStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodStatus) IsValid() bool {
	for _, existing := range allowedPaymentMethodStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentMethod.Status value
func (v PaymentMethodStatus) Ptr() *PaymentMethodStatus {
	return &v
}

type NullablePaymentMethodStatus struct {
	value *PaymentMethodStatus
	isSet bool
}

func (v NullablePaymentMethodStatus) Get() *PaymentMethodStatus {
	return v.value
}

func (v *NullablePaymentMethodStatus) Set(val *PaymentMethodStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodStatus(val *PaymentMethodStatus) *NullablePaymentMethodStatus {
	return &NullablePaymentMethodStatus{value: val, isSet: true}
}

func (v NullablePaymentMethodStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

