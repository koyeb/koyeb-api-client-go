# CreateService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**DeploymentDefinition**](DeploymentDefinition.md) |  | [optional] 
**LifeCycle** | Pointer to [**ServiceLifeCycle**](ServiceLifeCycle.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateService

`func NewCreateService() *CreateService`

NewCreateService instantiates a new CreateService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceWithDefaults

`func NewCreateServiceWithDefaults() *CreateService`

NewCreateServiceWithDefaults instantiates a new CreateService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *CreateService) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateService) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateService) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateService) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDefinition

`func (o *CreateService) GetDefinition() DeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateService) GetDefinitionOk() (*DeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateService) SetDefinition(v DeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *CreateService) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetLifeCycle

`func (o *CreateService) GetLifeCycle() ServiceLifeCycle`

GetLifeCycle returns the LifeCycle field if non-nil, zero value otherwise.

### GetLifeCycleOk

`func (o *CreateService) GetLifeCycleOk() (*ServiceLifeCycle, bool)`

GetLifeCycleOk returns a tuple with the LifeCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycle

`func (o *CreateService) SetLifeCycle(v ServiceLifeCycle)`

SetLifeCycle sets LifeCycle field to given value.

### HasLifeCycle

`func (o *CreateService) HasLifeCycle() bool`

HasLifeCycle returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateService) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateService) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateService) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateService) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


