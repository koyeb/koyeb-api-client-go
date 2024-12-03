# DeploymentScalingTargetSleepIdleDelay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int64** | Delay in seconds after which a service which received 0 request is scaled to 0. This is not configurable and must be set to 300 (5 minutes). Get in touch to tune it. | [optional] 

## Methods

### NewDeploymentScalingTargetSleepIdleDelay

`func NewDeploymentScalingTargetSleepIdleDelay() *DeploymentScalingTargetSleepIdleDelay`

NewDeploymentScalingTargetSleepIdleDelay instantiates a new DeploymentScalingTargetSleepIdleDelay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentScalingTargetSleepIdleDelayWithDefaults

`func NewDeploymentScalingTargetSleepIdleDelayWithDefaults() *DeploymentScalingTargetSleepIdleDelay`

NewDeploymentScalingTargetSleepIdleDelayWithDefaults instantiates a new DeploymentScalingTargetSleepIdleDelay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DeploymentScalingTargetSleepIdleDelay) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeploymentScalingTargetSleepIdleDelay) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeploymentScalingTargetSleepIdleDelay) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeploymentScalingTargetSleepIdleDelay) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


