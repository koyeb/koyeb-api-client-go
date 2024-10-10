# CatalogInstanceListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpu** | Pointer to **int64** | The number of cpus. Deprecated. Use vcpu_shares instead. | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **string** |  | [optional] 
**PricePerSecond** | Pointer to **string** |  | [optional] 
**PriceHourly** | Pointer to **string** |  | [optional] 
**PriceMonthly** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RequirePlan** | Pointer to **[]string** |  | [optional] 
**VcpuShares** | Pointer to **float32** | The number of vcpu shares reserved for the instance. | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to [**CatalogGPUDetails**](CatalogGPUDetails.md) |  | [optional] 
**ServiceTypes** | Pointer to **[]string** |  | [optional] 
**VolumesEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCatalogInstanceListItem

`func NewCatalogInstanceListItem() *CatalogInstanceListItem`

NewCatalogInstanceListItem instantiates a new CatalogInstanceListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogInstanceListItemWithDefaults

`func NewCatalogInstanceListItemWithDefaults() *CatalogInstanceListItem`

NewCatalogInstanceListItemWithDefaults instantiates a new CatalogInstanceListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogInstanceListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogInstanceListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogInstanceListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogInstanceListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogInstanceListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogInstanceListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogInstanceListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogInstanceListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpu

`func (o *CatalogInstanceListItem) GetVcpu() int64`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *CatalogInstanceListItem) GetVcpuOk() (*int64, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *CatalogInstanceListItem) SetVcpu(v int64)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *CatalogInstanceListItem) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetMemory

`func (o *CatalogInstanceListItem) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CatalogInstanceListItem) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CatalogInstanceListItem) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CatalogInstanceListItem) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDisk

`func (o *CatalogInstanceListItem) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CatalogInstanceListItem) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CatalogInstanceListItem) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *CatalogInstanceListItem) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetPricePerSecond

`func (o *CatalogInstanceListItem) GetPricePerSecond() string`

GetPricePerSecond returns the PricePerSecond field if non-nil, zero value otherwise.

### GetPricePerSecondOk

`func (o *CatalogInstanceListItem) GetPricePerSecondOk() (*string, bool)`

GetPricePerSecondOk returns a tuple with the PricePerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerSecond

`func (o *CatalogInstanceListItem) SetPricePerSecond(v string)`

SetPricePerSecond sets PricePerSecond field to given value.

### HasPricePerSecond

`func (o *CatalogInstanceListItem) HasPricePerSecond() bool`

HasPricePerSecond returns a boolean if a field has been set.

### GetPriceHourly

`func (o *CatalogInstanceListItem) GetPriceHourly() string`

GetPriceHourly returns the PriceHourly field if non-nil, zero value otherwise.

### GetPriceHourlyOk

`func (o *CatalogInstanceListItem) GetPriceHourlyOk() (*string, bool)`

GetPriceHourlyOk returns a tuple with the PriceHourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHourly

`func (o *CatalogInstanceListItem) SetPriceHourly(v string)`

SetPriceHourly sets PriceHourly field to given value.

### HasPriceHourly

`func (o *CatalogInstanceListItem) HasPriceHourly() bool`

HasPriceHourly returns a boolean if a field has been set.

### GetPriceMonthly

`func (o *CatalogInstanceListItem) GetPriceMonthly() string`

GetPriceMonthly returns the PriceMonthly field if non-nil, zero value otherwise.

### GetPriceMonthlyOk

`func (o *CatalogInstanceListItem) GetPriceMonthlyOk() (*string, bool)`

GetPriceMonthlyOk returns a tuple with the PriceMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthly

`func (o *CatalogInstanceListItem) SetPriceMonthly(v string)`

SetPriceMonthly sets PriceMonthly field to given value.

### HasPriceMonthly

`func (o *CatalogInstanceListItem) HasPriceMonthly() bool`

HasPriceMonthly returns a boolean if a field has been set.

### GetRegions

`func (o *CatalogInstanceListItem) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CatalogInstanceListItem) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CatalogInstanceListItem) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CatalogInstanceListItem) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogInstanceListItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogInstanceListItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogInstanceListItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogInstanceListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequirePlan

`func (o *CatalogInstanceListItem) GetRequirePlan() []string`

GetRequirePlan returns the RequirePlan field if non-nil, zero value otherwise.

### GetRequirePlanOk

`func (o *CatalogInstanceListItem) GetRequirePlanOk() (*[]string, bool)`

GetRequirePlanOk returns a tuple with the RequirePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePlan

`func (o *CatalogInstanceListItem) SetRequirePlan(v []string)`

SetRequirePlan sets RequirePlan field to given value.

### HasRequirePlan

`func (o *CatalogInstanceListItem) HasRequirePlan() bool`

HasRequirePlan returns a boolean if a field has been set.

### GetVcpuShares

`func (o *CatalogInstanceListItem) GetVcpuShares() float32`

GetVcpuShares returns the VcpuShares field if non-nil, zero value otherwise.

### GetVcpuSharesOk

`func (o *CatalogInstanceListItem) GetVcpuSharesOk() (*float32, bool)`

GetVcpuSharesOk returns a tuple with the VcpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpuShares

`func (o *CatalogInstanceListItem) SetVcpuShares(v float32)`

SetVcpuShares sets VcpuShares field to given value.

### HasVcpuShares

`func (o *CatalogInstanceListItem) HasVcpuShares() bool`

HasVcpuShares returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogInstanceListItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogInstanceListItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogInstanceListItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogInstanceListItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAliases

`func (o *CatalogInstanceListItem) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CatalogInstanceListItem) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CatalogInstanceListItem) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CatalogInstanceListItem) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetType

`func (o *CatalogInstanceListItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogInstanceListItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogInstanceListItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogInstanceListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGpu

`func (o *CatalogInstanceListItem) GetGpu() CatalogGPUDetails`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *CatalogInstanceListItem) GetGpuOk() (*CatalogGPUDetails, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *CatalogInstanceListItem) SetGpu(v CatalogGPUDetails)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *CatalogInstanceListItem) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetServiceTypes

`func (o *CatalogInstanceListItem) GetServiceTypes() []string`

GetServiceTypes returns the ServiceTypes field if non-nil, zero value otherwise.

### GetServiceTypesOk

`func (o *CatalogInstanceListItem) GetServiceTypesOk() (*[]string, bool)`

GetServiceTypesOk returns a tuple with the ServiceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypes

`func (o *CatalogInstanceListItem) SetServiceTypes(v []string)`

SetServiceTypes sets ServiceTypes field to given value.

### HasServiceTypes

`func (o *CatalogInstanceListItem) HasServiceTypes() bool`

HasServiceTypes returns a boolean if a field has been set.

### GetVolumesEnabled

`func (o *CatalogInstanceListItem) GetVolumesEnabled() bool`

GetVolumesEnabled returns the VolumesEnabled field if non-nil, zero value otherwise.

### GetVolumesEnabledOk

`func (o *CatalogInstanceListItem) GetVolumesEnabledOk() (*bool, bool)`

GetVolumesEnabledOk returns a tuple with the VolumesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEnabled

`func (o *CatalogInstanceListItem) SetVolumesEnabled(v bool)`

SetVolumesEnabled sets VolumesEnabled field to given value.

### HasVolumesEnabled

`func (o *CatalogInstanceListItem) HasVolumesEnabled() bool`

HasVolumesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


