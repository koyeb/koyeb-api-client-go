# DeploymentProvisioningInfoStageBuildAttemptBuildStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentProvisioningInfoStageStatus**](DeploymentProvisioningInfoStageStatus.md) |  | [optional] [default to DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN]
**Messages** | Pointer to **[]string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeploymentProvisioningInfoStageBuildAttemptBuildStep

`func NewDeploymentProvisioningInfoStageBuildAttemptBuildStep() *DeploymentProvisioningInfoStageBuildAttemptBuildStep`

NewDeploymentProvisioningInfoStageBuildAttemptBuildStep instantiates a new DeploymentProvisioningInfoStageBuildAttemptBuildStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProvisioningInfoStageBuildAttemptBuildStepWithDefaults

`func NewDeploymentProvisioningInfoStageBuildAttemptBuildStepWithDefaults() *DeploymentProvisioningInfoStageBuildAttemptBuildStep`

NewDeploymentProvisioningInfoStageBuildAttemptBuildStepWithDefaults instantiates a new DeploymentProvisioningInfoStageBuildAttemptBuildStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStatus() DeploymentProvisioningInfoStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetStatus(v DeploymentProvisioningInfoStageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetStartedAt

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DeploymentProvisioningInfoStageBuildAttemptBuildStep) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


