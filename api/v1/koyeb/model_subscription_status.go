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

// SubscriptionStatus the model 'SubscriptionStatus'
type SubscriptionStatus string

// List of Subscription.Status
const (
	SUBSCRIPTIONSTATUS_INVALID SubscriptionStatus = "INVALID"
	SUBSCRIPTIONSTATUS_CREATED SubscriptionStatus = "CREATED"
	SUBSCRIPTIONSTATUS_ACTIVE SubscriptionStatus = "ACTIVE"
	SUBSCRIPTIONSTATUS_WARNING SubscriptionStatus = "WARNING"
	SUBSCRIPTIONSTATUS_URGENT SubscriptionStatus = "URGENT"
	SUBSCRIPTIONSTATUS_CANCELING SubscriptionStatus = "CANCELING"
	SUBSCRIPTIONSTATUS_CANCELED SubscriptionStatus = "CANCELED"
	SUBSCRIPTIONSTATUS_TERMINATING SubscriptionStatus = "TERMINATING"
	SUBSCRIPTIONSTATUS_TERMINATED SubscriptionStatus = "TERMINATED"
)

var allowedSubscriptionStatusEnumValues = []SubscriptionStatus{
	"INVALID",
	"CREATED",
	"ACTIVE",
	"WARNING",
	"URGENT",
	"CANCELING",
	"CANCELED",
	"TERMINATING",
	"TERMINATED",
}

func (v *SubscriptionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionStatus(value)
	for _, existing := range allowedSubscriptionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionStatus", value)
}

// NewSubscriptionStatusFromValue returns a pointer to a valid SubscriptionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionStatusFromValue(v string) (*SubscriptionStatus, error) {
	ev := SubscriptionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionStatus: valid values are %v", v, allowedSubscriptionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionStatus) IsValid() bool {
	for _, existing := range allowedSubscriptionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Subscription.Status value
func (v SubscriptionStatus) Ptr() *SubscriptionStatus {
	return &v
}

type NullableSubscriptionStatus struct {
	value *SubscriptionStatus
	isSet bool
}

func (v NullableSubscriptionStatus) Get() *SubscriptionStatus {
	return v.value
}

func (v *NullableSubscriptionStatus) Set(val *SubscriptionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionStatus(val *SubscriptionStatus) *NullableSubscriptionStatus {
	return &NullableSubscriptionStatus{value: val, isSet: true}
}

func (v NullableSubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

