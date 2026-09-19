# PersistentVolumeRegionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**TotalSizeGbUsed** | Pointer to **int64** |  | [optional] 
**TotalSizeGbLimit** | Pointer to **int64** |  | [optional] 

## Methods

### NewPersistentVolumeRegionUsage

`func NewPersistentVolumeRegionUsage() *PersistentVolumeRegionUsage`

NewPersistentVolumeRegionUsage instantiates a new PersistentVolumeRegionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentVolumeRegionUsageWithDefaults

`func NewPersistentVolumeRegionUsageWithDefaults() *PersistentVolumeRegionUsage`

NewPersistentVolumeRegionUsageWithDefaults instantiates a new PersistentVolumeRegionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PersistentVolumeRegionUsage) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PersistentVolumeRegionUsage) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PersistentVolumeRegionUsage) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PersistentVolumeRegionUsage) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTotalSizeGbUsed

`func (o *PersistentVolumeRegionUsage) GetTotalSizeGbUsed() int64`

GetTotalSizeGbUsed returns the TotalSizeGbUsed field if non-nil, zero value otherwise.

### GetTotalSizeGbUsedOk

`func (o *PersistentVolumeRegionUsage) GetTotalSizeGbUsedOk() (*int64, bool)`

GetTotalSizeGbUsedOk returns a tuple with the TotalSizeGbUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeGbUsed

`func (o *PersistentVolumeRegionUsage) SetTotalSizeGbUsed(v int64)`

SetTotalSizeGbUsed sets TotalSizeGbUsed field to given value.

### HasTotalSizeGbUsed

`func (o *PersistentVolumeRegionUsage) HasTotalSizeGbUsed() bool`

HasTotalSizeGbUsed returns a boolean if a field has been set.

### GetTotalSizeGbLimit

`func (o *PersistentVolumeRegionUsage) GetTotalSizeGbLimit() int64`

GetTotalSizeGbLimit returns the TotalSizeGbLimit field if non-nil, zero value otherwise.

### GetTotalSizeGbLimitOk

`func (o *PersistentVolumeRegionUsage) GetTotalSizeGbLimitOk() (*int64, bool)`

GetTotalSizeGbLimitOk returns a tuple with the TotalSizeGbLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeGbLimit

`func (o *PersistentVolumeRegionUsage) SetTotalSizeGbLimit(v int64)`

SetTotalSizeGbLimit sets TotalSizeGbLimit field to given value.

### HasTotalSizeGbLimit

`func (o *PersistentVolumeRegionUsage) HasTotalSizeGbLimit() bool`

HasTotalSizeGbLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


