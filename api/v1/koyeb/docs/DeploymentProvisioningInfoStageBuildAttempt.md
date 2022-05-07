# DeploymentProvisioningInfoStageBuildAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**DeploymentProvisioningInfoStageStatus**](DeploymentProvisioningInfoStageStatus.md) |  | [optional] [default to DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN]
**Messages** | Pointer to **[]string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeploymentProvisioningInfoStageBuildAttempt

`func NewDeploymentProvisioningInfoStageBuildAttempt() *DeploymentProvisioningInfoStageBuildAttempt`

NewDeploymentProvisioningInfoStageBuildAttempt instantiates a new DeploymentProvisioningInfoStageBuildAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProvisioningInfoStageBuildAttemptWithDefaults

`func NewDeploymentProvisioningInfoStageBuildAttemptWithDefaults() *DeploymentProvisioningInfoStageBuildAttempt`

NewDeploymentProvisioningInfoStageBuildAttemptWithDefaults instantiates a new DeploymentProvisioningInfoStageBuildAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStatus() DeploymentProvisioningInfoStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetStatus(v DeploymentProvisioningInfoStageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetStartedAt

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


