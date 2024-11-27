/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// DeploymentStatus the model 'DeploymentStatus'
type DeploymentStatus string

// List of Deployment.Status
const (
	DEPLOYMENTSTATUS_PENDING DeploymentStatus = "PENDING"
	DEPLOYMENTSTATUS_PROVISIONING DeploymentStatus = "PROVISIONING"
	DEPLOYMENTSTATUS_SCHEDULED DeploymentStatus = "SCHEDULED"
	DEPLOYMENTSTATUS_CANCELING DeploymentStatus = "CANCELING"
	DEPLOYMENTSTATUS_CANCELED DeploymentStatus = "CANCELED"
	DEPLOYMENTSTATUS_ALLOCATING DeploymentStatus = "ALLOCATING"
	DEPLOYMENTSTATUS_STARTING DeploymentStatus = "STARTING"
	DEPLOYMENTSTATUS_HEALTHY DeploymentStatus = "HEALTHY"
	DEPLOYMENTSTATUS_DEGRADED DeploymentStatus = "DEGRADED"
	DEPLOYMENTSTATUS_UNHEALTHY DeploymentStatus = "UNHEALTHY"
	DEPLOYMENTSTATUS_STOPPING DeploymentStatus = "STOPPING"
	DEPLOYMENTSTATUS_STOPPED DeploymentStatus = "STOPPED"
	DEPLOYMENTSTATUS_ERRORING DeploymentStatus = "ERRORING"
	DEPLOYMENTSTATUS_ERROR DeploymentStatus = "ERROR"
	DEPLOYMENTSTATUS_STASHED DeploymentStatus = "STASHED"
	DEPLOYMENTSTATUS_SLEEPING DeploymentStatus = "SLEEPING"
)

// All allowed values of DeploymentStatus enum
var AllowedDeploymentStatusEnumValues = []DeploymentStatus{
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
	"STASHED",
	"SLEEPING",
}

func (v *DeploymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploymentStatus(value)
	for _, existing := range AllowedDeploymentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploymentStatus", value)
}

// NewDeploymentStatusFromValue returns a pointer to a valid DeploymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploymentStatusFromValue(v string) (*DeploymentStatus, error) {
	ev := DeploymentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploymentStatus: valid values are %v", v, AllowedDeploymentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploymentStatus) IsValid() bool {
	for _, existing := range AllowedDeploymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Deployment.Status value
func (v DeploymentStatus) Ptr() *DeploymentStatus {
	return &v
}

type NullableDeploymentStatus struct {
	value *DeploymentStatus
	isSet bool
}

func (v NullableDeploymentStatus) Get() *DeploymentStatus {
	return v.value
}

func (v *NullableDeploymentStatus) Set(val *DeploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStatus(val *DeploymentStatus) *NullableDeploymentStatus {
	return &NullableDeploymentStatus{value: val, isSet: true}
}

func (v NullableDeploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

