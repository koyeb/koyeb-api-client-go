# ServiceLifeCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAfterSleep** | Pointer to **int64** |  | [optional] 
**DeleteAfterCreate** | Pointer to **int64** |  | [optional] 

## Methods

### NewServiceLifeCycle

`func NewServiceLifeCycle() *ServiceLifeCycle`

NewServiceLifeCycle instantiates a new ServiceLifeCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLifeCycleWithDefaults

`func NewServiceLifeCycleWithDefaults() *ServiceLifeCycle`

NewServiceLifeCycleWithDefaults instantiates a new ServiceLifeCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAfterSleep

`func (o *ServiceLifeCycle) GetDeleteAfterSleep() int64`

GetDeleteAfterSleep returns the DeleteAfterSleep field if non-nil, zero value otherwise.

### GetDeleteAfterSleepOk

`func (o *ServiceLifeCycle) GetDeleteAfterSleepOk() (*int64, bool)`

GetDeleteAfterSleepOk returns a tuple with the DeleteAfterSleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterSleep

`func (o *ServiceLifeCycle) SetDeleteAfterSleep(v int64)`

SetDeleteAfterSleep sets DeleteAfterSleep field to given value.

### HasDeleteAfterSleep

`func (o *ServiceLifeCycle) HasDeleteAfterSleep() bool`

HasDeleteAfterSleep returns a boolean if a field has been set.

### GetDeleteAfterCreate

`func (o *ServiceLifeCycle) GetDeleteAfterCreate() int64`

GetDeleteAfterCreate returns the DeleteAfterCreate field if non-nil, zero value otherwise.

### GetDeleteAfterCreateOk

`func (o *ServiceLifeCycle) GetDeleteAfterCreateOk() (*int64, bool)`

GetDeleteAfterCreateOk returns a tuple with the DeleteAfterCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterCreate

`func (o *ServiceLifeCycle) SetDeleteAfterCreate(v int64)`

SetDeleteAfterCreate sets DeleteAfterCreate field to given value.

### HasDeleteAfterCreate

`func (o *ServiceLifeCycle) HasDeleteAfterCreate() bool`

HasDeleteAfterCreate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


