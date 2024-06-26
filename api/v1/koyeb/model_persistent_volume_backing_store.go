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

// PersistentVolumeBackingStore the model 'PersistentVolumeBackingStore'
type PersistentVolumeBackingStore string

// List of PersistentVolumeBackingStore
const (
	PERSISTENTVOLUMEBACKINGSTORE_INVALID PersistentVolumeBackingStore = "PERSISTENT_VOLUME_BACKING_STORE_INVALID"
	PERSISTENTVOLUMEBACKINGSTORE_LOCAL_BLK PersistentVolumeBackingStore = "PERSISTENT_VOLUME_BACKING_STORE_LOCAL_BLK"
)

// All allowed values of PersistentVolumeBackingStore enum
var AllowedPersistentVolumeBackingStoreEnumValues = []PersistentVolumeBackingStore{
	"PERSISTENT_VOLUME_BACKING_STORE_INVALID",
	"PERSISTENT_VOLUME_BACKING_STORE_LOCAL_BLK",
}

func (v *PersistentVolumeBackingStore) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersistentVolumeBackingStore(value)
	for _, existing := range AllowedPersistentVolumeBackingStoreEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersistentVolumeBackingStore", value)
}

// NewPersistentVolumeBackingStoreFromValue returns a pointer to a valid PersistentVolumeBackingStore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersistentVolumeBackingStoreFromValue(v string) (*PersistentVolumeBackingStore, error) {
	ev := PersistentVolumeBackingStore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersistentVolumeBackingStore: valid values are %v", v, AllowedPersistentVolumeBackingStoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersistentVolumeBackingStore) IsValid() bool {
	for _, existing := range AllowedPersistentVolumeBackingStoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersistentVolumeBackingStore value
func (v PersistentVolumeBackingStore) Ptr() *PersistentVolumeBackingStore {
	return &v
}

type NullablePersistentVolumeBackingStore struct {
	value *PersistentVolumeBackingStore
	isSet bool
}

func (v NullablePersistentVolumeBackingStore) Get() *PersistentVolumeBackingStore {
	return v.value
}

func (v *NullablePersistentVolumeBackingStore) Set(val *PersistentVolumeBackingStore) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistentVolumeBackingStore) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistentVolumeBackingStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistentVolumeBackingStore(val *PersistentVolumeBackingStore) *NullablePersistentVolumeBackingStore {
	return &NullablePersistentVolumeBackingStore{value: val, isSet: true}
}

func (v NullablePersistentVolumeBackingStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistentVolumeBackingStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

