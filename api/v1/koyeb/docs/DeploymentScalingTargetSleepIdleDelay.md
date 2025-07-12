# DeploymentScalingTargetSleepIdleDelay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int64** | DEPRECATED: use deep_sleep_value instead. Delay in seconds after which a service which received 0 request is put to deep sleep. | [optional] 
**DeepSleepValue** | Pointer to **int64** | Delay in seconds after which a service which received 0 request is put to deep sleep. | [optional] 
**LightSleepValue** | Pointer to **int64** | Delay in seconds after which a service which received 0 request is put to light sleep. | [optional] 

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

### GetDeepSleepValue

`func (o *DeploymentScalingTargetSleepIdleDelay) GetDeepSleepValue() int64`

GetDeepSleepValue returns the DeepSleepValue field if non-nil, zero value otherwise.

### GetDeepSleepValueOk

`func (o *DeploymentScalingTargetSleepIdleDelay) GetDeepSleepValueOk() (*int64, bool)`

GetDeepSleepValueOk returns a tuple with the DeepSleepValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepSleepValue

`func (o *DeploymentScalingTargetSleepIdleDelay) SetDeepSleepValue(v int64)`

SetDeepSleepValue sets DeepSleepValue field to given value.

### HasDeepSleepValue

`func (o *DeploymentScalingTargetSleepIdleDelay) HasDeepSleepValue() bool`

HasDeepSleepValue returns a boolean if a field has been set.

### GetLightSleepValue

`func (o *DeploymentScalingTargetSleepIdleDelay) GetLightSleepValue() int64`

GetLightSleepValue returns the LightSleepValue field if non-nil, zero value otherwise.

### GetLightSleepValueOk

`func (o *DeploymentScalingTargetSleepIdleDelay) GetLightSleepValueOk() (*int64, bool)`

GetLightSleepValueOk returns a tuple with the LightSleepValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightSleepValue

`func (o *DeploymentScalingTargetSleepIdleDelay) SetLightSleepValue(v int64)`

SetLightSleepValue sets LightSleepValue field to given value.

### HasLightSleepValue

`func (o *DeploymentScalingTargetSleepIdleDelay) HasLightSleepValue() bool`

HasLightSleepValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


