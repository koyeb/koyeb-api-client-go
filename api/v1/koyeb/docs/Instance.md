# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpu** | Pointer to **int64** |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **string** |  | [optional] 
**PriceHourly** | Pointer to **string** |  | [optional] 
**PriceMonthly** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *Instance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Instance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Instance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Instance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpu

`func (o *Instance) GetVcpu() int64`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *Instance) GetVcpuOk() (*int64, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *Instance) SetVcpu(v int64)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *Instance) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetMemory

`func (o *Instance) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Instance) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Instance) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Instance) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDisk

`func (o *Instance) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Instance) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Instance) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Instance) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetPriceHourly

`func (o *Instance) GetPriceHourly() string`

GetPriceHourly returns the PriceHourly field if non-nil, zero value otherwise.

### GetPriceHourlyOk

`func (o *Instance) GetPriceHourlyOk() (*string, bool)`

GetPriceHourlyOk returns a tuple with the PriceHourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHourly

`func (o *Instance) SetPriceHourly(v string)`

SetPriceHourly sets PriceHourly field to given value.

### HasPriceHourly

`func (o *Instance) HasPriceHourly() bool`

HasPriceHourly returns a boolean if a field has been set.

### GetPriceMonthly

`func (o *Instance) GetPriceMonthly() string`

GetPriceMonthly returns the PriceMonthly field if non-nil, zero value otherwise.

### GetPriceMonthlyOk

`func (o *Instance) GetPriceMonthlyOk() (*string, bool)`

GetPriceMonthlyOk returns a tuple with the PriceMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthly

`func (o *Instance) SetPriceMonthly(v string)`

SetPriceMonthly sets PriceMonthly field to given value.

### HasPriceMonthly

`func (o *Instance) HasPriceMonthly() bool`

HasPriceMonthly returns a boolean if a field has been set.

### GetRegions

`func (o *Instance) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Instance) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Instance) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Instance) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStatus

`func (o *Instance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Instance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


