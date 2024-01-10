# DeploymentScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to **[]string** |  | [optional] 
**Min** | Pointer to **int64** |  | [optional] 
**Max** | Pointer to **int64** |  | [optional] 
**Targets** | Pointer to [**[]DeploymentScalingTarget**](DeploymentScalingTarget.md) |  | [optional] 

## Methods

### NewDeploymentScaling

`func NewDeploymentScaling() *DeploymentScaling`

NewDeploymentScaling instantiates a new DeploymentScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentScalingWithDefaults

`func NewDeploymentScalingWithDefaults() *DeploymentScaling`

NewDeploymentScalingWithDefaults instantiates a new DeploymentScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *DeploymentScaling) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeploymentScaling) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeploymentScaling) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DeploymentScaling) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetMin

`func (o *DeploymentScaling) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DeploymentScaling) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DeploymentScaling) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *DeploymentScaling) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *DeploymentScaling) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DeploymentScaling) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DeploymentScaling) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *DeploymentScaling) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetTargets

`func (o *DeploymentScaling) GetTargets() []DeploymentScalingTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *DeploymentScaling) GetTargetsOk() (*[]DeploymentScalingTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *DeploymentScaling) SetTargets(v []DeploymentScalingTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *DeploymentScaling) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


