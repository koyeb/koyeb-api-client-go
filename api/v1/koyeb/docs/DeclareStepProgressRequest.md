# DeclareStepProgressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentProvisioningInfoStageStatus**](DeploymentProvisioningInfoStageStatus.md) |  | [optional] [default to DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN]
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeclareStepProgressRequest

`func NewDeclareStepProgressRequest() *DeclareStepProgressRequest`

NewDeclareStepProgressRequest instantiates a new DeclareStepProgressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclareStepProgressRequestWithDefaults

`func NewDeclareStepProgressRequestWithDefaults() *DeclareStepProgressRequest`

NewDeclareStepProgressRequestWithDefaults instantiates a new DeclareStepProgressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *DeclareStepProgressRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DeclareStepProgressRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DeclareStepProgressRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DeclareStepProgressRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *DeclareStepProgressRequest) GetStatus() DeploymentProvisioningInfoStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeclareStepProgressRequest) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeclareStepProgressRequest) SetStatus(v DeploymentProvisioningInfoStageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeclareStepProgressRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *DeclareStepProgressRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DeclareStepProgressRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DeclareStepProgressRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DeclareStepProgressRequest) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DeclareStepProgressRequest) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DeclareStepProgressRequest) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DeclareStepProgressRequest) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DeclareStepProgressRequest) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMessages

`func (o *DeclareStepProgressRequest) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeclareStepProgressRequest) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeclareStepProgressRequest) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeclareStepProgressRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


