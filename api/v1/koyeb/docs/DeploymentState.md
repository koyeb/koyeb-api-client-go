# DeploymentState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]ServiceInstance**](ServiceInstance.md) |  | [optional] 
**Status** | Pointer to [**DeploymentStateStatus**](DeploymentStateStatus.md) |  | [optional] [default to DEPLOYMENTSTATESTATUS_UNKNOWN]
**StatusMessage** | Pointer to **string** |  | [optional] 
**Datacenters** | Pointer to **[]string** |  | [optional] 
**BuildInfo** | Pointer to [**DeploymentStateBuildInfo**](DeploymentStateBuildInfo.md) |  | [optional] 

## Methods

### NewDeploymentState

`func NewDeploymentState() *DeploymentState`

NewDeploymentState instantiates a new DeploymentState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStateWithDefaults

`func NewDeploymentStateWithDefaults() *DeploymentState`

NewDeploymentStateWithDefaults instantiates a new DeploymentState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *DeploymentState) GetInstances() []ServiceInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *DeploymentState) GetInstancesOk() (*[]ServiceInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *DeploymentState) SetInstances(v []ServiceInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *DeploymentState) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentState) GetStatus() DeploymentStateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentState) GetStatusOk() (*DeploymentStateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentState) SetStatus(v DeploymentStateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DeploymentState) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DeploymentState) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DeploymentState) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DeploymentState) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDatacenters

`func (o *DeploymentState) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *DeploymentState) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *DeploymentState) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *DeploymentState) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetBuildInfo

`func (o *DeploymentState) GetBuildInfo() DeploymentStateBuildInfo`

GetBuildInfo returns the BuildInfo field if non-nil, zero value otherwise.

### GetBuildInfoOk

`func (o *DeploymentState) GetBuildInfoOk() (*DeploymentStateBuildInfo, bool)`

GetBuildInfoOk returns a tuple with the BuildInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfo

`func (o *DeploymentState) SetBuildInfo(v DeploymentStateBuildInfo)`

SetBuildInfo sets BuildInfo field to given value.

### HasBuildInfo

`func (o *DeploymentState) HasBuildInfo() bool`

HasBuildInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


