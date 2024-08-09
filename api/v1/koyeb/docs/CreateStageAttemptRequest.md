# CreateStageAttemptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentProvisioningInfoStageStatus**](DeploymentProvisioningInfoStageStatus.md) |  | [optional] [default to DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN]
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**Steps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateStageAttemptRequest

`func NewCreateStageAttemptRequest() *CreateStageAttemptRequest`

NewCreateStageAttemptRequest instantiates a new CreateStageAttemptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStageAttemptRequestWithDefaults

`func NewCreateStageAttemptRequestWithDefaults() *CreateStageAttemptRequest`

NewCreateStageAttemptRequestWithDefaults instantiates a new CreateStageAttemptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *CreateStageAttemptRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateStageAttemptRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateStageAttemptRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateStageAttemptRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *CreateStageAttemptRequest) GetStatus() DeploymentProvisioningInfoStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateStageAttemptRequest) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateStageAttemptRequest) SetStatus(v DeploymentProvisioningInfoStageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateStageAttemptRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *CreateStageAttemptRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CreateStageAttemptRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CreateStageAttemptRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CreateStageAttemptRequest) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetMessages

`func (o *CreateStageAttemptRequest) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CreateStageAttemptRequest) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CreateStageAttemptRequest) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CreateStageAttemptRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetSteps

`func (o *CreateStageAttemptRequest) GetSteps() []string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CreateStageAttemptRequest) GetStepsOk() (*[]string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CreateStageAttemptRequest) SetSteps(v []string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CreateStageAttemptRequest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


