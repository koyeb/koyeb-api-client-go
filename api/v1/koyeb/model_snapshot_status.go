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

// SnapshotStatus the model 'SnapshotStatus'
type SnapshotStatus string

// List of SnapshotStatus
const (
	SNAPSHOTSTATUS_INVALID SnapshotStatus = "SNAPSHOT_STATUS_INVALID"
	SNAPSHOTSTATUS_CREATING SnapshotStatus = "SNAPSHOT_STATUS_CREATING"
	SNAPSHOTSTATUS_AVAILABLE SnapshotStatus = "SNAPSHOT_STATUS_AVAILABLE"
	SNAPSHOTSTATUS_MIGRATING SnapshotStatus = "SNAPSHOT_STATUS_MIGRATING"
	SNAPSHOTSTATUS_DELETING SnapshotStatus = "SNAPSHOT_STATUS_DELETING"
	SNAPSHOTSTATUS_DELETED SnapshotStatus = "SNAPSHOT_STATUS_DELETED"
)

// All allowed values of SnapshotStatus enum
var AllowedSnapshotStatusEnumValues = []SnapshotStatus{
	"SNAPSHOT_STATUS_INVALID",
	"SNAPSHOT_STATUS_CREATING",
	"SNAPSHOT_STATUS_AVAILABLE",
	"SNAPSHOT_STATUS_MIGRATING",
	"SNAPSHOT_STATUS_DELETING",
	"SNAPSHOT_STATUS_DELETED",
}

func (v *SnapshotStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SnapshotStatus(value)
	for _, existing := range AllowedSnapshotStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SnapshotStatus", value)
}

// NewSnapshotStatusFromValue returns a pointer to a valid SnapshotStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnapshotStatusFromValue(v string) (*SnapshotStatus, error) {
	ev := SnapshotStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SnapshotStatus: valid values are %v", v, AllowedSnapshotStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnapshotStatus) IsValid() bool {
	for _, existing := range AllowedSnapshotStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SnapshotStatus value
func (v SnapshotStatus) Ptr() *SnapshotStatus {
	return &v
}

type NullableSnapshotStatus struct {
	value *SnapshotStatus
	isSet bool
}

func (v NullableSnapshotStatus) Get() *SnapshotStatus {
	return v.value
}

func (v *NullableSnapshotStatus) Set(val *SnapshotStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotStatus(val *SnapshotStatus) *NullableSnapshotStatus {
	return &NullableSnapshotStatus{value: val, isSet: true}
}

func (v NullableSnapshotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

