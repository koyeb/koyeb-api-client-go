# PersistentVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SnapshotId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**MaxSize** | Pointer to **int64** |  | [optional] 
**CurSize** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**PersistentVolumeStatus**](PersistentVolumeStatus.md) |  | [optional] [default to PERSISTENTVOLUMESTATUS_INVALID]
**BackingStore** | Pointer to [**PersistentVolumeBackingStore**](PersistentVolumeBackingStore.md) |  | [optional] [default to PERSISTENTVOLUMEBACKINGSTORE_INVALID]

## Methods

### NewPersistentVolume

`func NewPersistentVolume() *PersistentVolume`

NewPersistentVolume instantiates a new PersistentVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentVolumeWithDefaults

`func NewPersistentVolumeWithDefaults() *PersistentVolume`

NewPersistentVolumeWithDefaults instantiates a new PersistentVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersistentVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersistentVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersistentVolume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersistentVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PersistentVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersistentVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersistentVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersistentVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSnapshotId

`func (o *PersistentVolume) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *PersistentVolume) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *PersistentVolume) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *PersistentVolume) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PersistentVolume) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersistentVolume) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersistentVolume) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PersistentVolume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PersistentVolume) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PersistentVolume) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PersistentVolume) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PersistentVolume) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PersistentVolume) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PersistentVolume) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PersistentVolume) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PersistentVolume) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PersistentVolume) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PersistentVolume) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PersistentVolume) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PersistentVolume) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetServiceId

`func (o *PersistentVolume) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PersistentVolume) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PersistentVolume) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PersistentVolume) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRegion

`func (o *PersistentVolume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PersistentVolume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PersistentVolume) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PersistentVolume) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReadOnly

`func (o *PersistentVolume) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PersistentVolume) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PersistentVolume) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PersistentVolume) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetMaxSize

`func (o *PersistentVolume) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *PersistentVolume) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *PersistentVolume) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *PersistentVolume) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetCurSize

`func (o *PersistentVolume) GetCurSize() int64`

GetCurSize returns the CurSize field if non-nil, zero value otherwise.

### GetCurSizeOk

`func (o *PersistentVolume) GetCurSizeOk() (*int64, bool)`

GetCurSizeOk returns a tuple with the CurSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurSize

`func (o *PersistentVolume) SetCurSize(v int64)`

SetCurSize sets CurSize field to given value.

### HasCurSize

`func (o *PersistentVolume) HasCurSize() bool`

HasCurSize returns a boolean if a field has been set.

### GetStatus

`func (o *PersistentVolume) GetStatus() PersistentVolumeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PersistentVolume) GetStatusOk() (*PersistentVolumeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PersistentVolume) SetStatus(v PersistentVolumeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PersistentVolume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBackingStore

`func (o *PersistentVolume) GetBackingStore() PersistentVolumeBackingStore`

GetBackingStore returns the BackingStore field if non-nil, zero value otherwise.

### GetBackingStoreOk

`func (o *PersistentVolume) GetBackingStoreOk() (*PersistentVolumeBackingStore, bool)`

GetBackingStoreOk returns a tuple with the BackingStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingStore

`func (o *PersistentVolume) SetBackingStore(v PersistentVolumeBackingStore)`

SetBackingStore sets BackingStore field to given value.

### HasBackingStore

`func (o *PersistentVolume) HasBackingStore() bool`

HasBackingStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


