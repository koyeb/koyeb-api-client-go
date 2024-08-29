# DeploymentProvisioningInfoStageBuildAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**DeploymentProvisioningInfoStageStatus**](DeploymentProvisioningInfoStageStatus.md) |  | [optional] [default to DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN]
**Messages** | Pointer to **[]string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Steps** | Pointer to [**[]DeploymentProvisioningInfoStageBuildAttemptBuildStep**](DeploymentProvisioningInfoStageBuildAttemptBuildStep.md) |  | [optional] 
**ImagePushed** | Pointer to **bool** |  | [optional] 
**InternalFailure** | Pointer to **bool** |  | [optional] 
**RetryableFailure** | Pointer to **bool** |  | [optional] 
**WaitCompletion** | Pointer to **bool** | This flag is used to finalize the build, and continue the deployment in case of success, or cancel and potentially retry the build in case of failure. | [optional] 

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

### GetSteps

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetSteps() []DeploymentProvisioningInfoStageBuildAttemptBuildStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStepsOk() (*[]DeploymentProvisioningInfoStageBuildAttemptBuildStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetSteps(v []DeploymentProvisioningInfoStageBuildAttemptBuildStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetImagePushed

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetImagePushed() bool`

GetImagePushed returns the ImagePushed field if non-nil, zero value otherwise.

### GetImagePushedOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetImagePushedOk() (*bool, bool)`

GetImagePushedOk returns a tuple with the ImagePushed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePushed

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetImagePushed(v bool)`

SetImagePushed sets ImagePushed field to given value.

### HasImagePushed

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasImagePushed() bool`

HasImagePushed returns a boolean if a field has been set.

### GetInternalFailure

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetInternalFailure() bool`

GetInternalFailure returns the InternalFailure field if non-nil, zero value otherwise.

### GetInternalFailureOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetInternalFailureOk() (*bool, bool)`

GetInternalFailureOk returns a tuple with the InternalFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalFailure

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetInternalFailure(v bool)`

SetInternalFailure sets InternalFailure field to given value.

### HasInternalFailure

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasInternalFailure() bool`

HasInternalFailure returns a boolean if a field has been set.

### GetRetryableFailure

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetRetryableFailure() bool`

GetRetryableFailure returns the RetryableFailure field if non-nil, zero value otherwise.

### GetRetryableFailureOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetRetryableFailureOk() (*bool, bool)`

GetRetryableFailureOk returns a tuple with the RetryableFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryableFailure

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetRetryableFailure(v bool)`

SetRetryableFailure sets RetryableFailure field to given value.

### HasRetryableFailure

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasRetryableFailure() bool`

HasRetryableFailure returns a boolean if a field has been set.

### GetWaitCompletion

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetWaitCompletion() bool`

GetWaitCompletion returns the WaitCompletion field if non-nil, zero value otherwise.

### GetWaitCompletionOk

`func (o *DeploymentProvisioningInfoStageBuildAttempt) GetWaitCompletionOk() (*bool, bool)`

GetWaitCompletionOk returns a tuple with the WaitCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitCompletion

`func (o *DeploymentProvisioningInfoStageBuildAttempt) SetWaitCompletion(v bool)`

SetWaitCompletion sets WaitCompletion field to given value.

### HasWaitCompletion

`func (o *DeploymentProvisioningInfoStageBuildAttempt) HasWaitCompletion() bool`

HasWaitCompletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


