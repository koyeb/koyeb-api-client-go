/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// PersistentVolume struct for PersistentVolume
type PersistentVolume struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SnapshotId *string `json:"snapshot_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	Region *string `json:"region,omitempty"`
	ReadOnly *bool `json:"read_only,omitempty"`
	MaxSize *int64 `json:"max_size,omitempty"`
	CurSize *int64 `json:"cur_size,omitempty"`
	Status *PersistentVolumeStatus `json:"status,omitempty"`
	BackingStore *PersistentVolumeBackingStore `json:"backing_store,omitempty"`
}

// NewPersistentVolume instantiates a new PersistentVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersistentVolume() *PersistentVolume {
	this := PersistentVolume{}
	var status PersistentVolumeStatus = PERSISTENTVOLUMESTATUS_INVALID
	this.Status = &status
	var backingStore PersistentVolumeBackingStore = PERSISTENTVOLUMEBACKINGSTORE_INVALID
	this.BackingStore = &backingStore
	return &this
}

// NewPersistentVolumeWithDefaults instantiates a new PersistentVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersistentVolumeWithDefaults() *PersistentVolume {
	this := PersistentVolume{}
	var status PersistentVolumeStatus = PERSISTENTVOLUMESTATUS_INVALID
	this.Status = &status
	var backingStore PersistentVolumeBackingStore = PERSISTENTVOLUMEBACKINGSTORE_INVALID
	this.BackingStore = &backingStore
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersistentVolume) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersistentVolume) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PersistentVolume) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersistentVolume) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersistentVolume) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersistentVolume) SetName(v string) {
	o.Name = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *PersistentVolume) GetSnapshotId() string {
	if o == nil || isNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetSnapshotIdOk() (*string, bool) {
	if o == nil || isNil(o.SnapshotId) {
    return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *PersistentVolume) HasSnapshotId() bool {
	if o != nil && !isNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *PersistentVolume) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PersistentVolume) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PersistentVolume) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PersistentVolume) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PersistentVolume) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PersistentVolume) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PersistentVolume) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *PersistentVolume) GetDeletedAt() time.Time {
	if o == nil || isNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletedAt) {
    return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PersistentVolume) HasDeletedAt() bool {
	if o != nil && !isNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *PersistentVolume) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *PersistentVolume) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *PersistentVolume) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *PersistentVolume) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *PersistentVolume) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *PersistentVolume) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *PersistentVolume) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PersistentVolume) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PersistentVolume) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PersistentVolume) SetRegion(v string) {
	o.Region = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *PersistentVolume) GetReadOnly() bool {
	if o == nil || isNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetReadOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.ReadOnly) {
    return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *PersistentVolume) HasReadOnly() bool {
	if o != nil && !isNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *PersistentVolume) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *PersistentVolume) GetMaxSize() int64 {
	if o == nil || isNil(o.MaxSize) {
		var ret int64
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetMaxSizeOk() (*int64, bool) {
	if o == nil || isNil(o.MaxSize) {
    return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *PersistentVolume) HasMaxSize() bool {
	if o != nil && !isNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int64 and assigns it to the MaxSize field.
func (o *PersistentVolume) SetMaxSize(v int64) {
	o.MaxSize = &v
}

// GetCurSize returns the CurSize field value if set, zero value otherwise.
func (o *PersistentVolume) GetCurSize() int64 {
	if o == nil || isNil(o.CurSize) {
		var ret int64
		return ret
	}
	return *o.CurSize
}

// GetCurSizeOk returns a tuple with the CurSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetCurSizeOk() (*int64, bool) {
	if o == nil || isNil(o.CurSize) {
    return nil, false
	}
	return o.CurSize, true
}

// HasCurSize returns a boolean if a field has been set.
func (o *PersistentVolume) HasCurSize() bool {
	if o != nil && !isNil(o.CurSize) {
		return true
	}

	return false
}

// SetCurSize gets a reference to the given int64 and assigns it to the CurSize field.
func (o *PersistentVolume) SetCurSize(v int64) {
	o.CurSize = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PersistentVolume) GetStatus() PersistentVolumeStatus {
	if o == nil || isNil(o.Status) {
		var ret PersistentVolumeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetStatusOk() (*PersistentVolumeStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PersistentVolume) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PersistentVolumeStatus and assigns it to the Status field.
func (o *PersistentVolume) SetStatus(v PersistentVolumeStatus) {
	o.Status = &v
}

// GetBackingStore returns the BackingStore field value if set, zero value otherwise.
func (o *PersistentVolume) GetBackingStore() PersistentVolumeBackingStore {
	if o == nil || isNil(o.BackingStore) {
		var ret PersistentVolumeBackingStore
		return ret
	}
	return *o.BackingStore
}

// GetBackingStoreOk returns a tuple with the BackingStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolume) GetBackingStoreOk() (*PersistentVolumeBackingStore, bool) {
	if o == nil || isNil(o.BackingStore) {
    return nil, false
	}
	return o.BackingStore, true
}

// HasBackingStore returns a boolean if a field has been set.
func (o *PersistentVolume) HasBackingStore() bool {
	if o != nil && !isNil(o.BackingStore) {
		return true
	}

	return false
}

// SetBackingStore gets a reference to the given PersistentVolumeBackingStore and assigns it to the BackingStore field.
func (o *PersistentVolume) SetBackingStore(v PersistentVolumeBackingStore) {
	o.BackingStore = &v
}

func (o PersistentVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.ReadOnly) {
		toSerialize["read_only"] = o.ReadOnly
	}
	if !isNil(o.MaxSize) {
		toSerialize["max_size"] = o.MaxSize
	}
	if !isNil(o.CurSize) {
		toSerialize["cur_size"] = o.CurSize
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.BackingStore) {
		toSerialize["backing_store"] = o.BackingStore
	}
	return json.Marshal(toSerialize)
}

type NullablePersistentVolume struct {
	value *PersistentVolume
	isSet bool
}

func (v NullablePersistentVolume) Get() *PersistentVolume {
	return v.value
}

func (v *NullablePersistentVolume) Set(val *PersistentVolume) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistentVolume) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistentVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistentVolume(val *PersistentVolume) *NullablePersistentVolume {
	return &NullablePersistentVolume{value: val, isSet: true}
}

func (v NullablePersistentVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistentVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


