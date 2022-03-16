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

// RegionalDeploymentListItemStatus the model 'RegionalDeploymentListItemStatus'
type RegionalDeploymentListItemStatus string

// List of RegionalDeploymentListItem.Status
const (
	REGIONALDEPLOYMENTLISTITEMSTATUS_PENDING RegionalDeploymentListItemStatus = "PENDING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_PROVISIONING RegionalDeploymentListItemStatus = "PROVISIONING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_SCHEDULED RegionalDeploymentListItemStatus = "SCHEDULED"
	REGIONALDEPLOYMENTLISTITEMSTATUS_CANCELING RegionalDeploymentListItemStatus = "CANCELING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_CANCELED RegionalDeploymentListItemStatus = "CANCELED"
	REGIONALDEPLOYMENTLISTITEMSTATUS_ALLOCATING RegionalDeploymentListItemStatus = "ALLOCATING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_STARTING RegionalDeploymentListItemStatus = "STARTING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_HEALTHY RegionalDeploymentListItemStatus = "HEALTHY"
	REGIONALDEPLOYMENTLISTITEMSTATUS_DEGRADED RegionalDeploymentListItemStatus = "DEGRADED"
	REGIONALDEPLOYMENTLISTITEMSTATUS_UNHEALTHY RegionalDeploymentListItemStatus = "UNHEALTHY"
	REGIONALDEPLOYMENTLISTITEMSTATUS_STOPPING RegionalDeploymentListItemStatus = "STOPPING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_STOPPED RegionalDeploymentListItemStatus = "STOPPED"
	REGIONALDEPLOYMENTLISTITEMSTATUS_ERRORING RegionalDeploymentListItemStatus = "ERRORING"
	REGIONALDEPLOYMENTLISTITEMSTATUS_ERROR RegionalDeploymentListItemStatus = "ERROR"
)

var allowedRegionalDeploymentListItemStatusEnumValues = []RegionalDeploymentListItemStatus{
	"PENDING",
	"PROVISIONING",
	"SCHEDULED",
	"CANCELING",
	"CANCELED",
	"ALLOCATING",
	"STARTING",
	"HEALTHY",
	"DEGRADED",
	"UNHEALTHY",
	"STOPPING",
	"STOPPED",
	"ERRORING",
	"ERROR",
}

func (v *RegionalDeploymentListItemStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegionalDeploymentListItemStatus(value)
	for _, existing := range allowedRegionalDeploymentListItemStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegionalDeploymentListItemStatus", value)
}

// NewRegionalDeploymentListItemStatusFromValue returns a pointer to a valid RegionalDeploymentListItemStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegionalDeploymentListItemStatusFromValue(v string) (*RegionalDeploymentListItemStatus, error) {
	ev := RegionalDeploymentListItemStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegionalDeploymentListItemStatus: valid values are %v", v, allowedRegionalDeploymentListItemStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegionalDeploymentListItemStatus) IsValid() bool {
	for _, existing := range allowedRegionalDeploymentListItemStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegionalDeploymentListItem.Status value
func (v RegionalDeploymentListItemStatus) Ptr() *RegionalDeploymentListItemStatus {
	return &v
}

type NullableRegionalDeploymentListItemStatus struct {
	value *RegionalDeploymentListItemStatus
	isSet bool
}

func (v NullableRegionalDeploymentListItemStatus) Get() *RegionalDeploymentListItemStatus {
	return v.value
}

func (v *NullableRegionalDeploymentListItemStatus) Set(val *RegionalDeploymentListItemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentListItemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentListItemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentListItemStatus(val *RegionalDeploymentListItemStatus) *NullableRegionalDeploymentListItemStatus {
	return &NullableRegionalDeploymentListItemStatus{value: val, isSet: true}
}

func (v NullableRegionalDeploymentListItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentListItemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
