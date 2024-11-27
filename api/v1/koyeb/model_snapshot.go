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

// Snapshot The object that represents a snapshot. It can either be local, on a node, or remote, in a cold storage.
type Snapshot struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Size *int64 `json:"size,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	ParentVolumeId *string `json:"parent_volume_id,omitempty"`
	Region *string `json:"region,omitempty"`
	Status *SnapshotStatus `json:"status,omitempty"`
	Type *SnapshotType `json:"type,omitempty"`
}

// NewSnapshot instantiates a new Snapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshot() *Snapshot {
	this := Snapshot{}
	var status SnapshotStatus = SNAPSHOTSTATUS_INVALID
	this.Status = &status
	var type_ SnapshotType = SNAPSHOTTYPE_INVALID
	this.Type = &type_
	return &this
}

// NewSnapshotWithDefaults instantiates a new Snapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotWithDefaults() *Snapshot {
	this := Snapshot{}
	var status SnapshotStatus = SNAPSHOTSTATUS_INVALID
	this.Status = &status
	var type_ SnapshotType = SNAPSHOTTYPE_INVALID
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Snapshot) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Snapshot) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Snapshot) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Snapshot) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Snapshot) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Snapshot) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Snapshot) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Snapshot) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Snapshot) SetSize(v int64) {
	o.Size = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Snapshot) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Snapshot) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Snapshot) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Snapshot) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Snapshot) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Snapshot) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Snapshot) GetDeletedAt() time.Time {
	if o == nil || isNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletedAt) {
    return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Snapshot) HasDeletedAt() bool {
	if o != nil && !isNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Snapshot) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Snapshot) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Snapshot) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Snapshot) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetParentVolumeId returns the ParentVolumeId field value if set, zero value otherwise.
func (o *Snapshot) GetParentVolumeId() string {
	if o == nil || isNil(o.ParentVolumeId) {
		var ret string
		return ret
	}
	return *o.ParentVolumeId
}

// GetParentVolumeIdOk returns a tuple with the ParentVolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetParentVolumeIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentVolumeId) {
    return nil, false
	}
	return o.ParentVolumeId, true
}

// HasParentVolumeId returns a boolean if a field has been set.
func (o *Snapshot) HasParentVolumeId() bool {
	if o != nil && !isNil(o.ParentVolumeId) {
		return true
	}

	return false
}

// SetParentVolumeId gets a reference to the given string and assigns it to the ParentVolumeId field.
func (o *Snapshot) SetParentVolumeId(v string) {
	o.ParentVolumeId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Snapshot) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Snapshot) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Snapshot) SetRegion(v string) {
	o.Region = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Snapshot) GetStatus() SnapshotStatus {
	if o == nil || isNil(o.Status) {
		var ret SnapshotStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetStatusOk() (*SnapshotStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Snapshot) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SnapshotStatus and assigns it to the Status field.
func (o *Snapshot) SetStatus(v SnapshotStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Snapshot) GetType() SnapshotType {
	if o == nil || isNil(o.Type) {
		var ret SnapshotType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetTypeOk() (*SnapshotType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Snapshot) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SnapshotType and assigns it to the Type field.
func (o *Snapshot) SetType(v SnapshotType) {
	o.Type = &v
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
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
	if !isNil(o.ParentVolumeId) {
		toSerialize["parent_volume_id"] = o.ParentVolumeId
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshot struct {
	value *Snapshot
	isSet bool
}

func (v NullableSnapshot) Get() *Snapshot {
	return v.value
}

func (v *NullableSnapshot) Set(val *Snapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshot(val *Snapshot) *NullableSnapshot {
	return &NullableSnapshot{value: val, isSet: true}
}

func (v NullableSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

