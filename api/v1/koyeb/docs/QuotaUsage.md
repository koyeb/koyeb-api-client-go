# QuotaUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppsUsed** | Pointer to **int64** |  | [optional] 
**AppsLimit** | Pointer to **int64** |  | [optional] 
**ServicesUsed** | Pointer to **int64** |  | [optional] 
**ServicesLimit** | Pointer to **int64** |  | [optional] 
**MemoryMbUsed** | Pointer to **int64** |  | [optional] 
**MemoryMbLimit** | Pointer to **int64** |  | [optional] 
**CustomDomainsUsed** | Pointer to **int64** |  | [optional] 
**CustomDomainsLimit** | Pointer to **int64** |  | [optional] 
**KoyebLbDomainsUsed** | Pointer to **int64** |  | [optional] 
**KoyebLbDomainsLimit** | Pointer to **int64** |  | [optional] 
**ProxyPortsUsed** | Pointer to **int64** |  | [optional] 
**ProxyPortsLimit** | Pointer to **int64** |  | [optional] 
**InstancesByType** | Pointer to [**[]InstanceTypeUsage**](InstanceTypeUsage.md) |  | [optional] 
**PersistentVolumesByRegion** | Pointer to [**[]PersistentVolumeRegionUsage**](PersistentVolumeRegionUsage.md) |  | [optional] 
**InstanceSnapshotsByType** | Pointer to [**[]InstanceSnapshotTypeUsage**](InstanceSnapshotTypeUsage.md) |  | [optional] 

## Methods

### NewQuotaUsage

`func NewQuotaUsage() *QuotaUsage`

NewQuotaUsage instantiates a new QuotaUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaUsageWithDefaults

`func NewQuotaUsageWithDefaults() *QuotaUsage`

NewQuotaUsageWithDefaults instantiates a new QuotaUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsUsed

`func (o *QuotaUsage) GetAppsUsed() int64`

GetAppsUsed returns the AppsUsed field if non-nil, zero value otherwise.

### GetAppsUsedOk

`func (o *QuotaUsage) GetAppsUsedOk() (*int64, bool)`

GetAppsUsedOk returns a tuple with the AppsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsUsed

`func (o *QuotaUsage) SetAppsUsed(v int64)`

SetAppsUsed sets AppsUsed field to given value.

### HasAppsUsed

`func (o *QuotaUsage) HasAppsUsed() bool`

HasAppsUsed returns a boolean if a field has been set.

### GetAppsLimit

`func (o *QuotaUsage) GetAppsLimit() int64`

GetAppsLimit returns the AppsLimit field if non-nil, zero value otherwise.

### GetAppsLimitOk

`func (o *QuotaUsage) GetAppsLimitOk() (*int64, bool)`

GetAppsLimitOk returns a tuple with the AppsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsLimit

`func (o *QuotaUsage) SetAppsLimit(v int64)`

SetAppsLimit sets AppsLimit field to given value.

### HasAppsLimit

`func (o *QuotaUsage) HasAppsLimit() bool`

HasAppsLimit returns a boolean if a field has been set.

### GetServicesUsed

`func (o *QuotaUsage) GetServicesUsed() int64`

GetServicesUsed returns the ServicesUsed field if non-nil, zero value otherwise.

### GetServicesUsedOk

`func (o *QuotaUsage) GetServicesUsedOk() (*int64, bool)`

GetServicesUsedOk returns a tuple with the ServicesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesUsed

`func (o *QuotaUsage) SetServicesUsed(v int64)`

SetServicesUsed sets ServicesUsed field to given value.

### HasServicesUsed

`func (o *QuotaUsage) HasServicesUsed() bool`

HasServicesUsed returns a boolean if a field has been set.

### GetServicesLimit

`func (o *QuotaUsage) GetServicesLimit() int64`

GetServicesLimit returns the ServicesLimit field if non-nil, zero value otherwise.

### GetServicesLimitOk

`func (o *QuotaUsage) GetServicesLimitOk() (*int64, bool)`

GetServicesLimitOk returns a tuple with the ServicesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesLimit

`func (o *QuotaUsage) SetServicesLimit(v int64)`

SetServicesLimit sets ServicesLimit field to given value.

### HasServicesLimit

`func (o *QuotaUsage) HasServicesLimit() bool`

HasServicesLimit returns a boolean if a field has been set.

### GetMemoryMbUsed

`func (o *QuotaUsage) GetMemoryMbUsed() int64`

GetMemoryMbUsed returns the MemoryMbUsed field if non-nil, zero value otherwise.

### GetMemoryMbUsedOk

`func (o *QuotaUsage) GetMemoryMbUsedOk() (*int64, bool)`

GetMemoryMbUsedOk returns a tuple with the MemoryMbUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbUsed

`func (o *QuotaUsage) SetMemoryMbUsed(v int64)`

SetMemoryMbUsed sets MemoryMbUsed field to given value.

### HasMemoryMbUsed

`func (o *QuotaUsage) HasMemoryMbUsed() bool`

HasMemoryMbUsed returns a boolean if a field has been set.

### GetMemoryMbLimit

`func (o *QuotaUsage) GetMemoryMbLimit() int64`

GetMemoryMbLimit returns the MemoryMbLimit field if non-nil, zero value otherwise.

### GetMemoryMbLimitOk

`func (o *QuotaUsage) GetMemoryMbLimitOk() (*int64, bool)`

GetMemoryMbLimitOk returns a tuple with the MemoryMbLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbLimit

`func (o *QuotaUsage) SetMemoryMbLimit(v int64)`

SetMemoryMbLimit sets MemoryMbLimit field to given value.

### HasMemoryMbLimit

`func (o *QuotaUsage) HasMemoryMbLimit() bool`

HasMemoryMbLimit returns a boolean if a field has been set.

### GetCustomDomainsUsed

`func (o *QuotaUsage) GetCustomDomainsUsed() int64`

GetCustomDomainsUsed returns the CustomDomainsUsed field if non-nil, zero value otherwise.

### GetCustomDomainsUsedOk

`func (o *QuotaUsage) GetCustomDomainsUsedOk() (*int64, bool)`

GetCustomDomainsUsedOk returns a tuple with the CustomDomainsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomainsUsed

`func (o *QuotaUsage) SetCustomDomainsUsed(v int64)`

SetCustomDomainsUsed sets CustomDomainsUsed field to given value.

### HasCustomDomainsUsed

`func (o *QuotaUsage) HasCustomDomainsUsed() bool`

HasCustomDomainsUsed returns a boolean if a field has been set.

### GetCustomDomainsLimit

`func (o *QuotaUsage) GetCustomDomainsLimit() int64`

GetCustomDomainsLimit returns the CustomDomainsLimit field if non-nil, zero value otherwise.

### GetCustomDomainsLimitOk

`func (o *QuotaUsage) GetCustomDomainsLimitOk() (*int64, bool)`

GetCustomDomainsLimitOk returns a tuple with the CustomDomainsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomainsLimit

`func (o *QuotaUsage) SetCustomDomainsLimit(v int64)`

SetCustomDomainsLimit sets CustomDomainsLimit field to given value.

### HasCustomDomainsLimit

`func (o *QuotaUsage) HasCustomDomainsLimit() bool`

HasCustomDomainsLimit returns a boolean if a field has been set.

### GetKoyebLbDomainsUsed

`func (o *QuotaUsage) GetKoyebLbDomainsUsed() int64`

GetKoyebLbDomainsUsed returns the KoyebLbDomainsUsed field if non-nil, zero value otherwise.

### GetKoyebLbDomainsUsedOk

`func (o *QuotaUsage) GetKoyebLbDomainsUsedOk() (*int64, bool)`

GetKoyebLbDomainsUsedOk returns a tuple with the KoyebLbDomainsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKoyebLbDomainsUsed

`func (o *QuotaUsage) SetKoyebLbDomainsUsed(v int64)`

SetKoyebLbDomainsUsed sets KoyebLbDomainsUsed field to given value.

### HasKoyebLbDomainsUsed

`func (o *QuotaUsage) HasKoyebLbDomainsUsed() bool`

HasKoyebLbDomainsUsed returns a boolean if a field has been set.

### GetKoyebLbDomainsLimit

`func (o *QuotaUsage) GetKoyebLbDomainsLimit() int64`

GetKoyebLbDomainsLimit returns the KoyebLbDomainsLimit field if non-nil, zero value otherwise.

### GetKoyebLbDomainsLimitOk

`func (o *QuotaUsage) GetKoyebLbDomainsLimitOk() (*int64, bool)`

GetKoyebLbDomainsLimitOk returns a tuple with the KoyebLbDomainsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKoyebLbDomainsLimit

`func (o *QuotaUsage) SetKoyebLbDomainsLimit(v int64)`

SetKoyebLbDomainsLimit sets KoyebLbDomainsLimit field to given value.

### HasKoyebLbDomainsLimit

`func (o *QuotaUsage) HasKoyebLbDomainsLimit() bool`

HasKoyebLbDomainsLimit returns a boolean if a field has been set.

### GetProxyPortsUsed

`func (o *QuotaUsage) GetProxyPortsUsed() int64`

GetProxyPortsUsed returns the ProxyPortsUsed field if non-nil, zero value otherwise.

### GetProxyPortsUsedOk

`func (o *QuotaUsage) GetProxyPortsUsedOk() (*int64, bool)`

GetProxyPortsUsedOk returns a tuple with the ProxyPortsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPortsUsed

`func (o *QuotaUsage) SetProxyPortsUsed(v int64)`

SetProxyPortsUsed sets ProxyPortsUsed field to given value.

### HasProxyPortsUsed

`func (o *QuotaUsage) HasProxyPortsUsed() bool`

HasProxyPortsUsed returns a boolean if a field has been set.

### GetProxyPortsLimit

`func (o *QuotaUsage) GetProxyPortsLimit() int64`

GetProxyPortsLimit returns the ProxyPortsLimit field if non-nil, zero value otherwise.

### GetProxyPortsLimitOk

`func (o *QuotaUsage) GetProxyPortsLimitOk() (*int64, bool)`

GetProxyPortsLimitOk returns a tuple with the ProxyPortsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPortsLimit

`func (o *QuotaUsage) SetProxyPortsLimit(v int64)`

SetProxyPortsLimit sets ProxyPortsLimit field to given value.

### HasProxyPortsLimit

`func (o *QuotaUsage) HasProxyPortsLimit() bool`

HasProxyPortsLimit returns a boolean if a field has been set.

### GetInstancesByType

`func (o *QuotaUsage) GetInstancesByType() []InstanceTypeUsage`

GetInstancesByType returns the InstancesByType field if non-nil, zero value otherwise.

### GetInstancesByTypeOk

`func (o *QuotaUsage) GetInstancesByTypeOk() (*[]InstanceTypeUsage, bool)`

GetInstancesByTypeOk returns a tuple with the InstancesByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesByType

`func (o *QuotaUsage) SetInstancesByType(v []InstanceTypeUsage)`

SetInstancesByType sets InstancesByType field to given value.

### HasInstancesByType

`func (o *QuotaUsage) HasInstancesByType() bool`

HasInstancesByType returns a boolean if a field has been set.

### GetPersistentVolumesByRegion

`func (o *QuotaUsage) GetPersistentVolumesByRegion() []PersistentVolumeRegionUsage`

GetPersistentVolumesByRegion returns the PersistentVolumesByRegion field if non-nil, zero value otherwise.

### GetPersistentVolumesByRegionOk

`func (o *QuotaUsage) GetPersistentVolumesByRegionOk() (*[]PersistentVolumeRegionUsage, bool)`

GetPersistentVolumesByRegionOk returns a tuple with the PersistentVolumesByRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumesByRegion

`func (o *QuotaUsage) SetPersistentVolumesByRegion(v []PersistentVolumeRegionUsage)`

SetPersistentVolumesByRegion sets PersistentVolumesByRegion field to given value.

### HasPersistentVolumesByRegion

`func (o *QuotaUsage) HasPersistentVolumesByRegion() bool`

HasPersistentVolumesByRegion returns a boolean if a field has been set.

### GetInstanceSnapshotsByType

`func (o *QuotaUsage) GetInstanceSnapshotsByType() []InstanceSnapshotTypeUsage`

GetInstanceSnapshotsByType returns the InstanceSnapshotsByType field if non-nil, zero value otherwise.

### GetInstanceSnapshotsByTypeOk

`func (o *QuotaUsage) GetInstanceSnapshotsByTypeOk() (*[]InstanceSnapshotTypeUsage, bool)`

GetInstanceSnapshotsByTypeOk returns a tuple with the InstanceSnapshotsByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSnapshotsByType

`func (o *QuotaUsage) SetInstanceSnapshotsByType(v []InstanceSnapshotTypeUsage)`

SetInstanceSnapshotsByType sets InstanceSnapshotsByType field to given value.

### HasInstanceSnapshotsByType

`func (o *QuotaUsage) HasInstanceSnapshotsByType() bool`

HasInstanceSnapshotsByType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


