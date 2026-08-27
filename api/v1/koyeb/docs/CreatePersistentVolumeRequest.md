# CreatePersistentVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeType** | Pointer to [**PersistentVolumeBackingStore**](PersistentVolumeBackingStore.md) |  | [optional] [default to PERSISTENTVOLUMEBACKINGSTORE_INVALID]
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**MaxSize** | Pointer to **int64** |  | [optional] 
**SnapshotId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreatePersistentVolumeRequest

`func NewCreatePersistentVolumeRequest() *CreatePersistentVolumeRequest`

NewCreatePersistentVolumeRequest instantiates a new CreatePersistentVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePersistentVolumeRequestWithDefaults

`func NewCreatePersistentVolumeRequestWithDefaults() *CreatePersistentVolumeRequest`

NewCreatePersistentVolumeRequestWithDefaults instantiates a new CreatePersistentVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeType

`func (o *CreatePersistentVolumeRequest) GetVolumeType() PersistentVolumeBackingStore`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreatePersistentVolumeRequest) GetVolumeTypeOk() (*PersistentVolumeBackingStore, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CreatePersistentVolumeRequest) SetVolumeType(v PersistentVolumeBackingStore)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CreatePersistentVolumeRequest) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetName

`func (o *CreatePersistentVolumeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePersistentVolumeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePersistentVolumeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePersistentVolumeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CreatePersistentVolumeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreatePersistentVolumeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreatePersistentVolumeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreatePersistentVolumeRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReadOnly

`func (o *CreatePersistentVolumeRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *CreatePersistentVolumeRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *CreatePersistentVolumeRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *CreatePersistentVolumeRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetMaxSize

`func (o *CreatePersistentVolumeRequest) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *CreatePersistentVolumeRequest) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *CreatePersistentVolumeRequest) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *CreatePersistentVolumeRequest) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetSnapshotId

`func (o *CreatePersistentVolumeRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CreatePersistentVolumeRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CreatePersistentVolumeRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *CreatePersistentVolumeRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


