# DeploymentStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DeploymentStrategyType**](DeploymentStrategyType.md) |  | [optional] [default to DEPLOYMENTSTRATEGYTYPE_INVALID]

## Methods

### NewDeploymentStrategy

`func NewDeploymentStrategy() *DeploymentStrategy`

NewDeploymentStrategy instantiates a new DeploymentStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStrategyWithDefaults

`func NewDeploymentStrategyWithDefaults() *DeploymentStrategy`

NewDeploymentStrategyWithDefaults instantiates a new DeploymentStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentStrategy) GetType() DeploymentStrategyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentStrategy) GetTypeOk() (*DeploymentStrategyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentStrategy) SetType(v DeploymentStrategyType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


