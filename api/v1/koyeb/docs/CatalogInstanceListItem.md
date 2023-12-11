# CatalogInstanceListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpu** | Pointer to **int64** | The number of cpus. Deprecated. Use vcpu_shares instead. | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **string** |  | [optional] 
**PriceHourly** | Pointer to **string** |  | [optional] 
**PriceMonthly** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RequirePlan** | Pointer to **[]string** |  | [optional] 
**VcpuShares** | Pointer to **float32** | The number of vcpu shares reserved for the instance. | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


