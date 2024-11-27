# CatalogInstance

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

### NewCatalogInstance

`func NewCatalogInstance() *CatalogInstance`

NewCatalogInstance instantiates a new CatalogInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogInstanceWithDefaults

`func NewCatalogInstanceWithDefaults() *CatalogInstance`

NewCatalogInstanceWithDefaults instantiates a new CatalogInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpu

`func (o *CatalogInstance) GetVcpu() int64`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *CatalogInstance) GetVcpuOk() (*int64, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *CatalogInstance) SetVcpu(v int64)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *CatalogInstance) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetMemory

`func (o *CatalogInstance) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CatalogInstance) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CatalogInstance) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CatalogInstance) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDisk

`func (o *CatalogInstance) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CatalogInstance) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CatalogInstance) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *CatalogInstance) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetPricePerSecond

`func (o *CatalogInstance) GetPricePerSecond() string`

GetPricePerSecond returns the PricePerSecond field if non-nil, zero value otherwise.

### GetPricePerSecondOk

`func (o *CatalogInstance) GetPricePerSecondOk() (*string, bool)`

GetPricePerSecondOk returns a tuple with the PricePerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerSecond

`func (o *CatalogInstance) SetPricePerSecond(v string)`

SetPricePerSecond sets PricePerSecond field to given value.

### HasPricePerSecond

`func (o *CatalogInstance) HasPricePerSecond() bool`

HasPricePerSecond returns a boolean if a field has been set.

### GetPriceHourly

`func (o *CatalogInstance) GetPriceHourly() string`

GetPriceHourly returns the PriceHourly field if non-nil, zero value otherwise.

### GetPriceHourlyOk

`func (o *CatalogInstance) GetPriceHourlyOk() (*string, bool)`

GetPriceHourlyOk returns a tuple with the PriceHourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHourly

`func (o *CatalogInstance) SetPriceHourly(v string)`

SetPriceHourly sets PriceHourly field to given value.

### HasPriceHourly

`func (o *CatalogInstance) HasPriceHourly() bool`

HasPriceHourly returns a boolean if a field has been set.

### GetPriceMonthly

`func (o *CatalogInstance) GetPriceMonthly() string`

GetPriceMonthly returns the PriceMonthly field if non-nil, zero value otherwise.

### GetPriceMonthlyOk

`func (o *CatalogInstance) GetPriceMonthlyOk() (*string, bool)`

GetPriceMonthlyOk returns a tuple with the PriceMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthly

`func (o *CatalogInstance) SetPriceMonthly(v string)`

SetPriceMonthly sets PriceMonthly field to given value.

### HasPriceMonthly

`func (o *CatalogInstance) HasPriceMonthly() bool`

HasPriceMonthly returns a boolean if a field has been set.

### GetRegions

`func (o *CatalogInstance) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CatalogInstance) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CatalogInstance) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CatalogInstance) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequirePlan

`func (o *CatalogInstance) GetRequirePlan() []string`

GetRequirePlan returns the RequirePlan field if non-nil, zero value otherwise.

### GetRequirePlanOk

`func (o *CatalogInstance) GetRequirePlanOk() (*[]string, bool)`

GetRequirePlanOk returns a tuple with the RequirePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePlan

`func (o *CatalogInstance) SetRequirePlan(v []string)`

SetRequirePlan sets RequirePlan field to given value.

### HasRequirePlan

`func (o *CatalogInstance) HasRequirePlan() bool`

HasRequirePlan returns a boolean if a field has been set.

### GetVcpuShares

`func (o *CatalogInstance) GetVcpuShares() float32`

GetVcpuShares returns the VcpuShares field if non-nil, zero value otherwise.

### GetVcpuSharesOk

`func (o *CatalogInstance) GetVcpuSharesOk() (*float32, bool)`

GetVcpuSharesOk returns a tuple with the VcpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpuShares

`func (o *CatalogInstance) SetVcpuShares(v float32)`

SetVcpuShares sets VcpuShares field to given value.

### HasVcpuShares

`func (o *CatalogInstance) HasVcpuShares() bool`

HasVcpuShares returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAliases

`func (o *CatalogInstance) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CatalogInstance) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CatalogInstance) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CatalogInstance) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetType

`func (o *CatalogInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGpu

`func (o *CatalogInstance) GetGpu() CatalogGPUDetails`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *CatalogInstance) GetGpuOk() (*CatalogGPUDetails, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *CatalogInstance) SetGpu(v CatalogGPUDetails)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *CatalogInstance) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetServiceTypes

`func (o *CatalogInstance) GetServiceTypes() []string`

GetServiceTypes returns the ServiceTypes field if non-nil, zero value otherwise.

### GetServiceTypesOk

`func (o *CatalogInstance) GetServiceTypesOk() (*[]string, bool)`

GetServiceTypesOk returns a tuple with the ServiceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypes

`func (o *CatalogInstance) SetServiceTypes(v []string)`

SetServiceTypes sets ServiceTypes field to given value.

### HasServiceTypes

`func (o *CatalogInstance) HasServiceTypes() bool`

HasServiceTypes returns a boolean if a field has been set.

### GetVolumesEnabled

`func (o *CatalogInstance) GetVolumesEnabled() bool`

GetVolumesEnabled returns the VolumesEnabled field if non-nil, zero value otherwise.

### GetVolumesEnabledOk

`func (o *CatalogInstance) GetVolumesEnabledOk() (*bool, bool)`

GetVolumesEnabledOk returns a tuple with the VolumesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEnabled

`func (o *CatalogInstance) SetVolumesEnabled(v bool)`

SetVolumesEnabled sets VolumesEnabled field to given value.

### HasVolumesEnabled

`func (o *CatalogInstance) HasVolumesEnabled() bool`

HasVolumesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


