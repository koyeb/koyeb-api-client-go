# GithubInstallationRepoListReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]GithubRepo**](GithubRepo.md) |  | [optional] 
**InstallationId** | Pointer to **string** |  | [optional] 
**InstallationUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubInstallationRepoListReply

`func NewGithubInstallationRepoListReply() *GithubInstallationRepoListReply`

NewGithubInstallationRepoListReply instantiates a new GithubInstallationRepoListReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubInstallationRepoListReplyWithDefaults

`func NewGithubInstallationRepoListReplyWithDefaults() *GithubInstallationRepoListReply`

NewGithubInstallationRepoListReplyWithDefaults instantiates a new GithubInstallationRepoListReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *GithubInstallationRepoListReply) GetRepositories() []GithubRepo`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *GithubInstallationRepoListReply) GetRepositoriesOk() (*[]GithubRepo, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *GithubInstallationRepoListReply) SetRepositories(v []GithubRepo)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *GithubInstallationRepoListReply) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetInstallationId

`func (o *GithubInstallationRepoListReply) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *GithubInstallationRepoListReply) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *GithubInstallationRepoListReply) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *GithubInstallationRepoListReply) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetInstallationUrl

`func (o *GithubInstallationRepoListReply) GetInstallationUrl() string`

GetInstallationUrl returns the InstallationUrl field if non-nil, zero value otherwise.

### GetInstallationUrlOk

`func (o *GithubInstallationRepoListReply) GetInstallationUrlOk() (*string, bool)`

GetInstallationUrlOk returns a tuple with the InstallationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationUrl

`func (o *GithubInstallationRepoListReply) SetInstallationUrl(v string)`

SetInstallationUrl sets InstallationUrl field to given value.

### HasInstallationUrl

`func (o *GithubInstallationRepoListReply) HasInstallationUrl() bool`

HasInstallationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


