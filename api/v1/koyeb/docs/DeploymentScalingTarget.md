# DeploymentScalingTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageCpu** | Pointer to [**DeploymentScalingTargetAverageCPU**](DeploymentScalingTargetAverageCPU.md) |  | [optional] 
**AverageMem** | Pointer to [**DeploymentScalingTargetAverageMem**](DeploymentScalingTargetAverageMem.md) |  | [optional] 
**RequestsPerSecond** | Pointer to [**DeploymentScalingTargetRequestsPerSecond**](DeploymentScalingTargetRequestsPerSecond.md) |  | [optional] 
**ConcurrentRequests** | Pointer to [**DeploymentScalingTargetConcurrentRequests**](DeploymentScalingTargetConcurrentRequests.md) |  | [optional] 
**RequestsResponseTime** | Pointer to [**DeploymentScalingTargetRequestsResponseTime**](DeploymentScalingTargetRequestsResponseTime.md) |  | [optional] 
**SleepIdleDelay** | Pointer to [**DeploymentScalingTargetSleepIdleDelay**](DeploymentScalingTargetSleepIdleDelay.md) |  | [optional] 

## Methods

### NewDeploymentScalingTarget

`func NewDeploymentScalingTarget() *DeploymentScalingTarget`

NewDeploymentScalingTarget instantiates a new DeploymentScalingTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentScalingTargetWithDefaults

`func NewDeploymentScalingTargetWithDefaults() *DeploymentScalingTarget`

NewDeploymentScalingTargetWithDefaults instantiates a new DeploymentScalingTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageCpu

`func (o *DeploymentScalingTarget) GetAverageCpu() DeploymentScalingTargetAverageCPU`

GetAverageCpu returns the AverageCpu field if non-nil, zero value otherwise.

### GetAverageCpuOk

`func (o *DeploymentScalingTarget) GetAverageCpuOk() (*DeploymentScalingTargetAverageCPU, bool)`

GetAverageCpuOk returns a tuple with the AverageCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCpu

`func (o *DeploymentScalingTarget) SetAverageCpu(v DeploymentScalingTargetAverageCPU)`

SetAverageCpu sets AverageCpu field to given value.

### HasAverageCpu

`func (o *DeploymentScalingTarget) HasAverageCpu() bool`

HasAverageCpu returns a boolean if a field has been set.

### GetAverageMem

`func (o *DeploymentScalingTarget) GetAverageMem() DeploymentScalingTargetAverageMem`

GetAverageMem returns the AverageMem field if non-nil, zero value otherwise.

### GetAverageMemOk

`func (o *DeploymentScalingTarget) GetAverageMemOk() (*DeploymentScalingTargetAverageMem, bool)`

GetAverageMemOk returns a tuple with the AverageMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMem

`func (o *DeploymentScalingTarget) SetAverageMem(v DeploymentScalingTargetAverageMem)`

SetAverageMem sets AverageMem field to given value.

### HasAverageMem

`func (o *DeploymentScalingTarget) HasAverageMem() bool`

HasAverageMem returns a boolean if a field has been set.

### GetRequestsPerSecond

`func (o *DeploymentScalingTarget) GetRequestsPerSecond() DeploymentScalingTargetRequestsPerSecond`

GetRequestsPerSecond returns the RequestsPerSecond field if non-nil, zero value otherwise.

### GetRequestsPerSecondOk

`func (o *DeploymentScalingTarget) GetRequestsPerSecondOk() (*DeploymentScalingTargetRequestsPerSecond, bool)`

GetRequestsPerSecondOk returns a tuple with the RequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerSecond

`func (o *DeploymentScalingTarget) SetRequestsPerSecond(v DeploymentScalingTargetRequestsPerSecond)`

SetRequestsPerSecond sets RequestsPerSecond field to given value.

### HasRequestsPerSecond

`func (o *DeploymentScalingTarget) HasRequestsPerSecond() bool`

HasRequestsPerSecond returns a boolean if a field has been set.

### GetConcurrentRequests

`func (o *DeploymentScalingTarget) GetConcurrentRequests() DeploymentScalingTargetConcurrentRequests`

GetConcurrentRequests returns the ConcurrentRequests field if non-nil, zero value otherwise.

### GetConcurrentRequestsOk

`func (o *DeploymentScalingTarget) GetConcurrentRequestsOk() (*DeploymentScalingTargetConcurrentRequests, bool)`

GetConcurrentRequestsOk returns a tuple with the ConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentRequests

`func (o *DeploymentScalingTarget) SetConcurrentRequests(v DeploymentScalingTargetConcurrentRequests)`

SetConcurrentRequests sets ConcurrentRequests field to given value.

### HasConcurrentRequests

`func (o *DeploymentScalingTarget) HasConcurrentRequests() bool`

HasConcurrentRequests returns a boolean if a field has been set.

### GetRequestsResponseTime

`func (o *DeploymentScalingTarget) GetRequestsResponseTime() DeploymentScalingTargetRequestsResponseTime`

GetRequestsResponseTime returns the RequestsResponseTime field if non-nil, zero value otherwise.

### GetRequestsResponseTimeOk

`func (o *DeploymentScalingTarget) GetRequestsResponseTimeOk() (*DeploymentScalingTargetRequestsResponseTime, bool)`

GetRequestsResponseTimeOk returns a tuple with the RequestsResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsResponseTime

`func (o *DeploymentScalingTarget) SetRequestsResponseTime(v DeploymentScalingTargetRequestsResponseTime)`

SetRequestsResponseTime sets RequestsResponseTime field to given value.

### HasRequestsResponseTime

`func (o *DeploymentScalingTarget) HasRequestsResponseTime() bool`

HasRequestsResponseTime returns a boolean if a field has been set.

### GetSleepIdleDelay

`func (o *DeploymentScalingTarget) GetSleepIdleDelay() DeploymentScalingTargetSleepIdleDelay`

GetSleepIdleDelay returns the SleepIdleDelay field if non-nil, zero value otherwise.

### GetSleepIdleDelayOk

`func (o *DeploymentScalingTarget) GetSleepIdleDelayOk() (*DeploymentScalingTargetSleepIdleDelay, bool)`

GetSleepIdleDelayOk returns a tuple with the SleepIdleDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepIdleDelay

`func (o *DeploymentScalingTarget) SetSleepIdleDelay(v DeploymentScalingTargetSleepIdleDelay)`

SetSleepIdleDelay sets SleepIdleDelay field to given value.

### HasSleepIdleDelay

`func (o *DeploymentScalingTarget) HasSleepIdleDelay() bool`

HasSleepIdleDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


