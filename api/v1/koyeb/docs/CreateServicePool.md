# CreateServicePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Definition** | Pointer to [**DeploymentDefinition**](DeploymentDefinition.md) |  | [optional] 

## Methods

### NewCreateServicePool

`func NewCreateServicePool() *CreateServicePool`

NewCreateServicePool instantiates a new CreateServicePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServicePoolWithDefaults

`func NewCreateServicePoolWithDefaults() *CreateServicePool`

NewCreateServicePoolWithDefaults instantiates a new CreateServicePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateServicePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServicePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServicePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateServicePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *CreateServicePool) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateServicePool) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateServicePool) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateServicePool) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDefinition

`func (o *CreateServicePool) GetDefinition() DeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateServicePool) GetDefinitionOk() (*DeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateServicePool) SetDefinition(v DeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *CreateServicePool) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


