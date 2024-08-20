# DeclareStageProgressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentProvisioningInfoStageStatus**](DeploymentProvisioningInfoStageStatus.md) |  | [optional] [default to DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN]
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**ImagePushed** | Pointer to **bool** |  | [optional] 
**InternalFailure** | Pointer to **bool** |  | [optional] 
**RetryableFailure** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeclareStageProgressRequest

`func NewDeclareStageProgressRequest() *DeclareStageProgressRequest`

NewDeclareStageProgressRequest instantiates a new DeclareStageProgressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclareStageProgressRequestWithDefaults

`func NewDeclareStageProgressRequestWithDefaults() *DeclareStageProgressRequest`

NewDeclareStageProgressRequestWithDefaults instantiates a new DeclareStageProgressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *DeclareStageProgressRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DeclareStageProgressRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DeclareStageProgressRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DeclareStageProgressRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *DeclareStageProgressRequest) GetStatus() DeploymentProvisioningInfoStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeclareStageProgressRequest) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeclareStageProgressRequest) SetStatus(v DeploymentProvisioningInfoStageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeclareStageProgressRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFinishedAt

`func (o *DeclareStageProgressRequest) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *DeclareStageProgressRequest) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *DeclareStageProgressRequest) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *DeclareStageProgressRequest) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMessages

`func (o *DeclareStageProgressRequest) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeclareStageProgressRequest) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeclareStageProgressRequest) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeclareStageProgressRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetImagePushed

`func (o *DeclareStageProgressRequest) GetImagePushed() bool`

GetImagePushed returns the ImagePushed field if non-nil, zero value otherwise.

### GetImagePushedOk

`func (o *DeclareStageProgressRequest) GetImagePushedOk() (*bool, bool)`

GetImagePushedOk returns a tuple with the ImagePushed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePushed

`func (o *DeclareStageProgressRequest) SetImagePushed(v bool)`

SetImagePushed sets ImagePushed field to given value.

### HasImagePushed

`func (o *DeclareStageProgressRequest) HasImagePushed() bool`

HasImagePushed returns a boolean if a field has been set.

### GetInternalFailure

`func (o *DeclareStageProgressRequest) GetInternalFailure() bool`

GetInternalFailure returns the InternalFailure field if non-nil, zero value otherwise.

### GetInternalFailureOk

`func (o *DeclareStageProgressRequest) GetInternalFailureOk() (*bool, bool)`

GetInternalFailureOk returns a tuple with the InternalFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalFailure

`func (o *DeclareStageProgressRequest) SetInternalFailure(v bool)`

SetInternalFailure sets InternalFailure field to given value.

### HasInternalFailure

`func (o *DeclareStageProgressRequest) HasInternalFailure() bool`

HasInternalFailure returns a boolean if a field has been set.

### GetRetryableFailure

`func (o *DeclareStageProgressRequest) GetRetryableFailure() bool`

GetRetryableFailure returns the RetryableFailure field if non-nil, zero value otherwise.

### GetRetryableFailureOk

`func (o *DeclareStageProgressRequest) GetRetryableFailureOk() (*bool, bool)`

GetRetryableFailureOk returns a tuple with the RetryableFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryableFailure

`func (o *DeclareStageProgressRequest) SetRetryableFailure(v bool)`

SetRetryableFailure sets RetryableFailure field to given value.

### HasRetryableFailure

`func (o *DeclareStageProgressRequest) HasRetryableFailure() bool`

HasRetryableFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


