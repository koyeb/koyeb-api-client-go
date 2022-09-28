# DeploymentHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GracePeriod** | Pointer to **int64** |  | [optional] 
**Interval** | Pointer to **int64** |  | [optional] 
**RestartLimit** | Pointer to **int64** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Tcp** | Pointer to [**TCPHealthCheck**](TCPHealthCheck.md) |  | [optional] 
**Http** | Pointer to [**HTTPHealthCheck**](HTTPHealthCheck.md) |  | [optional] 

## Methods

### NewDeploymentHealthCheck

`func NewDeploymentHealthCheck() *DeploymentHealthCheck`

NewDeploymentHealthCheck instantiates a new DeploymentHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHealthCheckWithDefaults

`func NewDeploymentHealthCheckWithDefaults() *DeploymentHealthCheck`

NewDeploymentHealthCheckWithDefaults instantiates a new DeploymentHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGracePeriod

`func (o *DeploymentHealthCheck) GetGracePeriod() int64`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *DeploymentHealthCheck) GetGracePeriodOk() (*int64, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *DeploymentHealthCheck) SetGracePeriod(v int64)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *DeploymentHealthCheck) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetInterval

`func (o *DeploymentHealthCheck) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DeploymentHealthCheck) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DeploymentHealthCheck) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DeploymentHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRestartLimit

`func (o *DeploymentHealthCheck) GetRestartLimit() int64`

GetRestartLimit returns the RestartLimit field if non-nil, zero value otherwise.

### GetRestartLimitOk

`func (o *DeploymentHealthCheck) GetRestartLimitOk() (*int64, bool)`

GetRestartLimitOk returns a tuple with the RestartLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartLimit

`func (o *DeploymentHealthCheck) SetRestartLimit(v int64)`

SetRestartLimit sets RestartLimit field to given value.

### HasRestartLimit

`func (o *DeploymentHealthCheck) HasRestartLimit() bool`

HasRestartLimit returns a boolean if a field has been set.

### GetTimeout

`func (o *DeploymentHealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DeploymentHealthCheck) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DeploymentHealthCheck) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DeploymentHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTcp

`func (o *DeploymentHealthCheck) GetTcp() TCPHealthCheck`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *DeploymentHealthCheck) GetTcpOk() (*TCPHealthCheck, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *DeploymentHealthCheck) SetTcp(v TCPHealthCheck)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *DeploymentHealthCheck) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetHttp

`func (o *DeploymentHealthCheck) GetHttp() HTTPHealthCheck`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *DeploymentHealthCheck) GetHttpOk() (*HTTPHealthCheck, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *DeploymentHealthCheck) SetHttp(v HTTPHealthCheck)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *DeploymentHealthCheck) HasHttp() bool`

HasHttp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


