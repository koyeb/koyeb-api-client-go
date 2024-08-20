# PersistentVolumeQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTotalSize** | Pointer to **int64** | MaxTotalSize for all volumes on a region (in Gigabyte / GB). | [optional] 
**MaxVolumeSize** | Pointer to **int64** | MaxVolumeSize for one volume (in Gigabyte / GB). | [optional] 
**MaxPerInstanceSize** | Pointer to **int64** | MaxPerInstanceSize for all volumes on an instance (in Gigabyte / GB). | [optional] 

## Methods

### NewPersistentVolumeQuotas

`func NewPersistentVolumeQuotas() *PersistentVolumeQuotas`

NewPersistentVolumeQuotas instantiates a new PersistentVolumeQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentVolumeQuotasWithDefaults

`func NewPersistentVolumeQuotasWithDefaults() *PersistentVolumeQuotas`

NewPersistentVolumeQuotasWithDefaults instantiates a new PersistentVolumeQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTotalSize

`func (o *PersistentVolumeQuotas) GetMaxTotalSize() int64`

GetMaxTotalSize returns the MaxTotalSize field if non-nil, zero value otherwise.

### GetMaxTotalSizeOk

`func (o *PersistentVolumeQuotas) GetMaxTotalSizeOk() (*int64, bool)`

GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSize

`func (o *PersistentVolumeQuotas) SetMaxTotalSize(v int64)`

SetMaxTotalSize sets MaxTotalSize field to given value.

### HasMaxTotalSize

`func (o *PersistentVolumeQuotas) HasMaxTotalSize() bool`

HasMaxTotalSize returns a boolean if a field has been set.

### GetMaxVolumeSize

`func (o *PersistentVolumeQuotas) GetMaxVolumeSize() int64`

GetMaxVolumeSize returns the MaxVolumeSize field if non-nil, zero value otherwise.

### GetMaxVolumeSizeOk

`func (o *PersistentVolumeQuotas) GetMaxVolumeSizeOk() (*int64, bool)`

GetMaxVolumeSizeOk returns a tuple with the MaxVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumeSize

`func (o *PersistentVolumeQuotas) SetMaxVolumeSize(v int64)`

SetMaxVolumeSize sets MaxVolumeSize field to given value.

### HasMaxVolumeSize

`func (o *PersistentVolumeQuotas) HasMaxVolumeSize() bool`

HasMaxVolumeSize returns a boolean if a field has been set.

### GetMaxPerInstanceSize

`func (o *PersistentVolumeQuotas) GetMaxPerInstanceSize() int64`

GetMaxPerInstanceSize returns the MaxPerInstanceSize field if non-nil, zero value otherwise.

### GetMaxPerInstanceSizeOk

`func (o *PersistentVolumeQuotas) GetMaxPerInstanceSizeOk() (*int64, bool)`

GetMaxPerInstanceSizeOk returns a tuple with the MaxPerInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerInstanceSize

`func (o *PersistentVolumeQuotas) SetMaxPerInstanceSize(v int64)`

SetMaxPerInstanceSize sets MaxPerInstanceSize field to given value.

### HasMaxPerInstanceSize

`func (o *PersistentVolumeQuotas) HasMaxPerInstanceSize() bool`

HasMaxPerInstanceSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


