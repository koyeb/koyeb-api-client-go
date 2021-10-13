# GitDeploymentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**GitDeploymentMetadataProvider**](GitDeploymentMetadataProvider.md) |  | [optional] [default to GITDEPLOYMENTMETADATAPROVIDER_UNKNOWN]
**Repository** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SenderUsername** | Pointer to **string** |  | [optional] 
**SenderAvatarUrl** | Pointer to **string** |  | [optional] 
**SenderProfileUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewGitDeploymentMetadata

`func NewGitDeploymentMetadata() *GitDeploymentMetadata`

NewGitDeploymentMetadata instantiates a new GitDeploymentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitDeploymentMetadataWithDefaults

`func NewGitDeploymentMetadataWithDefaults() *GitDeploymentMetadata`

NewGitDeploymentMetadataWithDefaults instantiates a new GitDeploymentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *GitDeploymentMetadata) GetProvider() GitDeploymentMetadataProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitDeploymentMetadata) GetProviderOk() (*GitDeploymentMetadataProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitDeploymentMetadata) SetProvider(v GitDeploymentMetadataProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GitDeploymentMetadata) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRepository

`func (o *GitDeploymentMetadata) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GitDeploymentMetadata) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GitDeploymentMetadata) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GitDeploymentMetadata) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBranch

`func (o *GitDeploymentMetadata) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitDeploymentMetadata) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitDeploymentMetadata) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitDeploymentMetadata) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSha

`func (o *GitDeploymentMetadata) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitDeploymentMetadata) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitDeploymentMetadata) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitDeploymentMetadata) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetMessage

`func (o *GitDeploymentMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitDeploymentMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitDeploymentMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitDeploymentMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSenderUsername

`func (o *GitDeploymentMetadata) GetSenderUsername() string`

GetSenderUsername returns the SenderUsername field if non-nil, zero value otherwise.

### GetSenderUsernameOk

`func (o *GitDeploymentMetadata) GetSenderUsernameOk() (*string, bool)`

GetSenderUsernameOk returns a tuple with the SenderUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUsername

`func (o *GitDeploymentMetadata) SetSenderUsername(v string)`

SetSenderUsername sets SenderUsername field to given value.

### HasSenderUsername

`func (o *GitDeploymentMetadata) HasSenderUsername() bool`

HasSenderUsername returns a boolean if a field has been set.

### GetSenderAvatarUrl

`func (o *GitDeploymentMetadata) GetSenderAvatarUrl() string`

GetSenderAvatarUrl returns the SenderAvatarUrl field if non-nil, zero value otherwise.

### GetSenderAvatarUrlOk

`func (o *GitDeploymentMetadata) GetSenderAvatarUrlOk() (*string, bool)`

GetSenderAvatarUrlOk returns a tuple with the SenderAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAvatarUrl

`func (o *GitDeploymentMetadata) SetSenderAvatarUrl(v string)`

SetSenderAvatarUrl sets SenderAvatarUrl field to given value.

### HasSenderAvatarUrl

`func (o *GitDeploymentMetadata) HasSenderAvatarUrl() bool`

HasSenderAvatarUrl returns a boolean if a field has been set.

### GetSenderProfileUrl

`func (o *GitDeploymentMetadata) GetSenderProfileUrl() string`

GetSenderProfileUrl returns the SenderProfileUrl field if non-nil, zero value otherwise.

### GetSenderProfileUrlOk

`func (o *GitDeploymentMetadata) GetSenderProfileUrlOk() (*string, bool)`

GetSenderProfileUrlOk returns a tuple with the SenderProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderProfileUrl

`func (o *GitDeploymentMetadata) SetSenderProfileUrl(v string)`

SetSenderProfileUrl sets SenderProfileUrl field to given value.

### HasSenderProfileUrl

`func (o *GitDeploymentMetadata) HasSenderProfileUrl() bool`

HasSenderProfileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


