# TriggerGitDeploymentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**TriggerGitDeploymentMetadataProvider**](TriggerGitDeploymentMetadataProvider.md) |  | [optional] [default to TRIGGERGITDEPLOYMENTMETADATAPROVIDER_UNKNOWN]
**Repository** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SenderUsername** | Pointer to **string** |  | [optional] 
**SenderAvatarUrl** | Pointer to **string** |  | [optional] 
**SenderProfileUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewTriggerGitDeploymentMetadata

`func NewTriggerGitDeploymentMetadata() *TriggerGitDeploymentMetadata`

NewTriggerGitDeploymentMetadata instantiates a new TriggerGitDeploymentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerGitDeploymentMetadataWithDefaults

`func NewTriggerGitDeploymentMetadataWithDefaults() *TriggerGitDeploymentMetadata`

NewTriggerGitDeploymentMetadataWithDefaults instantiates a new TriggerGitDeploymentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *TriggerGitDeploymentMetadata) GetProvider() TriggerGitDeploymentMetadataProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TriggerGitDeploymentMetadata) GetProviderOk() (*TriggerGitDeploymentMetadataProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TriggerGitDeploymentMetadata) SetProvider(v TriggerGitDeploymentMetadataProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TriggerGitDeploymentMetadata) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRepository

`func (o *TriggerGitDeploymentMetadata) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TriggerGitDeploymentMetadata) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TriggerGitDeploymentMetadata) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *TriggerGitDeploymentMetadata) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBranch

`func (o *TriggerGitDeploymentMetadata) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *TriggerGitDeploymentMetadata) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *TriggerGitDeploymentMetadata) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *TriggerGitDeploymentMetadata) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSha

`func (o *TriggerGitDeploymentMetadata) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *TriggerGitDeploymentMetadata) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *TriggerGitDeploymentMetadata) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *TriggerGitDeploymentMetadata) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetMessage

`func (o *TriggerGitDeploymentMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TriggerGitDeploymentMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TriggerGitDeploymentMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TriggerGitDeploymentMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSenderUsername

`func (o *TriggerGitDeploymentMetadata) GetSenderUsername() string`

GetSenderUsername returns the SenderUsername field if non-nil, zero value otherwise.

### GetSenderUsernameOk

`func (o *TriggerGitDeploymentMetadata) GetSenderUsernameOk() (*string, bool)`

GetSenderUsernameOk returns a tuple with the SenderUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUsername

`func (o *TriggerGitDeploymentMetadata) SetSenderUsername(v string)`

SetSenderUsername sets SenderUsername field to given value.

### HasSenderUsername

`func (o *TriggerGitDeploymentMetadata) HasSenderUsername() bool`

HasSenderUsername returns a boolean if a field has been set.

### GetSenderAvatarUrl

`func (o *TriggerGitDeploymentMetadata) GetSenderAvatarUrl() string`

GetSenderAvatarUrl returns the SenderAvatarUrl field if non-nil, zero value otherwise.

### GetSenderAvatarUrlOk

`func (o *TriggerGitDeploymentMetadata) GetSenderAvatarUrlOk() (*string, bool)`

GetSenderAvatarUrlOk returns a tuple with the SenderAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAvatarUrl

`func (o *TriggerGitDeploymentMetadata) SetSenderAvatarUrl(v string)`

SetSenderAvatarUrl sets SenderAvatarUrl field to given value.

### HasSenderAvatarUrl

`func (o *TriggerGitDeploymentMetadata) HasSenderAvatarUrl() bool`

HasSenderAvatarUrl returns a boolean if a field has been set.

### GetSenderProfileUrl

`func (o *TriggerGitDeploymentMetadata) GetSenderProfileUrl() string`

GetSenderProfileUrl returns the SenderProfileUrl field if non-nil, zero value otherwise.

### GetSenderProfileUrlOk

`func (o *TriggerGitDeploymentMetadata) GetSenderProfileUrlOk() (*string, bool)`

GetSenderProfileUrlOk returns a tuple with the SenderProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderProfileUrl

`func (o *TriggerGitDeploymentMetadata) SetSenderProfileUrl(v string)`

SetSenderProfileUrl sets SenderProfileUrl field to given value.

### HasSenderProfileUrl

`func (o *TriggerGitDeploymentMetadata) HasSenderProfileUrl() bool`

HasSenderProfileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


