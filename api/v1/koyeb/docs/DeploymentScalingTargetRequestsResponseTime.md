# DeploymentScalingTargetRequestsResponseTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int64** |  | [optional] 
**Quantile** | Pointer to **int64** | The quantile to use for autoscaling. For example, set to 95 to use the 95th percentile (p95) for autoscaling.  Valid values are between 0 and 100. | [optional] 

## Methods

### NewDeploymentScalingTargetRequestsResponseTime

`func NewDeploymentScalingTargetRequestsResponseTime() *DeploymentScalingTargetRequestsResponseTime`

NewDeploymentScalingTargetRequestsResponseTime instantiates a new DeploymentScalingTargetRequestsResponseTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentScalingTargetRequestsResponseTimeWithDefaults

`func NewDeploymentScalingTargetRequestsResponseTimeWithDefaults() *DeploymentScalingTargetRequestsResponseTime`

NewDeploymentScalingTargetRequestsResponseTimeWithDefaults instantiates a new DeploymentScalingTargetRequestsResponseTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DeploymentScalingTargetRequestsResponseTime) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeploymentScalingTargetRequestsResponseTime) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeploymentScalingTargetRequestsResponseTime) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeploymentScalingTargetRequestsResponseTime) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetQuantile

`func (o *DeploymentScalingTargetRequestsResponseTime) GetQuantile() int64`

GetQuantile returns the Quantile field if non-nil, zero value otherwise.

### GetQuantileOk

`func (o *DeploymentScalingTargetRequestsResponseTime) GetQuantileOk() (*int64, bool)`

GetQuantileOk returns a tuple with the Quantile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantile

`func (o *DeploymentScalingTargetRequestsResponseTime) SetQuantile(v int64)`

SetQuantile sets Quantile field to given value.

### HasQuantile

`func (o *DeploymentScalingTargetRequestsResponseTime) HasQuantile() bool`

HasQuantile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


